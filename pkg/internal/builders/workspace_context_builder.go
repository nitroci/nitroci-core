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
package builders

import (
	pkgCCtx "github.com/nitroci/nitroci-core/pkg/core/contexts"
	pkgIntComp "github.com/nitroci/nitroci-core/pkg/internal/components"
	pkgIntConfigs "github.com/nitroci/nitroci-core/pkg/internal/configs"
	pkgIntCtx "github.com/nitroci/nitroci-core/pkg/internal/contexts"
	pkgIntTerminal "github.com/nitroci/nitroci-core/pkg/internal/terminal"
)

type workspaceBuilder struct {
	ctx pkgCCtx.CoreContexter
}

// Creational functions

func newWorkspaceBuilder() *workspaceBuilder {
	return &workspaceBuilder{}
}

// Builder specific functions

func (b *workspaceBuilder) createCoreContext(ctxInput pkgCCtx.CoreContextBuilderInput) {
	runtimeCtx, _ := pkgIntCtx.CreateContext(ctxInput, false)
	configuration, _ := pkgIntConfigs.CreateConfiguration(ctxInput, false)
	terminal, _ := pkgIntTerminal.CreateTerminal(configuration)
	b.ctx = &pkgIntCtx.CoreContext{
		RuntimeCtx: runtimeCtx,
		Configuration: configuration,
		Terminal: terminal,
	}
}

func (b *workspaceBuilder) initializeCoreContext(ctxInput pkgCCtx.CoreContextBuilderInput) {
	var next pkgIntComp.Component
	// Initialize folders
	first := &pkgIntComp.GlobalFoldersComponent{}
	next = first
	next = next.SetNext(&pkgIntComp.LocalFoldersComponent{})
	// Initialize cache
	next = next.SetNext(&pkgIntComp.GlobalCacheComponent{})
	next = next.SetNext(&pkgIntComp.LocalCacheComponent{})
	// Initialize plugins
	next = next.SetNext(&pkgIntComp.GlobalPluginsComponent{})
	next = next.SetNext(&pkgIntComp.LocalPluginsComponent{})
	// Initialize bits
	next = next.SetNext(&pkgIntComp.GlobalBitsComponent{})
	next.SetNext(&pkgIntComp.LocalBitsComponent{})
	runtimeCtx := b.ctx.GetRuntimeCtx()
	first.Execute(runtimeCtx)
}

func (b *workspaceBuilder) getCoreContext(ctxInput pkgCCtx.CoreContextBuilderInput) pkgCCtx.CoreContexter {
	return b.ctx
}
