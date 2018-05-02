package entities

type Entity interface {
	// GetID returns the public ID of the entity, eg. "https://example.com/@jsmith".
	GetID() string
	// GetType returns the type of the entity, eg. ["https://www.w3.org/ns/activitystreams#Person"].
	// An entity may have multiple types, in which case it will behave as all of them at once.
	GetType() []string
}
