package web

import (
  "strconv"
  "net/http"
  "github.com/jphastings/credence/lib/models"
  "github.com/jphastings/credence/lib/helpers"
)

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

  identityAssertion := helpers.AssertIdentity(user)
  helpers.ModelNegotiator().Negotiate(w, r, identityAssertion)
}

func UserDetailHTMLHandler(w http.ResponseWriter, user *models.User) {
  // TODO: Display page with details about this user, rather thatn just redirecting
  w.Header().Set("Location", user.IdentityUri)
  w.WriteHeader(http.StatusFound)
}
