package dep

import (
	newDep "github.com/x-izumin/gex/pkg/manager/dep"
)

// NewManager creates a manager.Interface instance to manage tools vendored with dep.
var NewManager = newDep.NewManager
