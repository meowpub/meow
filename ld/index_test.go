package ld

//func TestTypeIsSubClassOf(t *testing.T) {
//	assert.True(t, http_www_w3_org_ns_activitystreams_Note.IsSubClassOf(http_www_w3_org_ns_activitystreams_Object))
//	assert.False(t, http_www_w3_org_ns_activitystreams_Note.IsSubClassOf(http_www_w3_org_ns_activitystreams_Activity))
//}

//func TestQuacksLike(t *testing.T) {
//	obj := NewObject("https://example.com", "http://www.w3.org/ns/activitystreams#Person")
//	assert.True(t, QuacksLike(http_www_w3_org_ns_activitystreams_Object, obj))
//	assert.False(t, QuacksLike(http_www_w3_org_ns_activitystreams_Activity, obj))
//}

//func TestManifest(t *testing.T) {
//	t.Run("Unknown Type", func(t *testing.T) {
//		o, err := ParseObject([]byte(`{
//			"@id": "https://example.com/@jsmith",
//			"@type": ["http://schema.org/Person"],
//			"http://schema.org/name": [{"@value": "John Smith"}]
//		}`))
//		require.NoError(t, err)
//		assert.Empty(t, Manifest(o))
//	})
//	t.Run("One Type", func(t *testing.T) {
//		o, err := ParseObject([]byte(`{
//			"@id": "https://example.com/@jsmith",
//			"@type": ["http://www.w3.org/ns/activitystreams#Person"],
//			"http://www.w3.org/ns/activitystreams#name": [{"@value": "John Smith"}]
//		}`))
//		require.NoError(t, err)
//		assert.True(t, as.IsPerson(o))
//		if entities := Manifest(o); assert.Len(t, entities, 1) {
//			assert.IsType(t, as.Person{}, entities[0])
//		}
//	})
//	t.Run("Two Types", func(t *testing.T) {
//		// A person who is also a place... how mysterious.
//		o, err := ParseObject([]byte(`{
//			"@id": "https://example.com/@jsmith",
//			"@type": [
//				"http://www.w3.org/ns/activitystreams#Person",
//				"http://www.w3.org/ns/activitystreams#Place"
//			],
//			"http://www.w3.org/ns/activitystreams#name": [{"@value": "John Smith"}]
//		}`))
//		require.NoError(t, err)
//		assert.True(t, as.IsPerson(o))
//		assert.True(t, as.IsPlace(o))
//		if entities := Manifest(o); assert.Len(t, entities, 2) {
//			assert.IsType(t, as.Person{}, entities[0])
//			assert.IsType(t, as.Place{}, entities[1])
//		}
//	})
//}
