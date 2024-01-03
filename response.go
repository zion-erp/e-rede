package erede

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"
)

type ApiError struct {
	RequestDateTime *time.Time `json:"requestDateTime"`
	Code            string     `json:"returnCode"`
	Message         string     `json:"returnMessage"`
}

type Response struct {
	Payload   []byte      `json:"payload"`
	Code      int         `json:"-"`
	Error     error       `json:"error"`
	Message   string      `json:"message"`
	ApiErrors []*ApiError `json:"ApiErrors"`
}

func (r *Response) Ok() bool {
	return r.Code == 200
}

func (r *Response) Created() bool {
	return r.Code == 201
}

func (r *Response) OkCreated() bool {
	return r.Code == 200 || r.Code == 201
}

func (r *Response) ApiErrorsStr() string {
	resp := ""

	for _, err := range r.ApiErrors {
		resp += fmt.Sprintf("Code: %s Message: %s\n", err.Code, err.Message)
	}

	return resp
}

func (r *Response) HasApiErrors() bool {
	return len(r.ApiErrors) > 0
}

func (r *Response) UnmarshalJSON(b []byte) error {
	if len(b) == 0 {
		return fmt.Errorf("no bytes to unmarshal")
	}
	var f interface{}
	err := json.Unmarshal(b, &f)
	if err != nil {
		return err
	}

	m := f.(map[string]interface{})

	r.Message = m["message"].(string)
	if payload, ok := m["payload"]; ok {
		r.Payload, err = json.Marshal(payload)
		if err != nil {
			return err
		}
	}

	if errs, ok := m["error"]; ok {
		errs := errs.([]interface{})
		if len(errs) > 0 {
			ret := []string{
				r.Message,
			}
			for _, e := range errs {
				rs := e.(map[string]interface{})
				s := fmt.Sprintf("%s [param: %s]", rs["message"].(string), rs["param"].(string))
				ret = append(ret, s)
			}
			r.Error = errors.New(strings.Join(ret, " "))
		}
	}

	return nil
}
