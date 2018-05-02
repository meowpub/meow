package jsonld

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type Image struct {
	Type      Type   `json:"@type"`
	MediaType String `json:"https://www.w3.org/ns/activitystreams#mediaType"`
	URL       Ref    `json:"https://www.w3.org/ns/activitystreams#url"`
}

type Person struct {
	Meta
	ID   ID   `json:"@id"`
	Type Type `json:"@type"`

	PreferredUsername String `json:"https://www.w3.org/ns/activitystreams#preferredUsername"`
	Icon              Image  `json:"https://www.w3.org/ns/activitystreams#icon"`
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

	require.NoError(t, v.Meta.Unmarshal(data, &v))
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
	ID   ID   `json:"@id"`
	Type Type `json:"@type"`

	PreferredUsername String `json:"https://www.w3.org/ns/activitystreams#preferredUsername"`
}

func (s *marshalTestStruct) UnmarshalJSON(data []byte) error {
	return s.Meta.Unmarshal(data, s)
}

func (s marshalTestStruct) MarshalJSON() ([]byte, error) {
	return s.Meta.Marshal(s)
}

func TestMarshal(t *testing.T) {
	v := marshalTestStruct{
		ID:                ID("https://example.com/jsmith"),
		Type:              Type{"https://www.w3.org/ns/activitystreams#Person"},
		PreferredUsername: ToString("jsmith"),
	}
	data, err := v.Meta.Marshal(&v)
	require.NoError(t, err)
	assert.Equal(t, `{"@id":"https://example.com/jsmith","@type":["https://www.w3.org/ns/activitystreams#Person"],"https://www.w3.org/ns/activitystreams#preferredUsername":[{"@value":"jsmith"}]}`, string(data))
}
