package patches

import (
	"github.com/meowpub/meow/ld"
	"github.com/meowpub/meow/ld/ns/owl"
	"github.com/meowpub/meow/ld/ns/rdf"
	"github.com/meowpub/meow/ld/ns/sec"
)

var SecPatches = []*ld.Object{
	// This doesn't have any proper data on which properties are on what.
	// It's in the HTML, but not encoded properly as RDF. Thanks Manu Sporny.
	patch(sec.PropDigestAlgorithm, rdf.PropDomain, "@id", sec.TypeDigest),
	patch(sec.PropDigestValue, rdf.PropDomain, "@id", sec.TypeDigest),

	patch(sec.PropAuthenticationTag, rdf.PropDomain, "@id", sec.TypeEncryptedMessage),
	patch(sec.PropCipherAlgorithm, rdf.PropDomain, "@id", sec.TypeEncryptedMessage),
	patch(sec.PropCipherData, rdf.PropDomain, "@id", sec.TypeEncryptedMessage),
	patch(sec.PropCipherKey, rdf.PropDomain, "@id", sec.TypeEncryptedMessage),
	patch(sec.PropInitializationVector, rdf.PropDomain, "@id", sec.TypeEncryptedMessage),
	patch(sec.PropPublicKey, rdf.PropDomain, "@id", sec.TypeEncryptedMessage),

	patch(sec.PropSignatureValue, rdf.PropDomain, "@id", sec.TypeSignature),
	patch(sec.PropCreator, rdf.PropDomain, "@id", sec.TypeSignature),

	patch(sec.PropOwner, rdf.PropDomain, "@id", sec.TypeKey),
	patch(sec.PropPrivateKeyPem, rdf.PropDomain, "@id", sec.TypeKey),
	patch(sec.PropPublicKeyPem, rdf.PropDomain, "@id", sec.TypeKey),

	// Nothing has parent classes. Because of course they don't.
	patch(sec.TypeDigest, rdf.PropSubClassOf, "@id", owl.TypeThing),
	patch(sec.TypeEncryptedMessage, rdf.PropSubClassOf, "@id", owl.TypeThing),
	patch(sec.TypeGraphSignature2012, rdf.PropSubClassOf, "@id", sec.TypeSignature),
	patch(sec.TypeLinkedDataSignature2015, rdf.PropSubClassOf, "@id", sec.TypeSignature),
	patch(sec.TypeLinkedDataSignature2016, rdf.PropSubClassOf, "@id", sec.TypeSignature),
	patch(sec.TypeSignature, rdf.PropSubClassOf, "@id", owl.TypeThing),
	patch(sec.TypeKey, rdf.PropSubClassOf, "@id", owl.TypeThing),

	// Aaaaaaaaand there are properties that are just straight up never defined.
	patch("https://w3id.org/security#authenticationTag", rdf.PropType, "@id", rdf.TypeProperty),
	patch("https://w3id.org/security#creator", rdf.PropType, "@id", rdf.TypeProperty),

	// This entire *class* lacks RDF annotations, why are we even using this file.
	patch("https://w3id.org/security#LinkedDataSignature2016", rdf.PropType, "@id", rdf.TypeClass),

	// Fuck it, comments too.
	comment(sec.TypeDigest, `
This class represents a message digest that may be used for data integrity verification.
The digest algorithm used will determine the cryptographic properties of the digest.`[1:]),

	comment(sec.TypeEncryptedMessage, `
A class of messages that are obfuscated in some cryptographic manner.
These messages are incredibly difficult to decrypt without the proper decryption key.`[1:]),

	comment(sec.TypeGraphSignature2012, `
A graph signature is used for digital signatures on RDF graphs. The default canonicalization
mechanism is specified in the RDF Graph normalization specification, which effectively
deterministically names all unnamed nodes. The default signature mechanism uses a SHA-256 digest
and RSA to perform the digital signature.`[1:]),

	comment(sec.TypeLinkedDataSignature2015, `
A Linked Data signature is used for digital signatures on RDF Datasets. The default canonicalization
mechanism is specified in the RDF Dataset Normalization specification, which effectively
deterministically names all unnamed nodes. The default signature mechanism uses a SHA-256 digest and
RSA to perform the digital signature. This signature uses a algorithm for producing the data that it
signs and verifies that is different from other Linked Data signatures.`[1:]),

	comment(sec.TypeLinkedDataSignature2016, `
A Linked Data signature is used for digital signatures on RDF Datasets. The default canonicalization
mechanism is specified in the RDF Dataset Normalization specification, which effectively
deterministically names all unnamed nodes. The default signature mechanism uses a SHA-256 digest and
RSA to perform the digital signature.`[1:]),

	comment(sec.TypeKey, `
This class represents a cryptographic key that may be used for encryption, decryption, or
digitally signing data.`[1:]),

	comment(sec.TypeSignature, `
This class represents a digital signature on serialized data. It is an abstract class and should
not be used other than for Semantic Web reasoning purposes, such as by a reasoning agent.`[1:]),

	comment(sec.PropCipherAlgorithm, `
The cipher algorithm describes the mechanism used to encrypt a message. It is typically a string
expressing the cipher suite, the strength of the cipher, and a block cipher mode.`[1:]),

	comment(sec.PropCipherData, `
Cipher data an opaque blob of information that is used to specify an encrypted message.`[1:]),

	comment(sec.PropDigestAlgorithm, `
The digest algorithm is used to specify the cryptographic function to use when generating the
data to be digitally signed. Typically, data that is to be signed goes through three steps:
1) canonicalization, 2) digest, and 3) signature. This property is used to specify the algorithm
that should be used for step #2. A signature class typically specifies a default digest method,
so this property is typically used to specify information for a signature algorithm.`[1:]),

	comment(sec.PropDigestValue, `
The digest value is used to express the output of the digest algorithm expressed in Base-1
(hexadecimal) format.`[1:]),

	comment(sec.PropCipherKey, `
A cipher key is a symmetric key that is used to encrypt or decrypt a piece of information.
The key itself may be expressed in clear text or encrypted.`[1:]),

	comment(sec.PropExpires, `
The expiration time is typically associated with a Key and specifies when the validity of the key
will expire. It is considered a best practice to only create keys that have very definite
expiration periods. This period is typically set to between six months and two years. An digital
signature created using an expired key MUST be marked as invalid by any software attempting to
verify the signature.`[1:]),

	comment(sec.PropInitializationVector, `
The initialization vector (IV) is a byte stream that is typically used to initialize certain block
cipher encryption schemes. For a receiving application to be able to decrypt a message, it must
know the decryption key and the initialization vector. The value is typically base-64 encoded.`[1:]),

	comment(sec.PropNonce, `
This property is used in conjunction with the input to the signature hashing function in order to
protect against replay attacks. Typically, receivers need to track all nonce values used within a
certain time period in order to ensure that an attacker cannot merely re-send a compromised packet
in order to execute a privileged request.`[1:]),

	comment(sec.PropCanonicalizationAlgorithm, `
The canonicalization algorithm is used to transform the input data into a form that can be passed
to a cryptographic digest method. The digest is then digitally signed using a digital signature
algorithm. Canonicalization ensures that a piece of software that is generating a digital signature
is able to do so on the same set of information in a deterministic manner.`[1:]),

	comment(sec.PropOwner, `
An owner is an entity that claims control over a particular resource. Note that ownership is best
validated as a two-way relationship where the owner claims ownership over a particular resource,
and the resource clearly identifies its owner.`[1:]),

	comment(sec.PropPassword, `
A secret that is used to generate a key that can be used to encrypt or decrypt message.
It is typically a string value.`[1:]),

	comment(sec.PropPrivateKeyPem, `
A private key PEM property is used to specify the PEM-encoded version of the private key.
This encoding is compatible with almost every Secure Sockets Layer library implementation and
typically plugs directly into functions intializing private keys.`[1:]),

	comment(sec.PropPublicKey, `
A public key property is used to specify a URL that contains information about a public key.`[1:]),

	comment(sec.PropPublicKeyPem, `
A public key PEM property is used to specify the PEM-encoded version of the public key.
This encoding is compatible with almost every Secure Sockets Layer library implementation and
typically plugs directly into functions intializing public keys.`[1:]),

	comment(sec.PropPublicKeyService, `
The publicKeyService property is used to express the REST URL that provides public key management
services as defined by the Web Key [SECURE-MESSAGING] specification.`[1:]),

	comment(sec.PropRevoked, `
The revocation time is typically associated with a Key that has been marked as invalid as of the
date and time associated with the property. Key revocations are often used when a key is
compromised, such as the theft of the private key, or during the course of best-practice key
rotation schedules.`[1:]),

	comment(sec.PropSignature, `
The signature property is used to associate a signature with a graph of information.
The signature property is typically not included in the canonicalized graph that is then digested,
and digitally signed.`[1:]),

	comment(sec.PropSignatureValue, `
The signature value is used to express the output of the signature algorithm expressed in
base-64 format.`[1:]),

	comment(sec.PropSignatureAlgorithm, `
The signature algorithm is used to specify the cryptographic signature function to use when
digitally signing the digest data. Typically, text to be signed goes through three steps:
1) canonicalization, 2) digest, and 3) signature. This property is used to specify the algorithm
that should be used for step #3. A signature class typically specifies a default signature
algorithm, so this property rarely needs to be used in practice when specifying digital signatures.`[1:]),
}
