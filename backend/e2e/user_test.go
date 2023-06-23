package e2e

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"itl/e2e/fixtures"
	"itl/model"
	"testing"
)

type userTestProfile struct {
	act func(t *testing.T)
}

func TestUser(t *testing.T) {
	userFixture := fixtures.NewUserFixture()
	userTests := map[string]userTestProfile{
		"all": {
			act: func(t *testing.T) {
				client := fixtures.NewClient()
				user := model.User{
					Username: uuid.New().String(),
				}
				userResp := &model.User{}
				resp, err := client.R().SetBody(&user).SetResult(userResp).Post("/v1/user/create")
				assert.Nil(t, err)
				assert.Equal(t, 200, resp.StatusCode())
				assert.NotEmpty(t, resp.Body())
				assert.Equal(t, user.Username, userResp.Username)

				group := &model.Group{
					OwnerID: userResp.ID,
					Name:    uuid.New().String(),
				}
				groupId := userFixture.CreateGroup(group)
				resp, err = client.R().SetHeader("Authorization", userResp.ID.String()).Post("/v1/user/groups/join/" + groupId.String())
				assert.Nil(t, err)
				assert.Equal(t, 200, resp.StatusCode())

				var groupResp []model.Group
				resp, err = client.R().SetResult(&groupResp).SetHeader("Authorization", userResp.ID.String()).Get("/v1/user/groups")
				assert.Nil(t, err)
				assert.Equal(t, 200, resp.StatusCode())
				assert.NotEmpty(t, resp.Body())
				assert.Equal(t, group.ID, groupResp[0].ID)

			},
		},
		"userNoAuth": {
			act: func(t *testing.T) {
				client := fixtures.NewClient()
				resp, err := client.R().Get("/v1/user/groups")
				assert.Nil(t, err)
				assert.Equal(t, 400, resp.StatusCode())
				assert.NotEmpty(t, resp.Body())
			},
		},
		"userWithAuth": {
			act: func(t *testing.T) {
				client := fixtures.NewClient()
				resp, err := client.R().SetHeader("Authorization", "2326d0e8-18ba-4034-825b-8b8bdfc15353").Get("/v1/user/groups")
				assert.Nil(t, err)
				assert.Equal(t, 200, resp.StatusCode())
				assert.NotEmpty(t, resp.Body())
			},
		},
	}

	for name, c := range userTests {
		t.Run(name, func(t *testing.T) {
			c.act(t)
		})
	}

}
