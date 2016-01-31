package helpers

import (
  "net/http"
  "io/ioutil"
  "path/filepath"
  "html/template"
  "github.com/jphastings/credence/lib/web/view_models"
)

var templates = make(map[string]*template.Template)

func init() {
  paths, _ := filepath.Glob("templates/*.tpl.html")
  for _, path := range paths {
    data, _ := ioutil.ReadFile(path)
    name := path[10:len(path)-9]
    tpl, _ := template.New(name).Parse(string(data))
    templates[name] = tpl
  }
}

func RenderTemplate(w http.ResponseWriter, templateName string, props viewModels.Props) bool {
  tpl := templates[templateName]
  if tpl == nil {
    w.WriteHeader(http.StatusInternalServerError)
    return false
  }
  tpl.Execute(w, props)
  return true
}