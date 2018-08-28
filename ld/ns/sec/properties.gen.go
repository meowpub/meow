// GENERATED FILE, DO NOT EDIT.
// Please refer to: tools/nsgen/templates.go
package sec

import (
	"github.com/meowpub/meow/ld"
)

func GetAuthenticationTag(e ld.Entity) interface{} { return e.Get(PropAuthenticationTag) }

func SetAuthenticationTag(e ld.Entity, v interface{}) { e.Set(PropAuthenticationTag, v) }

// The canonicalization algorithm is used to transform the input data into a form that can be passed
// to a cryptographic digest method. The digest is then digitally signed using a digital signature
// algorithm. Canonicalization ensures that a piece of software that is generating a digital signature
// is able to do so on the same set of information in a deterministic manner.
func GetCanonicalizationAlgorithm(e ld.Entity) interface{} {
	return e.Get(PropCanonicalizationAlgorithm)
}

func SetCanonicalizationAlgorithm(e ld.Entity, v interface{}) { e.Set(PropCanonicalizationAlgorithm, v) }

// The cipher algorithm describes the mechanism used to encrypt a message. It is typically a string
// expressing the cipher suite, the strength of the cipher, and a block cipher mode.
func GetCipherAlgorithm(e ld.Entity) interface{} { return e.Get(PropCipherAlgorithm) }

func SetCipherAlgorithm(e ld.Entity, v interface{}) { e.Set(PropCipherAlgorithm, v) }

// Cipher data an opaque blob of information that is used to specify an encrypted message.
func GetCipherData(e ld.Entity) interface{} { return e.Get(PropCipherData) }

func SetCipherData(e ld.Entity, v interface{}) { e.Set(PropCipherData, v) }

// A cipher key is a symmetric key that is used to encrypt or decrypt a piece of information.
// The key itself may be expressed in clear text or encrypted.
func GetCipherKey(e ld.Entity) interface{} { return e.Get(PropCipherKey) }

func SetCipherKey(e ld.Entity, v interface{}) { e.Set(PropCipherKey, v) }

func GetCreator(e ld.Entity) interface{} { return e.Get(PropCreator) }

func SetCreator(e ld.Entity, v interface{}) { e.Set(PropCreator, v) }

// The digest algorithm is used to specify the cryptographic function to use when generating the
// data to be digitally signed. Typically, data that is to be signed goes through three steps:
// 1) canonicalization, 2) digest, and 3) signature. This property is used to specify the algorithm
// that should be used for step #2. A signature class typically specifies a default digest method,
// so this property is typically used to specify information for a signature algorithm.
func GetDigestAlgorithm(e ld.Entity) interface{} { return e.Get(PropDigestAlgorithm) }

func SetDigestAlgorithm(e ld.Entity, v interface{}) { e.Set(PropDigestAlgorithm, v) }

// The digest value is used to express the output of the digest algorithm expressed in Base-1
// (hexadecimal) format.
func GetDigestValue(e ld.Entity) interface{} { return e.Get(PropDigestValue) }

func SetDigestValue(e ld.Entity, v interface{}) { e.Set(PropDigestValue, v) }

// The expiration time is typically associated with a Key and specifies when the validity of the key
// will expire. It is considered a best practice to only create keys that have very definite
// expiration periods. This period is typically set to between six months and two years. An digital
// signature created using an expired key MUST be marked as invalid by any software attempting to
// verify the signature.
func GetExpires(e ld.Entity) interface{} { return e.Get(PropExpires) }

func SetExpires(e ld.Entity, v interface{}) { e.Set(PropExpires, v) }

// The initialization vector (IV) is a byte stream that is typically used to initialize certain block
// cipher encryption schemes. For a receiving application to be able to decrypt a message, it must
// know the decryption key and the initialization vector. The value is typically base-64 encoded.
func GetInitializationVector(e ld.Entity) interface{} { return e.Get(PropInitializationVector) }

func SetInitializationVector(e ld.Entity, v interface{}) { e.Set(PropInitializationVector, v) }

// This property is used in conjunction with the input to the signature hashing function in order to
// protect against replay attacks. Typically, receivers need to track all nonce values used within a
// certain time period in order to ensure that an attacker cannot merely re-send a compromised packet
// in order to execute a privileged request.
func GetNonce(e ld.Entity) interface{} { return e.Get(PropNonce) }

func SetNonce(e ld.Entity, v interface{}) { e.Set(PropNonce, v) }

// An owner is an entity that claims control over a particular resource. Note that ownership is best
// validated as a two-way relationship where the owner claims ownership over a particular resource,
// and the resource clearly identifies its owner.
func GetOwner(e ld.Entity) interface{} { return e.Get(PropOwner) }

func SetOwner(e ld.Entity, v interface{}) { e.Set(PropOwner, v) }

// A secret that is used to generate a key that can be used to encrypt or decrypt message.
// It is typically a string value.
func GetPassword(e ld.Entity) interface{} { return e.Get(PropPassword) }

func SetPassword(e ld.Entity, v interface{}) { e.Set(PropPassword, v) }

// A private key PEM property is used to specify the PEM-encoded version of the private key.
// This encoding is compatible with almost every Secure Sockets Layer library implementation and
// typically plugs directly into functions intializing private keys.
func GetPrivateKeyPem(e ld.Entity) interface{} { return e.Get(PropPrivateKeyPem) }

func SetPrivateKeyPem(e ld.Entity, v interface{}) { e.Set(PropPrivateKeyPem, v) }

// A public key property is used to specify a URL that contains information about a public key.
func GetPublicKey(e ld.Entity) interface{} { return e.Get(PropPublicKey) }

func SetPublicKey(e ld.Entity, v interface{}) { e.Set(PropPublicKey, v) }

// A public key PEM property is used to specify the PEM-encoded version of the public key.
// This encoding is compatible with almost every Secure Sockets Layer library implementation and
// typically plugs directly into functions intializing public keys.
func GetPublicKeyPem(e ld.Entity) interface{} { return e.Get(PropPublicKeyPem) }

func SetPublicKeyPem(e ld.Entity, v interface{}) { e.Set(PropPublicKeyPem, v) }

// The publicKeyService property is used to express the REST URL that provides public key management
// services as defined by the Web Key [SECURE-MESSAGING] specification.
func GetPublicKeyService(e ld.Entity) interface{} { return e.Get(PropPublicKeyService) }

func SetPublicKeyService(e ld.Entity, v interface{}) { e.Set(PropPublicKeyService, v) }

// The revocation time is typically associated with a Key that has been marked as invalid as of the
// date and time associated with the property. Key revocations are often used when a key is
// compromised, such as the theft of the private key, or during the course of best-practice key
// rotation schedules.
func GetRevoked(e ld.Entity) interface{} { return e.Get(PropRevoked) }

func SetRevoked(e ld.Entity, v interface{}) { e.Set(PropRevoked, v) }

// The signature property is used to associate a signature with a graph of information.
// The signature property is typically not included in the canonicalized graph that is then digested,
// and digitally signed.
func GetSignature(e ld.Entity) interface{} { return e.Get(PropSignature) }

func SetSignature(e ld.Entity, v interface{}) { e.Set(PropSignature, v) }

// The signature algorithm is used to specify the cryptographic signature function to use when
// digitally signing the digest data. Typically, text to be signed goes through three steps:
// 1) canonicalization, 2) digest, and 3) signature. This property is used to specify the algorithm
// that should be used for step #3. A signature class typically specifies a default signature
// algorithm, so this property rarely needs to be used in practice when specifying digital signatures.
func GetSignatureAlgorithm(e ld.Entity) interface{} { return e.Get(PropSignatureAlgorithm) }

func SetSignatureAlgorithm(e ld.Entity, v interface{}) { e.Set(PropSignatureAlgorithm, v) }

// The signature value is used to express the output of the signature algorithm expressed in
// base-64 format.
func GetSignatureValue(e ld.Entity) interface{} { return e.Get(PropSignatureValue) }

func SetSignatureValue(e ld.Entity, v interface{}) { e.Set(PropSignatureValue, v) }
