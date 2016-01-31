package models

type Peer struct {
  Server string
  IsBroadcatcher bool `sql:"default:false"`
}
