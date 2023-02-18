package host

import (
	"fmt"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v3"
)

const HOSTS_PATH = "/etc/kron/hosts.yaml"
var Hosts = HostList{};

type HostList []Host

type Host struct {
     Id string `yaml:"host"`
     Address string `yaml:"address"`
     User string `yaml:"user"`
     Password string `yaml:"password"`
     PrivateKeyPath string `yaml:"private-key"`
}

func (h Host) GetCronInformation() string {
     if h.Password != "" {
          fmt.Println(h.Id, "password is chosen")
     } else if h.PrivateKeyPath != "" {
          fmt.Println(h.Id, "private key method is chosen")
     }
     log.Fatalf("No valid authentication method was selected for '%s'!", h.Id)
     return ""
}

func (hosts *HostList ) load() {
     yfile, err := ioutil.ReadFile(HOSTS_PATH)
     if err != nil {
          log.Fatal(err)
     }
     err2 := yaml.Unmarshal(yfile, hosts)
     if err2 != nil {
          log.Fatal(err2)
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

func LoadHosts() {
     Hosts.load()
}
