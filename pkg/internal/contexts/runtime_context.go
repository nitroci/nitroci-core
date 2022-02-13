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
	pkgCCtx "github.com/nitroci/nitroci-core/pkg/core/contexts"
)

type RuntimeContext struct {
	workspaceLess bool
	Cli           *CliContext
	Virtual       *VirtualContext
}

// Creational functions

func (c *RuntimeContext) load() error {
	if err := c.Cli.load(); err != nil {
		return err
	}
	if !c.workspaceLess {
		if err := c.Virtual.load(); err != nil {
			return err
		}
	}
	return nil
}

func (c *RuntimeContext) validate() error {
	if err := c.Cli.validate(); err != nil {
		return err
	}
	if !c.workspaceLess {
		if err := c.Virtual.validate(); err != nil {
			return err
		}
	}
	return nil
}

func newRuntimeContext(coreContextBuilderInput pkgCCtx.CoreContextBuilderInput) *RuntimeContext {
	runtimeCtx := RuntimeContext{
		workspaceLess: true,
		Cli:           newCliContext(coreContextBuilderInput),
		Virtual:       newVirtualContext(coreContextBuilderInput),
	}
	return &runtimeCtx
}

func newRuntimeWorkspaceContext(coreContextBuilderInput pkgCCtx.CoreContextBuilderInput) *RuntimeContext {
	runtimeCtx := RuntimeContext{
		workspaceLess: false,
		Cli:           newCliContext(coreContextBuilderInput),
		Virtual:       newVirtualContext(coreContextBuilderInput),
	}
	return &runtimeCtx
}

// Contexter specific functions

func (r *RuntimeContext) IsWorkspaceRequired() bool {
	return !r.workspaceLess
}

func (r *RuntimeContext) GetWorkingDirectory() string {
	return r.Cli.WorkingDirectory
}

func (r *RuntimeContext) GetProfile() string {
	return r.Cli.Profile
}

func (r *RuntimeContext) GetEnvironment() string {
	return r.Cli.Environment
}

func (r *RuntimeContext) GetSettings(key string) (string, bool) {
	if val, ok := r.Cli.Settings[key]; ok {
		return val, true
	}
	return "", false
}

func (r *RuntimeContext) GetVerbose() bool {
	return r.Cli.Verbose
}

func (r *RuntimeContext) GetGoos() string {
	return r.Cli.Goos
}

func (r *RuntimeContext) GetGoarch() string {
	return r.Cli.Goarch
}

func (r *RuntimeContext) HasWorkspaces() bool {
	return r.Virtual.hasWorkspaces()
}

func (r *RuntimeContext) GetWorkspaces() ([]pkgCCtx.WorkspaceContexter, error) {
	return r.Virtual.getWorkspaces()
}

func (r *RuntimeContext) GetCurrentWorkspace() (pkgCCtx.WorkspaceContexter, error) {
	return r.Virtual.getWorkspace(r.Cli.WorkspaceDepth)
}

func (r *RuntimeContext) GetWorkspace(workspaceDepth int) (pkgCCtx.WorkspaceContexter, error) {
	return r.Virtual.getWorkspace(workspaceDepth)
}
