// ------------------------------------------------------------
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.
// ------------------------------------------------------------

package httpoauth

import (
	"testing"

	"github.com/dapr/components-contrib/bindings"
	"github.com/stretchr/testify/assert"
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
