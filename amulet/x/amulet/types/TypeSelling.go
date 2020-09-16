package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type Selling struct {
	Creator sdk.AccAddress `json:"creator" yaml:"creator"`
	ID      string         `json:"id" yaml:"id"`
  AmuletHash string `json:"AmuletHash" yaml:"AmuletHash"`
  SellingPrice string `json:"SellingPrice" yaml:"SellingPrice"`
  SellingBy string `json:"SellingBy" yaml:"SellingBy"`
}