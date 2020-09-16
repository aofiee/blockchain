package rest

import (
	"net/http"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/cosmos/cosmos-sdk/x/auth/client/utils"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/aofiee/amulet/x/amulet/types"
)

type createBuyingRequest struct {
	BaseReq rest.BaseReq `json:"base_req"`
	Creator string `json:"creator"`
	AmuletHash string `json:"AmuletHash"`
	BuyingPrice string `json:"BuyingPrice"`
	SellingBy string `json:"SellingBy"`
	BuyingBy string `json:"BuyingBy"`
	
}

func createBuyingHandler(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req createBuyingRequest
		if !rest.ReadRESTReq(w, r, cliCtx.Codec, &req) {
			rest.WriteErrorResponse(w, http.StatusBadRequest, "failed to parse request")
			return
		}
		baseReq := req.BaseReq.Sanitize()
		if !baseReq.ValidateBasic(w) {
			return
		}
		creator, err := sdk.AccAddressFromBech32(req.Creator)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}
		msg := types.NewMsgCreateBuying(creator,  req.AmuletHash,  req.BuyingPrice,  req.SellingBy,  req.BuyingBy, )
		utils.WriteGenerateStdTxResponse(w, cliCtx, baseReq, []sdk.Msg{msg})
	}
}
