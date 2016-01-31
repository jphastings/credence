package helpers

import (
  "github.com/jphastings/credence/lib/models"
)

func ConnectToPeers(broadcatchers bool, callback func(string) error) {
  db := models.DB()
  peers, _ := db.Model(models.Peer{}).Rows()

  for peers.Next() {
    var (
      peerUri string
      isBroadcatcher bool
    )
    peers.Scan(&peerUri, &isBroadcatcher)
    if broadcatchers == isBroadcatcher {
      // TODO: Catch error
      _ = callback(peerUri)
    }
  }
}