package messenger

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
)

var ErrUnmarshal = errors.New("unmarshal error")

type UnmarshalError struct {
	Content io.Reader
	Err     error
}

func (u *UnmarshalError) Error() string {
	content, err := ioutil.ReadAll(u.Content)
	if err != nil {
		content = []byte("[can not read content]")
	}
	return fmt.Sprintf("can not unmarshal content: %s", string(content))
}

func (u *UnmarshalError) Unwrap() error {
	return u.Err
}

func NewUnmarshalError() *UnmarshalError {
	return &UnmarshalError{Err: ErrUnmarshal}
}

func (u *UnmarshalError) WithReader(content io.Reader) *UnmarshalError {
	u.Content = content
	return u
}

func (u *UnmarshalError) WithReaderContent(content []byte) *UnmarshalError {
	u.Content = bytes.NewReader(content)
	return u
}

func (u *UnmarshalError) WithErr(err error) *UnmarshalError {
	u.Err = err
	return u
}
