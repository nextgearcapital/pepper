package salt

import (
	"os"
	"os/exec"
	"strings"

	"github.com/Sirupsen/logrus"
)

// Provision :
func Provision(profile string, host string) error {
	saltCloud := exec.Command("salt-cloud", "-y", "-p", profile, host)

	logrus.Info("Executing: " + strings.Join(saltCloud.Args, " "))

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

	logrus.Info("Executing: " + strings.Join(saltCloud.Args, " "))

	saltCloud.Stdout = os.Stdout
	saltCloud.Stderr = os.Stderr

	if err := saltCloud.Run(); err != nil {
		return err
	}
	return nil
}
