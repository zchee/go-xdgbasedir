// Copyright 2017 The go-xdgbasedir Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package xdgbasedir

import (
	"flag"
	"os"
	"path/filepath"
	"testing"
)

var testUserName = "gopher"
var testHomeDir = filepath.FromSlash(filepath.Join(string(filepath.Separator), "home", testUserName))

func TestMain(m *testing.M) {
	flag.Parse()

	usrHome = testHomeDir
	os.Exit(m.Run())
}

func TestDataHome(t *testing.T) {
	tests := []struct {
		name string
		env  string
		want string
	}{
		{
			name: "Set env based specification",
			env:  filepath.Join(testHomeDir, ".local", "share"),
			want: filepath.Join(testHomeDir, ".local", "share"),
		},
		{
			name: "Set env based different from specification",
			env:  filepath.Join(string(filepath.Separator), "tmp", "xdg", ".local", "share"),
			want: filepath.Join(string(filepath.Separator), "tmp", "xdg", ".local", "share"),
		},
		{
			name: "empty env",
			env:  "",
			want: filepath.Join(testHomeDir, ".local", "share"),
		},
	}
	for _, tt := range tests {
		os.Clearenv()
		if tt.env != "" {
			os.Setenv("XDG_DATA_HOME", tt.env)
		}
		t.Run(tt.name, func(t *testing.T) {
			if got := DataHome(); got != tt.want {
				t.Errorf("DataHome() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConfigHome(t *testing.T) {
	tests := []struct {
		name string
		env  string
		want string
	}{
		{
			name: "Set env based specification",
			env:  filepath.Join(testHomeDir, ".config"),
			want: filepath.Join(testHomeDir, ".config"),
		},
		{
			name: "Set env based different from specification",
			env:  filepath.Join(string(filepath.Separator), "tmp", "config"),
			want: filepath.Join(string(filepath.Separator), "tmp", "config"),
		},
		{
			name: "empty env",
			env:  "",
			want: filepath.Join(testHomeDir, ".config"),
		},
	}
	for _, tt := range tests {
		os.Clearenv()
		if tt.env != "" {
			os.Setenv("XDG_CONFIG_HOME", tt.env)
		}
		t.Run(tt.name, func(t *testing.T) {
			if got := ConfigHome(); got != tt.want {
				t.Errorf("ConfigHome() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDataDirs(t *testing.T) {
	defaultDataDirs := filepath.Join(string(filepath.Separator), "usr", "local", "share", string(filepath.ListSeparator), "usr", "share")
	tests := []struct {
		name string
		env  string
		want string
	}{
		{
			name: "Set env based specification",
			env:  defaultDataDirs,
			want: defaultDataDirs,
		},
		{
			name: "Set env based different from specification",
			env:  filepath.Join(string(filepath.Separator), "opt", "local", "share"),
			want: filepath.Join(string(filepath.Separator), "opt", "local", "share"),
		},
		{
			name: "empty env",
			env:  "",
			want: defaultDataDirs,
		},
	}
	for _, tt := range tests {
		os.Clearenv()
		if tt.env != "" {
			os.Setenv("XDG_DATA_DIRS", tt.env)
		}
		t.Run(tt.name, func(t *testing.T) {
			if got := DataDirs(); got != tt.want {
				t.Errorf("DataDirs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConfigDirs(t *testing.T) {
	defaultConfigDirs := filepath.Join(string(filepath.Separator), "etc", "xdg")
	tests := []struct {
		name string
		env  string
		want string
	}{
		{
			name: "Set env based specification",
			env:  defaultConfigDirs,
			want: defaultConfigDirs,
		},
		{
			name: "Set env based different from specification",
			env:  filepath.Join(string(filepath.Separator), "var", "etc", "xdg"),
			want: filepath.Join(string(filepath.Separator), "var", "etc", "xdg"),
		},
		{
			name: "empty env",
			env:  "",
			want: defaultConfigDirs,
		},
	}
	for _, tt := range tests {
		os.Clearenv()
		if tt.env != "" {
			os.Setenv("XDG_CONFIG_DIRS", tt.env)
		}
		t.Run(tt.name, func(t *testing.T) {
			if got := ConfigDirs(); got != tt.want {
				t.Errorf("ConfigDirs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCacheHome(t *testing.T) {
	defaultCacheHome := filepath.Join(testHomeDir, ".cache")
	tests := []struct {
		name string
		env  string
		want string
	}{
		{
			name: "Set env based specification",
			env:  defaultCacheHome,
			want: defaultCacheHome,
		},
		{
			name: "Set env based different from specification",
			env:  filepath.Join(string(filepath.Separator), "tmp", "cache"),
			want: filepath.Join(string(filepath.Separator), "tmp", "cache"),
		},
		{
			name: "empty env",
			env:  "",
			want: defaultCacheHome,
		},
	}
	for _, tt := range tests {
		os.Clearenv()
		if tt.env != "" {
			os.Setenv("XDG_CACHE_HOME", tt.env)
		}
		t.Run(tt.name, func(t *testing.T) {
			if got := CacheHome(); got != tt.want {
				t.Errorf("CacheHome() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRuntimeDir(t *testing.T) {
	usr.Uid = "1000"
	defaultRuntimeDir := filepath.Join(string(filepath.Separator), "run", "user", "1000")

	tests := []struct {
		name string
		env  string
		want string
	}{
		{
			name: "Set env based specification",
			env:  defaultRuntimeDir,
			want: defaultRuntimeDir,
		},
		{
			name: "Set env based different from specification",
			env:  filepath.Join(string(filepath.Separator), "tmp", "user", "1000"),
			want: filepath.Join(string(filepath.Separator), "tmp", "user", "1000"),
		},
		{
			name: "empty env",
			env:  "",
			want: defaultRuntimeDir,
		},
	}
	for _, tt := range tests {
		os.Clearenv()
		if tt.env != "" {
			os.Setenv("XDG_RUNTIME_DIR", tt.env)
		}
		t.Run(tt.name, func(t *testing.T) {
			if got := RuntimeDir(); got != tt.want {
				t.Errorf("RuntimeDir() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_homeDir(t *testing.T) {
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
			name: "assign different usrHome",
			env:  "",
			home: filepath.FromSlash(filepath.Join(string(filepath.Separator), "Users", testUserName)),
			want: filepath.FromSlash(filepath.Join(string(filepath.Separator), "Users", testUserName)),
		},
		{
			name: "empty HOME env",
			env:  "empty",
			home: filepath.FromSlash(filepath.Join(string(filepath.Separator), "home", "empty")),
			want: filepath.FromSlash(filepath.Join(string(filepath.Separator), "home", "empty")),
		},
	}
	for _, tt := range tests {
		os.Clearenv()

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
			if got := homeDir(); got != tt.want {
				t.Errorf("homeDir() = %v, want %v", got, tt.want)
			}
		})
	}
}
