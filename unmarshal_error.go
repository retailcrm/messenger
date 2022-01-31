package messenger

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
)

var ErrUnmarshal = errors.New("unmarshal error")

type UnmarshalError struct {
	Content   []byte
	ErrorText string
	Err       error
}

func (u *UnmarshalError) Error() string {
	return fmt.Sprintf("can not unmarshal content: %s; error: %s", string(u.Content), u.ErrorText)
}

func (u *UnmarshalError) Unwrap() error {
	return u.Err
}

func NewUnmarshalError(err error) *UnmarshalError {
	return &UnmarshalError{
		Err:       ErrUnmarshal,
		ErrorText: err.Error(),
	}
}

func (u *UnmarshalError) WithReader(reader io.Reader) *UnmarshalError {
	content, _ := ioutil.ReadAll(reader)
	u.Content = content
	return u
}

func (u *UnmarshalError) WithContent(content []byte) *UnmarshalError {
	u.Content = content
	return u
}

func (u *UnmarshalError) WithErr(err error) *UnmarshalError {
	u.Err = err
	return u
}
