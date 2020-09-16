package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type Buying struct {
	Creator sdk.AccAddress `json:"creator" yaml:"creator"`
	ID      string         `json:"id" yaml:"id"`
  AmuletHash string `json:"AmuletHash" yaml:"AmuletHash"`
  BuyingPrice string `json:"BuyingPrice" yaml:"BuyingPrice"`
  SellingBy string `json:"SellingBy" yaml:"SellingBy"`
  BuyingBy string `json:"BuyingBy" yaml:"BuyingBy"`
}