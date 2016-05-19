package salt

import (
	"os"
	"os/exec"
	"strings"

	"github.com/nextgearcapital/pepper/pkg/log"
)

// Provision :
func Provision(profile string, host string) error {
	saltCloud := exec.Command("salt-cloud", "-y", "-p", profile, host)

	log.Info("Executing: " + strings.Join(saltCloud.Args, " "))

	saltCloud.Stdout = os.Stdout
	saltCloud.Stderr = os.Stderr

	if err := saltCloud.Run(); err != nil {
		return err
	}
	return nil
}

// Destroy :
func Destroy(host string) error {
	saltCloud := exec.Command("salt-cloud", "-y", "-d", host)

	log.Info("Executing: " + strings.Join(saltCloud.Args, " "))

	saltCloud.Stdout = os.Stdout
	saltCloud.Stderr = os.Stderr

	if err := saltCloud.Run(); err != nil {
		return err
	}
	return nil
}
