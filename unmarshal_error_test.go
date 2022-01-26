package messenger

import (
	"bytes"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUnmarshalError(t *testing.T) {
	err := NewUnmarshalError()
	assert.True(t, errors.Is(err, ErrUnmarshal))
}

func TestUnmarshalError_Error(t *testing.T) {
	content := []byte("test content")
	actual := NewUnmarshalError().WithReaderContent(content).Error()
	expected := "can not unmarshal content: test content"
	assert.Equal(t, expected, actual)
}

func TestUnmarshalError_Unwrap(t *testing.T) {
	actual := NewUnmarshalError().Unwrap()
	expected := ErrUnmarshal
	assert.Equal(t, expected, actual)
}

func TestUnmarshalError_WithReaderContent(t *testing.T) {
	content := []byte("test content")
	reader := bytes.NewReader(content)

	actual := NewUnmarshalError().WithReaderContent(content)
	expected := &UnmarshalError{Err: ErrUnmarshal, Content: reader}
	assert.Equal(t, expected, actual)
}

func TestUnmarshalError_WithReader(t *testing.T) {
	content := []byte("test content")
	reader := bytes.NewReader(content)

	actual := NewUnmarshalError().WithReader(reader)
	expected := &UnmarshalError{Err: ErrUnmarshal, Content: reader}
	assert.Equal(t, expected, actual)
}

func TestUnmarshalError_WithErr(t *testing.T) {
	err := errors.New("some error")
	actual := NewUnmarshalError().WithErr(err)
	expected := &UnmarshalError{Err: err}
	assert.Equal(t, expected, actual)
}
