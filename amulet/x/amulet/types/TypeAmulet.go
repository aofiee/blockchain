package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type Amulet struct {
	Creator sdk.AccAddress `json:"creator" yaml:"creator"`
	ID      string         `json:"id" yaml:"id"`
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