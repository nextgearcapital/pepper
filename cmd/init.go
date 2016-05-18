// Copyright © 2016 Robert Deusser <robert.deusser@nextgearcapital.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"os"
	"text/template"

	"github.com/nextgearcapital/pepper/pkg/log"

	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize Pepper",
	Long:  `Creates the necessary directories and generates a basic profile config in /etc/pepper/config.d as a starting point.`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := os.MkdirAll("/etc/pepper/config.d", 0644); err != nil {
			log.Warn("%s", err)
		}
		log.Info("Created /etc/pepper/config.d")
		if err := os.MkdirAll("/etc/pepper/provider.d", 0644); err != nil {
			log.Warn("%s", err)
		}
		log.Info("Created /etc/pepper/provider.d")

		compiled, err := template.New("vsphere_profile").Parse(configTemplate)
		if err != nil {
			log.Die("%s", err)
		}

		f, err := os.OpenFile("/etc/pepper/config.d/template.yaml", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
		if err != nil {
			log.Die("%s", err)
		}

		if err := compiled.Execute(f, nil); err != nil {
			log.Die("%s", err)
		}
	},
}

func init() {
	RootCmd.AddCommand(initCmd)
}

const configTemplate = `
provider: vcenter01
dhcp: true
network: Development
gateway: 192.168.1.1
subnet: 255.255.255.0
domain: google.com
dns_servers:
- 8.8.8.8
- 8.8.4.4
cluster: Development
folder: Development
datastore: test
`
