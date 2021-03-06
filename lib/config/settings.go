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

func AssetDir(path string) string {
  // TODO: look in ../share/credence/ if not immediately visible
  return path
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
  DB struct {
    Type string
    ConnectionString string
  }
  Broadcaster struct {
    Host string
    Port int
  }
  Broadcatcher struct {
    Host string
    Port int
  }
  SearchRequests struct {
    ForwardProximityLimit int32
  }
  Application struct {
    OpenWebUIOnStart bool
  }
}

func WriteDefaultConfig(configPath string, config *Config) {
  cfgStr := `# Credence config file
[Server]
Host=127.0.0.1
Port=27339
[DB]
Type=sqlite3
ConnectionString=credence.db
[Broadcaster]
Host=0.0.0.0
Port=27336
[Broadcatcher]
Host=0.0.0.0
Port=27334
[SearchRequests]
ForwardProximityLimit=3
[Application]
OpenWebUIOnStart=true
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
