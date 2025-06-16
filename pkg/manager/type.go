package manager

import (
	newManager "github.com/x-izumin/gex/pkg/manager"
)

// Type represents the dependencies management tool that is used.
type Type = newManager.Type

// Type values
const (
	TypeUnknown = newManager.TypeUnknown
	TypeModules = newManager.TypeModules
	TypeDep     = newManager.TypeDep
)

// DetectType detects a current Mode and sets a root directory.
var DetectType = newManager.DetectType

// FindRoot gets a manifest file path.
var FindRoot = newManager.FindRoot
