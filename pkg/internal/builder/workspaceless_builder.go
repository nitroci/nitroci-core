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
	pkgCtx "github.com/nitroci/nitroci-core/pkg/core/contexts"
	pkgIntCtx "github.com/nitroci/nitroci-core/pkg/internal/contexts"
	pkgIntComponents "github.com/nitroci/nitroci-core/pkg/internal/components"
)

type workspacelessBuilder struct {
	ctx pkgCtx.CoreContexter
}

func newWorkspacelessBuilder() *workspacelessBuilder {
	return &workspacelessBuilder{}
}

func (b *workspacelessBuilder) createRuntimeContext() {
    ctxInput := pkgIntCtx.ContextInput{} 
    runtimeCtx, _ := pkgIntCtx.CreateContext(ctxInput, false)
	b.ctx = &pkgIntCtx.CoreContext{
		RuntimeCtx: runtimeCtx,
	}
}

func (b *workspacelessBuilder) ensureContextConfiguration() {
	runtimeCtx := b.ctx.GetRuntimeCtx()
    component := pkgIntComponents.CreateChainOfComponents(runtimeCtx)
	component.Execute(runtimeCtx)
}

func (b *workspacelessBuilder) getRuntimeContext() pkgCtx.CoreContexter {
    return b.ctx
}
