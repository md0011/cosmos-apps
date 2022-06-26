package keeper

import (
	"context"

	"book/x/book/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CreateBook(goCtx context.Context, msg *types.MsgCreateBook) (*types.MsgCreateBookResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	book := types.Book{
		Creator: msg.Creator,
		Author:  msg.Author,
		Title:   msg.Title,
	}

	count := k.AppendBook(ctx, book)
	// TODO: Handling the message
	_ = ctx

	return &types.MsgCreateBookResponse{Id: count}, nil
}
