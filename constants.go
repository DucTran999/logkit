package logkit

// Define a private type to avoid collisions
type contextKey string

const (
	// Env
	Development = "develop"
	Testing     = "testing"
	Staging     = "staging"
	Production  = "production"

	// key extract from context
	RequestIDKeyCtx = contextKey("request-id")
)
