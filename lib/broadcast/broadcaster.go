package broadcast

import (
  "log"
  "sync"
  "time"
  "encoding/hex"
  "github.com/spacemonkeygo/openssl"
  "github.com/jphastings/credence/lib/models"
)

func StartBroadcaster(wg sync.WaitGroup) {
  defer wg.Done()

  db := models.DB()

  msgBytes := make([]byte, 524288) // 0.5 Mb

  for {
    _, err := receiver.Read(msgBytes)
    if err != nil {
      log.Print(err)
      continue
    }

    hash, err := openssl.SHA1(msgBytes)
    if err != nil {
      panic(err)
    }

    messageHash := hex.EncodeToString(hash[:])
    log.Println("Checking message", messageHash)

    var previouslySent models.SentMessage
    db.Where("message_hash = ? AND sent_at > ?", messageHash, time.Now().Add(-5 * time.Minute)).First(&previouslySent)

    if db.NewRecord(previouslySent.MessageHash) {
      log.Println("Broadcasting message", messageHash)

      _, err = broadcaster.Write(msgBytes)
      if err != nil {
        panic(err)
      }

      log.Println("Pitching message", messageHash)

      _, err = pitcher.Write(msgBytes)
      if err != nil {
        panic(err)
      }

      previouslySent.SentAt = time.Now()
      previouslySent.MessageHash = messageHash
      db.Save(&previouslySent)
    } else {
      log.Println("Message already sent recently", messageHash)
    }
  }
}