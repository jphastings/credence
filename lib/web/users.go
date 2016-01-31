package web

import (
  "net/http"
  "github.com/jphastings/credence/lib/web/users"
)

func UsersHandler(w http.ResponseWriter, r *http.Request) {
  switch r.Method {
    case "GET": web.UserDetailHandler(w, r)
    case "POST": web.AddUserHandler(w, r)
    default: w.WriteHeader(http.StatusMethodNotAllowed)
  }
}
