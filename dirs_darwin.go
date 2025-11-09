//go:build darwin
// +build darwin

package appdirs

import (
	"os"
	"path"
)

func UserAppDir(appName, version string) string {
	home, _ := os.UserHomeDir()

	root := path.Join(home, "Applications")

	return buildAppDir(root, appName, version)
}

func UserDataDir(appName, version string) string {
	home, _ := os.UserHomeDir()

	root := path.Join(home, "Library", "Application Support")

	return buildAppDir(root, appName, version)
}

func UserCacheDir(appName, version string) string {
	home, _ := os.UserHomeDir()
	root := path.Join(home, "Library", "Caches")
	return buildAppDir(root, appName, version)
}

func UserConfigDir(appName, version string) string {
	home, _ := os.UserHomeDir()
	root := path.Join(home, "Library", "Preferences")
	return buildAppDir(root, appName, version)
}

func UserLogDir(appName, version string) string {
	home, _ := os.UserHomeDir()
	root := path.Join(home, "Library", "Logs")
	return buildAppDir(root, appName, version)
}

func SiteAppDir(appName, version string) string {
	root := path.Join("/Applications")
	return buildAppDir(root, appName, version)
}

func SiteDataDir(appName, version string) string {
	root := path.Join("/Library", "Application Support")
	return buildAppDir(root, appName, version)
}

func SiteConfigDir(appName, version string) string {
	root := path.Join("/Library", "Preferences")
	return buildAppDir(root, appName, version)
}

func SiteLogDir(appName, version string) string {
	root := path.Join("/Library", "Application Support")
	return buildAppDir(root, appName, version)
}
