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

type workspacelessBuilder struct {
	ctx pkgCCtx.CoreContexter
}

// Creational functions

func newWorkspacelessBuilder() *workspacelessBuilder {
	return &workspacelessBuilder{}
}

// Builder specific functions

func (b *workspacelessBuilder) createCoreContext(ctxInput pkgCCtx.CoreContextBuilderInput) {
	runtimeCtx, _ := pkgIntCtx.CreateContext(ctxInput, false)
	configuration, _ := pkgIntConfigs.CreateConfiguration(ctxInput)
	terminal, _ := pkgIntTerminal.CreateTerminal(configuration)
	b.ctx = &pkgIntCtx.CoreContext{
		RuntimeCtx: runtimeCtx,
		Configuration: configuration,
		Terminal: terminal,
	}
}

func (b *workspacelessBuilder) initializeCoreContext(ctxInput pkgCCtx.CoreContextBuilderInput) {
	var next pkgIntComp.Component
	// Initialize folders
	first := &pkgIntComp.GlobalFoldersComponent{}
	next = first
	// Initialize cache
	next = next.SetNext(&pkgIntComp.GlobalCacheComponent{})
	// Initialize plugins
	next = next.SetNext(&pkgIntComp.GlobalPluginsComponent{})
	// Initialize bits
	next.SetNext(&pkgIntComp.GlobalBitsComponent{})
	runtimeCtx := b.ctx.GetRuntimeCtx()
	first.Execute(runtimeCtx)
}

func (b *workspacelessBuilder) getCoreContext(ctxInput pkgCCtx.CoreContextBuilderInput) pkgCCtx.CoreContexter {
	return b.ctx
}
