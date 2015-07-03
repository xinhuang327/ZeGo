package zego

import (
	"fmt"
	"net/url"
)

type Search_Results struct {
	Count    int       `json:"count"`
	NextPage string    `json:"next_page"`
	PrevPage string    `json:"prev_page"`
	Results  []*Result `json:"results"`
}

type Result struct {
	Name       string `json:"name"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
	Id         int    `json:"id"`
	ResultType string `json:"result_type"`
	Url        string `json:"url"`
}

func (a Auth) Search(queryAndParams ...string) (*Resource, error) {

	path := "/search.json?query=" + url.QueryEscape(queryAndParams[0])
	for i := 1; i < len(queryAndParams); i++ {
		path += "&" + url.QueryEscape(queryAndParams[i])
	}
	fmt.Println(path)
	resource, err := api(a, "GET", path, "")
	if err != nil {
		return nil, err
	}

	return resource, nil

}
