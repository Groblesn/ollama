package taintflow

import "os/exec"

func archive(name string) ([]byte, error) {
	return exec.Command("sh", "-c", "tar czf /var/backups/snapshot.tgz "+name).CombinedOutput()
}
