package params

import "github.com/btcsuite/btcd/chaincfg"

// DashParams defines the chopped down Dash parameters.
var DashParams = chaincfg.Params{
	Name: "dash",

	PubKeyHashAddrID: 0x4C, // starts with X
	PrivateKeyID:     0xCC, // starts with 7 (uncompressed) or X (compressed)
}
