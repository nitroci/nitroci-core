/*
Copyright 2021 The NitroCI Authors.

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
package http

import (
	"encoding/base64"
	"encoding/json"
	"io"
	"net/http"
	"path"
	"strings"

	extJson "github.com/nitroci/nitroci-core/pkg/core/extensions/json"
)

func basicAuth(username, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}

func redirectPolicyFunc(req *http.Request, via []*http.Request) error {
	req.Header.Add("Authorization", "Basic "+basicAuth("username1", "password123"))
	return nil
}

type HttpResult struct {
	FileName      string
	FileExtension string
	StatusCode    int
	Body          []byte
}

func (httpResult *HttpResult) ToString() string {
	return string(httpResult.Body)
}

func (httpResult HttpResult) ToJson(target *interface{}) error {
	bodyStr := httpResult.ToString()
	_, err := extJson.IsJSON(bodyStr)
	if err != nil {
		return err
	}
	return json.NewDecoder(strings.NewReader(bodyStr)).Decode(*target)
}

func HttpGet(url string) (httpResult *HttpResult, err error) {
	return HttpGetWithAuth(url, "", "")
}

func HttpGetWithAuth(url string, username string, password string) (httpResult *HttpResult, err error) {
	client := &http.Client{
		CheckRedirect: redirectPolicyFunc,
	}
	req, _ := http.NewRequest("GET", url, nil)
	if len(username) > 0 && len(password) > 0 {
		req.Header.Add("Authorization", "Basic "+basicAuth(username, password))
	}
	req.Header.Add("Cache-Control", "private, no-store, max-age=0")
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	fileName := path.Base(req.URL.Path)
	fileExtensions := path.Ext(req.URL.Path)
	if strings.HasSuffix(fileName, "tar.gz") {
		fileExtensions = "tar.gz"
	}
	httpResult = &HttpResult{
		FileName:      fileName,
		FileExtension: fileExtensions,
		StatusCode:    resp.StatusCode,
		Body:          bodyBytes,
	}
	return httpResult, nil
}
