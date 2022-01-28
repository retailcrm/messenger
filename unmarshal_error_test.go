package messenger

import (
	"bytes"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUnmarshalError(t *testing.T) {
	err := errors.New("some error")
	unmarshalError := NewUnmarshalError(err)
	assert.True(t, errors.Is(unmarshalError, ErrUnmarshal))
}

func TestUnmarshalError_Error(t *testing.T) {
	err := errors.New("some error")
	content := []byte("test content")
	actual := NewUnmarshalError(err).WithContent(content).Error()
	expected := "can not unmarshal content: test content; error: some error"
	assert.Equal(t, expected, actual)
}

func TestUnmarshalError_Unwrap(t *testing.T) {
	err := errors.New("some error")
	actual := NewUnmarshalError(err).Unwrap()
	expected := ErrUnmarshal
	assert.Equal(t, expected, actual)
}

func TestUnmarshalError_WithContent(t *testing.T) {
	err := errors.New("some error")
	content := []byte("test content")

	actual := NewUnmarshalError(err).WithContent(content)
	expected := &UnmarshalError{Err: ErrUnmarshal, Content: content, ErrorText: err.Error()}
	assert.Equal(t, expected, actual)
}

func TestUnmarshalError_WithReader(t *testing.T) {
	err := errors.New("some error")
	content := []byte("test content")
	reader := bytes.NewReader(content)

	actual := NewUnmarshalError(err).WithReader(reader)
	expected := &UnmarshalError{Err: ErrUnmarshal, Content: content, ErrorText: err.Error()}
	assert.Equal(t, expected, actual)
}

func TestUnmarshalError_WithErr(t *testing.T) {
	someError := errors.New("some error")
	otherError := errors.New("other error")
	actual := NewUnmarshalError(someError).WithErr(otherError)
	expected := &UnmarshalError{Err: otherError, ErrorText: someError.Error()}
	assert.Equal(t, expected, actual)
}
