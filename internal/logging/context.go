// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package logging

import (
	"context"
	"strings"
)

type pathstring string

const path pathstring = "path"

type Path []string

// SetPathInContext is used to maintain a path indicating the current
// location within the schema that is being processed.
func SetPathInContext(ctx context.Context, pathStep string) context.Context {
	if v, ok := ctx.Value(path).(Path); ok {
		return context.WithValue(ctx, path, append(v, pathStep))
	}

	return context.WithValue(ctx, path, Path{pathStep})
}

// GetPathFromContext returns a dot-separated path.
func GetPathFromContext(ctx context.Context) string {
	if v, ok := ctx.Value(path).(Path); ok {
		return strings.Join(v, ".")
	}

	return ""
}
