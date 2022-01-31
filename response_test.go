package messenger

import (
	"bytes"
	"encoding/json"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_MarshalStructuredMessageElement(t *testing.T) {
	data, err := json.Marshal(StructuredMessageElement{
		Title: "Title",
	})
	require.NoError(t, err)
	assert.JSONEq(t, string(data), `{"image_url":"", "subtitle":"", "title": "Title"}`)
}

func TestResponse_checkFacebookError_UnmarshalError(t *testing.T) {
	r := bytes.NewReader([]byte("test error text"))
	err := checkFacebookError(r)
	assert.True(t, errors.Is(err, ErrUnmarshal))
	assert.Contains(t, err.Error(), "test error text")
}

func TestResponse_getFacebookQueryResponse_UnmarshalError(t *testing.T) {
	r := bytes.NewReader([]byte("test error text"))
	_, err := getFacebookQueryResponse(r)
	assert.True(t, errors.Is(err, ErrUnmarshal))
	assert.Contains(t, err.Error(), "test error text")
}
