// GENERATED FILE, DO NOT EDIT.
// Please refer to: tools/nsgen/templates.go
package sec

import (
	"github.com/meowpub/meow/ld"
)

type Digest struct{ o *ld.Object }

// Ducktypes the object into a(n) Digest.
func AsDigest(e ld.Entity) Digest { return Digest{o: e.Obj()} }

// Does the object quack like a(n) Digest?
func IsDigest(e ld.Entity) bool { return ld.Is(e, TypeDigest) }

// Returns the wrapped plain ld.Object. Implements ld.Entity.
func (obj Digest) Obj() *ld.Object { return obj.o }

// Returns the object's @id. Implements ld.Entity.
func (obj Digest) ID() string { return obj.o.ID() }

// Returns the object's @value. Implements ld.Entity.
func (obj Digest) Value() string { return obj.o.Value() }

// Returns the object's @type. Implements ld.Entity.
func (obj Digest) Type() []string { return obj.o.Type() }

// Returns the named attribute. Implements ld.Entity.
func (obj Digest) Get(key string) interface{} { return obj.o.Get(key) }

// Applies another object as a patch to this one. Implements ld.Entity.
func (obj Digest) Apply(other ld.Entity, mergeArrays bool) error {
	return obj.o.Apply(other, mergeArrays)
}

type EncryptedMessage struct{ o *ld.Object }

// Ducktypes the object into a(n) EncryptedMessage.
func AsEncryptedMessage(e ld.Entity) EncryptedMessage { return EncryptedMessage{o: e.Obj()} }

// Does the object quack like a(n) EncryptedMessage?
func IsEncryptedMessage(e ld.Entity) bool { return ld.Is(e, TypeEncryptedMessage) }

// Returns the wrapped plain ld.Object. Implements ld.Entity.
func (obj EncryptedMessage) Obj() *ld.Object { return obj.o }

// Returns the object's @id. Implements ld.Entity.
func (obj EncryptedMessage) ID() string { return obj.o.ID() }

// Returns the object's @value. Implements ld.Entity.
func (obj EncryptedMessage) Value() string { return obj.o.Value() }

// Returns the object's @type. Implements ld.Entity.
func (obj EncryptedMessage) Type() []string { return obj.o.Type() }

// Returns the named attribute. Implements ld.Entity.
func (obj EncryptedMessage) Get(key string) interface{} { return obj.o.Get(key) }

// Applies another object as a patch to this one. Implements ld.Entity.
func (obj EncryptedMessage) Apply(other ld.Entity, mergeArrays bool) error {
	return obj.o.Apply(other, mergeArrays)
}

type GraphSignature2012 struct{ o *ld.Object }

// Ducktypes the object into a(n) GraphSignature2012.
func AsGraphSignature2012(e ld.Entity) GraphSignature2012 { return GraphSignature2012{o: e.Obj()} }

// Does the object quack like a(n) GraphSignature2012?
func IsGraphSignature2012(e ld.Entity) bool { return ld.Is(e, TypeGraphSignature2012) }

// Returns the wrapped plain ld.Object. Implements ld.Entity.
func (obj GraphSignature2012) Obj() *ld.Object { return obj.o }

// Returns the object's @id. Implements ld.Entity.
func (obj GraphSignature2012) ID() string { return obj.o.ID() }

// Returns the object's @value. Implements ld.Entity.
func (obj GraphSignature2012) Value() string { return obj.o.Value() }

// Returns the object's @type. Implements ld.Entity.
func (obj GraphSignature2012) Type() []string { return obj.o.Type() }

// Returns the named attribute. Implements ld.Entity.
func (obj GraphSignature2012) Get(key string) interface{} { return obj.o.Get(key) }

// Applies another object as a patch to this one. Implements ld.Entity.
func (obj GraphSignature2012) Apply(other ld.Entity, mergeArrays bool) error {
	return obj.o.Apply(other, mergeArrays)
}

type Key struct{ o *ld.Object }

// Ducktypes the object into a(n) Key.
func AsKey(e ld.Entity) Key { return Key{o: e.Obj()} }

// Does the object quack like a(n) Key?
func IsKey(e ld.Entity) bool { return ld.Is(e, TypeKey) }

// Returns the wrapped plain ld.Object. Implements ld.Entity.
func (obj Key) Obj() *ld.Object { return obj.o }

// Returns the object's @id. Implements ld.Entity.
func (obj Key) ID() string { return obj.o.ID() }

// Returns the object's @value. Implements ld.Entity.
func (obj Key) Value() string { return obj.o.Value() }

// Returns the object's @type. Implements ld.Entity.
func (obj Key) Type() []string { return obj.o.Type() }

// Returns the named attribute. Implements ld.Entity.
func (obj Key) Get(key string) interface{} { return obj.o.Get(key) }

// Applies another object as a patch to this one. Implements ld.Entity.
func (obj Key) Apply(other ld.Entity, mergeArrays bool) error { return obj.o.Apply(other, mergeArrays) }

type LinkedDataSignature2015 struct{ o *ld.Object }

// Ducktypes the object into a(n) LinkedDataSignature2015.
func AsLinkedDataSignature2015(e ld.Entity) LinkedDataSignature2015 {
	return LinkedDataSignature2015{o: e.Obj()}
}

// Does the object quack like a(n) LinkedDataSignature2015?
func IsLinkedDataSignature2015(e ld.Entity) bool { return ld.Is(e, TypeLinkedDataSignature2015) }

// Returns the wrapped plain ld.Object. Implements ld.Entity.
func (obj LinkedDataSignature2015) Obj() *ld.Object { return obj.o }

// Returns the object's @id. Implements ld.Entity.
func (obj LinkedDataSignature2015) ID() string { return obj.o.ID() }

// Returns the object's @value. Implements ld.Entity.
func (obj LinkedDataSignature2015) Value() string { return obj.o.Value() }

// Returns the object's @type. Implements ld.Entity.
func (obj LinkedDataSignature2015) Type() []string { return obj.o.Type() }

// Returns the named attribute. Implements ld.Entity.
func (obj LinkedDataSignature2015) Get(key string) interface{} { return obj.o.Get(key) }

// Applies another object as a patch to this one. Implements ld.Entity.
func (obj LinkedDataSignature2015) Apply(other ld.Entity, mergeArrays bool) error {
	return obj.o.Apply(other, mergeArrays)
}

type Signature struct{ o *ld.Object }

// Ducktypes the object into a(n) Signature.
func AsSignature(e ld.Entity) Signature { return Signature{o: e.Obj()} }

// Does the object quack like a(n) Signature?
func IsSignature(e ld.Entity) bool { return ld.Is(e, TypeSignature) }

// Returns the wrapped plain ld.Object. Implements ld.Entity.
func (obj Signature) Obj() *ld.Object { return obj.o }

// Returns the object's @id. Implements ld.Entity.
func (obj Signature) ID() string { return obj.o.ID() }

// Returns the object's @value. Implements ld.Entity.
func (obj Signature) Value() string { return obj.o.Value() }

// Returns the object's @type. Implements ld.Entity.
func (obj Signature) Type() []string { return obj.o.Type() }

// Returns the named attribute. Implements ld.Entity.
func (obj Signature) Get(key string) interface{} { return obj.o.Get(key) }

// Applies another object as a patch to this one. Implements ld.Entity.
func (obj Signature) Apply(other ld.Entity, mergeArrays bool) error {
	return obj.o.Apply(other, mergeArrays)
}

var (
	_ ld.Entity = Digest{}
	_ ld.Entity = EncryptedMessage{}
	_ ld.Entity = GraphSignature2012{}
	_ ld.Entity = Key{}
	_ ld.Entity = LinkedDataSignature2015{}
	_ ld.Entity = Signature{}
)
