package api

import (
  "fmt"
  "net/http"
  "github.com/jphastings/credence/lib/helpers"
)

func RawCredHandler(w http.ResponseWriter, r *http.Request) {
  if r.Method != "GET" {
    w.WriteHeader(http.StatusMethodNotAllowed)
    return
  }

  b64Cred := r.URL.Query()["cred"][0]

  if len(b64Cred) > 0 {
    cred, err := helpers.CredFromBase64(b64Cred)
    if err != nil {
      w.WriteHeader(http.StatusBadRequest)
      return
    }
    _, credHash := helpers.StoreCredUnknownAuthor(cred)
    w.Header().Set("Location", fmt.Sprintf("/creds/info/%s", credHash))
    w.WriteHeader(http.StatusSeeOther)
    return
  }

  w.WriteHeader(http.StatusBadRequest)
  return
}