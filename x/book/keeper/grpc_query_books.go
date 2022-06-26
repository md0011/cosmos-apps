package keeper

import (
	"context"

	"book/x/book/types"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) Books(goCtx context.Context, req *types.QueryBooksRequest) (*types.QueryBooksResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var books []*types.Book
	ctx := sdk.UnwrapSDKContext(goCtx)
	store := ctx.KVStore(k.storeKey)
	bookStore := prefix.NewStore(store, []byte(types.BookKey))
	pageRes, err := query.Paginate(bookStore, req.Pagination, func(key, value []byte) error {
		var book types.Book
		if err := k.cdc.Unmarshal(value, &book); err != nil {
			return err
		}
		books = append(books, &book)
		return nil
	})

	// Throw an error if pagination failed
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	// TODO: Process the query
	_ = ctx

	return &types.QueryBooksResponse{Book: books, Pagination: pageRes}, nil
}
