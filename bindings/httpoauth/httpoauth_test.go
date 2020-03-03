// ------------------------------------------------------------
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.
// ------------------------------------------------------------

package httpoauth

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/dapr/components-contrib/bindings"
	"github.com/stretchr/testify/assert"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
)

func TestInit(t *testing.T) {
	m := bindings.Metadata{}
	m.Properties = map[string]string{"url": "a", "method": "a", "clientID": "c", "clientSecret": "d", "tokenURL": "e"}
	hs := HTTPSource{}
	err := hs.Init(m)
	assert.Nil(t, err)
	assert.Equal(t, "a", hs.metadata.URL)
	assert.Equal(t, "a", hs.metadata.Method)
	assert.Equal(t, "c", hs.metadata.ClientID)
	assert.Equal(t, "d", hs.metadata.ClientSecret)
	assert.Equal(t, "e", hs.metadata.TokenURL)
}

func TestToken(t *testing.T) {

	conf := clientcredentials.Config{
		ClientID:     "sb-a251f995-3572-47fd-bd4a-1c2b0486a6b4!b10804|one-mds-master!b9046",
		ClientSecret: "HZcfOrP3lcPkZbZfOGbhVRLNtJ0=",
		AuthStyle:    oauth2.AuthStyleInParams,
		TokenURL:     "https://kerneltest.authentication.sap.hana.ondemand.com/oauth/token",
	}
	ctx := context.Background()
	client := conf.Client(ctx)

	data := `{ "test":"cde"}`
	req, err := http.NewRequest("POST", "https://entyijlef0jes.x.pipedream.net/", bytes.NewBuffer([]byte(data)))
	if err != nil {
		println(fmt.Sprintf("Error! %+v", err))
	}

	resp, err := client.Do(req)
	if err != nil {
		println(fmt.Sprintf("Error! %+v", err))
	} else {
		println("Success")
	}

	if resp != nil && resp.Body != nil {
		resp.Body.Close()
	}
}

// func TestWrite(t *testing.T) {
// 	m := bindings.Metadata{}
// 	m.Properties = map[string]string{"url": "https://google.com", "method": "a", "username": "c", "password": "d"}
// 	hs := HTTPSource{}
// 	err := hs.Init(m)
// 	assert.Nil(t, err)

// 	data := []byte("{ \"key\": \"value\" }")

// 	err = hs.Write(&bindings.WriteRequest{Data: data})
// 	assert.Nil(t, err)
// }

// func TestRead(t *testing.T) {
// 	m := bindings.Metadata{}
// 	m.Properties = map[string]string{"url": "https://google.com", "method": "a", "username": "c", "password": "d"}
// 	hs := HTTPSource{}
// 	err := hs.Init(m)
// 	assert.Nil(t, err)

// 	err = hs.Read(func(rr *bindings.ReadResponse) error {
// 		assert.True(t, len(rr.Data) > 0)
// 		return nil
// 	})
// 	assert.Nil(t, err)
// }
