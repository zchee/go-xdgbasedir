// Copyright 2017 The go-xdgbasedir Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package xdgbasedir

import (
	"os"
	"path/filepath"
	"strconv"
	"testing"

	"github.com/zchee/go-xdgbasedir/home"
)

func TestDataHome(t *testing.T) {
	defaultDataHome := filepath.Join(home.Dir(), ".local", "share")

	tests := []struct {
		name string
		env  string
		want string
	}{
		{
			name: "Set env based specification",
			env:  defaultDataHome,
			want: defaultDataHome,
		},
		{
			name: "Set env based different from specification",
			env:  filepath.Join(string(filepath.Separator), "tmp", "xdg", ".local", "share"),
			want: filepath.Join(string(filepath.Separator), "tmp", "xdg", ".local", "share"),
		},
		{
			name: "Empty env",
			env:  "",
			want: defaultDataHome,
		},
	}
	for _, tt := range tests {
		os.Unsetenv("XDG_DATA_HOME")
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
	defaultConfigHome := filepath.Join(home.Dir(), ".config")

	tests := []struct {
		name string
		env  string
		want string
	}{
		{
			name: "Set env based specification",
			env:  defaultConfigHome,
			want: defaultConfigHome,
		},
		{
			name: "Set env based different from specification",
			env:  filepath.Join(string(filepath.Separator), "tmp", "config"),
			want: filepath.Join(string(filepath.Separator), "tmp", "config"),
		},
		{
			name: "Empty env",
			env:  "",
			want: defaultConfigHome,
		},
	}
	for _, tt := range tests {
		os.Unsetenv("XDG_CONFIG_HOME")
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
			name: "Empty env",
			env:  "",
			want: defaultDataDirs,
		},
	}
	for _, tt := range tests {
		os.Unsetenv("XDG_DATA_DIRS")
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
			name: "Empty env",
			env:  "",
			want: defaultConfigDirs,
		},
	}
	for _, tt := range tests {
		os.Unsetenv("XDG_CONFIG_DIRS")
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
	defaultCacheHome := filepath.Join(home.Dir(), ".cache")
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
			name: "Empty env",
			env:  "",
			want: defaultCacheHome,
		},
	}
	for _, tt := range tests {
		os.Unsetenv("XDG_CACHE_HOME")
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
	defaultRuntimeDir := filepath.Join(string(filepath.Separator), "run", "user", strconv.Itoa(os.Getuid()))

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
			name: "Empty env",
			env:  "",
			want: defaultRuntimeDir,
		},
	}
	for _, tt := range tests {
		os.Unsetenv("XDG_RUNTIME_DIR")
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
