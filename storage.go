package osin

import (
	"golang.org/x/net/context"
)

// Storage interface
type Storage interface {
	// Clone the storage if needed. For example, using mgo, you can clone the session with session.Clone
	// to avoid concurrent access problems.
	// This is to avoid cloning the connection at each method access.
	// Can return itself if not a problem.
	Clone() Storage

	// Close the resources the Storage potentially holds (using Clone for example)
	Close()

	// GetClient loads the client by id (client_id)
	GetClient(c context.Context, id string) (Client, error)

	// SaveAuthorize saves authorize data.
	SaveAuthorize(c context.Context, data *AuthorizeData) error

	// LoadAuthorize looks up AuthorizeData by a code.
	// Client information MUST be loaded together.
	// Optionally can return error if expired.
	LoadAuthorize(c context.Context, code string) (*AuthorizeData, error)

	// RemoveAuthorize revokes or deletes the authorization code.
	RemoveAuthorize(c context.Context, code string) error

	// SaveAccess writes AccessData.
	// If RefreshToken is not blank, it must save in a way that can be loaded using LoadRefresh.
	SaveAccess(c context.Context, data *AccessData) error

	// LoadAccess retrieves access data by token. Client information MUST be loaded together.
	// AuthorizeData and AccessData DON'T NEED to be loaded if not easily available.
	// Optionally can return error if expired.
	LoadAccess(c context.Context, token string) (*AccessData, error)

	// RemoveAccess revokes or deletes an AccessData.
	RemoveAccess(c context.Context, token string) error

	// LoadRefresh retrieves refresh AccessData. Client information MUST be loaded together.
	// AuthorizeData and AccessData DON'T NEED to be loaded if not easily available.
	// Optionally can return error if expired.
	LoadRefresh(c context.Context, token string) (*AccessData, error)

	// RemoveRefresh revokes or deletes refresh AccessData.
	RemoveRefresh(c context.Context, token string) error
}
