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

func init() {
  usr, _ := user.Current()
  defaultDir := filepath.Join(usr.HomeDir, ".credence")
  flag.StringVar(&configDir, "config", defaultDir, "the directory config is stored in")
  flag.Parse()
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
}

func WriteDefaultConfig(configPath string, config *Config) {
  cfgStr := `# Credence config file
[Server]
Host=127.0.0.1
Port=8808
`

  os.MkdirAll(filepath.Dir(configPath), 0700)

  err := ioutil.WriteFile(configPath, []byte(cfgStr), 0600)
  if err != nil {
    panic(err)
  }

  err = gcfg.ReadFileInto(config, configPath)
  if err != nil {
    panic("The default config settings are invalid.")
  }
}
