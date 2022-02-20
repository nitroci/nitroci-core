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

	pkgCCtx "github.com/nitroci/nitroci-core/pkg/core/contexts"
	pkgWorkspaces "github.com/nitroci/nitroci-core/pkg/core/workspaces"
	pkgFilesearch "github.com/nitroci/nitroci-core/pkg/extensions/filesearch"
	pkgYaml "github.com/nitroci/nitroci-core/pkg/extensions/yaml"
)

type VirtualContext struct {
	Workspaces []*WorkspaceContext
}

// Creational functions

func (c *VirtualContext) load(settings *map[string]string) error {
	wksFolder := (*settings)[CFG_NAME_WKS_FILE_FOLDER]
	wksFileName := (*settings)[CFG_NAME_WKS_FILE_NAME]
	workingDirectory := (*settings)[CFG_NAME_WORKING_DIRECTORY]
	prjFiles := pkgFilesearch.InverseRecursiveFindFiles(workingDirectory, wksFolder, wksFileName)
	prjFilesCount := len(prjFiles)
	c.Workspaces = make([]*WorkspaceContext, prjFilesCount)
	if prjFilesCount == 0 {
		return nil
	}
	for i, prjFile := range prjFiles {
		var wksModel = &pkgWorkspaces.WorkspaceModel{}
		pkgYaml.LoadYamlFile(prjFile, &wksModel)
		var wksContext = WorkspaceContext{}
		wksContext.WorkspacePath = prjFile
		wksContext.WorkspaceHome = filepath.Dir(prjFile)
		wksContext.WorkspaceFileFolder = wksContext.WorkspaceHome
		if strings.HasSuffix(wksContext.WorkspaceHome, (*settings)[CFG_NAME_WKS_FILE_FOLDER]) {
			wksContext.WorkspaceHome = filepath.Dir(wksContext.WorkspaceFileFolder)
		}
		wksContext.WorkspaceFile = filepath.Base(prjFile)
		wksContext.Version = wksModel.Version
		wksContext.Id = wksModel.Workspace.ID
		wksContext.Name = wksModel.Workspace.Name
		c.Workspaces[i] = &wksContext
	}
	return nil
}

func (c *VirtualContext) validate() error {
	for _, wspace := range c.Workspaces {
		err := wspace.validateWorkspaceContext()
		if err != nil {
			return err
		}
	}
	return nil
}

func newVirtualContext(coreContextBuilderInput pkgCCtx.CoreContextBuilderInput) *VirtualContext {
	return &VirtualContext{}
}

// Contexts specific functions

func (v *VirtualContext) hasWorkspaces() bool {
	return v.Workspaces != nil && len(v.Workspaces) > 0
}

func (v *VirtualContext) getWorkspaces() ([]pkgCCtx.WorkspaceContexter, error) {
	if v.Workspaces == nil || len(v.Workspaces) == 0 {
		return nil, errors.New("please initialize a valid workspace")
	}
	var wctx = []pkgCCtx.WorkspaceContexter{}
	for _, element := range v.Workspaces {
		wctx = append(wctx, element)
	}
	return wctx, nil
}

func (v *VirtualContext) getWorkspace(workspaceDepth int) (pkgCCtx.WorkspaceContexter, error) {
	if v.Workspaces == nil || len(v.Workspaces) == 0 {
		return nil, errors.New("please initialize a valid workspace")
	}
	if workspaceDepth < 0 || len(v.Workspaces) <= workspaceDepth {
		return nil, errors.New("invalid workspace depth")
	}
	return v.Workspaces[workspaceDepth], nil
}
