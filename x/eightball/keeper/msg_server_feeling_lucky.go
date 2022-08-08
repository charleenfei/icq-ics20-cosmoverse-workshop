package keeper

import (
	"context"

	"github.com/charleenfei/icq-ics20-cosmoverse-workshop/x/eightball/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) FeelingLucky(goCtx context.Context, msg *types.MsgFeelingLucky) (*types.MsgFeelingLuckyResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// check if the offering is enough
		// send icq request (SendPacket)
	// if offering is enough
	_ = ctx

	return &types.MsgFeelingLuckyResponse{}, nil
}
