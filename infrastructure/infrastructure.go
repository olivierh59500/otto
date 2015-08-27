package infrastructure

import (
	"github.com/hashicorp/otto/appfile"
	"github.com/hashicorp/otto/directory"
	"github.com/hashicorp/otto/ui"
)

// Infrastructure is an interface that must be implemented by each
// infrastructure type with a method of creating it.
type Infrastructure interface {
	// Creds is called when Otto determines that it needs credentials
	// for this infrastructure provider. The Infra should query the
	// user (or environment) for creds and return them. Otto will
	// handle encrypting, storing, and retrieving the credentials.
	Creds(*Context) (map[string]string, error)

	Execute(*Context) error
	Compile(*Context) (*CompileResult, error)
	Flavors() []string
}

// Context is the context for operations on infrastructures. Some of
// the fields in this struct are only available for certain operations.
type Context struct {
	// Action is the sub-action to take when being executed.
	//
	// ActionArgs is the list of arguments for this action.
	//
	// Both of these fields will only be set for the Execute call.
	Action     string
	ActionArgs []string

	// Dir is the directory that the compilation is allowed to write to
	// for persistant storage of data. For other tasks, this will be the
	// directory that was already populated by compilation.
	Dir string

	// The infrastructure configuration itself from the Appfile. This includes
	// the flavor of the infrastructure we want to launch.
	Infra *appfile.Infrastructure

	// Ui is the Ui object that can be used to communicate with the user.
	Ui ui.Ui

	// Directory is the directory service. This is available during
	// both execution and compilation and can be used to view the
	// global data prior to doing anything.
	Directory directory.Backend
}

// CompileResult is the structure containing compilation result values.
type CompileResult struct {
}
