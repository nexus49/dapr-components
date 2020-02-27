// ------------------------------------------------------------
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.
// ------------------------------------------------------------

package httpbasic

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/dapr/components-contrib/bindings"
)

// HTTPSource is a binding for an http url endpoint invocation
// nolint:golint
type HTTPSource struct {
	metadata httpMetadata
}

type httpMetadata struct {
	URL      string `json:"url"`
	Method   string `json:"method"`
	Username string `json:"username"`
	Password string `json:"password"`
}

// NewHTTP returns a new HTTPSource
func NewHTTP() *HTTPSource {
	return &HTTPSource{}
}

// Init performs metadata parsing
func (h *HTTPSource) Init(metadata bindings.Metadata) error {
	b, err := json.Marshal(metadata.Properties)
	if err != nil {
		return err
	}

	var m httpMetadata
	err = json.Unmarshal(b, &m)
	if err != nil {
		return err
	}

	h.metadata = m
	return nil
}

func (h *HTTPSource) get(url string) ([]byte, error) {
	client := http.Client{Timeout: time.Second * 60}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.SetBasicAuth(h.metadata.Username, h.metadata.Password)

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp != nil && resp.Body != nil {
		resp.Body.Close()
	}
	return b, nil
}

func (h *HTTPSource) Read(handler func(*bindings.ReadResponse) error) error {
	b, err := h.get(h.metadata.URL)
	if err != nil {
		return err
	}

	handler(&bindings.ReadResponse{
		Data: b,
	})
	return nil
}

func (h *HTTPSource) Write(wq *bindings.WriteRequest) error {
	client := http.Client{Timeout: time.Second * 5}
	req, err := http.NewRequest(h.metadata.Method, h.metadata.URL, bytes.NewBuffer(wq.Data))
	if err != nil {
		return err
	}
	req.SetBasicAuth(h.metadata.Username, h.metadata.Password)

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	if resp != nil && resp.Body != nil {
		resp.Body.Close()
	}
	return nil
}
