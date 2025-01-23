package e2e

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"itl/e2e/fixtures"
	"itl/model/domain"
	"testing"
)

type groupTestProfile struct {
	act func(t *testing.T)
}

func TestGroup(t *testing.T) {

	groupFixture := fixtures.NewGroupFixture()
	groupTests := map[string]groupTestProfile{
		"create": {
			act: func(t *testing.T) {
				client := fixtures.NewClient()
				user := groupFixture.GetDefaultUser(&domain.User{
					Username: uuid.New().String(),
				})

				group := domain.Group{
					Name: uuid.New().String(),
				}
				//create
				groupResp := &domain.Group{}
				resp, err := client.R().SetHeader("Authorization", user.ID.String()).SetBody(&group).SetResult(groupResp).Post("/v1/group/create")
				assert.Nil(t, err)
				assert.Equal(t, 200, resp.StatusCode())
				assert.NotEmpty(t, resp.Body())
				assert.Equal(t, group.Name, groupResp.Name)
				assert.Equal(t, user.ID, groupResp.OwnerID)

				//get
				groupResp2 := &domain.Group{}
				resp, err = client.R().SetHeader("Authorization", user.ID.String()).SetResult(groupResp2).Get("/v1/group/" + groupResp.ID.String())
				assert.Nil(t, err)
				assert.Equal(t, 200, resp.StatusCode())
				assert.NotEmpty(t, resp.Body())
				assert.Equal(t, group.Name, groupResp2.Name)
				assert.Equal(t, user.ID, groupResp2.OwnerID)
				assert.Empty(t, groupResp2.Posts)

				post := domain.Post{
					UserID:      user.ID,
					GroupID:     groupResp.ID,
					Uri:         "http://testuri",
					Description: "test desc",
				}
				postCreated := groupFixture.CreatePostInGroup(&post)

				var postResp []domain.Post
				resp, err = client.R().SetHeader("Authorization", user.ID.String()).SetResult(&postResp).Get("/v1/group/" + groupResp.ID.String() + "/posts")
				assert.Nil(t, err)
				assert.Equal(t, 200, resp.StatusCode())
				assert.NotEmpty(t, resp.Body())
				assert.NotEmpty(t, postResp)
				assert.Equal(t, postCreated.GroupID, postResp[0].GroupID)

				//Add user to group
				user2 := groupFixture.GetDefaultUser(&domain.User{
					Username: uuid.New().String(),
				})
				groupFixture.AddUserToGroup(user2.ID, groupResp.ID)
				var users []domain.User
				resp, err = client.R().SetHeader("Authorization", user.ID.String()).SetResult(&users).Get("/v1/group/" + groupResp.ID.String() + "/users")
				assert.Nil(t, err)
				assert.Equal(t, 200, resp.StatusCode())
				assert.NotEmpty(t, resp.Body())
				assert.NotEmpty(t, users)
				assert.Equal(t, user2.ID, users[0].ID)

			},
		},
	}

	for name, c := range groupTests {
		t.Run(name, func(t *testing.T) {
			c.act(t)
		})
	}
}
