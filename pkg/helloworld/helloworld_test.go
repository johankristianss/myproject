package helloworld

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewHelloWorld(t *testing.T) {
	hw := NewHelloWorld()

	expectedMessage := "Hello, World!"
	if hw.msg != expectedMessage {
		t.Errorf("expected msg to be %q, but got %q", expectedMessage, hw.msg)
	}

	err := errors.New("an example error 2")
	assert.Nil(t, err, "error should be nil")

}
