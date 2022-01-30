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
	Cli     *CliContext
	Virtual *VirtualContext
}

func LoadRuntimeContext(profile string, environment string, workspaceDepth int, verbose bool) (*RuntimeContext, error) {
	var ctx = &RuntimeContext{
		Cli:     &CliContext{},
		Virtual: &VirtualContext{},
	}
	ctx.Cli.loadCliContext(profile, environment, workspaceDepth, verbose)
	_, err := ctx.Cli.loadCliContextSettings()
	if err != nil {
		return nil, err
	}
	_, err = ctx.Virtual.loadVirtualContext(ctx, workspaceDepth)
	if err != nil {
		return nil, err
	}
	err = ctx.Cli.validateCliContext()
	if err != nil {
		return nil, err
	}
	err = ctx.Virtual.validateVirtualContext()
	if err != nil {
		return nil, err
	}
	return ctx, nil
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
