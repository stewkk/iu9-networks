package terminals

import (
	"os/exec"

	"github.com/google/uuid"
)

var (
	sessions map[string]session
)

type session struct {

}

func New() (string, error) {
	return uuid.New().String(), nil
}

func GetLog(id string) (string, error) {
	return "TODO!", nil
}

func Cmd(cmd string) (output string) {
	log, err := exec.Command("sh", "-c", cmd).CombinedOutput()
	output = string(log)
	if err != nil {
		output = output+err.Error()
	}
	return output
}
