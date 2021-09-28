package logs

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

// Message extracts the message from the provided log line.
func Message(line string) string {
	_, message := getMessageAndLevel(line)
	return message
}

// MessageLen counts the amount of characters (runes) in the message of the log line.
func MessageLen(line string) int {
	message := Message(line)
	return utf8.RuneCountInString(message)
}

// LogLevel extracts the log level string from the provided log line.
func LogLevel(line string) string {
	level, _ := getMessageAndLevel(line)
	return strings.ToLower(strings.Trim(level, "[]"))
}

// Reformat reformats the log line in the format `message (logLevel)`.
func Reformat(line string) string {
	return fmt.Sprintf("%s (%s)", Message(line), LogLevel(line))
}

func getMessageAndLevel(line string) (string, string) {
	v := strings.Split(line, ": ")
	return strings.TrimSpace(v[0]), strings.TrimSpace(v[1])
}
