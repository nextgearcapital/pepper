package salt

import (
	"fmt"
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

	err := saltCloud.Run()
	if err != nil {
		log.Die("%s", err)
	}
	return nil
}

// Destroy :
func Destroy(host string) error {
	saltCloud := exec.Command("salt-cloud", "-y", "-d", host)

	log.Info("Executing: " + strings.Join(saltCloud.Args, " "))

	saltCloud.Stdout = os.Stdout
	saltCloud.Stderr = os.Stderr

	var response int

	fmt.Printf("Are you sure you want to destroy %s", host)

	fmt.Scanf("%c", &response)
	switch response {
	default:
		fmt.Println("Aborted!")
	case 'y':
		fmt.Println("Let's get this party started!")
	case 'Y':
		fmt.Println("Let's get this party started!")
	}

	err := saltCloud.Run()
	if err != nil {
		log.Die("%s", err)
	}
	return nil
}
