package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/google/uuid"
)

var _ sdk.Msg = &MsgCreateSelling{}

type MsgCreateSelling struct {
  ID      string
  Creator sdk.AccAddress `json:"creator" yaml:"creator"`
  AmuletHash string `json:"AmuletHash" yaml:"AmuletHash"`
  SellingPrice string `json:"SellingPrice" yaml:"SellingPrice"`
  SellingBy string `json:"SellingBy" yaml:"SellingBy"`
}

func NewMsgCreateSelling(creator sdk.AccAddress, AmuletHash string, SellingPrice string, SellingBy string) MsgCreateSelling {
  return MsgCreateSelling{
    ID: uuid.New().String(),
		Creator: creator,
    AmuletHash: AmuletHash,
    SellingPrice: SellingPrice,
    SellingBy: SellingBy,
	}
}

func (msg MsgCreateSelling) Route() string {
  return RouterKey
}

func (msg MsgCreateSelling) Type() string {
  return "CreateSelling"
}

func (msg MsgCreateSelling) GetSigners() []sdk.AccAddress {
  return []sdk.AccAddress{sdk.AccAddress(msg.Creator)}
}

func (msg MsgCreateSelling) GetSignBytes() []byte {
  bz := ModuleCdc.MustMarshalJSON(msg)
  return sdk.MustSortJSON(bz)
}

func (msg MsgCreateSelling) ValidateBasic() error {
  if msg.Creator.Empty() {
    return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator can't be empty")
  }
  return nil
}