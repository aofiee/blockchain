package rest

import (
	"net/http"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/cosmos/cosmos-sdk/x/auth/client/utils"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/aofiee/amulet/x/amulet/types"
)

type createAmuletRequest struct {
	BaseReq rest.BaseReq `json:"base_req"`
	Creator string `json:"creator"`
	Name string `json:"Name"`
	Identity string `json:"Identity"`
	Address string `json:"Address"`
	Telephon string `json:"Telephon"`
	AmuletType string `json:"AmuletType"`
	AmuletName string `json:"AmuletName"`
	From string `json:"From"`
	Model string `json:"Model"`
	Surface string `json:"Surface"`
	Year string `json:"Year"`
	Province string `json:"Province"`
	Description string `json:"Description"`
	Remark string `json:"Remark"`
	
}

func createAmuletHandler(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req createAmuletRequest
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
		msg := types.NewMsgCreateAmulet(creator,  req.Name,  req.Identity,  req.Address,  req.Telephon,  req.AmuletType,  req.AmuletName,  req.From,  req.Model,  req.Surface,  req.Year,  req.Province,  req.Description,  req.Remark, )
		utils.WriteGenerateStdTxResponse(w, cliCtx, baseReq, []sdk.Msg{msg})
	}
}
