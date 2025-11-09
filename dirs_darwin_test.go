//go:build darwin
// +build darwin

package appdirs

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func homePath(pathExtension string) string {
	home, _ := os.UserHomeDir()
	return filepath.Join(home, pathExtension)
}

func TestUserDataDir(t *testing.T) {
	expected_results := [][]string{
		{"appie", "1.0", homePath("Library/Application Support/appie/1.0")},
		{"appie", "", homePath("Library/Application Support/appie")},
	}
	for _, args := range expected_results {
		app, _ := NewAppDirs(args[0], args[1])
		assert.Equal(t, args[2], app.UserDataDir())
	}
}

func TestUserCacheDir(t *testing.T) {
	expected_results := [][]string{
		{"appie", "1.0", homePath("Library/Caches/appie/1.0")},
		{"appie", "", homePath("Library/Caches/appie")},
	}
	for _, args := range expected_results {
		app, _ := NewAppDirs(args[0], args[1])
		assert.Equal(t, args[2], app.UserCacheDir())
	}
}

func TestUserConfigDir(t *testing.T) {
	expected_results := [][]string{
		{"appie", "1.0", homePath("Library/Preferences/appie/1.0")},
		{"appie", "", homePath("Library/Preferences/appie")},
	}
	for _, args := range expected_results {
		app, _ := NewAppDirs(args[0], args[1])
		assert.Equal(t, args[2], app.UserConfigDir())
	}
}

func TestUserLogDir(t *testing.T) {
	expected_results := [][]string{
		{"appie", "1.0", homePath("Library/Logs/appie/1.0")},
		{"appie", "", homePath("Library/Logs/appie")},
	}
	for _, args := range expected_results {
		app, _ := NewAppDirs(args[0], args[1])
		assert.Equal(t, args[2], app.UserLogDir())
	}
}

type DirTest struct {
	Name     string
	Version  string
	Expected string
}

func NewDirTest(name, version, expected string) DirTest {
	return DirTest{
		Name:     name,
		Version:  version,
		Expected: expected,
	}
}

func TestSiteAppDir(t *testing.T) {
	expected_results := []DirTest{
		NewDirTest("appie", "1.0", "/Applications/appie/1.0"),
		NewDirTest("appie", "", "/Applications/appie"),
	}
	for _, app := range expected_results {
		appDirs, _ := NewAppDirs(app.Name, app.Version)
		assert.Equal(t, app.Expected, appDirs.SiteAppDir())
	}
}

func TestSiteConfigDir(t *testing.T) {
	expected_results := []DirTest{
		NewDirTest("appie", "1.0", "/Library/Preferences/appie/1.0"),
		NewDirTest("appie", "", "/Library/Preferences/appie"),
	}
	for _, app := range expected_results {
		appDirs, _ := NewAppDirs(app.Name, app.Version)
		assert.Equal(t, app.Expected, appDirs.SiteConfigDir())
	}
}

func TestSiteDataDir(t *testing.T) {
	expected_results := []DirTest{
		NewDirTest("appie", "1.0", "/Library/Application Support/appie/1.0"),
		NewDirTest("appie", "", "/Library/Application Support/appie"),
	}
	for _, app := range expected_results {
		appDirs, _ := NewAppDirs(app.Name, app.Version)
		assert.Equal(t, app.Expected, appDirs.SiteDataDir())
	}
}

func TestSiteLogDir(t *testing.T) {
	expected_results := []DirTest{
		NewDirTest("appie", "1.0", "/Library/Application Support/appie/1.0"),
		NewDirTest("appie", "", "/Library/Application Support/appie"),
	}
	for _, app := range expected_results {
		appDirs, _ := NewAppDirs(app.Name, app.Version)
		assert.Equal(t, app.Expected, appDirs.SiteLogDir())
	}
}
