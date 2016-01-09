package config

import (
  "os"
  "flag"
  "os/user"
  "path/filepath"
  "io/ioutil"
  "gopkg.in/gcfg.v1"
)

var config Config
var configDir string

func Setup() {
  usr, _ := user.Current()
  defaultDir := filepath.Join(usr.HomeDir, ".credence")
  flag.StringVar(&configDir, "config", defaultDir, "the directory config is stored in")
  flag.Parse()

  os.MkdirAll(configDir, 0700)
  Read()
}

func ConfigFile(filename string) string {
  return filepath.Join(configDir, filename)
}

func Read() Config {
  emptyConfig := Config{}
  if config == emptyConfig {
    configPath := ConfigFile("settings.config")

    err := gcfg.ReadFileInto(&config, configPath)
    if err != nil {
      WriteDefaultConfig(configPath, &config)
    }
  }
  return config
}

type Config struct {
  Server struct {
    Host string
    Port int
  }
  Broadcaster struct {
    Host string
    Port int
  }
  SearchRequests struct {
    ForwardProximityLimit int32
  }
}

func WriteDefaultConfig(configPath string, config *Config) {
  cfgStr := `# Credence config file
[Server]
Host=127.0.0.1
Port=8808
[Broadcaster]
Host=0.0.0.0
Port=9099
[SearchRequests]
ForwardProximityLimit=3
`

  err := ioutil.WriteFile(configPath, []byte(cfgStr), 0600)
  if err != nil {
    panic(err)
  }

  err = gcfg.ReadFileInto(config, configPath)
  if err != nil {
    panic("The default config settings are invalid.")
  }
}
