package services

import(
	"fmt"
	"os/exec"
)

func Service(serviceName string) {
	cmd := exec.Command("sudo", "systemctl", "restart", serviceName)

	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error restarting service:", err)
		return
	}

	fmt.Println("Service restarted successfully:", string(output))
}