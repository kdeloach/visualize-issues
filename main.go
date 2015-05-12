package main

import (
    "fmt"
    "time"
    "encoding/json"
    "github.com/google/go-github/github"
    "golang.org/x/oauth2"
)

type repoTuple struct {
    owner string
    name string
}

type issueJsonObj struct {
    Number int `json:"number,omitempty"`
    Title string `json:"title,omitempty"`
    Url string `json:"url,omitempty"`
    HtmlUrl string `json:"html_url,omitempty"`
    Milestone string `json:"milestone,omitempty"`
    Status string `json:"status,omitempty"`
}

// tokenSource is an oauth2.TokenSource which returns a static access token
type tokenSource struct {
    token *oauth2.Token
}

// Token implements the oauth2.TokenSource interface
func (t *tokenSource) Token() (*oauth2.Token, error){
    return t.token, nil
}

func makeClient() *github.Client {
    ts := &tokenSource{
        &oauth2.Token{AccessToken: "813a87727d1bc0862cd3b1908fb1060fe4b2f3cf"},
    }
    tc := oauth2.NewClient(oauth2.NoContext, ts)
    client := github.NewClient(tc)
    client.UserAgent = "kdeloach@azavea.com"
    return client
}

func fetchRepoIssues(client *github.Client, repos []repoTuple) []github.Issue {
    result := make([]github.Issue, 0)
    for _, repo := range repos {
        issues := fetchIssues(client, repo)
        result = append(result, issues...)
    }
    return result
}

func fetchIssues(client *github.Client, repo repoTuple) []github.Issue {
    result := make([]github.Issue, 0)
    opt := &github.IssueListByRepoOptions{
        Milestone: "*",
        ListOptions: github.ListOptions{
            PerPage: 30,
        },
    }
    for {
        issues, resp, err := client.Issues.ListByRepo(repo.owner, repo.name, opt)
        if err != nil {
            break
        }
        result = append(result, issues...)
        if resp.NextPage == 0 {
            break
        }
        opt.ListOptions.Page = resp.NextPage
        time.Sleep(500 * time.Millisecond)
    }
    return result
}

func makeIssuesJson(issues []github.Issue) []issueJsonObj {
    result := make([]issueJsonObj, 0)
    for _, issue := range issues {
        jsonObj := issueJsonObj{
            Number: *issue.Number,
            Title: *issue.Title,
            Url: *issue.URL,
            HtmlUrl: *issue.HTMLURL,
            Milestone: *issue.Milestone.Title,
            Status: getIssueStatus(issue),
        }
        result = append(result, jsonObj)
    }
    return result
}

func getIssueStatus(issue github.Issue) string {
    for _, label := range issue.Labels {
        name := *label.Name
        if name == "queue" || name == "in progress" || name == "in review" {
            return name
        }
    }
    return "backlog";
}

func main() {
    client := makeClient()
    repos := []repoTuple{
        repoTuple{"azavea", "backbone.hashmodels"},
        repoTuple{"azavea", "civic-apps"},
        repoTuple{"azavea", "django-queryset-csv"},
        repoTuple{"azavea", "django-tinsel"},
        repoTuple{"azavea", "dor-parcel-explorer"},
        repoTuple{"azavea", "dor-philly-history-blog"},
        repoTuple{"azavea", "lr-common"},
        repoTuple{"azavea", "model-my-watershed"},
        repoTuple{"azavea", "nyc-trees"},
        repoTuple{"azavea", "oit-ulrs"},
        repoTuple{"azavea", "OpenTreeMap-iOS-skin"},
        repoTuple{"azavea", "pwd-stormwater-allocation"},
        repoTuple{"azavea", "pwd-stormwater-interactive"},
        repoTuple{"azavea", "pwd-waterworks-revealed"},
        repoTuple{"azavea", "stanford-campus-map"},
        repoTuple{"azavea", "tr-55"},
        repoTuple{"azavea", "usace-wisdm"},
        repoTuple{"azavea", "usace-wisdm-filter-viewer"},
        repoTuple{"azavea", "usace-wisdm-filter-viewer-data"},
        repoTuple{"azavea", "usace-wisdm-symbolizer"},
        repoTuple{"azavea", "usace-wisdm-symbolizer-data"},
        repoTuple{"CoastalResilienceNetwork", "GeositeFramework"},
        repoTuple{"OpenTreeMap", "clients"},
        repoTuple{"OpenTreeMap", "cloudbuild"},
        repoTuple{"OpenTreeMap", "ecobenefits"},
        repoTuple{"OpenTreeMap", "OpenTreeMap"},
        repoTuple{"OpenTreeMap", "OpenTreeMap-Android"},
        repoTuple{"OpenTreeMap", "OpenTreeMap-iOS"},
        repoTuple{"OpenTreeMap", "OpenTreeMap-Modeling"},
        repoTuple{"OpenTreeMap", "opentreemap.github.com"},
        repoTuple{"OpenTreeMap", "otm-mobile-skins"},
        repoTuple{"OpenTreeMap", "otm-wordpress"},
        repoTuple{"OpenTreeMap", "otm-wordpress-release-scripts"},
        repoTuple{"OpenTreeMap", "OTM2"},
        repoTuple{"OpenTreeMap", "otm2-addons"},
        repoTuple{"OpenTreeMap", "OTM2-tiler"},
        repoTuple{"OpenTreeMap", "otm2-vagrant"},
    }

    issues := fetchRepoIssues(client, repos)
    json_issues := makeIssuesJson(issues)
    output, err := json.Marshal(json_issues)
    if err != nil {
        fmt.Println(err)
        return
    }

    fmt.Println(string(output))
}
