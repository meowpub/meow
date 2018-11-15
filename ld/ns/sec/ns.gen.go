// GENERATED FILE, DO NOT EDIT.
// Please refer to: tools/nsgen/templates.go
package sec

import (
	"github.com/meowpub/meow/ld"
	"github.com/meowpub/meow/ld/ns/meta"
	"github.com/meowpub/meow/ld/ns/owl"
)

var (
	SEC = &meta.Namespace{
		ID:    "https://w3id.org/security#",
		Short: "sec",
		Props: []*meta.Prop{
			Prop_AuthenticationTag,
			Prop_CanonicalizationAlgorithm,
			Prop_CipherAlgorithm,
			Prop_CipherData,
			Prop_CipherKey,
			Prop_Creator,
			Prop_DigestAlgorithm,
			Prop_DigestValue,
			Prop_Expires,
			Prop_InitializationVector,
			Prop_Nonce,
			Prop_Owner,
			Prop_Password,
			Prop_PrivateKeyPem,
			Prop_PublicKey,
			Prop_PublicKeyPem,
			Prop_PublicKeyService,
			Prop_Revoked,
			Prop_Signature,
			Prop_SignatureAlgorithm,
			Prop_SignatureValue,
		},
		Types: map[string]*meta.Type{
			"https://w3id.org/security#Digest":                  Class_Digest,
			"https://w3id.org/security#EncryptedMessage":        Class_EncryptedMessage,
			"https://w3id.org/security#GraphSignature2012":      Class_GraphSignature2012,
			"https://w3id.org/security#Key":                     Class_Key,
			"https://w3id.org/security#LinkedDataSignature2015": Class_LinkedDataSignature2015,
			"https://w3id.org/security#LinkedDataSignature2016": Class_LinkedDataSignature2016,
			"https://w3id.org/security#Signature":               Class_Signature,
		},
	}

	Prop_AuthenticationTag = &meta.Prop{
		ID:      "https://w3id.org/security#authenticationTag",
		Short:   "authenticationTag",
		Comment: "",
	}

	// The canonicalization algorithm is used to transform the input data into a form that can be passed
	// to a cryptographic digest method. The digest is then digitally signed using a digital signature
	// algorithm. Canonicalization ensures that a piece of software that is generating a digital signature
	// is able to do so on the same set of information in a deterministic manner.
	Prop_CanonicalizationAlgorithm = &meta.Prop{
		ID:      "https://w3id.org/security#canonicalizationAlgorithm",
		Short:   "canonicalizationAlgorithm",
		Comment: "The canonicalization algorithm is used to transform the input data into a form that can be passed to a cryptographic digest method. The digest is then digitally signed using a digital signature algorithm. Canonicalization ensures that a piece of software that is generating a digital signature is able to do so on the same set of information in a deterministic manner.",
	}

	// The cipher algorithm describes the mechanism used to encrypt a message. It is typically a string
	// expressing the cipher suite, the strength of the cipher, and a block cipher mode.
	Prop_CipherAlgorithm = &meta.Prop{
		ID:      "https://w3id.org/security#cipherAlgorithm",
		Short:   "cipherAlgorithm",
		Comment: "The cipher algorithm describes the mechanism used to encrypt a message. It is typically a string expressing the cipher suite, the strength of the cipher, and a block cipher mode.",
	}

	// Cipher data an opaque blob of information that is used to specify an encrypted message.
	Prop_CipherData = &meta.Prop{
		ID:      "https://w3id.org/security#cipherData",
		Short:   "cipherData",
		Comment: "Cipher data an opaque blob of information that is used to specify an encrypted message.",
	}

	// A cipher key is a symmetric key that is used to encrypt or decrypt a piece of information.
	// The key itself may be expressed in clear text or encrypted.
	Prop_CipherKey = &meta.Prop{
		ID:      "https://w3id.org/security#cipherKey",
		Short:   "cipherKey",
		Comment: "A cipher key is a symmetric key that is used to encrypt or decrypt a piece of information. The key itself may be expressed in clear text or encrypted.",
	}

	Prop_Creator = &meta.Prop{
		ID:      "https://w3id.org/security#creator",
		Short:   "creator",
		Comment: "",
	}

	// The digest algorithm is used to specify the cryptographic function to use when generating the
	// data to be digitally signed. Typically, data that is to be signed goes through three steps:
	// 1) canonicalization, 2) digest, and 3) signature. This property is used to specify the algorithm
	// that should be used for step #2. A signature class typically specifies a default digest method,
	// so this property is typically used to specify information for a signature algorithm.
	Prop_DigestAlgorithm = &meta.Prop{
		ID:      "https://w3id.org/security#digestAlgorithm",
		Short:   "digestAlgorithm",
		Comment: "The digest algorithm is used to specify the cryptographic function to use when generating the data to be digitally signed. Typically, data that is to be signed goes through three steps: 1) canonicalization, 2) digest, and 3) signature. This property is used to specify the algorithm that should be used for step #2. A signature class typically specifies a default digest method, so this property is typically used to specify information for a signature algorithm.",
	}

	// The digest value is used to express the output of the digest algorithm expressed in Base-1
	// (hexadecimal) format.
	Prop_DigestValue = &meta.Prop{
		ID:      "https://w3id.org/security#digestValue",
		Short:   "digestValue",
		Comment: "The digest value is used to express the output of the digest algorithm expressed in Base-1 (hexadecimal) format.",
	}

	// The expiration time is typically associated with a Key and specifies when the validity of the key
	// will expire. It is considered a best practice to only create keys that have very definite
	// expiration periods. This period is typically set to between six months and two years. An digital
	// signature created using an expired key MUST be marked as invalid by any software attempting to
	// verify the signature.
	Prop_Expires = &meta.Prop{
		ID:      "https://w3id.org/security#expires",
		Short:   "expires",
		Comment: "The expiration time is typically associated with a Key and specifies when the validity of the key will expire. It is considered a best practice to only create keys that have very definite expiration periods. This period is typically set to between six months and two years. An digital signature created using an expired key MUST be marked as invalid by any software attempting to verify the signature.",
	}

	// The initialization vector (IV) is a byte stream that is typically used to initialize certain block
	// cipher encryption schemes. For a receiving application to be able to decrypt a message, it must
	// know the decryption key and the initialization vector. The value is typically base-64 encoded.
	Prop_InitializationVector = &meta.Prop{
		ID:      "https://w3id.org/security#initializationVector",
		Short:   "initializationVector",
		Comment: "The initialization vector (IV) is a byte stream that is typically used to initialize certain block cipher encryption schemes. For a receiving application to be able to decrypt a message, it must know the decryption key and the initialization vector. The value is typically base-64 encoded.",
	}

	// This property is used in conjunction with the input to the signature hashing function in order to
	// protect against replay attacks. Typically, receivers need to track all nonce values used within a
	// certain time period in order to ensure that an attacker cannot merely re-send a compromised packet
	// in order to execute a privileged request.
	Prop_Nonce = &meta.Prop{
		ID:      "https://w3id.org/security#nonce",
		Short:   "nonce",
		Comment: "This property is used in conjunction with the input to the signature hashing function in order to protect against replay attacks. Typically, receivers need to track all nonce values used within a certain time period in order to ensure that an attacker cannot merely re-send a compromised packet in order to execute a privileged request.",
	}

	// An owner is an entity that claims control over a particular resource. Note that ownership is best
	// validated as a two-way relationship where the owner claims ownership over a particular resource,
	// and the resource clearly identifies its owner.
	Prop_Owner = &meta.Prop{
		ID:      "https://w3id.org/security#owner",
		Short:   "owner",
		Comment: "An owner is an entity that claims control over a particular resource. Note that ownership is best validated as a two-way relationship where the owner claims ownership over a particular resource, and the resource clearly identifies its owner.",
	}

	// A secret that is used to generate a key that can be used to encrypt or decrypt message.
	// It is typically a string value.
	Prop_Password = &meta.Prop{
		ID:      "https://w3id.org/security#password",
		Short:   "password",
		Comment: "A secret that is used to generate a key that can be used to encrypt or decrypt message. It is typically a string value.",
	}

	// A private key PEM property is used to specify the PEM-encoded version of the private key.
	// This encoding is compatible with almost every Secure Sockets Layer library implementation and
	// typically plugs directly into functions intializing private keys.
	Prop_PrivateKeyPem = &meta.Prop{
		ID:      "https://w3id.org/security#privateKeyPem",
		Short:   "privateKeyPem",
		Comment: "A private key PEM property is used to specify the PEM-encoded version of the private key. This encoding is compatible with almost every Secure Sockets Layer library implementation and typically plugs directly into functions intializing private keys.",
	}

	// A public key property is used to specify a URL that contains information about a public key.
	Prop_PublicKey = &meta.Prop{
		ID:      "https://w3id.org/security#publicKey",
		Short:   "publicKey",
		Comment: "A public key property is used to specify a URL that contains information about a public key.",
	}

	// A public key PEM property is used to specify the PEM-encoded version of the public key.
	// This encoding is compatible with almost every Secure Sockets Layer library implementation and
	// typically plugs directly into functions intializing public keys.
	Prop_PublicKeyPem = &meta.Prop{
		ID:      "https://w3id.org/security#publicKeyPem",
		Short:   "publicKeyPem",
		Comment: "A public key PEM property is used to specify the PEM-encoded version of the public key. This encoding is compatible with almost every Secure Sockets Layer library implementation and typically plugs directly into functions intializing public keys.",
	}

	// The publicKeyService property is used to express the REST URL that provides public key management
	// services as defined by the Web Key [SECURE-MESSAGING] specification.
	Prop_PublicKeyService = &meta.Prop{
		ID:      "https://w3id.org/security#publicKeyService",
		Short:   "publicKeyService",
		Comment: "The publicKeyService property is used to express the REST URL that provides public key management services as defined by the Web Key [SECURE-MESSAGING] specification.",
	}

	// The revocation time is typically associated with a Key that has been marked as invalid as of the
	// date and time associated with the property. Key revocations are often used when a key is
	// compromised, such as the theft of the private key, or during the course of best-practice key
	// rotation schedules.
	Prop_Revoked = &meta.Prop{
		ID:      "https://w3id.org/security#revoked",
		Short:   "revoked",
		Comment: "The revocation time is typically associated with a Key that has been marked as invalid as of the date and time associated with the property. Key revocations are often used when a key is compromised, such as the theft of the private key, or during the course of best-practice key rotation schedules.",
	}

	// The signature property is used to associate a signature with a graph of information.
	// The signature property is typically not included in the canonicalized graph that is then digested,
	// and digitally signed.
	Prop_Signature = &meta.Prop{
		ID:      "https://w3id.org/security#signature",
		Short:   "signature",
		Comment: "The signature property is used to associate a signature with a graph of information. The signature property is typically not included in the canonicalized graph that is then digested, and digitally signed.",
	}

	// The signature algorithm is used to specify the cryptographic signature function to use when
	// digitally signing the digest data. Typically, text to be signed goes through three steps:
	// 1) canonicalization, 2) digest, and 3) signature. This property is used to specify the algorithm
	// that should be used for step #3. A signature class typically specifies a default signature
	// algorithm, so this property rarely needs to be used in practice when specifying digital signatures.
	Prop_SignatureAlgorithm = &meta.Prop{
		ID:      "https://w3id.org/security#signatureAlgorithm",
		Short:   "signatureAlgorithm",
		Comment: "The signature algorithm is used to specify the cryptographic signature function to use when digitally signing the digest data. Typically, text to be signed goes through three steps: 1) canonicalization, 2) digest, and 3) signature. This property is used to specify the algorithm that should be used for step #3. A signature class typically specifies a default signature algorithm, so this property rarely needs to be used in practice when specifying digital signatures.",
	}

	// The signature value is used to express the output of the signature algorithm expressed in
	// base-64 format.
	Prop_SignatureValue = &meta.Prop{
		ID:      "https://w3id.org/security#signatureValue",
		Short:   "signatureValue",
		Comment: "The signature value is used to express the output of the signature algorithm expressed in base-64 format.",
	}

	// This class represents a message digest that may be used for data integrity verification.
	// The digest algorithm used will determine the cryptographic properties of the digest.
	Class_Digest = &meta.Type{
		ID:    "https://w3id.org/security#Digest",
		Short: "Digest",
		Cast:  func(e ld.Entity) ld.Entity { return AsDigest(e) },
		SubClassOf: []*meta.Type{
			owl.Class_Thing,
		},
		Props: []*meta.Prop{
			Prop_DigestAlgorithm,
			Prop_DigestValue,
		},
	}

	// A class of messages that are obfuscated in some cryptographic manner.
	// These messages are incredibly difficult to decrypt without the proper decryption key.
	Class_EncryptedMessage = &meta.Type{
		ID:    "https://w3id.org/security#EncryptedMessage",
		Short: "EncryptedMessage",
		Cast:  func(e ld.Entity) ld.Entity { return AsEncryptedMessage(e) },
		SubClassOf: []*meta.Type{
			owl.Class_Thing,
		},
		Props: []*meta.Prop{
			Prop_AuthenticationTag,
			Prop_CipherAlgorithm,
			Prop_CipherData,
			Prop_CipherKey,
			Prop_InitializationVector,
			Prop_PublicKey,
		},
	}

	// A graph signature is used for digital signatures on RDF graphs. The default canonicalization
	// mechanism is specified in the RDF Graph normalization specification, which effectively
	// deterministically names all unnamed nodes. The default signature mechanism uses a SHA-256 digest
	// and RSA to perform the digital signature.
	Class_GraphSignature2012 = &meta.Type{
		ID:    "https://w3id.org/security#GraphSignature2012",
		Short: "GraphSignature2012",
		Cast:  func(e ld.Entity) ld.Entity { return AsGraphSignature2012(e) },
		SubClassOf: []*meta.Type{
			Class_Signature,
		},
		Props: []*meta.Prop{},
	}

	// This class represents a cryptographic key that may be used for encryption, decryption, or
	// digitally signing data.
	Class_Key = &meta.Type{
		ID:    "https://w3id.org/security#Key",
		Short: "Key",
		Cast:  func(e ld.Entity) ld.Entity { return AsKey(e) },
		SubClassOf: []*meta.Type{
			owl.Class_Thing,
		},
		Props: []*meta.Prop{
			Prop_Owner,
			Prop_PrivateKeyPem,
			Prop_PublicKeyPem,
		},
	}

	// A Linked Data signature is used for digital signatures on RDF Datasets. The default canonicalization
	// mechanism is specified in the RDF Dataset Normalization specification, which effectively
	// deterministically names all unnamed nodes. The default signature mechanism uses a SHA-256 digest and
	// RSA to perform the digital signature. This signature uses a algorithm for producing the data that it
	// signs and verifies that is different from other Linked Data signatures.
	Class_LinkedDataSignature2015 = &meta.Type{
		ID:    "https://w3id.org/security#LinkedDataSignature2015",
		Short: "LinkedDataSignature2015",
		Cast:  func(e ld.Entity) ld.Entity { return AsLinkedDataSignature2015(e) },
		SubClassOf: []*meta.Type{
			Class_Signature,
		},
		Props: []*meta.Prop{},
	}

	// A Linked Data signature is used for digital signatures on RDF Datasets. The default canonicalization
	// mechanism is specified in the RDF Dataset Normalization specification, which effectively
	// deterministically names all unnamed nodes. The default signature mechanism uses a SHA-256 digest and
	// RSA to perform the digital signature.
	Class_LinkedDataSignature2016 = &meta.Type{
		ID:    "https://w3id.org/security#LinkedDataSignature2016",
		Short: "LinkedDataSignature2016",
		Cast:  func(e ld.Entity) ld.Entity { return AsLinkedDataSignature2016(e) },
		SubClassOf: []*meta.Type{
			Class_Signature,
		},
		Props: []*meta.Prop{},
	}

	// This class represents a digital signature on serialized data. It is an abstract class and should
	// not be used other than for Semantic Web reasoning purposes, such as by a reasoning agent.
	Class_Signature = &meta.Type{
		ID:    "https://w3id.org/security#Signature",
		Short: "Signature",
		Cast:  func(e ld.Entity) ld.Entity { return AsSignature(e) },
		SubClassOf: []*meta.Type{
			owl.Class_Thing,
		},
		Props: []*meta.Prop{
			Prop_Creator,
			Prop_SignatureValue,
		},
	}
)
