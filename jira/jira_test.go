package jira

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	api "gopkg.in/andygrunwald/go-jira.v1"
	"github.com/erdembozdg/erdem-go/jira/mocks"
	"go.uber.org/zap"
)

func TestSearchJiras(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockJiraer := mocks.NewMockJiraer(ctrl)

	i1 := api.Issue{
		ID:  "111",
		Key: "Test-1",
		Fields: &api.IssueFields{
			Summary: "Test-1: bla bla",
			Labels:  []string{"fake1", "fake2"},
		},
	}

	i2 := api.Issue{
		ID:   "222",
		Key:  "Test-2",
		Self: "",
		Fields: &api.IssueFields{
			Summary: "Test-2: bla bla",
			Labels:  []string{"fake1", "fake2"},
		},
	}

	jql := "fake jql"
	baseUrl := "https://fake.base"
	jira := Jira{
		conf: Configuration{
			BaseURL: baseUrl,
			BasicAuth: BasicAuthConfiguration{
				User:  "ebzdag",
				Token: "111",
			},
		},
		client: mockJiraer,
		logger: zap.NewNop(),
	}

	searchResult := []api.Issue{i1, i2}

	mockJiraer.EXPECT().Search(jql).Times(1).Return(&searchResult, nil)
	issues, err := jira.Search(jql)
	assert.NoError(t, err)
	assert.Len(t, issues, 2)
	assert.Equal(t, i1.Key, issues[0].Key)
	assert.Equal(t, i2.Key, issues[1].Key)
	assert.Equal(t, i1.ID, issues[0].ID)
	assert.Equal(t, i2.ID, issues[1].ID)
	assert.Equal(t, i1.Fields.Summary, issues[0].Summary)
	assert.Equal(t, i2.Fields.Summary, issues[1].Summary)
	assert.Equal(t, []string{"fake1", "fake2"}, issues[0].Labels)
	assert.Equal(t, []string{"fake1", "fake2"}, issues[1].Labels)
}
