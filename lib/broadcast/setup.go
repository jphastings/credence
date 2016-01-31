package broadcast

import (
  "fmt"
  "log"
  "github.com/zeromq/goczmq"
  "github.com/jphastings/credence/lib/config"
  "github.com/jphastings/credence/lib/helpers"
)

var receiver *goczmq.Sock
var broadcaster *goczmq.Sock
var pitcher *goczmq.Sock

func Setup() {
  config := config.Read()
  broadcastUri := fmt.Sprintf("tcp://%s:%d", config.Broadcaster.Host, config.Broadcaster.Port)

  var err error
  broadcaster, err = goczmq.NewPub(broadcastUri)
  if err != nil {
    panic(err)
  }
  log.Println("Broadcaster will start on", broadcastUri)

  pitcher = goczmq.NewSock(goczmq.Push)

  receiver, err = goczmq.NewPull("inproc://broadcast")
  if err != nil {
    panic(err)
  }

  helpers.ConnectToPeers(true, ConnectToBroadcatcher)
}

func ConnectToBroadcatcher(uri string) error {
  broadcatcherUri := fmt.Sprintf("tcp://%s", uri)
  log.Println("Connecting to broadcatcher at", broadcatcherUri)
  err := pitcher.Connect(broadcatcherUri)
  return err
}