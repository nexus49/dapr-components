// ------------------------------------------------------------
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.
// ------------------------------------------------------------

package httpoauth

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/dapr/components-contrib/bindings"
	"github.com/dapr/dapr/pkg/logger"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
)

// HTTPSource is a binding for an http url endpoint invocation
// nolint:golint
type HTTPSource struct {
	metadata httpMetadata
	logger   logger.Logger
}

type httpMetadata struct {
	URL          string `json:"url"`
	Method       string `json:"method"`
	ClientID     string `json:"clientID"`
	ClientSecret string `json:"clientSecret"`
	TokenURL     string `json:"tokenURL"`
}

var log = logger.NewLogger("dapr.httpoauth")

// NewHTTP returns a new HTTPSource
func NewHTTPOAuth(logger logger.Logger) *HTTPSource {
	return &HTTPSource{
		logger: logger,
	}
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
	conf := getConfig(h)
	ctx := context.Background()
	client := conf.Client(ctx)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

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
	h.logger.Info("Received new Write call.")
	conf := getConfig(h)
	ctx := context.Background()
	client := conf.Client(ctx)

	req, err := http.NewRequest(h.metadata.Method, h.metadata.URL, bytes.NewBuffer(wq.Data))
	h.logger.Info(fmt.Sprintf("%+v", string(wq.Data)))
	if err != nil {
		return err
	}
	req.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return err
	} else {
		h.logger.Info(fmt.Sprintf("Response: %+v", resp))
	}

	if resp != nil && resp.Body != nil {
		resp.Body.Close()
	}
	return nil
}

func getConfig(h *HTTPSource) clientcredentials.Config {
	return clientcredentials.Config{
		ClientID:     h.metadata.ClientID,
		ClientSecret: h.metadata.ClientSecret,
		AuthStyle:    oauth2.AuthStyleInParams,
		TokenURL:     h.metadata.TokenURL,
	}
}
