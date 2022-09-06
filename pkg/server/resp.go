package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"reflect"

	"github.com/NganJason/Unsplash-BE/pkg/cerr"
)

type HandlerResp struct {
	Payload interface{}
	Err     error
}

func NewHandlerResp(payload interface{}, err error) *HandlerResp {
	return &HandlerResp{
		Payload: payload,
		Err:     err,
	}
}

func JsonResponse(
	writer http.ResponseWriter,
	payload interface{}, err error,
) {
	if err != nil {
		setDebugMsg(payload, err)
	}

	p, marshalErr := json.Marshal(payload)
	if marshalErr != nil {
		JsonResponse(
			writer,
			payload,
			cerr.New(
				fmt.Sprintf("marshal json resp err=%s", err.Error()),
				http.StatusInternalServerError,
			),
		)

		return
	}

	if err != nil {
		code := cerr.Code(err)
		writer.WriteHeader(code)
	} else {
		writer.WriteHeader(http.StatusOK)
	}

	writer.Header().Set("Content-Type", "application/json")
	if _, err := writer.Write(p); err != nil {
		log.Printf("Failed to handle request :%s", err)
	}
}

func setDebugMsg(
	payload interface{},
	err error,
) {
	if payload == nil || err == nil {
		return
	}

	msg := err.Error()

	debugMsgField := "DebugMsg"
	structField, found := reflect.TypeOf(payload).Elem().FieldByName(debugMsgField)
	if !found {
		return
	}

	fieldType := structField.Type
	if fieldType.Kind() != reflect.Ptr || fieldType.Elem().Kind() != reflect.String {
		return
	}

	requiredField := reflect.ValueOf(payload).Elem().FieldByName(debugMsgField)

	if requiredField.CanSet() {
		var finalMsg string

		elem := requiredField.Elem()
		if elem.IsValid() && len(elem.String()) != 0 {
			finalMsg = elem.String() + ": " + msg
		} else {
			finalMsg = msg
		}

		requiredField.Set(reflect.ValueOf(&finalMsg))
	}
}
