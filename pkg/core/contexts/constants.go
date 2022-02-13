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

// Keys for calculated configurations

const (
	CFG_NAME_WORKING_DIRECTORY = "WORKING_DIRECTORY"
	CFG_NAME_CONFIG_PATH = "NITROCI_CONFIG"
	CFG_NAME_CONFIG_HOME = "NITROCI_CONFIG_HOME"
	CFG_NAME_CONFIG_FILE = "NITROCI_CONFIG_FILE"
	CFG_NAME_CONFIG_TYPE = "NITROCI_CONFIG_TYPE"

	CFG_NAME_CACHE_PATH         = "NITROCI_CACHE"
	CFG_NAME_CACHE_PLUGINS_PATH = "NITROCI_CACHE_PLUGINS"
	CFG_NAME_CACHE_BITS_PATH    = "NITROCI_CACHE_BITS"

	CFG_NAME_PLUGINS_REGISTRY = "NITROCI_PLUGINS_REGISTRY_URI"

	CFG_NAME_WKS_FILE_FOLDER = "NITROCI_WKS_FILE_FOLDER"
	CFG_NAME_WKS_FILE_NAME   = "NITROCI_WKS_FILE_NAME"

	CFG_NAME_BITS_REGISTRY = "NITROCI_BITS_REGISTRY_URI"
)

const (
	CORE_BUILDER_WORKSPACE_TYPE     = "workspace"
	CORE_BUILDER_WORKSPACELESS_TYPE = "workspaceless"
)
