package messenger

import (
	"bytes"
	"errors"
	"fmt"
	"io"
)

var UnmarshalError = errors.New("unmarshal error")

type ErrUnmarshalJSON struct {
	Content io.Reader
	Err     error
}

func (u *ErrUnmarshalJSON) Error() string {
	content, err := io.ReadAll(u.Content)
	if err != nil {
		content = []byte("[can not read content]")
	}
	return fmt.Sprintf("can not unmarshal content: %s", string(content))
}

func (u *ErrUnmarshalJSON) Unwrap() error {
	return u.Err
}

func NewErrUnmarshalJSON() *ErrUnmarshalJSON {
	return &ErrUnmarshalJSON{Err: UnmarshalError}
}

func (u *ErrUnmarshalJSON) WithReader(content io.Reader) *ErrUnmarshalJSON {
	u.Content = content
	return u
}

func (u *ErrUnmarshalJSON) WithReaderContent(content []byte) *ErrUnmarshalJSON {
	u.Content = bytes.NewReader(content)
	return u
}

func (u *ErrUnmarshalJSON) WithErr(err error) *ErrUnmarshalJSON {
	u.Err = err
	return u
}
