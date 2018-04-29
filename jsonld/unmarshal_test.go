package jsonld

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUnmarshalPerson(t *testing.T) {
	data := []byte(`{
		"@id": "https://example.com/@jsmith",
		"@type": ["https://www.w3.org/ns/activitystreams#Person"],
		"https://www.w3.org/ns/activitystreams#name": [{"@value": "John Smith"}],
		"https://www.w3.org/ns/activitystreams#preferredUsername": [{"@value": "jsmith"}],
		"https://www.w3.org/ns/activitystreams#icon": {
			"@type": ["https://www.w3.org/ns/activitystreams#Image"],
			"https://www.w3.org/ns/activitystreams#mediaType": [{"@value": "image/jpeg"}],
			"https://www.w3.org/ns/activitystreams#url": [{"@id": "https://example.com/icon.jpg"}]
		}
	}`)
	var v struct {
		Meta
		ID   ID   `json:"@id"`
		Type Type `json:"@type"`

		PreferredUsername String `json:"https://www.w3.org/ns/activitystreams#preferredUsername"`
		Icon              struct {
			Type      Type   `json:"@type"`
			MediaType String `json:"https://www.w3.org/ns/activitystreams#mediaType"`
			URL       Ref    `json:"https://www.w3.org/ns/activitystreams#url"`
		} `json:"https://www.w3.org/ns/activitystreams#icon"`
	}
	require.NoError(t, Unmarshal(data, &v))
	assert.Equal(t, ID("https://example.com/@jsmith"), v.ID)
	assert.Equal(t, Type{"https://www.w3.org/ns/activitystreams#Person"}, v.Type)
	assert.Equal(t, ToString("jsmith"), v.PreferredUsername)
	{
		assert.Equal(t, Type{"https://www.w3.org/ns/activitystreams#Image"}, v.Icon.Type)
		assert.Equal(t, ToString("image/jpeg"), v.Icon.MediaType)
		assert.Equal(t, ToRef(ID("https://example.com/icon.jpg"), nil), v.Icon.URL)
	}
}
