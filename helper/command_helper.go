package helper

import (
	"github.com/Agilistikmal/venti/config"
	"strings"
)

func GetCommand(message string) (string, []string) {
	if strings.HasPrefix(message, config.Prefix) == true {
		fullCommand := strings.Split(strings.ToLower(message[1:]), " ")
		command := fullCommand[0]
		args := fullCommand[1:]
		return command, args
	}
	return "", nil
}
