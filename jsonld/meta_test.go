package jsonld

import (
	"encoding/json"

	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type Image struct {
	Meta
	Type      Type   `json:"@type"`
	MediaType String `json:"https://www.w3.org/ns/activitystreams#mediaType"`
	URL       Ref    `json:"https://www.w3.org/ns/activitystreams#url"`
}

type Person struct {
	Meta
	ID   ID   `json:"@id"`
	Type Type `json:"@type"`

	Name              LangString `json:"https://www.w3.org/ns/activitystreams#name"`
	PreferredUsername String     `json:"https://www.w3.org/ns/activitystreams#preferredUsername"`
	Icon              Image      `json:"https://www.w3.org/ns/activitystreams#icon"`
}

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
	var v Person
	var m map[string]interface{}

	require.NoError(t, json.Unmarshal(data, &m), "JSON unmarshal")
	require.NoError(t, Unmarshal(m, &v), "Unmarshal")

	assert.Equal(t, ID("https://example.com/@jsmith"), v.ID)
	assert.Equal(t, Type{"https://www.w3.org/ns/activitystreams#Person"}, v.Type)
	assert.Equal(t, ToString("jsmith"), v.PreferredUsername)
	{
		assert.Equal(t, Type{"https://www.w3.org/ns/activitystreams#Image"}, v.Icon.Type)
		assert.Equal(t, ToString("image/jpeg"), v.Icon.MediaType)
		assert.Equal(t, ToRef(ID("https://example.com/icon.jpg"), nil), v.Icon.URL)
	}
}

type marshalTestStruct struct {
	Meta
	ID   ID         `json:"@id"`
	Type Type       `json:"@type"`
	Name LangString `json:"https://www.w3.org/ns/activitystreams#name,omitempty"`

	PreferredUsername String `json:"https://www.w3.org/ns/activitystreams#preferredUsername,omitempty"`
}

func TestMarshal(t *testing.T) {
	v := marshalTestStruct{
		ID:                ID("https://example.com/jsmith"),
		Type:              Type{"https://www.w3.org/ns/activitystreams#Person"},
		PreferredUsername: ToString("jsmith"),
	}

	map_, err := Marshal(&v)
	require.NoError(t, err, "Marshal")

	buf, err := json.Marshal(map_)
	require.NoError(t, err, "json.Marshal")

	assert.Equal(t, `{"@id":"https://example.com/jsmith","@type":["https://www.w3.org/ns/activitystreams#Person"],"https://www.w3.org/ns/activitystreams#preferredUsername":[{"@value":"jsmith"}]}`, string(buf))
}
