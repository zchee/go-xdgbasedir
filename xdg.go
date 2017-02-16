// Copyright 2017 The go-xdg Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package xdg implements a freedesktop.org XDG Base Directory Specification.
// XDG Base Directory Specification:
//  https://specifications.freedesktop.org/basedir-spec/basedir-spec-0.8.html
//
// The XDG Base Directory Specification is based on the following concepts:
//
// There is a single base directory relative to which user-specific data files should be written. This directory is defined by the environment variable $XDG_DATA_HOME.
//
// There is a single base directory relative to which user-specific configuration files should be written. This directory is defined by the environment variable $XDG_CONFIG_HOME.
//
// There is a set of preference ordered base directories relative to which data files should be searched. This set of directories is defined by the environment variable $XDG_DATA_DIRS.
//
// There is a set of preference ordered base directories relative to which configuration files should be searched. This set of directories is defined by the environment variable $XDG_CONFIG_DIRS.
//
// There is a single base directory relative to which user-specific non-essential (cached) data should be written. This directory is defined by the environment variable $XDG_CACHE_HOME.
package xdg

import (
	"log"
	"os"
	"os/user"
	"path/filepath"
	"runtime"
)

var usr = &user.User{}

func init() {
	cUser, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	usr = cUser
}

// DataHome return the XDG_DATA_HOME based directory path.
//
// $XDG_DATA_HOME defines the base directory relative to which user specific data files should be stored.
// If $XDG_DATA_HOME is either not set or empty, a default equal to $HOME/.local/share should be used.
func DataHome() string {
	dataHome := os.Getenv("XDG_DATA_HOME")
	if dataHome == "" {
		dataHome = filepath.Join(homeDir(), ".local", "share")
	}
	return dataHome
}

// ConfigHome return the XDG_CONFIG_HOME based directory path.
//
// $XDG_CONFIG_HOME defines the base directory relative to which user specific configuration files should be stored.
// If $XDG_CONFIG_HOME is either not set or empty, a default equal to $HOME/.config should be used.
func ConfigHome() string {
	configHome := os.Getenv("XDG_CONFIG_HOME")
	if configHome == "" {
		configHome = filepath.Join(homeDir(), ".config")
	}
	return configHome
}

// DataDirs return the XDG_DATA_DIRS based directory path.
//
// $XDG_DATA_DIRS defines the preference-ordered set of base directories to search for data files in addition
// to the $XDG_DATA_HOME base directory. The directories in $XDG_DATA_DIRS should be seperated with a colon ':'.
// If $XDG_DATA_DIRS is either not set or empty, a value equal to /usr/local/share/:/usr/share/ should be used.
func DataDirs() string {
	dataDirs := os.Getenv("XDG_DATA_DIRS")
	if dataDirs == "" {
		dataDirs = filepath.Join("usr", "local", "share", string(filepath.ListSeparator), "usr", "share")
	}
	return dataDirs
}

// ConfigDirs return the XDG_CONFIG_DIRS based directory path.
//
// $XDG_CONFIG_DIRS defines the preference-ordered set of base directories to search for configuration files in addition
// to the $XDG_CONFIG_HOME base directory. The directories in $XDG_CONFIG_DIRS should be seperated with a colon ':'.
// If $XDG_CONFIG_DIRS is either not set or empty, a value equal to /etc/xdg should be used.
func ConfigDirs() string {
	configDirs := os.Getenv("XDG_CONFIG_DIRS")
	if configDirs == "" {
		configDirs = filepath.Join("etc", "xdg")
	}
	return configDirs
}

// CacheHome return the XDG_CACHE_HOME based directory path.
//
// $XDG_CACHE_HOME defines the base directory relative to which user specific non-essential data files should be stored.
// If $XDG_CACHE_HOME is either not set or empty, a default equal to $HOME/.cache should be used.
func CacheHome() string {
	cacheHome := os.Getenv("XDG_CACHE_HOME")
	if cacheHome == "" {
		cacheHome = filepath.Join(homeDir(), ".cache")
	}
	return cacheHome
}

// RuntimeDir return the XDG_RUNTIME_DIR based directory path.
//
// $XDG_RUNTIME_DIR defines the base directory relative to which user-specific non-essential runtime files and
// other file objects (such as sockets, named pipes, ...) should be stored. The directory MUST be owned by the user,
// and he MUST be the only one having read and write access to it. Its Unix access mode MUST be 0700.
func RuntimeDir() string {
	runtimeDir := os.Getenv("XDG_RUNTIME_DIR")
	if runtimeDir == "" {
		runtimeDir = filepath.Join("run", "user", usr.Uid)
	}
	return runtimeDir
}

func homeDir() string {
	if runtime.GOOS == "windows" {
		home := filepath.Join(os.Getenv("HOMEDRIVE"), os.Getenv("HOMEPATH"))
		if home == "" {
			home = os.Getenv("USERPROFILE")
		}
		return home
	}

	homeDir := os.Getenv("HOME")
	if homeDir != "" {
		homeDir = usr.HomeDir
	}

	return homeDir
}
