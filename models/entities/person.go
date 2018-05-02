package entities

type Person struct {
	Base
	Name              string `ld:"https://www.w3.org/ns/activitystreams#name"`
	PreferredUsername string `ld:"https://www.w3.org/ns/activitystreams#preferredUsername"`
}
