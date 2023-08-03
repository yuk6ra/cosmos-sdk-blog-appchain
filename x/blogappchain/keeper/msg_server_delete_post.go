package keeper

import (
	"context"

	"blog-appchain/x/blogappchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) DeletePost(goCtx context.Context, msg *types.MsgDeletePost) (*types.MsgDeletePostResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgDeletePostResponse{}, nil
}
