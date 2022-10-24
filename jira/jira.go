package jira

import (
	"go.uber.org/zap"
	api "gopkg.in/andygrunwald/go-jira.v1"
)

// Issue is a Jira issue.
type Issue struct {
	ID      string
	Key     string
	Labels  []string
	Summary string
}

// Configuration configures a Jira client.
type Configuration struct {
	BaseURL   string
	BasicAuth BasicAuthConfiguration
}

// BasicAuthConfiguration configures basic authentication for Jira.
type BasicAuthConfiguration struct {
	User  string
	Token string
}

//go:generate mockgen -destination=mocks/jira.go -package=mocks -source=jira.go
type issuer interface {
	Search(jql string, options *api.SearchOptions) ([]api.Issue, *api.Response, error)
}

// Jira allows querying Jira issues.
type Jira struct {
	conf   Configuration
	issue issuer
	logger *zap.Logger
}

// New creates a Jira or panics if it fails to.
func NewOrFail(conf Configuration, logger *zap.Logger) *Jira {
	jira, err := New(conf, logger)
	if err != nil {
		panic(err)
	}
	return jira
}

// New creates a new Jira.
func New(conf Configuration, logger *zap.Logger) (*Jira, error) {

	tp := api.BasicAuthTransport{
		Username: conf.BasicAuth.User,
		Password: conf.BasicAuth.Token,
	}

	client, err := api.NewClient(tp.Client(), conf.BaseURL)

	if err != nil {
		return nil, err
	}

	return &Jira{
		conf:   conf,
		issue: client.Issue,
		logger: logger,
	}, nil
}

// Search will query Jira API using the provided JQL string
func (j *Jira) Search(jql string) ([]Issue, error) {

	opt := &api.SearchOptions{
		MaxResults: 50, // Set this arbitrarily high so we don't need to deal with pagination yet.
		StartAt:    0,  // Make sure we start grabbing issues from last checkpoint
	}

	results, res, err := j.issue.Search(jql, opt)
	if err != nil || res == nil {
		return nil, err
	}

	j.logger.Debug("found jira issues", zap.Any("jira.search.total", res.Total))

	// If we encounter this error it's time to implement pagination.
	if res.Total > res.MaxResults {
		j.logger.Error("jira search exceeded max results",
			zap.String("jira.search.query", jql),
			zap.Int("jira.search.total", res.Total),
		)
	}

	var issues []Issue
	for _, r := range results {
		issues = append(issues, Issue{
			Key:     r.Key,
			ID:      r.ID,
			Labels:  r.Fields.Labels,
			Summary: r.Fields.Summary,
		})
	}

	return issues, nil
}
