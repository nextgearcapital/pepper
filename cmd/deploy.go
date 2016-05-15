package cmd

import (
	"strings"

	"github.com/nextgearcapital/pepper/pkg/device42"
	"github.com/nextgearcapital/pepper/pkg/log"
	"github.com/nextgearcapital/pepper/pkg/salt"
	"github.com/nextgearcapital/pepper/template/vsphere"

	"github.com/spf13/cobra"
)

var (
	profile    string
	roles      string
	osTemplate string
	ipam       bool
)

func init() {
	RootCmd.AddCommand(deployCmd)

	deployCmd.Flags().StringVarP(&profile, "profile", "p", "", "Profile to generate and output to /etc/salt/cloud.profiles.d for salt-cloud to use")
	deployCmd.Flags().StringVarP(&roles, "roles", "r", "", "List of roles to assign to the host in D42 [eg: dcos,dcos-master]")
	deployCmd.Flags().StringVarP(&osTemplate, "template", "t", "", "Which OS template you want to use [eg: Ubuntu, CentOS, someothertemplatename]")
	deployCmd.Flags().BoolVarP(&ipam, "no-ipam", "", false, "Whether or not to use Device42 IPAM [This is only used internally]")
	deployCmd.Flags().BoolVarP(&log.IsDebugging, "debug", "d", false, "Turn debugging on")
}

var deployCmd = &cobra.Command{
	Use:   "deploy",
	Short: "Deploy VM's via salt-cloud",
	Long: `pepper is a wrapper around salt-cloud that will generate salt-cloud profiles based on information you provide in profile configs.
Profile configs live in "/etc/pepper/config.d/{platform}/{environment}. Pepper is opinionated and looks at the profile you pass in as it's source
of truth. For example: If you pass in "vmware-dev-large" as the profile, it will look for your profile config in "/etc/pepper/config.d/vmware/large.yaml".
This allows for maximum flexibility due to the fact that everyone has different environments and may have some sort of naming scheme associated with them
so Pepper makes no assumptions on that. Pepper does however make assumptions on your instance type. [eg: nano, micro, small, medium, etc] Although these
options are available to you, you are free to override them as you see fit.
For example:

Provision new host web01 (Ubuntu) in the dev environment from the nano profile using vmware as a provider:

$ pepper deploy -p vmware-dev-nano -t Ubuntu web01

Or alternatively:

$ pepper deploy --profile vmware-dev-nano --template Ubuntu web01

Provision new host web02 (CentOS) in the prd environment from the large profile using vmware as a provider:

$ pepper deploy -p vmware-prd-large -t CentOS web02

Provision new host web03 (Ubuntu) in the uat environment from the hyper profile using vmware as a provider:

$ pepper deploy -p vmware-uat-hyper -t Ubuntu web03

Are you getting this yet?

$ pepper deploy -p vmware-prd-mid -t Ubuntu -r dcos,dcos-master dcos01 dcos02 dcos03`,
	Run: func(cmd *cobra.Command, args []string) {
		if profile == "" {
			log.Die("You didn't specify a profile.")
		} else if osTemplate == "" {
			log.Die("You didn't specify an OS template.")
		} else if len(args) == 0 {
			log.Die("You didn't specify any hosts.")
		}

		splitProfile := strings.Split(profile, "-")

		// These will be the basis for how the profile gets generated.
		platform := splitProfile[0]
		environment := splitProfile[1]
		instancetype := splitProfile[2]

		// Nothing really gained here it just makes the code more readable.
		hosts := args

		var ipAddress string

		for _, host := range hosts {
			if ipam != true {
				if err := device42.ReadConfig(); err != nil {
					log.Die("%s", err)
				}
				if environment == "prd" {
					// Get a new IP
					newIP, err := device42.GetNextIP(device42.PrdRange)
					if err != nil {
						log.Die("%s", err)
					}
					ipAddress = newIP
					// Create the Device
					if err := device42.CreateDevice(host, "Production"); err != nil {
						if err = device42.DeleteDevice(host); err != nil {
							log.Die("%s", err)
						}
						log.Die("%s", err)
					}
					// Reserve IP
					if err := device42.ReserveIP(newIP, host); err != nil {
						if err = device42.MakeIPAvailable(newIP); err != nil {
							log.Die("%s", err)
						}
						if err = device42.DeleteDevice(host); err != nil {
							log.Die("%s", err)
						}
						log.Die("%s", err)
					}
					// Update custom fields
					if err := device42.UpdateCustomFields(host, "roles", roles); err != nil {
						if err = device42.MakeIPAvailable(newIP); err != nil {
							log.Die("%s", err)
						}
						if err = device42.DeleteDevice(host); err != nil {
							log.Die("%s", err)
						}
						log.Die("%s", err)
					}
				} else if environment == "dev" {
					// Get a new IP
					newIP, err := device42.GetNextIP(device42.DevRange)
					if err != nil {
						log.Die("%s", err)
					}
					ipAddress = newIP
					// Create the Device
					if err := device42.CreateDevice(host, "Development"); err != nil {
						if err = device42.DeleteDevice(host); err != nil {
							log.Die("%s", err)
						}
						log.Die("%s", err)
					}
					// Reserve IP
					if err := device42.ReserveIP(newIP, host); err != nil {
						if err = device42.MakeIPAvailable(newIP); err != nil {
							log.Die("%s", err)
						}
						if err = device42.DeleteDevice(host); err != nil {
							log.Die("%s", err)
						}
						log.Die("%s", err)
					}
					// Update custom fields
					if err := device42.UpdateCustomFields(host, "roles", roles); err != nil {
						if err = device42.MakeIPAvailable(newIP); err != nil {
							log.Die("%s", err)
						}
						if err = device42.DeleteDevice(host); err != nil {
							log.Die("%s", err)
						}
						log.Die("%s", err)
					}
				}
			}
			switch platform {
			case "vmware":
				var vsphere vsphere.ProfileConfig
				if err := vsphere.Prepare(platform, environment, instancetype, osTemplate, ipAddress); err != nil {
					log.Die("%s", err)
				}
				if err := vsphere.Generate(); err != nil {
					log.Die("%s", err)
				}
				if err := salt.Provision(profile, host); err != nil {
					log.Die("%s", err)
				}
				if err := vsphere.Remove(); err != nil {
					log.Die("%s", err)
				}
			default:
				log.Die("I don't recognize this platform!")
			}
		}
	},
}
