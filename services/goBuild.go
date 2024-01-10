package services

import(
	"fmt"
	"os"
	"os/exec"
)

func GoBuild(dirPath string, subject string) {
	cmd := exec.Command("go", "build", "-o", subject)

	cmd.Dir = dirPath

	cmd.Stdout = os.Stdout

	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Build successful")
}