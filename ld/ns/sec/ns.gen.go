// GENERATED FILE, DO NOT EDIT.
// Please refer to: tools/nsgen/templates.go
package sec

const (
	PropAuthenticationTag = "https://w3id.org/security#authenticationTag"

	// The canonicalization algorithm is used to transform the input data into a form that can be passed
	// to a cryptographic digest method. The digest is then digitally signed using a digital signature
	// algorithm. Canonicalization ensures that a piece of software that is generating a digital signature
	// is able to do so on the same set of information in a deterministic manner.
	PropCanonicalizationAlgorithm = "https://w3id.org/security#canonicalizationAlgorithm"

	// The cipher algorithm describes the mechanism used to encrypt a message. It is typically a string
	// expressing the cipher suite, the strength of the cipher, and a block cipher mode.
	PropCipherAlgorithm = "https://w3id.org/security#cipherAlgorithm"

	// Cipher data an opaque blob of information that is used to specify an encrypted message.
	PropCipherData = "https://w3id.org/security#cipherData"

	// A cipher key is a symmetric key that is used to encrypt or decrypt a piece of information.
	// The key itself may be expressed in clear text or encrypted.
	PropCipherKey = "https://w3id.org/security#cipherKey"

	PropCreator = "https://w3id.org/security#creator"

	// The digest algorithm is used to specify the cryptographic function to use when generating the
	// data to be digitally signed. Typically, data that is to be signed goes through three steps:
	// 1) canonicalization, 2) digest, and 3) signature. This property is used to specify the algorithm
	// that should be used for step #2. A signature class typically specifies a default digest method,
	// so this property is typically used to specify information for a signature algorithm.
	PropDigestAlgorithm = "https://w3id.org/security#digestAlgorithm"

	// The digest value is used to express the output of the digest algorithm expressed in Base-1
	// (hexadecimal) format.
	PropDigestValue = "https://w3id.org/security#digestValue"

	// The expiration time is typically associated with a Key and specifies when the validity of the key
	// will expire. It is considered a best practice to only create keys that have very definite
	// expiration periods. This period is typically set to between six months and two years. An digital
	// signature created using an expired key MUST be marked as invalid by any software attempting to
	// verify the signature.
	PropExpires = "https://w3id.org/security#expires"

	// The initialization vector (IV) is a byte stream that is typically used to initialize certain block
	// cipher encryption schemes. For a receiving application to be able to decrypt a message, it must
	// know the decryption key and the initialization vector. The value is typically base-64 encoded.
	PropInitializationVector = "https://w3id.org/security#initializationVector"

	// This property is used in conjunction with the input to the signature hashing function in order to
	// protect against replay attacks. Typically, receivers need to track all nonce values used within a
	// certain time period in order to ensure that an attacker cannot merely re-send a compromised packet
	// in order to execute a privileged request.
	PropNonce = "https://w3id.org/security#nonce"

	// An owner is an entity that claims control over a particular resource. Note that ownership is best
	// validated as a two-way relationship where the owner claims ownership over a particular resource,
	// and the resource clearly identifies its owner.
	PropOwner = "https://w3id.org/security#owner"

	// A secret that is used to generate a key that can be used to encrypt or decrypt message.
	// It is typically a string value.
	PropPassword = "https://w3id.org/security#password"

	// A private key PEM property is used to specify the PEM-encoded version of the private key.
	// This encoding is compatible with almost every Secure Sockets Layer library implementation and
	// typically plugs directly into functions intializing private keys.
	PropPrivateKeyPem = "https://w3id.org/security#privateKeyPem"

	// A public key property is used to specify a URL that contains information about a public key.
	PropPublicKey = "https://w3id.org/security#publicKey"

	// A public key PEM property is used to specify the PEM-encoded version of the public key.
	// This encoding is compatible with almost every Secure Sockets Layer library implementation and
	// typically plugs directly into functions intializing public keys.
	PropPublicKeyPem = "https://w3id.org/security#publicKeyPem"

	// The publicKeyService property is used to express the REST URL that provides public key management
	// services as defined by the Web Key [SECURE-MESSAGING] specification.
	PropPublicKeyService = "https://w3id.org/security#publicKeyService"

	// The revocation time is typically associated with a Key that has been marked as invalid as of the
	// date and time associated with the property. Key revocations are often used when a key is
	// compromised, such as the theft of the private key, or during the course of best-practice key
	// rotation schedules.
	PropRevoked = "https://w3id.org/security#revoked"

	// The signature property is used to associate a signature with a graph of information.
	// The signature property is typically not included in the canonicalized graph that is then digested,
	// and digitally signed.
	PropSignature = "https://w3id.org/security#signature"

	// The signature algorithm is used to specify the cryptographic signature function to use when
	// digitally signing the digest data. Typically, text to be signed goes through three steps:
	// 1) canonicalization, 2) digest, and 3) signature. This property is used to specify the algorithm
	// that should be used for step #3. A signature class typically specifies a default signature
	// algorithm, so this property rarely needs to be used in practice when specifying digital signatures.
	PropSignatureAlgorithm = "https://w3id.org/security#signatureAlgorithm"

	// The signature value is used to express the output of the signature algorithm expressed in
	// base-64 format.
	PropSignatureValue = "https://w3id.org/security#signatureValue"
)

const (
	// This class represents a message digest that may be used for data integrity verification.
	// The digest algorithm used will determine the cryptographic properties of the digest.
	TypeDigest = "https://w3id.org/security#Digest"

	// A class of messages that are obfuscated in some cryptographic manner.
	// These messages are incredibly difficult to decrypt without the proper decryption key.
	TypeEncryptedMessage = "https://w3id.org/security#EncryptedMessage"

	// A graph signature is used for digital signatures on RDF graphs. The default canonicalization
	// mechanism is specified in the RDF Graph normalization specification, which effectively
	// deterministically names all unnamed nodes. The default signature mechanism uses a SHA-256 digest
	// and RSA to perform the digital signature.
	TypeGraphSignature2012 = "https://w3id.org/security#GraphSignature2012"

	// This class represents a cryptographic key that may be used for encryption, decryption, or
	// digitally signing data.
	TypeKey = "https://w3id.org/security#Key"

	// A Linked Data signature is used for digital signatures on RDF Datasets. The default canonicalization
	// mechanism is specified in the RDF Dataset Normalization specification, which effectively
	// deterministically names all unnamed nodes. The default signature mechanism uses a SHA-256 digest and
	// RSA to perform the digital signature. This signature uses a algorithm for producing the data that it
	// signs and verifies that is different from other Linked Data signatures.
	TypeLinkedDataSignature2015 = "https://w3id.org/security#LinkedDataSignature2015"

	// A Linked Data signature is used for digital signatures on RDF Datasets. The default canonicalization
	// mechanism is specified in the RDF Dataset Normalization specification, which effectively
	// deterministically names all unnamed nodes. The default signature mechanism uses a SHA-256 digest and
	// RSA to perform the digital signature.
	TypeLinkedDataSignature2016 = "https://w3id.org/security#LinkedDataSignature2016"

	// This class represents a digital signature on serialized data. It is an abstract class and should
	// not be used other than for Semantic Web reasoning purposes, such as by a reasoning agent.
	TypeSignature = "https://w3id.org/security#Signature"
)
