package host

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/cemanaral/kron/pkg/sshutil"
	"gopkg.in/yaml.v3"
)

const HOSTS_PATH = "/etc/kron/hosts.yaml"

var Hosts = []Host{}

type HostConfig struct {
	Id             string `yaml:"host"`
	Address        string `yaml:"address"`
	User           string `yaml:"user"`
	Password       string `yaml:"password"`
	PrivateKeyPath string `yaml:"private-key"`
}

type Host struct {
	SshConn sshutil.SshConnection
	HostConfig `yaml:",inline"`
}

func (h Host) GetCronInformation() string {
	return h.SshConn.ExecuteCommand("cat /etc/crontab")
}

func LoadHosts() {
	yfile, err := ioutil.ReadFile(HOSTS_PATH)
	if err != nil {
		log.Fatal(err)
	}
	err2 := yaml.Unmarshal(yfile, &Hosts)
	if err2 != nil {
		log.Fatal(err2)
	}
	for i := 0; i < len(Hosts); i++ {
		Hosts[i].determineAuthMethod()
	}

}

func (h *Host) determineAuthMethod() {
	if h.Password != "" {
		h.SshConn = sshutil.SshConnectionWithPassword{}
	} else if h.PrivateKeyPath != "" {
		h.SshConn = sshutil.SshConnectionWithPrivateKey{}
	}
}

func (host Host) String() string {
	return fmt.Sprintf(`host -> %s
address -> %s
user -> %s
password -> %s
private-key -> %s
`, host.Id, host.Address, host.User, host.Password, host.PrivateKeyPath)
}
