//go:build !go1.13
// +build !go1.13

package main

import (
	"github.com/x-izumin/gex/pkg/tool"
)

func asBuildErrors(err error) *tool.BuildErrors {
	if errs, ok := err.(*tool.BuildErrors); ok {
		return errs
	}
	return nil
}
