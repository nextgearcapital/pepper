// vSphere template
// The purpose of this is to abstract away all the values in a vSphere Salt Profile to be used
// elsewhere in pepper.

package vsphere

import (
	"os"
	"path/filepath"
	"text/template"

	"github.com/nextgearcapital/pepper/pkg/log"

	"github.com/spf13/viper"
)

// ProfileConfig :
type ProfileConfig struct {
	Platform     string   `mapstructure:"platform"`
	Environment  string   `mapstructure:"environment"`
	InstanceType string   `mapstructure:"instance_type"`
	Provider     string   `mapstructure:"provider"`
	Template     string   `mapstructure:"template"`
	CPU          int      `mapstructure:"cpu"`
	Memory       int      `mapstructure:"memory"`
	DiskSize     int      `mapstructure:"disksize"`
	DHCP         bool     `mapstructure:"dhcp"`
	Network      string   `mapstructure:"network"`
	IPAddress    string   `mapstructure:"ip_address"`
	Gateway      string   `mapstructure:"gateway"`
	Subnet       string   `mapstructure:"subnet"`
	Domain       string   `mapstructure:"domain"`
	DNSServers   []string `mapstructure:"dns_servers"`
	Cluster      string   `mapstructure:"cluster"`
	Folder       string   `mapstructure:"folder"`
	Datastore    string   `mapstructure:"datastore"`
	IsCoreOS     bool     `mapstructure:"is_coreos"`
	ConfigData   string   `mapstructure:"config_data"`
}

var (
	configPath   = "/etc/pepper/config.d/vmware"
	saltProfiles = "/etc/salt/cloud.profiles.d"
)

// Prepare :
func (profile *ProfileConfig) Prepare(platform, environment, instancetype, template, ipAddress string) error {
	if err := os.MkdirAll(configPath, 0644); err != nil {
		log.Warn("%s", err)
	}

	profile.Platform = platform
	profile.Environment = environment
	profile.InstanceType = instancetype
	profile.Template = template
	profile.IPAddress = ipAddress

	switch instancetype {
	case "nano":
		profile.CPU = 1
		profile.Memory = 512
		profile.DiskSize = 20
	case "micro":
		profile.CPU = 1
		profile.Memory = 1024
		profile.DiskSize = 20
	case "small":
		profile.CPU = 1
		profile.Memory = 2048
		profile.DiskSize = 40
	case "medium":
		profile.CPU = 2
		profile.Memory = 4096
		profile.DiskSize = 60
	case "large":
		profile.CPU = 2
		profile.Memory = 8192
		profile.DiskSize = 80
	case "xlarge":
		profile.CPU = 4
		profile.Memory = 16384
		profile.DiskSize = 100
	case "ultra":
		profile.CPU = 8
		profile.Memory = 32768
		profile.DiskSize = 160
	case "mega":
		profile.CPU = 16
		profile.Memory = 65536
		profile.DiskSize = 200
	}

	viper.SetConfigName(environment)
	viper.SetConfigType("yaml")
	viper.AddConfigPath(configPath)

	if err := viper.ReadInConfig(); err != nil {
		log.Die("Couldn't read the config! Did you put it in /etc/pepper/config.d/vmware? %s", err)
	}

	if err := viper.Unmarshal(profile); err != nil {
		log.Die("%s", err)
	}
	return nil
}

// Generate :
func (profile *ProfileConfig) Generate() error {
	compiled, err := template.New("vsphere_profile").Parse(vsphereTemplate)
	if err != nil {
		log.Die("%s", err)
	}

	profileName := profile.Platform + "-" + profile.Environment + "-" + profile.InstanceType + ".conf"
	f, err := os.OpenFile(filepath.Join(saltProfiles, profileName), os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		log.Die("%s", err)
	}

	if err := compiled.Execute(f, profile); err != nil {
		log.Die("%s", err)
	}
	return nil
}

// Remove :
func (profile *ProfileConfig) Remove() error {
	profilePath := saltProfiles + "/" + profile.Platform + "-" + profile.Environment + "-" + profile.InstanceType + ".conf"
	if err := os.Remove(profilePath); err != nil {
		log.Die("%s", err)
	}
	return nil
}

const vsphereTemplate = `
{{- .Platform}}-{{.Environment}}-{{.InstanceType}}:
  provider: {{.Provider}}
  clonefrom: {{.Template}}
  num_cpus: {{.CPU}}
  memory: {{.Memory}}
  devices:
    cd:
      CD/DVD drive 1:
        device_type: client_device
        mode: atapi
    disk:
      Hard disk 1:
        size: {{.DiskSize}}
    network:
      Network adapter 1:
        name: {{.Network}}
        adapter_type: vmxnet3
        switch_type: distributed
		{{- if eq .DHCP false}}
        ip: {{.IPAddress}}
        gateway: {{.Gateway}}
        subnet_mask: {{.Subnet}}
	    {{- end}}
    scsi:
      SCSI controller 0:
        type: lsilogic
  {{- if .Domain}}
  domain: {{.Domain}}
  {{- end}}
  {{- if .DNSServers}}
  dns_servers:
  {{- range $value := .DNSServers}}
  - {{$value}}
  {{- end}}
  {{- end}}
  cluster: {{.Cluster}}
  folder: {{.Folder}}
  datastore: {{.Datastore}}
  template: False
  power_on: True
  deploy: False
  extra_config:
    cpu.hotadd: 'yes'
    mem.hotadd: 'yes'
	{{- if .IsCoreOS}}
	guestinfo.coreos.config.data: {{.ConfigData}}
	guestinfo.coreos.config.data.encoding: base64
	{{- end}}
`
