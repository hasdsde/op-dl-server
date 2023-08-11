package schedule

import (
	"fmt"
	"log"
	"os/exec"
)

func SwagInit() {
	command := exec.Command("swag", "init")

	_, err := command.CombinedOutput()
	if err != nil {
		log.Println("swag init error ", err)
	} else {
		fmt.Println("swag init finish..")
	}
}
