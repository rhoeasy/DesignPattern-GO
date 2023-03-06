package bridge

import (
	"testing"

	"github.com/longbridgeapp/assert"
)

func TestErrorNOtification_Notify(t *testing.T) {
	sender := NewEmailMsgSender([]string{"test@test.com"})
	n := NewErrorNotification(sender)
	err := n.Notify("test msg")

	assert.Nil(t, err)
}
