package messenger

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUnmarshalGenericTemplatePayload(t *testing.T) {
	t.Parallel()

	var payload Payload
	err := json.Unmarshal([]byte(`{
		"generic": {"elements": [
			{
				"title": "Your points!\nMentions: 1",
				"subtitle": "Card description",
				"image_url": "https://example.com/card.jpg",
				"buttons": [
					{"title": "Table", "type": "open_url", "url": "https://example.com/table"},
					{"title": "Points", "subtitle": "Button description", "type": "postback", "payload": "inline_button_id:123"}
				]
			},
			{"title": "Second card"}
		]}
	}`), &payload)
	require.NoError(t, err)
	require.NotNil(t, payload.Generic)
	assert.Equal(t, []GenericTemplateElement{
		{
			Title:    "Your points!\nMentions: 1",
			Subtitle: "Card description",
			ImageURL: "https://example.com/card.jpg",
			Buttons: []Button{
				{Title: "Table", Type: "open_url", URL: "https://example.com/table"},
				{Title: "Points", Subtitle: "Button description", Type: "postback", Payload: "inline_button_id:123"},
			},
		},
		{Title: "Second card"},
	}, payload.Generic.Elements)
}
