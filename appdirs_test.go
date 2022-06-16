//go:build darwin
// +build darwin

package appdirs

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserDataDir(t *testing.T) {
	expected_results := [][]string{
		[]string{"", "", "/Users/mark/Library/Application Support"},
		[]string{"", "1.0", "/Users/mark/Library/Application Support"},
		[]string{"appie", "1.0", "/Users/mark/Library/Application Support/appie/1.0"},
		[]string{"appie", "", "/Users/mark/Library/Application Support/appie"},
	}
	for _, args := range expected_results {
		app := NewAppDirs(args[0], args[1])
		assert.Equal(t, args[2], app.UserDataDir())
	}
}

func TestUserCacheDir(t *testing.T) {

	expected_results := [][]string{
		[]string{"", "", "/Users/mark/Library/Caches"},
		[]string{"", "1.0", "/Users/mark/Library/Caches"},
		[]string{"appie", "1.0", "/Users/mark/Library/Caches/appie/1.0"},
		[]string{"appie", "", "/Users/mark/Library/Caches/appie"},
	}
	for _, args := range expected_results {
		app := NewAppDirs(args[0], args[1])
		assert.Equal(t, args[2], app.UserCacheDir())
	}
}

func TestUserConfigDir(t *testing.T) {

	expected_results := [][]string{
		[]string{"", "", "/Users/mark/Library/Preferences"},
		[]string{"", "1.0", "/Users/mark/Library/Preferences"},
		[]string{"appie", "1.0", "/Users/mark/Library/Preferences/appie/1.0"},
		[]string{"appie", "", "/Users/mark/Library/Preferences/appie"},
	}
	for _, args := range expected_results {
		app := NewAppDirs(args[0], args[1])
		assert.Equal(t, args[2], app.UserConfigDir())
	}
}

func TestUserLogDir(t *testing.T) {
	expected_results := [][]string{
		[]string{"", "", "/Users/mark/Library/Logs"},
		[]string{"", "1.0", "/Users/mark/Library/Logs"},
		[]string{"appie", "1.0", "/Users/mark/Library/Logs/appie/1.0"},
		[]string{"appie", "", "/Users/mark/Library/Logs/appie"},
	}
	for _, args := range expected_results {
		app := NewAppDirs(args[0], args[1])
		assert.Equal(t, args[2], app.UserLogDir())
	}
}
