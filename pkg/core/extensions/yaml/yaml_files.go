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
package yaml

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

func LoadYamlFile(fileName string, data interface{}) (*interface{}, error) {
	yamlFile, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(yamlFile, data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func LoadYamlBytes(yamlBytes []byte, data interface{}) (*interface{}, error) {
	err := yaml.Unmarshal(yamlBytes, data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func SaveYamlFile(fileName string, data interface{}) (string, error){
	yamlData, _ := yaml.Marshal(data)
	fmt.Println(string(yamlData))
	return fileName, nil
}
