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
package contexts

import(
	"errors"

	pkgYaml "github.com/nitroci/nitroci-core/pkg/core/extensions/yaml"
	pkgWorkspaces "github.com/nitroci/nitroci-core/pkg/core/workspaces"
)

type WorkspaceContext struct {
	WorkspacePath       string
	WorkspaceHome       string
	WorkspaceFileFolder string
	WorkspaceFile       string
	Version             int
	Id                  string
	Name                string
}

// Contexts specific functions

func (v *WorkspaceContext) validateWorkspaceContext() error {
	if len(v.Id) == 0 {
		return errors.New("invalid workspace")
	}
	return nil
}

func (v *WorkspaceContext) GetWorkspacePath() string {
	return v.WorkspacePath
}

func (v *WorkspaceContext) GetWorkspaceHome() string {
	return v.WorkspaceHome
}

func (v *WorkspaceContext) GetWorkspaceFileFolder() string {
	return v.GetWorkspaceFileFolder()
}

func (v *WorkspaceContext) GetWorkspaceFile() string {
	return v.WorkspaceFile
}

func (v *WorkspaceContext) GetVersion() int {
	return v.Version
}

func (v *WorkspaceContext) GetId() string {
	return v.WorkspacePath
}

func (v *WorkspaceContext) GetName() string {
	return v.Name
}

func (v *WorkspaceContext) CreateWorkspaceInstance() (*pkgWorkspaces.WorkspaceModel, error) {
	var wks = &pkgWorkspaces.WorkspaceModel{}
	_, err := pkgYaml.LoadYamlFile(v.WorkspacePath, &wks)
	if err != nil {
		return nil, err
	}
	return wks, nil
}