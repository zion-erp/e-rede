package erede

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"sync"

	"github.com/hashicorp/go-hclog"
)

var (
	lock = sync.Mutex{}
)

type ERedeClient struct {
	headers map[string]string
	Config  Config
	log     hclog.Logger
}

func NewClient(config Config, log hclog.Logger) *ERedeClient {
	c := &ERedeClient{
		Config: config,
		log:    log,
	}
	return c
}

func (c *ERedeClient) Delete(url string) *Response {
	return c.Send(http.MethodDelete, url, nil)
}

func (c *ERedeClient) Get(url string) *Response {
	return c.Send(http.MethodGet, url, nil)
}

func (c *ERedeClient) Post(url string, data any) *Response {
	return c.Send(http.MethodPost, url, data)
}

func (c *ERedeClient) Put(url string, data any) *Response {
	return c.Send(http.MethodPut, url, data)
}

func (c *ERedeClient) basicAuth(username, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}

func (c *ERedeClient) initNewRequest(hasBody bool) {
	c.headers = make(map[string]string)
	c.AddHeader("User-Agent", fmt.Sprintf("ERede/%s Golang SDK", c.Config.Version))
	c.AddHeader("Authorization", fmt.Sprintf("Basic %s", c.basicAuth(c.Config.Filiation, c.Config.Token)))
	if hasBody {
		c.AddHeader("Content-Type", "application/json")
	} else {
		c.AddHeader("Content-Length", "0")
	}
	c.AddHeader("Transaction-Response", "brand-return-opened")
}

func (c *ERedeClient) Send(method, url string, data any) *Response {
	var body io.Reader

	hasBody := data != nil
	if hasBody {
		dt, err := json.Marshal(data)
		if err != nil {
			return &Response{
				Error: errInRequestBody(url, dt),
			}
		}
		body = bytes.NewReader(dt)
		c.log.Debug(fmt.Sprintf("%s [url] = %s [body] = %s \r\n", method, url, string(dt[:])))
	} else {
		c.log.Debug(fmt.Sprintf("%s [url] = %s \r\n", method, url))
	}

	c.initNewRequest(hasBody)

	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return &Response{Error: err}
	}
	c.setHeaders(req)

	httpClient := &http.Client{}
	res, err := httpClient.Do(req)
	if err != nil {
		c.log.Debug(fmt.Sprintf("error do: %s", err.Error()))
		return &Response{
			Error: err,
			Code:  res.StatusCode,
		}
	}
	defer res.Body.Close()

	content, err := io.ReadAll(res.Body)
	if err != nil {
		c.log.Debug(fmt.Sprintf("error readAll: %s\r\nbody: %s", err.Error(), string(content[:])))
		return &Response{
			Error: err,
			Code:  res.StatusCode,
		}
	}

	c.log.Debug(fmt.Sprintf("response: %d: %s", res.StatusCode, string(content[:])))

	response := &Response{
		Code:    res.StatusCode,
		Payload: content,
	}

	if res.StatusCode >= http.StatusBadRequest {
		err = json.Unmarshal(content, &response.ApiErrors)
		if err != nil {
			er := &ApiError{}
			err = json.Unmarshal(content, &er)
			if err == nil {
				response.ApiErrors = append(response.ApiErrors, er)
			}
		}
		response.Error = errors.New(string(content[:]))
	}

	return response
}

func (c *ERedeClient) AddHeader(key, value string) {
	lock.Lock()
	defer lock.Unlock()
	if c.headers == nil {
		c.headers = make(map[string]string)
	}
	c.headers[key] = value
}

func (c *ERedeClient) Header(key string) string {
	return c.headers[key]
}

func (c *ERedeClient) setHeaders(req *http.Request) {
	lock.Lock()
	defer lock.Unlock()
	for k, v := range c.headers {
		req.Header.Add(k, v)
	}
	req.Header.Add("Accept", "application/json")
}

func ByteArrayToStruct[T any](b []byte) (*T, error) {
	r := new(T)
	if len(b) > 0 {
		err := json.Unmarshal(b, r)
		if err != nil {
			return nil, err
		}
	}
	return r, nil
}

func UpdateStructWithByteArray[T any](b []byte, c *T) error {
	if len(b) > 0 {
		err := json.Unmarshal(b, c)
		if err != nil {
			return err
		}
	}
	return nil
}

func errInRequestBody(method string, obj any) error {
	return fmt.Errorf("error in marshal request body method [%s]\r\nbody: %v", method, obj)
}
