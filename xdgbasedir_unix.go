// Copyright 2017 The go-xdgbasedir Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build !darwin
// +build !windows

package xdgbasedir

import (
	"os"
	"path/filepath"
	"strconv"
)

var home string

func init() {
	home, _ = os.UserHomeDir()
}

var (
	defaultDataHome   = filepath.Join(home, ".local", "share")
	defaultConfigHome = filepath.Join(home, ".config")
	defaultDataDirs   = filepath.Join("/usr", "local", "share") + string(filepath.ListSeparator) + filepath.Join("/usr", "share")
	defaultConfigDirs = filepath.Join("/etc", "xdg")
	defaultCacheHome  = filepath.Join(home, ".cache")
	defaultRuntimeDir = filepath.Join("/run", "user", strconv.Itoa(os.Getuid()))
)

func dataHome() string {
	return defaultDataHome
}

func configHome() string {
	return defaultConfigHome
}

func dataDirs() string {
	return defaultDataDirs
}

func configDirs() string {
	return defaultConfigDirs
}

func cacheHome() string {
	return defaultCacheHome
}

func runtimeDir() string {
	return defaultRuntimeDir
}
