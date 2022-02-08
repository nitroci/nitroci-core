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

// Configuration keys

const (
	ENV_NAME_CONFIG           = "NITROCI_CONFIG"
	ENV_NAME_CACHE_HOME       = "NITROCI_CACHE"
	ENV_NAME_PLUGINS_REGISTRY = "NITROCI_PLUGINS_REGISTRY" // GITHUB+https://raw.githubusercontent.com/nitroci/nitroci-plugins/main
	ENV_NAME_WKS_FILE_FOLDER  = "NITROCI_WKS_FILE_FOLDER"  // GITHUB+https://raw.githubusercontent.com/nitroci/nitroci-bits/main
	ENV_NAME_WKS_FILE_NAME    = "NITROCI_WKS_FILE_NAME"
	ENV_NAME_BITS_REGISTRY    = "NITROCI_BITS_REGISTRY"
)

// Default configuration values

const (
	CFG_DEFVAL_PLUGINS_REGISTRY_GITHUB_URL = "GITHUB+https://raw.githubusercontent.com/nitroci/nitroci-plugins/main"
	CFG_DEFVAL_WKS_FILE_FOLDER             = ".nitroci"
	CFG_DEFVAL_WKS_FILE_NAME               = "workspace.yml"
	CFG_DEFVAL_BITS_REGISTRY_GITHUB_URL    = "GITHUB+https://raw.githubusercontent.com/nitroci/nitroci-bits/main"
)
