package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/google/uuid"
)

var _ sdk.Msg = &MsgCreateAmulet{}

type MsgCreateAmulet struct {
  ID      string
  Creator sdk.AccAddress `json:"creator" yaml:"creator"`
  Name string `json:"Name" yaml:"Name"`
  Identity string `json:"Identity" yaml:"Identity"`
  Address string `json:"Address" yaml:"Address"`
  Telephon string `json:"Telephon" yaml:"Telephon"`
  AmuletType string `json:"AmuletType" yaml:"AmuletType"`
  AmuletName string `json:"AmuletName" yaml:"AmuletName"`
  From string `json:"From" yaml:"From"`
  Model string `json:"Model" yaml:"Model"`
  Surface string `json:"Surface" yaml:"Surface"`
  Year string `json:"Year" yaml:"Year"`
  Province string `json:"Province" yaml:"Province"`
  Description string `json:"Description" yaml:"Description"`
  Remark string `json:"Remark" yaml:"Remark"`
}

func NewMsgCreateAmulet(creator sdk.AccAddress, Name string, Identity string, Address string, Telephon string, AmuletType string, AmuletName string, From string, Model string, Surface string, Year string, Province string, Description string, Remark string) MsgCreateAmulet {
  return MsgCreateAmulet{
    ID: uuid.New().String(),
		Creator: creator,
    Name: Name,
    Identity: Identity,
    Address: Address,
    Telephon: Telephon,
    AmuletType: AmuletType,
    AmuletName: AmuletName,
    From: From,
    Model: Model,
    Surface: Surface,
    Year: Year,
    Province: Province,
    Description: Description,
    Remark: Remark,
	}
}

func (msg MsgCreateAmulet) Route() string {
  return RouterKey
}

func (msg MsgCreateAmulet) Type() string {
  return "CreateAmulet"
}

func (msg MsgCreateAmulet) GetSigners() []sdk.AccAddress {
  return []sdk.AccAddress{sdk.AccAddress(msg.Creator)}
}

func (msg MsgCreateAmulet) GetSignBytes() []byte {
  bz := ModuleCdc.MustMarshalJSON(msg)
  return sdk.MustSortJSON(bz)
}

func (msg MsgCreateAmulet) ValidateBasic() error {
  if msg.Creator.Empty() {
    return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator can't be empty")
  }
  return nil
}