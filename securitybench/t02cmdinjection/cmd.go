package cmdinjection

import "os/exec"

func PingHost(host string) ([]byte, error) {
	cmd := exec.Command("sh", "-c", "ping -c 1 "+host)
	return cmd.CombinedOutput()
}
