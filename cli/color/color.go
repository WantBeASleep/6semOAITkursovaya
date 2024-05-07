package color

import (
	color "github.com/fatih/color"
)

type defaultColorSprintFuncs struct {
	Red func(a ...interface{}) string
	Green func(a ...interface{}) string
	Yellow func(a ...interface{}) string
	Cyan func(a ...interface{}) string
	Blue func(a ...interface{}) string
}

func GetColorSprintFuncs() defaultColorSprintFuncs {
	return defaultColorSprintFuncs{
		Red: color.New(color.FgRed).SprintFunc(),
		Green: color.New(color.FgGreen).SprintFunc(),
		Yellow: color.New(color.FgYellow).SprintFunc(),
		Cyan: color.New(color.FgCyan).SprintFunc(),
		Blue: color.New(color.FgBlue).SprintFunc(),
	}
}