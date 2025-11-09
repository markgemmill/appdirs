//go:build darwin
// +build darwin

package appdirs

import (
	"os"
	"path/filepath"
)

func UserAppDir(appName, version string) string {
	home, _ := os.UserHomeDir()

	root := filepath.Join(home, "Applications")

	return buildAppDir(root, appName, version)
}

func UserDataDir(appName, version string) string {
	home, _ := os.UserHomeDir()

	root := filepath.Join(home, "Library", "Application Support")

	return buildAppDir(root, appName, version)
}

func UserCacheDir(appName, version string) string {
	home, _ := os.UserHomeDir()
	root := filepath.Join(home, "Library", "Caches")
	return buildAppDir(root, appName, version)
}

func UserConfigDir(appName, version string) string {
	home, _ := os.UserHomeDir()
	root := filepath.Join(home, "Library", "Preferences")
	return buildAppDir(root, appName, version)
}

func UserLogDir(appName, version string) string {
	home, _ := os.UserHomeDir()
	root := filepath.Join(home, "Library", "Logs")
	return buildAppDir(root, appName, version)
}

func SiteAppDir(appName, version string) string {
	root := filepath.Join("/Applications")
	return buildAppDir(root, appName, version)
}

func SiteDataDir(appName, version string) string {
	root := filepath.Join("/Library", "Application Support")
	return buildAppDir(root, appName, version)
}

func SiteConfigDir(appName, version string) string {
	root := filepath.Join("/Library", "Preferences")
	return buildAppDir(root, appName, version)
}

func SiteLogDir(appName, version string) string {
	root := filepath.Join("/Library", "Application Support")
	return buildAppDir(root, appName, version)
}
