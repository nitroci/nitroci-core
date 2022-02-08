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

func newRuntimeContext(contextInput ContextInput) *RuntimeContext {
	runtimeCtx := &RuntimeContext{
		workspaceLess: true,
		Cli: newCliContext(contextInput),
	}
	return runtimeCtx
}

func newRuntimeWorkspaceContext(contextInput ContextInput) *RuntimeContext {
	runtimeCtx := newRuntimeContext(contextInput)
	runtimeCtx.workspaceLess = false
	runtimeCtx.Virtual = newVirtualContext(contextInput)
	return runtimeCtx
}

// Contexter specific functions

type RuntimeContexter interface {
	HasWorkspaces() bool
	GetWorkspaces() ([]*WorkspaceContext, error)
	GetCurrentWorkspace() (*WorkspaceContext, error)
	GetWorkspace(workspaceDepth int) (*WorkspaceContext, error)
}

func (r *RuntimeContext) HasWorkspaces() bool {
	return r.Virtual.hasWorkspaces()
}

func (r *RuntimeContext) GetWorkspaces() ([]*WorkspaceContext, error) {
	return r.Virtual.getWorkspaces()
}

func (r *RuntimeContext) GetCurrentWorkspace() (*WorkspaceContext, error) {
	return r.Virtual.getWorkspace(r.Cli.WorkspaceDepth)
}

func (r *RuntimeContext) GetWorkspace(workspaceDepth int) (*WorkspaceContext, error) {
	return r.Virtual.getWorkspace(workspaceDepth)
}
