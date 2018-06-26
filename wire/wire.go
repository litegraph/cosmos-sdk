package wire

import (
	"bytes"
	"encoding/json"

	amino "github.com/tendermint/go-amino"
	"github.com/tendermint/tendermint/crypto"
)

// amino codec to marshal/unmarshal
type Codec = amino.Codec

func NewCodec() *Codec {
	cdc := amino.NewCodec()
	return cdc
}

// Register the go-crypto to the codec
func RegisterCrypto(cdc *Codec) {
	crypto.RegisterAmino(cdc)
}

// attempt to make some pretty json
func MarshalJSONIndent(cdc *Codec, obj interface{}) ([]byte, error) {
	bz, err := cdc.MarshalJSON(obj)
	if err != nil {
		return nil, err
	}

	var out bytes.Buffer
	err = json.Indent(&out, bz, "", "  ")
	if err != nil {
		return nil, err
	}
	return out.Bytes(), nil
}
