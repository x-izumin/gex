package manager

import (
	newManager "github.com/x-izumin/gex/pkg/manager"
)

// Executor is an interface for executing managers.
type Executor = newManager.Executor

// NewExecutor creates a new Executor instance.
var NewExecutor = newManager.NewExecutor
