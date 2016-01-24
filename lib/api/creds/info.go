package api

import (
  "fmt"
  "io/ioutil"
  "net/url"
  "net/http"
  "html/template"
  "github.com/jphastings/credence/lib/helpers"
  "github.com/jphastings/credence/lib/definitions/credence"
)

var AssertionStrings = map[credence.Cred_AssertionType]string{
  0: "not commentable upon",
  1: "true",
  2: "false",
  3: "ambiguous",
}
var tpl *template.Template

type KeyDetails struct {
  String string
  Url string
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

  cred, err := helpers.CredFromBase64(r.URL.Query()["cred"][0])
  if err != nil {
    w.WriteHeader(http.StatusBadRequest)
    return
  }

  props := struct {
    Assertion string
    Statement string
    ProofUri string
    Keys []KeyDetails
  }{
    Assertion: AssertionStrings[cred.Assertion],
    Statement: cred.GetHumanReadable().Statement,
    ProofUri: cred.ProofUri,
  }

  for _, key := range cred.Keys {
    keyDetails := KeyDetails{String: key}

    u, err := url.Parse(key)
    if err == nil {
      keyDetails.String = fmt.Sprintf("%s/…", u.Host)
      keyDetails.Url = key
    }

    props.Keys = append(props.Keys, keyDetails)
  }

  w.WriteHeader(http.StatusOK)
  tpl.Execute(w, props)
}