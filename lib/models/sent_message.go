package models

import "time"

type SentMessage struct {
  MessageHash string
  SentAt time.Time
}
