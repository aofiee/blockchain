package rest

import (
	"github.com/gorilla/mux"

	"github.com/cosmos/cosmos-sdk/client/context"
)

// RegisterRoutes registers amulet-related REST handlers to a router
func RegisterRoutes(cliCtx context.CLIContext, r *mux.Router) {
  // this line is used by starport scaffolding
	r.HandleFunc("/amulet/buying", listBuyingHandler(cliCtx, "amulet")).Methods("GET")
	r.HandleFunc("/amulet/buying", createBuyingHandler(cliCtx)).Methods("POST")
	r.HandleFunc("/amulet/selling", listSellingHandler(cliCtx, "amulet")).Methods("GET")
	r.HandleFunc("/amulet/selling", createSellingHandler(cliCtx)).Methods("POST")
	r.HandleFunc("/amulet/amulet", listAmuletHandler(cliCtx, "amulet")).Methods("GET")
	r.HandleFunc("/amulet/amulet", createAmuletHandler(cliCtx)).Methods("POST")
}
