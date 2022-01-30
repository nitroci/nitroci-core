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
package filepath

import (
	"errors"
	"os"
	"path/filepath"
	"strings"

	pkgStrings "github.com/nitroci/nitroci-core/pkg/core/extensions/strings"
)

type PathDescription struct {
	Path          string
	Home          string
	FileName      string
	FileExtension string
}

func GetDirPathDescription(filePath string, checkExist bool) (*PathDescription, error) {
	if len(filePath) == 0 {
		return nil, errors.New("invalid file path")
	}
	if checkExist {
		if _, err := os.Stat(filePath); os.IsNotExist(err) {
			return nil, err
		}
	}
	if len(filepath.Ext(filePath)) != 0 {
		return nil, errors.New("invalid directory")
	}
	pathDesc := &PathDescription{}
	pathDesc.Path = filePath
	pathDesc.Home = filePath
	pathDesc.FileName = ""
	pathDesc.FileExtension = ""
	return pathDesc, nil
}

func GetFilePathDescription(filePath string, extensions []string, checkExist bool) (*PathDescription, error) {
	if len(filePath) == 0 {
		return nil, errors.New("invalid file path")
	}
	if checkExist {
		if _, err := os.Stat(filePath); os.IsNotExist(err) {
			return nil, err
		}
	}
	pathDesc := &PathDescription{}
	ext := strings.Replace(strings.ToUpper(filepath.Ext(filePath)), ".", "", 1)
	if !pkgStrings.StringInSlice(ext, extensions) {
		return nil, errors.New("extension is not valid")
	}
	if len(ext) == 0 {
		pathDesc.Path = filePath
		pathDesc.Home = filePath
		pathDesc.FileName = filePath
		pathDesc.FileExtension = filePath
	} else {
		pathDesc.Path = filePath
		pathDesc.Home = filepath.Dir(filePath)
		pathDesc.FileName = filepath.Base(filePath)
		pathDesc.FileExtension = ext
	}
	return pathDesc, nil
}
