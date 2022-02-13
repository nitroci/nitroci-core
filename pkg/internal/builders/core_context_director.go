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
)

type Director struct {
	CoreContextBuilder
}

func NewDirector(b CoreContextBuilder) *Director {
	return &Director{
		CoreContextBuilder: b,
	}
}

func (d *Director) SetBuilder(b CoreContextBuilder) {
	d.CoreContextBuilder = b
}

func (d *Director) BuildContext(ctxInput pkgCCtx.CoreContextBuilderInput) pkgCCtx.CoreContexter {
	d.createCoreContext(ctxInput)
	d.initializeCoreContext(ctxInput)
	return d.getCoreContext(ctxInput)
}
