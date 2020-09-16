package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/google/uuid"
)

var _ sdk.Msg = &MsgCreateBuying{}

type MsgCreateBuying struct {
  ID      string
  Creator sdk.AccAddress `json:"creator" yaml:"creator"`
  AmuletHash string `json:"AmuletHash" yaml:"AmuletHash"`
  BuyingPrice string `json:"BuyingPrice" yaml:"BuyingPrice"`
  SellingBy string `json:"SellingBy" yaml:"SellingBy"`
  BuyingBy string `json:"BuyingBy" yaml:"BuyingBy"`
}

func NewMsgCreateBuying(creator sdk.AccAddress, AmuletHash string, BuyingPrice string, SellingBy string, BuyingBy string) MsgCreateBuying {
  return MsgCreateBuying{
    ID: uuid.New().String(),
		Creator: creator,
    AmuletHash: AmuletHash,
    BuyingPrice: BuyingPrice,
    SellingBy: SellingBy,
    BuyingBy: BuyingBy,
	}
}

func (msg MsgCreateBuying) Route() string {
  return RouterKey
}

func (msg MsgCreateBuying) Type() string {
  return "CreateBuying"
}

func (msg MsgCreateBuying) GetSigners() []sdk.AccAddress {
  return []sdk.AccAddress{sdk.AccAddress(msg.Creator)}
}

func (msg MsgCreateBuying) GetSignBytes() []byte {
  bz := ModuleCdc.MustMarshalJSON(msg)
  return sdk.MustSortJSON(bz)
}

func (msg MsgCreateBuying) ValidateBasic() error {
  if msg.Creator.Empty() {
    return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator can't be empty")
  }
  return nil
}