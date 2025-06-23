package messenger

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/h2non/gock"
	"net/http"
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

func TestMarshalSendMessage(t *testing.T) {
	t.Parallel()
	sm := SendMessage{
		MessagingType: ResponseType,
		Message:       MessageData{Text: "Hello World"},
		ThreadControl: &ThreadControl{Payload: PassThreadControl},
	}

	data, err := json.Marshal(sm)
	require.NoError(t, err)
	assert.JSONEq(t, `{
	  "messaging_type" : "RESPONSE",
	  "recipient" : { },
	  "message" : {
		"text" : "Hello World"
	  },
	  "thread_control" : {
		"payload" : "pass_thread_control"
	  }
	}`, string(data))
}

//nolint:paralleltest
func TestTextWithReplies(t *testing.T) {
	r := Response{
		token:          "blabla",
		to:             Recipient{ID: 154},
		sendAPIVersion: DefaultSendAPIVersion,
	}

	sm := SendMessage{
		MessagingType: ResponseType,
		Message:       MessageData{Text: "Hello World"},
		ThreadControl: &ThreadControl{Payload: PassThreadControl},
	}

	data, err := json.Marshal(sm)
	require.NoError(t, err)

	defer gock.Off()
	gock.New("https://graph.facebook.com").
		Post(fmt.Sprintf("%s/me/messages", DefaultSendAPIVersion)).
		BodyString(string(data)).
		Reply(http.StatusOK).
		JSON(`{"message_id": "ABCD"}`)

	resp, err := r.DispatchMessage(sm)
	require.NoError(t, err)
	require.Equal(t, "ABCD", resp.MessageID)
}
