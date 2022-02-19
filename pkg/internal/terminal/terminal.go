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
	"fmt"
	"strings"

	pkgConfigs "github.com/nitroci/nitroci-core/pkg/core/configs"
	pkgCTerminal "github.com/nitroci/nitroci-core/pkg/core/terminal"
	"github.com/manifoldco/promptui"
)

const (
	COLOR_RESET  = "\033[0m"
	COLOR_RED    = "\033[31m"
	COLOR_GREEN  = "\033[32m"
	COLOR_YELLOW = "\033[33m"
	COLOR_BLUE   = "\033[34m"
	COLOR_PURPLE = "\033[35m"
	COLOR_CYAN   = "\033[36m"
	COLOR_WHITE  = "\033[37m"
)

func color(terminalMsg pkgCTerminal.TerminalMessageType, isChange bool) string {
	switch terminalMsg {
	case pkgCTerminal.Info:
		if isChange {
			return string(COLOR_GREEN)
		}
		return string(COLOR_RESET)
	case pkgCTerminal.Warning:
		return string(COLOR_YELLOW)
	case pkgCTerminal.Error:
		return string(COLOR_RED)
	}
	return string(COLOR_RESET)
}

func printStrings(color string, strings []string) {
	printStringsWithTabs(color, 0, strings)
}

func printStringsWithTabs(color string, tabs int, strs []string) {
	for _, m := range strs {
		fmt.Print(string(color), strings.Repeat("  ", tabs)+m)
		fmt.Println()
	}
}

type Terminal struct {
	configurer pkgConfigs.Configurer
}

func CreateTerminal(configurer pkgConfigs.Configurer) (pkgCTerminal.Terminal, error) {
	return &Terminal{
		configurer: configurer,
	}, nil
}

func (t *Terminal) ConvertToCyanColor(text string) string {
	return fmt.Sprintf("%v%v%v", string(COLOR_CYAN), text, string(COLOR_RESET))
}

func (t *Terminal) ConvertToYellowColor(text string) string {
	return fmt.Sprintf("%v%v%v", string(COLOR_YELLOW), text, string(COLOR_RESET))
}

func (t *Terminal) ConvertToRedColor(text string) string {
	return fmt.Sprintf("%v%v%v", string(COLOR_RED), text, string(COLOR_RESET))
}

func (t *Terminal) Print(output *pkgCTerminal.TerminalOutput) {
	if output.Messages != nil && len(output.Messages) > 0 {
		printStrings(color(output.MessageType, false), output.Messages)
		fmt.Print(string(COLOR_RESET))
	}
	if output.ItemsOutput != nil && len(output.ItemsOutput) > 0 {
		for _, m := range output.ItemsOutput {
			fmt.Println()
			printStrings(string(COLOR_RESET), m.Messages)
			printStringsWithTabs(string(COLOR_RESET), 1, m.Suggestions)
			printStringsWithTabs(color(m.ItemsType, true), 4, m.Items)
		}
		fmt.Println(string(COLOR_RESET))
	}
	if output != nil && len(output.Output) > 0 {
		fmt.Println(output.Output)
	}
}

type terminalActionOutputCursor struct {
	stepPrinted bool
	outputIndex int
}

func (t terminalActionOutputCursor) GetStepPrinted() bool {
	return t.stepPrinted
}

func (t terminalActionOutputCursor) GetOutputIndex() int{
	return t.outputIndex
}

func (t *Terminal) PrintActions(output *pkgCTerminal.TerminalActionOutput) {
	if output.Cursor == nil {
		output.Cursor = &terminalActionOutputCursor{
			stepPrinted: false,
			outputIndex: -1,
		}
	}
	cursor, _ := output.Cursor.(terminalActionOutputCursor)
	if !cursor.stepPrinted {
		m := output.Step
		if len(output.StepId) > 0 {
			m = fmt.Sprintf("%s %s", output.StepId, m)
		}
		fmt.Print(string(COLOR_GREEN), m)
		fmt.Print(string(COLOR_RESET))
		cursor.stepPrinted = true
		fmt.Println()
	}
	if output.Outputs != nil && len(output.Outputs) > 0 {
		for i, m := range output.Outputs {
			if cursor.outputIndex >= i {
				continue
			}
			fmt.Println(m)
			cursor.outputIndex = i
			fmt.Print(string(COLOR_RESET))
		}
	}
}

func (t *Terminal) Println(a ...interface{}) (n int, err error) {
	return fmt.Println(a...)
}

func (t *Terminal) Printf(format string, a ...interface{}) (n int, err error) {
	return fmt.Printf(format, a...)
}

func (t *Terminal) PromptGlobalConfigKeyAndSave(profile string, label string, secret bool, key string, save bool) (string, error) {
	prompt := promptui.Prompt{
		Label:       label,
		HideEntered: true,
	}
	if secret {
		prompt.Mask = '*'
	}
	value, err := prompt.Run()
	if err != nil {
		t.Printf("Prompt failed %v\n", err)
		return "", err
	}
	if save {
		t.configurer.SetGlobalValue(key, value)
	}
	return value, nil
}

func (t *Terminal) PromptGlobalConfigKey(profile string, label string, secret bool) (string, error) {
	return t.PromptGlobalConfigKeyAndSave(profile, label, secret, "", false)
}
