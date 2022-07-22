package proxy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_proxy(t *testing.T) {
	p := NewProxy()
	p.SetService(&Service{})
	assert.NoError(t, p.CheckAccess())
	p.Operation()
}
