package web

import (
  "io"
  "net/http"
  "github.com/golang/protobuf/jsonpb"
  "github.com/jphastings/credence/lib/definitions/credence"
)

func AddUserHandler(w http.ResponseWriter, r *http.Request) {
  confirmAdd := len(r.URL.Query()["confirm"]) > 0
  
  identityAssertion := &credence.IdentityAssertion{}
  if err := jsonpb.Unmarshal(r.Body, identityAssertion); err != nil {
    w.WriteHeader(http.StatusBadRequest)
    return
  }

  if confirmAdd {
    // TODO: Implement
    w.WriteHeader(http.StatusNotImplemented)
    return
  } else {
    // TODO: Proper CORS
    w.Header().Set("Access-Control-Allow-Origin", "*")

    io.WriteString(w, identityAssertion.String())
    return
  }

}