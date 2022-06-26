package keeper

import (
	"book/x/book/types"
	"encoding/binary"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) GetBookCount(ctx sdk.Context) uint64 {
	byteKey := []byte(types.BookCountKey)
	store := prefix.NewStore(ctx.KVStore(k.storeKey), byteKey)
	bz := store.Get(byteKey)

	if bz == nil {
		return 0
	}

	return binary.BigEndian.Uint64(bz)
}

func (k Keeper) SetBookCount(ctx sdk.Context, count uint64) {
	byteKey := []byte(types.BookCountKey)
	store := prefix.NewStore(ctx.KVStore(k.storeKey), byteKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

func (k Keeper) AppendBook(ctx sdk.Context, book types.Book) uint64 {
	count := k.GetBookCount(ctx)
	book.Id = count
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.BookKey))
	byteKey := make([]byte, 8)

	binary.BigEndian.PutUint64(byteKey, book.Id)

	appendedValue := k.cdc.MustMarshal(&book)
	store.Set(byteKey, appendedValue)
	k.SetBookCount(ctx, count+1)
	return count
}
