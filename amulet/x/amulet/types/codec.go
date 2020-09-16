package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
)

// RegisterCodec registers concrete types on codec
func RegisterCodec(cdc *codec.Codec) {
  // this line is used by starport scaffolding
		cdc.RegisterConcrete(MsgCreateBuying{}, "amulet/CreateBuying", nil)
		cdc.RegisterConcrete(MsgCreateSelling{}, "amulet/CreateSelling", nil)
		cdc.RegisterConcrete(MsgCreateAmulet{}, "amulet/CreateAmulet", nil)
}

// ModuleCdc defines the module codec
var ModuleCdc *codec.Codec

func init() {
	ModuleCdc = codec.New()
	RegisterCodec(ModuleCdc)
	codec.RegisterCrypto(ModuleCdc)
	ModuleCdc.Seal()
}
