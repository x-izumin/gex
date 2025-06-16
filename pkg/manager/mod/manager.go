package mod

import (
	newMod "github.com/x-izumin/gex/pkg/manager/mod"
)

// NewManager creates a manager.Interface instance to build tools vendored with Modules.
var NewManager = newMod.NewManager
