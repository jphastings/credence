package web

import (
  "io/ioutil"
  "net/url"
  "net/http"
  "html/template"
  "github.com/jphastings/credence/lib/models"
  "github.com/jphastings/credence/lib/helpers"
)

var tpl *template.Template

type SourceDetails struct {
  String string
  Url string
}

type TemplateProps struct {
  Assertion string
  Statement string
  ProofUri string
  SourceUri SourceDetails
  Author models.User
}

func init() {
  data, _ := ioutil.ReadFile("templates/cred_info.tpl.html")
  tpl, _ = template.New("cred_info").Parse(string(data))
}

func InfoCredHandler(w http.ResponseWriter, r *http.Request) {
  if r.Method != "GET" {
    w.WriteHeader(http.StatusMethodNotAllowed)
    return
  }

  // TODO: This is a hack, need proper routing
  credHash := r.URL.Path[12:]

  credRecord, found := helpers.CredRecordFromCredHash(credHash)
  if !found {
    w.WriteHeader(http.StatusNotFound)
    // TODO: Pretty 404
    return
  }

  cred := credRecord.Cred()

  sourceUri, _ := url.Parse(cred.SourceUri)

  props := TemplateProps{
    Assertion: cred.Assertion.String(),
    Statement: cred.GetHumanReadable().Statement,
    ProofUri: cred.ProofUri,
    SourceUri: SourceDetails{
      String: sourceUri.Host,
      Url: cred.SourceUri,
    },
    Author: credRecord.Author,
  }

  w.WriteHeader(http.StatusOK)
  tpl.Execute(w, props)
}