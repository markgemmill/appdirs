//go:build windows
// +build windows

package appdirs

import (
	"os"
	"path/filepath"
)

func UserAppData() string {
	home, _ := os.UserHomeDir()
	return filepath.Join(home, "AppData", "Local")
}

func ProgramData() string {
	return os.Getenv("ProgramData")
}

func UserAppDir(appName, version string) string {
	home, _ := os.UserHomeDir()

	root := filepath.Join(home, "Applications")
	return buildAppDir(root, appName, version)
}

func UserDataDir(appName, version string) string {
	root := UserAppData()
	return buildAppDir(root, appName, version)
}

func UserCacheDir(appName, version string) string {
	root := UserAppData()
	return buildAppDir(root, appName, version, "cache")
}

func UserConfigDir(appName, version string) string {
	root := UserAppData()
	return buildAppDir(root, appName, version)
}

func UserLogDir(appName, version string) string {
	root := UserAppData()
	return buildAppDir(root, appName, version, "logs")
}

func SiteAppDir(appName, version string) string {
	root := ProgramData()
	return buildAppDir(root, appName, version)
}

func SiteDataDir(appName, version string) string {
	root := ProgramData()
	return buildAppDir(root, appName, version)
}

func SiteConfigDir(appName, version string) string {
	root := ProgramData()
	return buildAppDir(root, appName, version, "config")
}

func SiteLogDir(appName, version string) string {
	root := ProgramData()
	return buildAppDir(root, appName, version, "logs")
}
