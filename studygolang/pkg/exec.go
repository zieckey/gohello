package pkg

import (
	"fmt"
	"os/exec"
)

func TestExec_date() {
	fmt.Printf("\n\n")
	out, err := exec.Command("date").Output()
	if err != nil {
		fmt.Printf("%v\n", err.Error())
		return
	}
	fmt.Printf("The date is %s\n", out)
}
