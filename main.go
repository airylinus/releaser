/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package main //import github.com/airylinus/releaser

import (
	"flag"

	"github.com/airylinus/releaser/cmd"
)

var host *string = flag.String("h", "host", "gitlab service Domain: etc https://gitlab.example.com")
var project *int64 = flag.Int64("p", 0, "File to receive sorted values: etc 123")
var tag *string = flag.String("t", "tag", "gitlab tag asset name etc : v1.2.34")
var note *string = flag.String("n", "note", "note , etc : this publish version is doing bla bla")
var link *string = flag.String("l", "https://github.com/airylinus/releaser", "link etc: https://google.com")
var token *string = flag.String("o", "", "token etc: xxx123ccc private token from gitlab")

func main() {
	flag.Parse()
	cmd.Run(cmd.PublishConfig{Host: *host, Project: *project, Tag: *tag, Note: *note, Link: *link, Token: *token})
}
