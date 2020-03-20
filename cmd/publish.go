package cmd

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

type PublishConfig struct {
	Host string
	Project int64
	Tag    string
	Link   string
	Note    string
	Token   string
}

func Run(config PublishConfig) bool {
	if config.Host == "" {
		fmt.Println("error host empty ")
		os.Exit(1)
	}
	url := fmt.Sprintf("%s/api/v4/projects/%d/releases", config.Host, config.Project)

	//url := "https://hi.hxxbaby.com/api/v4/projects/301/releases"
	method := "POST"
	client := http.Client{Timeout:time.Duration(5)*time.Second}
	payload := strings.NewReader(fmt.Sprintf(`{
	"name": "%s",
		"tag_name": "%s",
		"description": "%s",
		"assets": {
		"links": [
		{
			"name": "%s",
			"url": "%s"
		}]}
	}`, config.Tag, config.Tag, config.Note, config.Note, config.Link))
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
	}
	req.Header.Add("PRIVATE-TOKEN", config.Token)
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)

	fmt.Println(string(body))
	return false
}