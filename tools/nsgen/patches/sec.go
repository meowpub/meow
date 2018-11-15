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
	patch(sec.Prop_DigestAlgorithm.ID, rdf.Prop_Domain.ID, "@id", sec.Class_Digest.ID),
	patch(sec.Prop_DigestValue.ID, rdf.Prop_Domain.ID, "@id", sec.Class_Digest.ID),

	patch(sec.Prop_AuthenticationTag.ID, rdf.Prop_Domain.ID, "@id", sec.Class_EncryptedMessage.ID),
	patch(sec.Prop_CipherAlgorithm.ID, rdf.Prop_Domain.ID, "@id", sec.Class_EncryptedMessage.ID),
	patch(sec.Prop_CipherData.ID, rdf.Prop_Domain.ID, "@id", sec.Class_EncryptedMessage.ID),
	patch(sec.Prop_CipherKey.ID, rdf.Prop_Domain.ID, "@id", sec.Class_EncryptedMessage.ID),
	patch(sec.Prop_InitializationVector.ID, rdf.Prop_Domain.ID, "@id", sec.Class_EncryptedMessage.ID),
	patch(sec.Prop_PublicKey.ID, rdf.Prop_Domain.ID, "@id", sec.Class_EncryptedMessage.ID),

	patch(sec.Prop_SignatureValue.ID, rdf.Prop_Domain.ID, "@id", sec.Class_Signature.ID),
	patch(sec.Prop_Creator.ID, rdf.Prop_Domain.ID, "@id", sec.Class_Signature.ID),

	patch(sec.Prop_Owner.ID, rdf.Prop_Domain.ID, "@id", sec.Class_Key.ID),
	patch(sec.Prop_PrivateKeyPem.ID, rdf.Prop_Domain.ID, "@id", sec.Class_Key.ID),
	patch(sec.Prop_PublicKeyPem.ID, rdf.Prop_Domain.ID, "@id", sec.Class_Key.ID),

	// Nothing has parent classes. Because of course they don't.
	patch(sec.Class_Digest.ID, rdf.Prop_SubClassOf.ID, "@id", owl.Class_Thing.ID),
	patch(sec.Class_EncryptedMessage.ID, rdf.Prop_SubClassOf.ID, "@id", owl.Class_Thing.ID),
	patch(sec.Class_GraphSignature2012.ID, rdf.Prop_SubClassOf.ID, "@id", sec.Class_Signature.ID),
	patch(sec.Class_LinkedDataSignature2015.ID, rdf.Prop_SubClassOf.ID, "@id", sec.Class_Signature.ID),
	patch(sec.Class_LinkedDataSignature2016.ID, rdf.Prop_SubClassOf.ID, "@id", sec.Class_Signature.ID),
	patch(sec.Class_Signature.ID, rdf.Prop_SubClassOf.ID, "@id", owl.Class_Thing.ID),
	patch(sec.Class_Key.ID, rdf.Prop_SubClassOf.ID, "@id", owl.Class_Thing.ID),

	// Aaaaaaaaand there are properties that are just straight up never defined.
	patch("https://w3id.org/security#authenticationTag", rdf.Prop_Type.ID, "@id", rdf.Class_Property.ID),
	patch("https://w3id.org/security#creator", rdf.Prop_Type.ID, "@id", rdf.Class_Property.ID),

	// This entire *class* lacks RDF annotations, why are we even using this file.
	patch("https://w3id.org/security#LinkedDataSignature2016", rdf.Prop_Type.ID, "@id", rdf.Class_Class.ID),

	// Fuck it, comments too.
	comment(sec.Class_Digest.ID, `
This class represents a message digest that may be used for data integrity verification.
The digest algorithm used will determine the cryptographic properties of the digest.`[1:]),

	comment(sec.Class_EncryptedMessage.ID, `
A class of messages that are obfuscated in some cryptographic manner.
These messages are incredibly difficult to decrypt without the proper decryption key.`[1:]),

	comment(sec.Class_GraphSignature2012.ID, `
A graph signature is used for digital signatures on RDF graphs. The default canonicalization
mechanism is specified in the RDF Graph normalization specification, which effectively
deterministically names all unnamed nodes. The default signature mechanism uses a SHA-256 digest
and RSA to perform the digital signature.`[1:]),

	comment(sec.Class_LinkedDataSignature2015.ID, `
A Linked Data signature is used for digital signatures on RDF Datasets. The default canonicalization
mechanism is specified in the RDF Dataset Normalization specification, which effectively
deterministically names all unnamed nodes. The default signature mechanism uses a SHA-256 digest and
RSA to perform the digital signature. This signature uses a algorithm for producing the data that it
signs and verifies that is different from other Linked Data signatures.`[1:]),

	comment(sec.Class_LinkedDataSignature2016.ID, `
A Linked Data signature is used for digital signatures on RDF Datasets. The default canonicalization
mechanism is specified in the RDF Dataset Normalization specification, which effectively
deterministically names all unnamed nodes. The default signature mechanism uses a SHA-256 digest and
RSA to perform the digital signature.`[1:]),

	comment(sec.Class_Key.ID, `
This class represents a cryptographic key that may be used for encryption, decryption, or
digitally signing data.`[1:]),

	comment(sec.Class_Signature.ID, `
This class represents a digital signature on serialized data. It is an abstract class and should
not be used other than for Semantic Web reasoning purposes, such as by a reasoning agent.`[1:]),

	comment(sec.Prop_CipherAlgorithm.ID, `
The cipher algorithm describes the mechanism used to encrypt a message. It is typically a string
expressing the cipher suite, the strength of the cipher, and a block cipher mode.`[1:]),

	comment(sec.Prop_CipherData.ID, `
Cipher data an opaque blob of information that is used to specify an encrypted message.`[1:]),

	comment(sec.Prop_DigestAlgorithm.ID, `
The digest algorithm is used to specify the cryptographic function to use when generating the
data to be digitally signed. Typically, data that is to be signed goes through three steps:
1) canonicalization, 2) digest, and 3) signature. This property is used to specify the algorithm
that should be used for step #2. A signature class typically specifies a default digest method,
so this property is typically used to specify information for a signature algorithm.`[1:]),

	comment(sec.Prop_DigestValue.ID, `
The digest value is used to express the output of the digest algorithm expressed in Base-1
(hexadecimal) format.`[1:]),

	comment(sec.Prop_CipherKey.ID, `
A cipher key is a symmetric key that is used to encrypt or decrypt a piece of information.
The key itself may be expressed in clear text or encrypted.`[1:]),

	comment(sec.Prop_Expires.ID, `
The expiration time is typically associated with a Key and specifies when the validity of the key
will expire. It is considered a best practice to only create keys that have very definite
expiration periods. This period is typically set to between six months and two years. An digital
signature created using an expired key MUST be marked as invalid by any software attempting to
verify the signature.`[1:]),

	comment(sec.Prop_InitializationVector.ID, `
The initialization vector (IV) is a byte stream that is typically used to initialize certain block
cipher encryption schemes. For a receiving application to be able to decrypt a message, it must
know the decryption key and the initialization vector. The value is typically base-64 encoded.`[1:]),

	comment(sec.Prop_Nonce.ID, `
This property is used in conjunction with the input to the signature hashing function in order to
protect against replay attacks. Typically, receivers need to track all nonce values used within a
certain time period in order to ensure that an attacker cannot merely re-send a compromised packet
in order to execute a privileged request.`[1:]),

	comment(sec.Prop_CanonicalizationAlgorithm.ID, `
The canonicalization algorithm is used to transform the input data into a form that can be passed
to a cryptographic digest method. The digest is then digitally signed using a digital signature
algorithm. Canonicalization ensures that a piece of software that is generating a digital signature
is able to do so on the same set of information in a deterministic manner.`[1:]),

	comment(sec.Prop_Owner.ID, `
An owner is an entity that claims control over a particular resource. Note that ownership is best
validated as a two-way relationship where the owner claims ownership over a particular resource,
and the resource clearly identifies its owner.`[1:]),

	comment(sec.Prop_Password.ID, `
A secret that is used to generate a key that can be used to encrypt or decrypt message.
It is typically a string value.`[1:]),

	comment(sec.Prop_PrivateKeyPem.ID, `
A private key PEM property is used to specify the PEM-encoded version of the private key.
This encoding is compatible with almost every Secure Sockets Layer library implementation and
typically plugs directly into functions intializing private keys.`[1:]),

	comment(sec.Prop_PublicKey.ID, `
A public key property is used to specify a URL that contains information about a public key.`[1:]),

	comment(sec.Prop_PublicKeyPem.ID, `
A public key PEM property is used to specify the PEM-encoded version of the public key.
This encoding is compatible with almost every Secure Sockets Layer library implementation and
typically plugs directly into functions intializing public keys.`[1:]),

	comment(sec.Prop_PublicKeyService.ID, `
The publicKeyService property is used to express the REST URL that provides public key management
services as defined by the Web Key [SECURE-MESSAGING] specification.`[1:]),

	comment(sec.Prop_Revoked.ID, `
The revocation time is typically associated with a Key that has been marked as invalid as of the
date and time associated with the property. Key revocations are often used when a key is
compromised, such as the theft of the private key, or during the course of best-practice key
rotation schedules.`[1:]),

	comment(sec.Prop_Signature.ID, `
The signature property is used to associate a signature with a graph of information.
The signature property is typically not included in the canonicalized graph that is then digested,
and digitally signed.`[1:]),

	comment(sec.Prop_SignatureValue.ID, `
The signature value is used to express the output of the signature algorithm expressed in
base-64 format.`[1:]),

	comment(sec.Prop_SignatureAlgorithm.ID, `
The signature algorithm is used to specify the cryptographic signature function to use when
digitally signing the digest data. Typically, text to be signed goes through three steps:
1) canonicalization, 2) digest, and 3) signature. This property is used to specify the algorithm
that should be used for step #3. A signature class typically specifies a default signature
algorithm, so this property rarely needs to be used in practice when specifying digital signatures.`[1:]),
}
