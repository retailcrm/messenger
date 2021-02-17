package messenger

import (
	"encoding/json"
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
