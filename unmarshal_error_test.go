package messenger

import (
	"bytes"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewErrUnmarshalJSON(t *testing.T) {
	err := NewErrUnmarshalJSON()
	assert.True(t, errors.Is(err, UnmarshalError))
}

func TestUnmarshalError_Error(t *testing.T) {
	content := []byte("test content")
	actual := NewErrUnmarshalJSON().WithReaderContent(content).Error()
	expected := "can not unmarshal content: test content"
	assert.Equal(t, expected, actual)
}

func TestErrUnmarshalJSON_Unwrap(t *testing.T) {
	actual := NewErrUnmarshalJSON().Unwrap()
	expected := UnmarshalError
	assert.Equal(t, expected, actual)
}

func TestUnmarshalError_WithReaderContent(t *testing.T) {
	content := []byte("test content")
	reader := bytes.NewReader(content)

	actual := NewErrUnmarshalJSON().WithReaderContent(content)
	expected := &ErrUnmarshalJSON{Err: UnmarshalError, Content: reader}
	assert.Equal(t, expected, actual)
}

func TestUnmarshalError_WithReader(t *testing.T) {
	content := []byte("test content")
	reader := bytes.NewReader(content)

	actual := NewErrUnmarshalJSON().WithReader(reader)
	expected := &ErrUnmarshalJSON{Err: UnmarshalError, Content: reader}
	assert.Equal(t, expected, actual)
}

func TestUnmarshalError_WithErr(t *testing.T) {
	err := errors.New("some error")
	actual := NewErrUnmarshalJSON().WithErr(err)
	expected := &ErrUnmarshalJSON{Err: err}
	assert.Equal(t, expected, actual)
}
