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
        &oauth2.Token{AccessToken: "<GitHub Access Token here>"},
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
        // State: "all",
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

func main() {
    client := makeClient()
    repos := []repoTuple{
        repoTuple{"azavea", "civic-apps"},
        repoTuple{"azavea", "azavea-data-hub"},
        repoTuple{"azavea", "dor-city-owned-property"},
        repoTuple{"azavea", "dor-icris-extracts"},
        repoTuple{"azavea", "dor-parcel-explorer"},
        repoTuple{"azavea", "dor-property-reporting"},
        repoTuple{"CoastalResilienceNetwork", "GeositeFramework"},
        repoTuple{"azavea", "lr-common"},
        repoTuple{"WikiWatershed", "docker-rwd"},
        repoTuple{"WikiWatershed", "mmw-geoprocessing"},
        repoTuple{"WikiWatershed", "rapid-watershed-delineation"},
        repoTuple{"WikiWatershed", "model-my-watershed"},
        repoTuple{"azavea", "modellab"},
        repoTuple{"azavea", "modellab-geoprocessing"},
        repoTuple{"azavea", "pwd-parcel-viewer"},
        repoTuple{"azavea", "pwd-stormwater-allocation"},
        repoTuple{"azavea", "pwd-stormwater-interactive"},
        repoTuple{"azavea", "pwd-waterworks-revealed"},
        repoTuple{"azavea", "raster-foundry"},
        repoTuple{"azavea", "raster-foundry-tiler"},
        repoTuple{"azavea", "stanford-campus-map"},
        repoTuple{"azavea", "tnc-coastal-resilience-build"},
        repoTuple{"WikiWatershed", "tr-55"},
        repoTuple{"azavea", "oit-ulrs"},
        repoTuple{"azavea", "UrbanInstitute-DCPreservation"},
        repoTuple{"azavea", "usace-embed-viewer"},
        repoTuple{"azavea", "usace-flood-model"},
        repoTuple{"azavea", "usace-flood-geoprocessing"},
        repoTuple{"azavea", "usace-wisdm"},
        repoTuple{"azavea", "usace-wisdm-filter-viewer"},
        repoTuple{"azavea", "usace-wisdm-filter-viewer-data"},
        repoTuple{"azavea", "usace-wisdm-symbolizer"},
        repoTuple{"azavea", "usace-wisdm-symbolizer-data"},
    }

    issues := fetchRepoIssues(client, repos)
    output, err := json.Marshal(issues)
    if err != nil {
        fmt.Println(err)
        return
    }

    fmt.Println(string(output))
}
