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

var testHomeDir string

func TestMain(m *testing.M) {
	flag.Parse()

	switch runtime.GOOS {
	case "darwin":
		testHomeDir = filepath.Join("/Users", os.Getenv("USER"))
	case "linux":
		testHomeDir = filepath.Join("/home", os.Getenv("USER"))
	case "windows":
		testHomeDir = filepath.FromSlash(filepath.Join("C:/Users", os.Getenv("USERNAME")))
	}

	os.Exit(m.Run())
}

func TestDir(t *testing.T) {
	tests := []struct {
		name string
		env  string
		want string
	}{
		{
			name: "default usrHome",
			env:  "",
			want: testHomeDir,
		},
		{
			name: "Set different $HOME env",
			env:  filepath.FromSlash(filepath.Join("/tmp", "home")),
			want: filepath.FromSlash(filepath.Join("/tmp", "home")),
		},
		{
			name: "Empty HOME env",
			env:  "empty",
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

		t.Run(tt.name, func(t *testing.T) {
			if got := Dir(); got != tt.want {
				t.Errorf("Dir() = %v, want %v", got, tt.want)
			}
		})
	}
}
