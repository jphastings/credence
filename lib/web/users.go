package web

import (
  "io"
  "strconv"
  "net/http"
  "github.com/golang/protobuf/proto"
  "github.com/golang/protobuf/jsonpb"
  "github.com/jphastings/credence/lib/models"
  "github.com/jphastings/credence/lib/helpers"
  "github.com/jphastings/credence/lib/definitions/credence"
)

func UsersHandler(w http.ResponseWriter, r *http.Request) {
  switch r.Method {
    case "GET": UserDetailHandler(w, r)
    case "POST": AddUserHandler(w, r)
    default: w.WriteHeader(http.StatusMethodNotAllowed)
  }
}

func UserDetailHandler(w http.ResponseWriter, r *http.Request) {
  // TODO: This is a hack, need proper routing
  userParam := r.URL.Path[7:]

  userId, err := strconv.ParseUint(userParam, 10, 64)
  if err != nil {
    w.WriteHeader(http.StatusBadRequest)
    return
  }
  db := models.DB()

  user := &models.User{}
  db.Where(&models.User{ID: uint(userId)}).First(user)
  if db.NewRecord(user) {
    w.WriteHeader(http.StatusNotFound)
    return
  }

  // Respond over HTTP
  switch r.Header.Get("Accept") {
  case "application/vnd.google.protobuf":
    UserDetailProtoBufHandler(w, helpers.AssertIdentity(user))
  case "application/json":
    UserDetailJSONHandler(w, helpers.AssertIdentity(user))
    return
  case "text/html":
    UserDetailHTMLHandler(w, user)
    return
  default:
    w.WriteHeader(http.StatusNotAcceptable)
    return
  }
}

func UserDetailHTMLHandler(w http.ResponseWriter, user *models.User) {
  // TODO: Display page with details about this user, rather thatn just redirecting
  w.Header().Set("Location", user.IdentityUri)
  w.WriteHeader(http.StatusFound)
}

func UserDetailProtoBufHandler(w http.ResponseWriter, identityAssertion *credence.IdentityAssertion) {
  w.Header().Set("Content-Type", "application/vnd.google.protobuf")
  marshaled, _ := proto.Marshal(identityAssertion)

  w.WriteHeader(http.StatusOK)
  io.WriteString(w, string(marshaled))
}

func UserDetailJSONHandler(w http.ResponseWriter, identityAssertion *credence.IdentityAssertion) {
  w.Header().Set("Content-Type", "application/json")
  marshaler := jsonpb.Marshaler{}
  marshaled, _ := marshaler.MarshalToString(identityAssertion)

  w.WriteHeader(http.StatusOK)
  io.WriteString(w, marshaled)
}

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