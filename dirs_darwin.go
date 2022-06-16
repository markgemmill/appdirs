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

	return MakeAppDir(root, appName, version)
}

func UserDataDir(appName, version string) string {
	home, _ := os.UserHomeDir()

	root := path.Join(home, "Library", "Application Support")

	return MakeAppDir(root, appName, version)
}

func UserCacheDir(appName, version string) string {
	home, _ := os.UserHomeDir()
	root := path.Join(home, "Library", "Caches")
	return MakeAppDir(root, appName, version)
}

func UserConfigDir(appName, version string) string {
	home, _ := os.UserHomeDir()
	root := path.Join(home, "Library", "Preferences")
	return MakeAppDir(root, appName, version)
}

func UserLogDir(appName, version string) string {
	home, _ := os.UserHomeDir()
	root := path.Join(home, "Library", "Logs")
	return MakeAppDir(root, appName, version)
}

func SiteAppDir(appName, version string) string {
	root := path.Join("/Applications")
	return MakeAppDir(root, appName, version)
}

func SiteDataDir(appName, version string) string {
	root := path.Join("/Library", "Application Support")
	return MakeAppDir(root, appName, version)
}

func SiteConfigDir(appName, version string) string {
	root := path.Join("/Library", "Preferences")
	return MakeAppDir(root, appName, version)
}

func SiteLogDir(appName, version string) string {
	root := path.Join("/Library", "Application Support")
	return MakeAppDir(root, appName, version)
}
