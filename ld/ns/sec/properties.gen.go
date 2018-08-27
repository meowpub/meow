// GENERATED FILE, DO NOT EDIT.
// Please refer to: tools/nsgen/templates.go
package sec

import (
	"github.com/meowpub/meow/ld"
)

func GetCanonicalizationAlgorithm(e ld.Entity) interface{} {
	return e.Get(PropCanonicalizationAlgorithm)
}

func GetCipherAlgorithm(e ld.Entity) interface{} { return e.Get(PropCipherAlgorithm) }

func GetCipherData(e ld.Entity) interface{} { return e.Get(PropCipherData) }

func GetCipherKey(e ld.Entity) interface{} { return e.Get(PropCipherKey) }

func GetDigestAlgorithm(e ld.Entity) interface{} { return e.Get(PropDigestAlgorithm) }

func GetDigestValue(e ld.Entity) interface{} { return e.Get(PropDigestValue) }

func GetExpires(e ld.Entity) interface{} { return e.Get(PropExpires) }

func GetInitializationVector(e ld.Entity) interface{} { return e.Get(PropInitializationVector) }

func GetNonce(e ld.Entity) interface{} { return e.Get(PropNonce) }

func GetOwner(e ld.Entity) interface{} { return e.Get(PropOwner) }

func GetPassword(e ld.Entity) interface{} { return e.Get(PropPassword) }

func GetPrivateKeyPem(e ld.Entity) interface{} { return e.Get(PropPrivateKeyPem) }

func GetPublicKey(e ld.Entity) interface{} { return e.Get(PropPublicKey) }

func GetPublicKeyPem(e ld.Entity) interface{} { return e.Get(PropPublicKeyPem) }

func GetPublicKeyService(e ld.Entity) interface{} { return e.Get(PropPublicKeyService) }

func GetRevoked(e ld.Entity) interface{} { return e.Get(PropRevoked) }

func GetSignature(e ld.Entity) interface{} { return e.Get(PropSignature) }

func GetSignatureAlgorithm(e ld.Entity) interface{} { return e.Get(PropSignatureAlgorithm) }

func GetSignatureValue(e ld.Entity) interface{} { return e.Get(PropSignatureValue) }
