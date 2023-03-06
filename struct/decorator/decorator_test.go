package decorator

import (
	"testing"

	"github.com/longbridgeapp/assert"
)

func TestColorSquare_Draw(t *testing.T) {
	sq := Square{}
	csq := NewColorSqure(sq, "red")
	got := csq.Draw()
	assert.Equal(t, "this is a square, color is red", got)
}
