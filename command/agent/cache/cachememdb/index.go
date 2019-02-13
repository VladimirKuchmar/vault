package cachememdb

import "context"

type ContextInfo struct {
	Ctx        context.Context
	CancelFunc context.CancelFunc
	DoneCh     chan struct{}
}

// Index holds the response to be cached along with multiple other values that
// serve as pointers to refer back to this index.
type Index struct {
	// ID is a value that uniquely represents the request held by this
	// index. This is computed by serializing and hashing the response object.
	// Required: true, Unique: true
	ID string

	// Token is the token that fetched the response held by this index
	// Required: true, Unique: false
	Token string

	// TokenParent is the parent token of the token held by this index
	// Required: false, Unique: false
	TokenParent string

	// TokenAccessor is the accessor of the token being cached in this index
	// Required: true, Unique: false
	TokenAccessor string

	// Namespace is the namespace that was provided in the request path as the
	// Vault namespace to query
	Namespace string

	// RequestPath is the path of the request that resulted in the response
	// held by this index.
	// Required: true, Unique: false
	RequestPath string

	// Lease is the identifier of the lease in Vault, that belongs to the
	// response held by this index.
	// Required: false, Unique: true
	Lease string

	// Response is the serialized response object that the agent is caching.
	Response []byte

	// RenewCtxInfo holds the context and the corresponding cancel func for the
	// goroutine that manages the renewal of the secret belonging to the
	// response in this index.
	RenewCtxInfo *ContextInfo
}

type IndexName uint32

const (
	// IndexNameID is the ID of the index constructed from the serialized request.
	IndexNameID = "id"

	// IndexNameLease is the lease of the index.
	IndexNameLease = "lease"

	// IndexNameRequestPath is the request path of the index.
	IndexNameRequestPath = "request_path"

	// IndexNameToken is the token of the index.
	IndexNameToken = "token"

	// IndexNameTokenAccessor is the token accessor of the index.
	IndexNameTokenAccessor = "token_accessor"

	// IndexNameTokenParent is the token parent of the index.
	IndexNameTokenParent = "token_parent"
)

func validIndexName(indexName string) bool {
	switch indexName {
	case "id":
	case "lease":
	case "request_path":
	case "token":
	case "token_accessor":
	case "token_parent":
	default:
		return false
	}
	return true
}
