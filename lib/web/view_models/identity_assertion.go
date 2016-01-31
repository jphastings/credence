package viewModels

import (
  "net/url"
  "github.com/jphastings/credence/lib/models"
  "github.com/jphastings/credence/lib/definitions/credence"
)

type IdentityAssertionProps struct {
  Props

  Name string
  Identity DisplayedUri
  LatestCreds []Props
}

func RetrieveIdentityAssertion(identityAssertion *credence.IdentityAssertion) IdentityAssertionProps {
	var creds []Props
	latestCreds := models.CredRecordSample(identityAssertion.Fingerprint, 15)
	for _, credRecord := range latestCreds {
		creds = append(creds, RetrieveCredRecord(credRecord))
	}

	identityUri, _ := url.Parse(identityAssertion.IdentityUri)
  props := IdentityAssertionProps{
    Name: identityAssertion.Name,
    Identity: DisplayedUri{
    	String: identityUri.Host,
    	Url: identityAssertion.IdentityUri,
    },
    LatestCreds: creds,
  }
  return props
}