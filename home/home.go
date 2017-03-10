// Copyright 2017 The go-xdgbasedir Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package home implements a detecting the user home directory.
package home

import (
	"os"
	"runtime"
)

var usrHome = os.Getenv("HOME")

// Dir detects and returns the user home directory.
func Dir() string {
	// At first, Check the $HOME environment variable
	if usrHome != "" {
		return usrHome
	}

	// TODO(zchee): In Windows OS, which of $HOME and these checks has priority?
	if runtime.GOOS == "windows" {
		// Respect the USERPROFILE environment variable because Go stdlib uses it for default GOPATH in the "go/build" package.
		if usrHome = os.Getenv("USERPROFILE"); usrHome == "" {
			usrHome = os.Getenv("HOMEDRIVE") + os.Getenv("HOMEPATH")
		}
	}

	return usrHome
}
