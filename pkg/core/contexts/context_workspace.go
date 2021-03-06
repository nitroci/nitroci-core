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

import (
	"errors"
	"path/filepath"
	"strings"

	pkgFilesearch "github.com/nitroci/nitroci-core/pkg/core/extensions/filesearch"
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

type VirtualContext struct {
	Workspaces []*WorkspaceContext
}

func findWorkspaceFiles(runtimeContext *RuntimeContext) (workspaceFiles []string) {
	wksFolder := runtimeContext.Cli.Settings[CFG_NAME_WKS_FILE_FOLDER]
	wksFileName := runtimeContext.Cli.Settings[CFG_NAME_WKS_FILE_NAME]
	return pkgFilesearch.InverseRecursiveFindFiles(runtimeContext.Cli.WorkingDirectory, wksFolder, wksFileName)
}

func (v *WorkspaceContext) CreateWorkspaceInstance() (*pkgWorkspaces.WorkspaceModel, error) {
	var wks = &pkgWorkspaces.WorkspaceModel{}
	_, err := pkgYaml.LoadYamlFile(v.WorkspacePath, &wks)
	if err != nil {
		return nil, err
	}
	return wks, nil
}

func (v *VirtualContext) loadVirtualContext(runtimeContext *RuntimeContext, workspaceDepth int) (*VirtualContext, error) {
	prjFiles := findWorkspaceFiles(runtimeContext)
	prjFilesCount := len(prjFiles)
	v.Workspaces = make([]*WorkspaceContext, prjFilesCount)
	if prjFilesCount == 0 {
		return v, nil
	}
	for i, prjFile := range prjFiles {
		var wksModel = &pkgWorkspaces.WorkspaceModel{}
		pkgYaml.LoadYamlFile(prjFile, &wksModel)
		var wksContext = WorkspaceContext{}
		wksContext.WorkspacePath = prjFile
		wksContext.WorkspaceHome = filepath.Dir(prjFile)
		wksContext.WorkspaceFileFolder = wksContext.WorkspaceHome
		if strings.HasSuffix(wksContext.WorkspaceHome, runtimeContext.Cli.Settings[CFG_NAME_WKS_FILE_FOLDER]) {
			wksContext.WorkspaceHome = filepath.Dir(wksContext.WorkspaceFileFolder)
		}
		wksContext.WorkspaceFile = filepath.Base(prjFile)
		wksContext.Version = wksModel.Version
		wksContext.Id = wksModel.Workspace.ID
		wksContext.Name = wksModel.Workspace.Name
		v.Workspaces[i] = &wksContext
	}
	return v, nil
}

func (v *VirtualContext) hasWorkspaces() bool {
	return v.Workspaces != nil && len(v.Workspaces) > 0
}

func (v *VirtualContext) getWorkspaces() ([]*WorkspaceContext, error) {
	if v.Workspaces == nil || len(v.Workspaces) == 0 {
		return nil, errors.New("please initialize a valid workspace")
	}
	return v.Workspaces, nil
}

func (v *VirtualContext) getWorkspace(workspaceDepth int) (*WorkspaceContext, error) {
	if v.Workspaces == nil || len(v.Workspaces) == 0 {
		return nil, errors.New("please initialize a valid workspace")
	}
	if workspaceDepth < 0 || len(v.Workspaces) <= workspaceDepth {
		return nil, errors.New("invalid workspace depth")
	}
	return v.Workspaces[workspaceDepth], nil
}

func (v *WorkspaceContext) validateWorkspaceContext() error {
	if len(v.Id) == 0 {
		return errors.New("invalid workspace")
	}
	return nil
}

func (v *VirtualContext) validateVirtualContext() error {
	for _, wspace := range v.Workspaces {
		err := wspace.validateWorkspaceContext()
		if err != nil {
			return err
		}
	}
	return nil
}
