package e2e

import (
	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/assert"
	"itl/e2e/fixtures"
	"testing"
)

type userTestProfile struct {
	arrange func(t *testing.T)
	act     func(t *testing.T) (*resty.Response, error)
	assert  func(resp *resty.Response, err error)
}

func TestUser(t *testing.T) {
	fixtures.NewSession()
	userTests := map[string]userTestProfile{
		"userNoAuth": {
			arrange: func(t *testing.T) {
			},
			act: func(t *testing.T) (*resty.Response, error) {
				client := fixtures.NewClient()
				resp, err := client.R().Get("/v1/user/groups")
				return resp, err
			},
			assert: func(resp *resty.Response, err error) {
				assert.Nil(t, err)
				assert.Equal(t, 400, resp.StatusCode())
				assert.NotEmpty(t, resp.Body())
			},
		},
		"userWithAuth": {
			arrange: func(t *testing.T) {
			},
			act: func(t *testing.T) (*resty.Response, error) {
				client := fixtures.NewClient()
				resp, err := client.R().SetHeader("Authorization", "2326d0e8-18ba-4034-825b-8b8bdfc15353").Get("/v1/user/groups")
				return resp, err
			},
			assert: func(resp *resty.Response, err error) {
				assert.Nil(t, err)
				assert.Equal(t, 200, resp.StatusCode())
				assert.NotEmpty(t, resp.Body())
			},
		},
	}

	for name, c := range userTests {
		t.Run(name, func(t *testing.T) {
			c.arrange(t)
			resp, err := c.act(t)
			c.assert(resp, err)
		})
	}

}
