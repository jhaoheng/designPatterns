package adapterclass

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_adapter_class(t *testing.T) {
	adapter := NewAdapter()
	adapter.SetData("123")
	output := adapter.ConvertToServiceFormat()
	assert.Equal(t, 123, output)
}
