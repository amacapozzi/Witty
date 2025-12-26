package logger

import (
	"fmt"

	"github.com/fatih/color"
)

type LogType int

const (
	Error   LogType = iota
	Warning LogType = iota
	Info    LogType = iota
	Success LogType = iota
)

var getEmoji = map[LogType]string{
	Success: "✅",
	Warning: "⚠️",
	Error:   "❌",
	Info:    "ℹ️",
}

var getColor = map[LogType]color.Attribute{
	Success: color.FgGreen,
	Warning: color.FgYellow,
	Error:   color.FgRed,
	Info:    color.BgBlue,
}

func CreateLog(log LogType, msg string) {
	c := color.New(getColor[log])
	e := getEmoji[log]
	logMessage := fmt.Sprintf("[%s] - %s", e, msg)
	c.Println(logMessage)

}
