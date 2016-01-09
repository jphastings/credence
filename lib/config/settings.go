package config

import (
  "os"
  "os/user"
  "path/filepath"
  "io/ioutil"
  "gopkg.in/gcfg.v1"
)

var config Config

func Read() Config {
  emptyConfig := Config{}
  if config == emptyConfig {
    usr, _ := user.Current()
    configPath := filepath.Join(usr.HomeDir, ".credence", "settings.config")

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
