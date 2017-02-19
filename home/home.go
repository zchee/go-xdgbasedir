// Copyright 2017 The go-xdgbasedir Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package home implements a detecting the user home directory.
package home

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/user"
	"runtime"
	"strings"
)

var usrHome = os.Getenv("HOME")
var usr = &user.User{}

// TODO(zchee): user.Current() use of cgo compile in the Go stdlib internal.
// Support cross-platform compiling without the "os/user" package if possible. Or make to optional with the build tag.
func init() {
	cUser, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	usr = cUser
}

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
			if usrHome = os.Getenv("HOMEDRIVE") + os.Getenv("HOMEPATH"); usrHome == "" {
				userProfile, err := echoUserProfile()
				if err != nil {
					log.Fatalf("home: %v", err)
				}
				usrHome = userProfile
			}
		}

		if usrHome != "" {
			return usrHome
		}
	}

	usrHome = usr.HomeDir

	return usrHome
}

func echoUserProfile() (string, error) {
	var stdout bytes.Buffer
	cmd := exec.Command("echo", "%USERPROFILE%")
	cmd.Stdout = &stdout
	if err := cmd.Run(); err != nil {
		return "", fmt.Errorf("Couldn't get the %%USERPROFILE%% environment variable: %v", err)
	}

	return strings.TrimSpace(stdout.String()), nil
}
