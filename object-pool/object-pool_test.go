package objectpool

import "testing"

func Test_(t *testing.T) {
	p := NewObject(2)

	select {
	case obj := <-p:
		obj.Do()
		p <- obj
	default:
	}
}
