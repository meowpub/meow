// GENERATED FILE, DO NOT EDIT.
// Please refer to: tools/nsgen/templates.go
package sec

import (
	"github.com/meowpub/meow/ld"
	"github.com/meowpub/meow/ld/ns/owl"
)

// This class represents a message digest that may be used for data integrity verification.
// The digest algorithm used will determine the cryptographic properties of the digest.
type Digest struct{ owl.Thing }

// Ducktypes the object into a(n) Digest.
func AsDigest(e ld.Entity) Digest { return Digest{owl.AsThing(e)} }

// Does the object quack like a(n) Digest?
func IsDigest(e ld.Entity) bool { return ld.Is(e, TypeDigest) }

// The digest algorithm is used to specify the cryptographic function to use when generating the
// data to be digitally signed. Typically, data that is to be signed goes through three steps:
// 1) canonicalization, 2) digest, and 3) signature. This property is used to specify the algorithm
// that should be used for step #2. A signature class typically specifies a default digest method,
// so this property is typically used to specify information for a signature algorithm.
func (obj Digest) DigestAlgorithm() interface{} { return obj.Get(PropDigestAlgorithm) }

// The digest value is used to express the output of the digest algorithm expressed in Base-1
// (hexadecimal) format.
func (obj Digest) DigestValue() interface{} { return obj.Get(PropDigestValue) }

// A class of messages that are obfuscated in some cryptographic manner.
// These messages are incredibly difficult to decrypt without the proper decryption key.
type EncryptedMessage struct{ owl.Thing }

// Ducktypes the object into a(n) EncryptedMessage.
func AsEncryptedMessage(e ld.Entity) EncryptedMessage { return EncryptedMessage{owl.AsThing(e)} }

// Does the object quack like a(n) EncryptedMessage?
func IsEncryptedMessage(e ld.Entity) bool { return ld.Is(e, TypeEncryptedMessage) }

func (obj EncryptedMessage) AuthenticationTag() interface{} { return obj.Get(PropAuthenticationTag) }

// The cipher algorithm describes the mechanism used to encrypt a message. It is typically a string
// expressing the cipher suite, the strength of the cipher, and a block cipher mode.
func (obj EncryptedMessage) CipherAlgorithm() interface{} { return obj.Get(PropCipherAlgorithm) }

// Cipher data an opaque blob of information that is used to specify an encrypted message.
func (obj EncryptedMessage) CipherData() interface{} { return obj.Get(PropCipherData) }

// A cipher key is a symmetric key that is used to encrypt or decrypt a piece of information.
// The key itself may be expressed in clear text or encrypted.
func (obj EncryptedMessage) CipherKey() interface{} { return obj.Get(PropCipherKey) }

// The initialization vector (IV) is a byte stream that is typically used to initialize certain block
// cipher encryption schemes. For a receiving application to be able to decrypt a message, it must
// know the decryption key and the initialization vector. The value is typically base-64 encoded.
func (obj EncryptedMessage) InitializationVector() interface{} {
	return obj.Get(PropInitializationVector)
}

// A public key property is used to specify a URL that contains information about a public key.
func (obj EncryptedMessage) PublicKey() interface{} { return obj.Get(PropPublicKey) }

// A graph signature is used for digital signatures on RDF graphs. The default canonicalization
// mechanism is specified in the RDF Graph normalization specification, which effectively
// deterministically names all unnamed nodes. The default signature mechanism uses a SHA-256 digest
// and RSA to perform the digital signature.
type GraphSignature2012 struct{ Signature }

// Ducktypes the object into a(n) GraphSignature2012.
func AsGraphSignature2012(e ld.Entity) GraphSignature2012 { return GraphSignature2012{AsSignature(e)} }

// Does the object quack like a(n) GraphSignature2012?
func IsGraphSignature2012(e ld.Entity) bool { return ld.Is(e, TypeGraphSignature2012) }

// This class represents a cryptographic key that may be used for encryption, decryption, or
// digitally signing data.
type Key struct{ owl.Thing }

// Ducktypes the object into a(n) Key.
func AsKey(e ld.Entity) Key { return Key{owl.AsThing(e)} }

// Does the object quack like a(n) Key?
func IsKey(e ld.Entity) bool { return ld.Is(e, TypeKey) }

// An owner is an entity that claims control over a particular resource. Note that ownership is best
// validated as a two-way relationship where the owner claims ownership over a particular resource,
// and the resource clearly identifies its owner.
func (obj Key) Owner() interface{} { return obj.Get(PropOwner) }

// A private key PEM property is used to specify the PEM-encoded version of the private key.
// This encoding is compatible with almost every Secure Sockets Layer library implementation and
// typically plugs directly into functions intializing private keys.
func (obj Key) PrivateKeyPem() interface{} { return obj.Get(PropPrivateKeyPem) }

// A public key PEM property is used to specify the PEM-encoded version of the public key.
// This encoding is compatible with almost every Secure Sockets Layer library implementation and
// typically plugs directly into functions intializing public keys.
func (obj Key) PublicKeyPem() interface{} { return obj.Get(PropPublicKeyPem) }

// A Linked Data signature is used for digital signatures on RDF Datasets. The default canonicalization
// mechanism is specified in the RDF Dataset Normalization specification, which effectively
// deterministically names all unnamed nodes. The default signature mechanism uses a SHA-256 digest and
// RSA to perform the digital signature. This signature uses a algorithm for producing the data that it
// signs and verifies that is different from other Linked Data signatures.
type LinkedDataSignature2015 struct{ Signature }

// Ducktypes the object into a(n) LinkedDataSignature2015.
func AsLinkedDataSignature2015(e ld.Entity) LinkedDataSignature2015 {
	return LinkedDataSignature2015{AsSignature(e)}
}

// Does the object quack like a(n) LinkedDataSignature2015?
func IsLinkedDataSignature2015(e ld.Entity) bool { return ld.Is(e, TypeLinkedDataSignature2015) }

// A Linked Data signature is used for digital signatures on RDF Datasets. The default canonicalization
// mechanism is specified in the RDF Dataset Normalization specification, which effectively
// deterministically names all unnamed nodes. The default signature mechanism uses a SHA-256 digest and
// RSA to perform the digital signature.
type LinkedDataSignature2016 struct{ Signature }

// Ducktypes the object into a(n) LinkedDataSignature2016.
func AsLinkedDataSignature2016(e ld.Entity) LinkedDataSignature2016 {
	return LinkedDataSignature2016{AsSignature(e)}
}

// Does the object quack like a(n) LinkedDataSignature2016?
func IsLinkedDataSignature2016(e ld.Entity) bool { return ld.Is(e, TypeLinkedDataSignature2016) }

// This class represents a digital signature on serialized data. It is an abstract class and should
// not be used other than for Semantic Web reasoning purposes, such as by a reasoning agent.
type Signature struct{ owl.Thing }

// Ducktypes the object into a(n) Signature.
func AsSignature(e ld.Entity) Signature { return Signature{owl.AsThing(e)} }

// Does the object quack like a(n) Signature?
func IsSignature(e ld.Entity) bool { return ld.Is(e, TypeSignature) }

func (obj Signature) Creator() interface{} { return obj.Get(PropCreator) }

// The signature value is used to express the output of the signature algorithm expressed in
// base-64 format.
func (obj Signature) SignatureValue() interface{} { return obj.Get(PropSignatureValue) }

var (
	_ ld.Entity = Digest{}
	_ ld.Entity = EncryptedMessage{}
	_ ld.Entity = GraphSignature2012{}
	_ ld.Entity = Key{}
	_ ld.Entity = LinkedDataSignature2015{}
	_ ld.Entity = LinkedDataSignature2016{}
	_ ld.Entity = Signature{}
)
