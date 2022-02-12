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
	pkgWks "github.com/nitroci/nitroci-core/pkg/core/workspaces"
)

type RuntimeContexter interface {
	IsWorkspaceRequired() bool
	GetWorkingDirectory() string
	GetProfile() string
	GetEnvironment() string
	GetSettings(key string) (string, bool)
	GetVerbose() bool
	GetGoos() string
	GetGoarch() string
	HasWorkspaces() bool
	GetWorkspaces() ([]WorkspaceContexter, error)
	GetCurrentWorkspace() (WorkspaceContexter, error)
	GetWorkspace(workspaceDepth int) (WorkspaceContexter, error)
}

type WorkspaceContexter interface {
	GetWorkspacePath() string
	GetWorkspaceHome() string
	GetWorkspaceFileFolder() string
	GetWorkspaceFile() string
	GetVersion() int
	GetId() string
	GetName() string
	CreateWorkspaceInstance() (*pkgWks.WorkspaceModel, error)
}
