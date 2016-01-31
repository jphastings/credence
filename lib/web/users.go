package web

import (
  "strconv"
  "net/http"
  "github.com/jphastings/credence/lib/models"
)

func UsersHandler(w http.ResponseWriter, r *http.Request) {
  switch r.Method {
    case "GET":
      // TODO: This is a hack, need proper routing
      userId, err := strconv.ParseUint(r.URL.Path[7:], 10, 64)
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

      // TODO: Display page with details about this user, rather thatn just redirecting
      w.Header().Set("Location", user.IdentityUri)
      w.WriteHeader(http.StatusSeeOther)
      return
    default: w.WriteHeader(http.StatusMethodNotAllowed)
  }
}
