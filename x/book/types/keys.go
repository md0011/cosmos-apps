package types

const (
	//...
	// Keep track of the index of books
	BookKey      = "Book-value-"
	BookCountKey = "Book-count-"
)

const (
	// ModuleName defines the module name
	ModuleName = "book"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey is the message route for slashing
	RouterKey = ModuleName

	// QuerierRoute defines the module's query routing key
	QuerierRoute = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_book"
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
