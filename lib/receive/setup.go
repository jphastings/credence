package receive

import (
	"log"
	"fmt"
  "github.com/zeromq/goczmq"
  "github.com/jphastings/credence/lib/config"
  "github.com/jphastings/credence/lib/helpers"
)

var receiver *goczmq.Sock
var broadcaster *goczmq.Sock
var broadcatcher *goczmq.Sock

func Setup() {
	config := config.Read()
  broadcatchUri := fmt.Sprintf("tcp://%s:%d", config.Broadcatcher.Host, config.Broadcatcher.Port)

  var err error
  broadcaster, err = goczmq.NewPush("inproc://broadcast")
  if err != nil {
    panic(err)
  }

  broadcatcher, err = goczmq.NewPull(broadcatchUri)
  if err != nil {
    panic(err)
  }
  log.Println("Broadcatcher will start on", broadcatchUri)

  receiver = goczmq.NewSock(goczmq.Sub)
  receiver.SetSubscribe("")

  helpers.ConnectToPeers(false, ConnectToBroadcaster)
}

func ConnectToBroadcaster(uri string) error {
  broadcasterUri := fmt.Sprintf("tcp://%s", uri)
  log.Println("Connecting to broadcaster at", broadcasterUri)
  err := receiver.Connect(broadcasterUri)
  return err
}