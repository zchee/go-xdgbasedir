// Copyright 2017 The go-xdgbasedir Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package home

import (
	"flag"
	"os"
	"path/filepath"
	"runtime"
	"testing"
)

var testUserName = "gopher"
var testHomeDir = filepath.FromSlash(filepath.Join("/Users", testUserName))

func TestMain(m *testing.M) {
	flag.Parse()

	if runtime.GOOS == "windows" {
		// TODO(zchee): what?
		testUserName = os.Getenv("USERNAME")
		testHomeDir = filepath.FromSlash(filepath.Join("C:/Users", testUserName))
	}

	usrHome = testHomeDir
	os.Exit(m.Run())
}

func TestDir(t *testing.T) {
	tests := []struct {
		name string
		env  string
		home string
		want string
	}{
		{
			name: "default usrHome",
			env:  "",
			home: "",
			want: testHomeDir,
		},
		{
			name: "Assign different usrHome",
			env:  "",
			home: filepath.FromSlash(filepath.Join("/Users", testUserName)),
			want: filepath.FromSlash(filepath.Join("/Users", testUserName)),
		},
		{
			name: "Set different $HOME env",
			env:  filepath.FromSlash(filepath.Join("/tmp", "home")),
			home: filepath.FromSlash(filepath.Join("/Users", testUserName)),
			want: filepath.FromSlash(filepath.Join("/Users", testUserName)),
		},
		{
			name: "Empty HOME env",
			env:  "",
			home: "empty",
			want: testHomeDir,
		},
	}
	for _, tt := range tests {
		switch tt.env {
		case "":
			// nothing to do
		case "empty":
			os.Unsetenv("HOME")
		default:
			os.Setenv("HOME", tt.env)
		}

		switch tt.home {
		case "":
			// nothing to do
		case "empty":
			usrHome = ""
			usr.HomeDir = testHomeDir
		default:
			usrHome = tt.home
		}

		t.Run(tt.name, func(t *testing.T) {
			if got := Dir(); got != tt.want {
				t.Errorf("Dir() = %v, want %v", got, tt.want)
			}
		})
	}
}
