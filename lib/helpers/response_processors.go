package helpers

import (
	"strings"
  "net/http"
  "github.com/jchannon/negotiator"
  "github.com/golang/protobuf/proto"
  "github.com/golang/protobuf/jsonpb"
  "github.com/jphastings/credence/lib/models"
  "github.com/jphastings/credence/lib/web/view_models"
)

type JSONPBResponseProcessor struct {}
type RawPBResponseProcessor struct {}
type HTMLTemplatePBResponseProcessor struct {}

func ModelNegotiator() *negotiator.Negotiator {
	return negotiator.New(
    &HTMLTemplatePBResponseProcessor{},
    &JSONPBResponseProcessor{},
    &RawPBResponseProcessor{},
  )
}

func (*JSONPBResponseProcessor) CanProcess(mediaRange string) bool {
  return strings.EqualFold(mediaRange, "application/json")
}

func (*JSONPBResponseProcessor) Process(w http.ResponseWriter, model interface{}) error {
  msg := model.(proto.Message)
  w.Header().Set("Content-Type", "application/json")
  marshaler := jsonpb.Marshaler{}

  w.WriteHeader(http.StatusOK)
  marshaler.Marshal(w, msg)
  return nil
}

func (*RawPBResponseProcessor) CanProcess(mediaRange string) bool {
  return strings.EqualFold(mediaRange, "application/vnd.google.protobuf")
}

func (*RawPBResponseProcessor) Process(w http.ResponseWriter, model interface{}) error {
  msg := model.(proto.Message)
  w.Header().Set("Content-Type", "application/vnd.google.protobuf")
  marshaled, err := proto.Marshal(msg)
  if err != nil {
    return err
  }

  w.WriteHeader(http.StatusOK)
  w.Write(marshaled)
  return nil
}

func (*HTMLTemplatePBResponseProcessor) CanProcess(mediaRange string) bool {
  return strings.EqualFold(mediaRange, "text/html")
}

func (*HTMLTemplatePBResponseProcessor) Process(w http.ResponseWriter, model interface{}) error {
  w.Header().Set("Content-Type", "text/html")

  var (
    templateName string
    props viewModels.Props
  )

  msg, isMessage := model.(proto.Message)
  record, isCredRecord := model.(*models.CredRecord)
  if isMessage {
    templateName, props = viewModels.Retrieve(msg)
  } else if isCredRecord {
    templateName, props = viewModels.Retrieve(record)
  } else {
    panic("Don't how know to HTML render this thing")
  }
  
  if RenderTemplate(w, templateName, props) {
    w.WriteHeader(http.StatusOK)
  }
  return nil
}