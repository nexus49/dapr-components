package oauth2ClientCredentials

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/dapr/components-contrib/middleware"
	"github.com/fasthttp-contrib/sessions"
	"github.com/valyala/fasthttp"
)

// Metadata is the oAuth middleware config
type oAuth2MiddlewareMetadata struct {
	ClientID       string `json:"clientID"`
	ClientSecret   string `json:"clientSecret"`
	TokenURL       string `json:"tokenURL"`
	AuthHeaderName string `json:"authHeaderName"`
}

type token struct {
	AccessToken string `json:"access_token"`
	Type        string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
	Scope       string `json:"scope"`
}

// NewOAuth2Middleware returns a new oAuth2 middleware
func NewOAuth2ClientCredentialsMiddleware() *Middleware {
	return &Middleware{}
}

// Middleware is an oAuth2 authentication middleware
type Middleware struct {
}

const (
	stateParam   = "state"
	savedState   = "auth-state"
	redirectPath = "redirect-url"
	codeParam    = "code"
)

// GetHandler retruns the HTTP handler provided by the middleware
func (m *Middleware) GetHandler(metadata middleware.Metadata) (func(h fasthttp.RequestHandler) fasthttp.RequestHandler, error) {
	meta, err := m.getNativeMetadata(metadata)
	if err != nil {
		return nil, err
	}

	return func(h fasthttp.RequestHandler) fasthttp.RequestHandler {
		return func(ctx *fasthttp.RequestCtx) {

			session := sessions.StartFasthttp(ctx)
			if session.GetString(meta.AuthHeaderName) != "" {
				ctx.Request.Header.Add(meta.AuthHeaderName, session.GetString(meta.AuthHeaderName))
				h(ctx)
				return
			}

			client := http.Client{Timeout: time.Second * 5}
			req, err := http.NewRequest(http.MethodPost, meta.TokenURL, nil)
			if err != nil {
				ctx.Error("invalid state", fasthttp.StatusBadRequest)
			}
			req.SetBasicAuth(meta.ClientID, meta.ClientSecret)

			resp, err := client.Do(req)
			if err != nil {
				ctx.Error("invalid state", fasthttp.StatusBadRequest)
			}

			if resp != nil && resp.Body != nil {
				defer resp.Body.Close()
				body, err := ioutil.ReadAll(resp.Body)
				var token token
				err = json.Unmarshal(body, &token)
				if err == nil {
					session.Set(meta.AuthHeaderName, token.Type+" "+token.AccessToken)
					ctx.Request.Header.Add(meta.AuthHeaderName, token.Type+" "+token.AccessToken)
					h(ctx)
				} else {
					ctx.Error("invalid state", fasthttp.StatusBadRequest)
				}
			} else {
				ctx.Error("invalid state", fasthttp.StatusBadRequest)
			}
		}
	}, nil
}

func (m *Middleware) getNativeMetadata(metadata middleware.Metadata) (*oAuth2MiddlewareMetadata, error) {
	b, err := json.Marshal(metadata.Properties)
	if err != nil {
		return nil, err
	}

	var middlewareMetadata oAuth2MiddlewareMetadata
	err = json.Unmarshal(b, &middlewareMetadata)
	if err != nil {
		return nil, err
	}
	return &middlewareMetadata, nil
}
