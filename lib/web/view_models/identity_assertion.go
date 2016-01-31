package viewModels

import (
  "net/url"
  "encoding/hex"
  "github.com/jphastings/credence/lib/models"
  "github.com/jphastings/credence/lib/definitions/credence"
)

type IdentityAssertionProps struct {
  Props

  Name string
  Identity DisplayedUri
  Proof DisplayedUri
  LatestCreds []Props
  Fingerprint string
}

func RetrieveIdentityAssertion(identityAssertion *credence.IdentityAssertion) IdentityAssertionProps {
  var creds []Props
  // TODO: I don't think this fingerprint is actually the correct identity…
  latestCreds := models.CredRecordSample(identityAssertion.Fingerprint, 15)
  for _, credRecord := range latestCreds {
    creds = append(creds, RetrieveCredRecord(credRecord))
  }

  proofUri, _ := url.Parse(identityAssertion.ProofUri)
  identityUri, _ := url.Parse(identityAssertion.IdentityUri)
  props := IdentityAssertionProps{
    Name: identityAssertion.Name,
    Identity: DisplayedUri{
      String: identityUri.Host,
      Url: identityAssertion.IdentityUri,
    },
    Proof: DisplayedUri{
      String: proofUri.Host,
      Url: identityAssertion.ProofUri,
    },
    // TODO: Find better way of representing this
    Fingerprint: hex.EncodeToString(identityAssertion.Fingerprint),
    
    LatestCreds: creds,
  }
  return props
}