// Copyright © 2016 NAME HERE <EMAIL ADDRESS>
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
	"github.com/spf13/cobra"

	"github.com/nextgearcapital/pepper/pkg/device42"
	"github.com/nextgearcapital/pepper/pkg/log"
	"github.com/nextgearcapital/pepper/pkg/salt"
)

func init() {
	RootCmd.AddCommand(destroyCmd)
}

var destroyCmd = &cobra.Command{
	Use:   "destroy",
	Short: "Destroys VM's",
	Long: `This command will call salt-cloud with the -d parameter.
	Multiple VM's can be specified for destruction. For example:

You can destroy a single VM:

$ pepper destroy web01

Or multiple VM's:

$ pepper destroy web01 web02 web03`,
	Run: func(cmd *cobra.Command, args []string) {
		// This isn't necessary. It simply makes it easier to understand.
		hosts := args

		if len(hosts) == 0 {
			log.Die("You didn't specify any hosts to destroy.")
		}

		for _, host := range hosts {
			if err := salt.Destroy(host); err != nil {
				log.Die("%s", err)
			}
			if ipam == true {
				if err := device42.DeleteDevice(host); err != nil {
					log.Warn("%s", err)
				}
			}
		}
		log.CleanExit("Success!")
	},
}
