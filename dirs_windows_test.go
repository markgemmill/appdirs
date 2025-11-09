//go:build windows
// +build windows

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
		{"", "", homePath("AppData\\Local")},
		{"", "1.0", homePath("AppData\\Local")},
		{"appie", "1.0", homePath("AppData\\Local\\appie\\1.0")},
		{"appie", "", homePath("AppData\\Local\\appie")},
	}
	for _, args := range expected_results {
		app := NewAppDirs(args[0], args[1])
		assert.Equal(t, args[2], app.UserDataDir())
	}
}

func TestUserCacheDir(t *testing.T) {
	expected_results := [][]string{
		{"", "", homePath("AppData\\Local")},
		{"", "1.0", homePath("AppData\\Local")},
		{"appie", "1.0", homePath("AppData\\Local\\appie\\1.0")},
		{"appie", "", homePath("AppData\\Local\\appie")},
	}
	for _, args := range expected_results {
		app := NewAppDirs(args[0], args[1])
		assert.Equal(t, args[2], app.UserCacheDir())
	}
}

func TestUserConfigDir(t *testing.T) {
	expected_results := [][]string{
		{"", "", homePath("AppData\\Local")},
		{"", "1.0", homePath("AppData\\Local")},
		{"appie", "1.0", homePath("AppData\\Local\\appie\\1.0")},
		{"appie", "", homePath("AppData\\Local\\appie")},
	}
	for _, args := range expected_results {
		app := NewAppDirs(args[0], args[1])
		assert.Equal(t, args[2], app.UserConfigDir())
	}
}

func TestUserLogDir(t *testing.T) {
	expected_results := [][]string{
		{"", "", homePath("AppData\\Local")},
		{"", "1.0", homePath("AppData\\Local")},
		{"appie", "1.0", homePath("AppData\\Local\\appie\\1.0\\logs")},
		{"appie", "", homePath("AppData\\Local\\appie\\logs")},
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

func sitePath(pathExtension string) string {
	site := ProgramData()
	return filepath.Join(site, pathExtension)
}

func TestSiteAppDir(t *testing.T) {
	expected_results := []DirTest{
		NewDirTest("", "", ProgramData()),
		NewDirTest("", "1.0", ProgramData()),
		NewDirTest("appie", "1.0", sitePath("\\appie\\1.0")),
		NewDirTest("appie", "", sitePath("appie")),
	}
	for _, app := range expected_results {
		appDirs := NewAppDirs(app.Name, app.Version)
		assert.Equal(t, app.Expected, appDirs.SiteAppDir())
	}
}

func TestSiteConfigDir(t *testing.T) {
	expected_results := []DirTest{
		NewDirTest("", "", sitePath("config")),
		NewDirTest("", "1.0", sitePath("config")),
		NewDirTest("appie", "1.0", sitePath("appie\\1.0\\config")),
		NewDirTest("appie", "", sitePath("appie\\config")),
	}
	for _, app := range expected_results {
		appDirs := NewAppDirs(app.Name, app.Version)
		assert.Equal(t, app.Expected, appDirs.SiteConfigDir())
	}
}

func TestSiteDataDir(t *testing.T) {
	expected_results := []DirTest{
		NewDirTest("", "", ProgramData()),
		NewDirTest("", "1.0", ProgramData()),
		NewDirTest("appie", "1.0", sitePath("appie\\1.0")),
		NewDirTest("appie", "", sitePath("appie")),
	}
	for _, app := range expected_results {
		appDirs := NewAppDirs(app.Name, app.Version)
		assert.Equal(t, app.Expected, appDirs.SiteDataDir())
	}
}

func TestSiteLogDir(t *testing.T) {
	expected_results := []DirTest{
		NewDirTest("", "", sitePath("logs")),
		NewDirTest("", "1.0", sitePath("logs")),
		NewDirTest("appie", "1.0", sitePath("appie\\1.0\\logs")),
		NewDirTest("appie", "", sitePath("appie\\logs")),
	}
	for _, app := range expected_results {
		appDirs := NewAppDirs(app.Name, app.Version)
		assert.Equal(t, app.Expected, appDirs.SiteLogDir())
	}
}
