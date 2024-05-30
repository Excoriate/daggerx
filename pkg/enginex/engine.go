package enginex

import (
	"context"
	"os"

	"dagger.io/dagger"
	"github.com/Excoriate/daggerx/pkg/logger"
)

// Engine is an interface that represents a Dagger engine.
// It provides methods for starting the engine and getting the engine instance.
//
// The Start method is used to start the engine and returns a pointer to a dagger.Client instance.
// The GetEngine method is used to get the engine instance.
type Engine interface {
	Start(ctx context.Context, options ...dagger.ClientOpt) (*dagger.Client, error)
	GetEngine() *dagger.Client
}

// DaggerEngine is a struct that implements the Engine interface.
//
// It contains a logger.Log instance and a dagger.Client instance.
// The logger.Log instance is used for logging, and the dagger.Client instance is used for
// interacting with the Dagger engine.
//
// The Start method is used to start the Dagger engine and returns a pointer to a dagger.Client
// instance.
type DaggerEngine struct {
	l logger.Log
	c *dagger.Client
}

// New returns a new DaggerEngine instance.
//
// Parameters:
//   - l: A logger instance to use for logging.
//
// Returns:
//   - A new DaggerEngine instance.
func New() Engine {
	return &DaggerEngine{
		l: logger.NewLogger().Logger,
	}
}

// Start starts the Dagger engine and returns a pointer to a dagger.Client instance.
// It takes a context and a list of options as input.
//
// Parameters:
//   - ctx: A context to use for the operation.
//   - options: A list of options to pass to the Dagger engine.
//
// Returns:
//   - A pointer to a dagger.Client instance.
func (d *DaggerEngine) Start(ctx context.Context, options ...dagger.ClientOpt) (*dagger.
	Client, error) {
	var c context.Context
	if ctx == nil {
		c = context.Background()
	} else {
		c = ctx
	}

	var daggerOptions []dagger.ClientOpt

	if len(options) == 0 {
		daggerOptions = append(daggerOptions, dagger.WithLogOutput(os.Stdout))
	} else {
		daggerOptions = append(daggerOptions, options...)
	}

	daggerClient, err := dagger.Connect(c, daggerOptions...)
	if err != nil {
		return nil, err
	}

	d.c = daggerClient
	return daggerClient, nil
}

// GetEngine returns the Dagger engine instance.
//
// Returns:
//   - A pointer to a dagger.Client instance.
//
// Example:
//
//	engine := enginex.New()
//	client, err := engine.Start(context.Background())
//	if err != nil {
//	    // handle error
//	}
//	engine.GetEngine() // Returns a pointer to a dagger.Client instance
func (d *DaggerEngine) GetEngine() *dagger.Client {
	return d.c
}
