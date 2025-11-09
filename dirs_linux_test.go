//go:build linux
// +build linux

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
		{"", "", homePath(".local/share")},
		{"", "1.0", homePath(".local/share")},
		{"appie", "1.0", homePath(".local/share/appie/1.0")},
		{"appie", "", homePath(".local/share/appie")},
	}
	for _, args := range expected_results {
		app := NewAppDirs(args[0], args[1])
		assert.Equal(t, args[2], app.UserDataDir())
	}
}

func TestUserCacheDir(t *testing.T) {
	expected_results := [][]string{
		{"", "", homePath(".cache")},
		{"", "1.0", homePath(".cache")},
		{"appie", "1.0", homePath(".cache/appie/1.0")},
		{"appie", "", homePath(".cache/appie")},
	}
	for _, args := range expected_results {
		app := NewAppDirs(args[0], args[1])
		assert.Equal(t, args[2], app.UserCacheDir())
	}
}

func TestUserConfigDir(t *testing.T) {
	expected_results := [][]string{
		{"", "", homePath(".config")},
		{"", "1.0", homePath(".config")},
		{"appie", "1.0", homePath(".config/appie/1.0")},
		{"appie", "", homePath(".config/appie")},
	}
	for _, args := range expected_results {
		app := NewAppDirs(args[0], args[1])
		assert.Equal(t, args[2], app.UserConfigDir())
	}
}

func TestUserLogDir(t *testing.T) {
	expected_results := [][]string{
		{"", "", homePath(".cache/log")},
		{"", "1.0", homePath(".cache/log")},
		{"appie", "1.0", homePath(".cache/appie/1.0/log")},
		{"appie", "", homePath(".cache/appie/log")},
	}
	for _, args := range expected_results {
		app := NewAppDirs(args[0], args[1])
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
		NewDirTest("", "", "/usr/bin"),
		NewDirTest("", "1.0", "/usr/bin"),
		NewDirTest("appie", "1.0", "/usr/bin/appie/1.0"),
		NewDirTest("appie", "", "/usr/bin/appie"),
	}
	for _, app := range expected_results {
		appDirs := NewAppDirs(app.Name, app.Version)
		assert.Equal(t, app.Expected, appDirs.SiteAppDir())
	}
}

func TestSiteConfigDir(t *testing.T) {
	expected_results := []DirTest{
		NewDirTest("", "", "/etc/xdg"),
		NewDirTest("", "1.0", "/etc/xdg"),
		NewDirTest("appie", "1.0", "/etc/xdg/appie/1.0"),
		NewDirTest("appie", "", "/etc/xdg/appie"),
	}
	for _, app := range expected_results {
		appDirs := NewAppDirs(app.Name, app.Version)
		assert.Equal(t, app.Expected, appDirs.SiteConfigDir())
	}
}

func TestSiteDataDir(t *testing.T) {
	expected_results := []DirTest{
		NewDirTest("", "", "/usr/local/share"),
		NewDirTest("", "1.0", "/usr/local/share"),
		NewDirTest("appie", "1.0", "/usr/local/share/appie/1.0"),
		NewDirTest("appie", "", "/usr/local/share/appie"),
	}
	for _, app := range expected_results {
		appDirs := NewAppDirs(app.Name, app.Version)
		assert.Equal(t, app.Expected, appDirs.SiteDataDir())
	}
}

func TestSiteLogDir(t *testing.T) {
	expected_results := []DirTest{
		NewDirTest("", "", "/var/log"),
		NewDirTest("", "1.0", "/var/log"),
		NewDirTest("appie", "1.0", "/var/log/appie/1.0"),
		NewDirTest("appie", "", "/var/log/appie"),
	}
	for _, app := range expected_results {
		appDirs := NewAppDirs(app.Name, app.Version)
		assert.Equal(t, app.Expected, appDirs.SiteLogDir())
	}
}
