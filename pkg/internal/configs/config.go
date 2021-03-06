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
package configs

import (
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

func EnsureConfiguration(configFile string) (string, error) {
	configHome := filepath.Dir(configFile)
	configName := filepath.Base(configFile)
	configType := filepath.Ext(configFile)
	configPath := filepath.Join(configHome, configName)
	viper.AddConfigPath(configHome)
	viper.SetConfigName(configName)
	viper.SetConfigType(configType)
	abs, _ := filepath.Abs(configHome)
	if _, err := os.Stat(abs); os.IsNotExist(err) {
		os.MkdirAll(configHome, 0700)
	}
	_, err := os.Stat(configPath)
	if err!= nil && !os.IsExist(err) { 
		if err := viper.SafeWriteConfig(); err != nil {
			return "", err
		}
	}
	viper.AutomaticEnv()
	err = viper.ReadInConfig()
	if err != nil {
		return "", err
	}
	return configFile, nil
}
