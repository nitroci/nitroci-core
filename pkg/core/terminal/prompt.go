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
package terminal

import (
	"github.com/manifoldco/promptui"
	"github.com/nitroci/nitroci-core/pkg/core/configs"
)

func PromptGlobalConfigKeyAndSave(profile string, label string, secret bool, key string, save bool) (string, error) {
	prompt := promptui.Prompt{
		Label:       label,
		HideEntered: true,
	}
	if secret {
		prompt.Mask = '*'
	}
	value, err := prompt.Run()
	if err != nil {
		Printf("Prompt failed %v\n", err)
		return "", err
	}
	if save {
		configs.SetGlobalConfigString(profile, key, value)
	}
	return value, nil
}

func PromptGlobalConfigKey(profile string, label string, secret bool) (string, error) {
	return PromptGlobalConfigKeyAndSave(profile, label, secret, "", false)
}
