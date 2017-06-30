// Copyright 2017 The go-xdgbasedir Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package xdgbasedir

import (
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"sync"
	"testing"

	"github.com/zchee/go-xdgbasedir/home"
)

func TestDataHome(t *testing.T) {
	var testDefaultDataHome string
	switch runtime.GOOS {
	case "windows":
		testDefaultDataHome = filepath.Join(home.Dir(), "AppData", "Local")
	default:
		testDefaultDataHome = filepath.Join(home.Dir(), ".local", "share")
	}

	tests := []struct {
		name string
		env  string
		want string
		mode mode
	}{
		{
			name: "set env based specification",
			env:  testDefaultDataHome,
			want: testDefaultDataHome,
		},
		{
			name: "set env based different from specification",
			env:  filepath.Join(string(filepath.Separator), "tmp", "xdg", ".local", "share"),
			want: filepath.Join(string(filepath.Separator), "tmp", "xdg", ".local", "share"),
		},
		{
			name: "empty env",
			env:  "",
			want: testDefaultDataHome,
		},
	}
	for _, tt := range tests {
		os.Setenv("XDG_DATA_HOME", tt.env)
		t.Run(tt.name, func(t *testing.T) {
			if got := DataHome(); got != tt.want {
				t.Errorf("DataHome() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConfigHome(t *testing.T) {
	var testDefaultConfigHome string
	switch runtime.GOOS {
	case "windows":
		testDefaultConfigHome = filepath.Join(home.Dir(), "AppData", "Local")
	default:
		testDefaultConfigHome = filepath.Join(home.Dir(), ".config")
	}

	tests := []struct {
		name string
		env  string
		want string
	}{
		{
			name: "set env based specification",
			env:  testDefaultConfigHome,
			want: testDefaultConfigHome,
		},
		{
			name: "set env based different from specification",
			env:  filepath.Join(string(filepath.Separator), "tmp", "config"),
			want: filepath.Join(string(filepath.Separator), "tmp", "config"),
		},
		{
			name: "empty env",
			env:  "",
			want: testDefaultConfigHome,
		},
	}
	for _, tt := range tests {
		os.Setenv("XDG_CONFIG_HOME", tt.env)
		t.Run(tt.name, func(t *testing.T) {
			if got := ConfigHome(); got != tt.want {
				t.Errorf("ConfigHome() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDataDirs(t *testing.T) {
	var testDefaultDataDirs string
	switch runtime.GOOS {
	case "windows":
		testDefaultDataDirs = filepath.Join(home.Dir(), "AppData", "Local")
	default:
		testDefaultDataDirs = filepath.Join("/usr", "local", "share") + string(filepath.ListSeparator) + filepath.Join("/usr", "share")
	}

	tests := []struct {
		name string
		env  string
		want string
	}{
		{
			name: "set env based specification",
			env:  testDefaultDataDirs,
			want: testDefaultDataDirs,
		},
		{
			name: "set env based different from specification",
			env:  filepath.Join("/opt", "local", "share"),
			want: filepath.Join("/opt", "local", "share"),
		},
		{
			name: "empty env",
			env:  "",
			want: testDefaultDataDirs,
		},
	}
	for _, tt := range tests {
		os.Setenv("XDG_DATA_DIRS", tt.env)
		t.Run(tt.name, func(t *testing.T) {
			if got := DataDirs(); got != tt.want {
				t.Errorf("DataDirs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConfigDirs(t *testing.T) {
	var testDefaultConfigDirs string
	switch runtime.GOOS {
	case "windows":
		testDefaultConfigDirs = filepath.Join(home.Dir(), "AppData", "Local")
	default:
		testDefaultConfigDirs = filepath.Join("/etc", "xdg")
	}

	tests := []struct {
		name string
		env  string
		want string
	}{
		{
			name: "set env based specification",
			env:  testDefaultConfigDirs,
			want: testDefaultConfigDirs,
		},
		{
			name: "set env based different from specification",
			env:  filepath.Join(string(filepath.Separator), "var", "etc", "xdg"),
			want: filepath.Join(string(filepath.Separator), "var", "etc", "xdg"),
		},
		{
			name: "empty env",
			env:  "",
			want: testDefaultConfigDirs,
		},
	}
	for _, tt := range tests {
		os.Setenv("XDG_CONFIG_DIRS", tt.env)
		t.Run(tt.name, func(t *testing.T) {
			if got := ConfigDirs(); got != tt.want {
				t.Errorf("ConfigDirs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCacheHome(t *testing.T) {
	var testDefaultCacheHome string
	switch runtime.GOOS {
	case "windows":
		testDefaultCacheHome = filepath.Join(home.Dir(), "AppData", "Local", "cache")
	default:
		testDefaultCacheHome = filepath.Join(home.Dir(), ".cache")
	}

	tests := []struct {
		name string
		env  string
		want string
	}{
		{
			name: "set env based specification",
			env:  testDefaultCacheHome,
			want: testDefaultCacheHome,
		},
		{
			name: "set env based different from specification",
			env:  filepath.Join(string(filepath.Separator), "tmp", "cache"),
			want: filepath.Join(string(filepath.Separator), "tmp", "cache"),
		},
		{
			name: "empty env",
			env:  "",
			want: testDefaultCacheHome,
		},
	}
	for _, tt := range tests {
		os.Setenv("XDG_CACHE_HOME", tt.env)
		t.Run(tt.name, func(t *testing.T) {
			if got := CacheHome(); got != tt.want {
				t.Errorf("CacheHome() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRuntimeDir(t *testing.T) {
	var testDefaultRuntimeDir string
	switch runtime.GOOS {
	case "windows":
		testDefaultRuntimeDir = home.Dir()
	default:
		testDefaultRuntimeDir = filepath.Join("/run", "user", strconv.Itoa(os.Getuid()))
	}

	tests := []struct {
		name string
		env  string
		want string
	}{
		{
			name: "set env based specification",
			env:  testDefaultRuntimeDir,
			want: testDefaultRuntimeDir,
		},
		{
			name: "set env based different from specification",
			env:  filepath.Join(string(filepath.Separator), "tmp", "user", "1000"),
			want: filepath.Join(string(filepath.Separator), "tmp", "user", "1000"),
		},
		{
			name: "empty env",
			env:  "",
			want: testDefaultRuntimeDir,
		},
	}
	for _, tt := range tests {
		os.Setenv("XDG_RUNTIME_DIR", tt.env)
		t.Run(tt.name, func(t *testing.T) {
			if got := RuntimeDir(); got != tt.want {
				t.Errorf("RuntimeDir() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNativeMode(t *testing.T) {
	// skip test if not darwin
	if runtime.GOOS != "darwin" {
		t.Skip("native mode for darwin only")
	}

	Mode = Native
	initOnce = sync.Once{}

	tests := []struct {
		name string
		fn   string
		want string
	}{
		{
			name: "DataHome",
			fn:   DataHome(),
			want: filepath.Join(home.Dir(), "Library", "Application Support"),
		},
		{
			name: "ConfigHome",
			fn:   ConfigHome(),
			want: filepath.Join(home.Dir(), "Library", "Preferences"),
		},
		{
			name: "DataDirs",
			fn:   DataDirs(),
			want: filepath.Join(home.Dir(), "Library", "Application Support"),
		},
		{
			name: "ConfigDirs",
			fn:   ConfigDirs(),
			want: filepath.Join(home.Dir(), "Library", "Preferences"),
		},
		{
			name: "CacheHome",
			fn:   CacheHome(),
			want: filepath.Join(home.Dir(), "Library", "Caches"),
		},
		{
			name: "RuntimeDir",
			fn:   RuntimeDir(),
			want: filepath.Join(home.Dir(), "Library", "Application Support"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.fn; got != tt.want {
				t.Errorf("NativeMode() = %v, want %v", got, tt.want)
			}
		})
	}
}
