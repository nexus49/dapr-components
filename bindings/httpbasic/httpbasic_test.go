// ------------------------------------------------------------
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.
// ------------------------------------------------------------

package httpbasic

import (
	"testing"

	"github.com/nexus49/components-contrib/bindings"
	"github.com/stretchr/testify/assert"
)

func TestInit(t *testing.T) {
	m := bindings.Metadata{}
	m.Properties = map[string]string{"url": "a", "method": "a", "username": "c", "password": "d"}
	hs := HTTPSource{}
	err := hs.Init(m)
	assert.Nil(t, err)
	assert.Equal(t, "a", hs.metadata.URL)
	assert.Equal(t, "a", hs.metadata.Method)
	assert.Equal(t, "c", hs.metadata.Username)
	assert.Equal(t, "d", hs.metadata.Password)
}
