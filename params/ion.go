package params

import "github.com/btcsuite/btcd/chaincfg"

// IonParams defines the chopped down Ion parameters.
var IonParams = chaincfg.Params{
	Name: "ion",

	PubKeyHashAddrID: 0x67, // starts with i
	PrivateKeyID:     0x99, // starts with ? (uncompressed) or ? (compressed)
}
