// Package pipeline provides a generic pipe interface.
package pipeline

import "github.com/lordnynex/goreleaser/context"

// Pipe interface
type Pipe interface {
	// Name of the pipe
	Description() string

	// Run the pipe
	Run(ctx *context.Context) error
}
