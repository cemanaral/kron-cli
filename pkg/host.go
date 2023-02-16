package pkg

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v3"
)

const HOSTS_PATH = "/etc/kron/hosts.yaml"
var Hosts = HostList{};

type HostList []Host

type Host struct {
     Id string
     Address string
	User string
	Password string
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

func LoadHosts() {
     Hosts.load()
}
