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

type TerminalMessageType int64

const (
	Info TerminalMessageType = iota
	Warning
	Error
)

type TerminalItemsOutput struct {
	Messages    []string
	Suggestions []string
	Items       []string
	ItemsType   TerminalMessageType
}

type TerminalOutput struct {
	Messages    []string
	MessageType TerminalMessageType
	ItemsOutput []TerminalItemsOutput
	Output      string
}

type TerminalActionOutputCursorer interface {
	GetStepPrinted() bool
	GetOutputIndex() int
}

type TerminalActionOutput struct {
	StepId  string
	Step    string
	Outputs []string
	Cursor  TerminalActionOutputCursorer
}

type Terminal interface {
	ConvertToCyanColor(text string) string
	ConvertToYellowColor(text string) string
	ConvertToRedColor(text string) string
	Print(output *TerminalOutput)
	PrintActions(output *TerminalActionOutput)
	Println(a ...interface{}) (n int, err error)
	Printf(format string, a ...interface{}) (n int, err error)
}
