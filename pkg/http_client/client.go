package http_client

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	net_url "net/url"
	"strings"
)

type JSON map[string]interface{}

func BuildQuery(query map[string]string) net_url.Values {
	q, _ := net_url.ParseQuery("")
	for k, v := range query {
		q.Add(k, v)
	}
	return q
}

func BuildURL(url string, query map[string]string) (*net_url.URL, error) {
	u, err := net_url.Parse(url)
	if err != nil {
		return nil, err
	}

	q := BuildQuery(query)
	u.RawQuery = q.Encode()
	return u, nil
}

func Get(url string, query map[string]string) (JSON, error) {
	u, err := net_url.Parse(url)
	if err != nil {
		return nil, err
	}

	q := u.Query()
	for k, v := range query {
		q.Add(k, v)
	}
	u.RawQuery = q.Encode()

	resp, err := http.Get(u.String())
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	res := make(JSON)
	if err := json.Unmarshal(body, &res); err != nil {
		return nil, err
	}
	return res, nil
}

func Post(url string, form map[string]string) (JSON, error) {
	u, err := net_url.Parse(url)
	if err != nil {
		return nil, err
	}

	q, _ := net_url.ParseQuery("")
	for k, v := range form {
		q.Add(k, v)
	}

	resp, err := http.Post(
		u.String(),
		"application/x-www-form-urlencoded",
		strings.NewReader(q.Encode()))
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	res := make(JSON)
	if err := json.Unmarshal(body, &res); err != nil {
		return nil, err
	}
	return res, nil
}
