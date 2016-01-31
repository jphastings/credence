package viewModels

import (
  "net/url"
  "github.com/jphastings/credence/lib/models"
)

type CredRecordProps struct {
  Props

  Assertion string
  Statement string
  ProofUri string
  Source DisplayedUri
  Author models.User
  CredHash string
}

func RetrieveCredRecord(credRecord *models.CredRecord) CredRecordProps {
  cred := credRecord.Cred()
  sourceUri, _ := url.Parse(cred.SourceUri)
  props := CredRecordProps{
    Assertion: cred.Assertion.String(),
    Statement: cred.GetHumanReadable().Statement,
    ProofUri: cred.ProofUri,
    Source: DisplayedUri{
      String: sourceUri.Host,
      Url: cred.SourceUri,
    },
    Author: credRecord.Author,
    CredHash: credRecord.CredHash,
  }
  return props
}