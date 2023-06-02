// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package input

import (
	"io"
	"os"
)

func Read(path string) ([]byte, error) {
	if path == "" {
		stdin, err := io.ReadAll(os.Stdin)
		if err != nil {
			return nil, err
		}

		return stdin, nil
	}

	src, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	return src, nil
}
