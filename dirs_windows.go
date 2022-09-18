//go:build windows
// +build windows

package appdirs

import (
	"os"
	"path"
)

func UserAppData() string {
	home, _ := os.UserHomeDir()
	return path.Join(home, "AppData", "Local")
}

func ProgramData() string {
	return path.Join("C:", "ProgramData")
}

func UserAppDir(appName, version string) string {
	home, _ := os.UserHomeDir()

	root := path.Join(home, "Applications")
	return MakeAppDir(root, appName, version)
}

func UserDataDir(appName, version string) string {
	root := UserAppData()
	return MakeAppDir(root, appName, version)
}

func UserCacheDir(appName, version string) string {
	root := UserAppData()
	return MakeAppDir(root, appName, version, "cache")
}

func UserConfigDir(appName, version string) string {
	root := UserAppData()
	return MakeAppDir(root, appName, version)
}

func UserLogDir(appName, version string) string {
	root := UserAppData()
	return MakeAppDir(root, appName, version, "logs")
}

func SiteAppDir(appName, version string) string {
	root := ProgramData()
	return MakeAppDir(root, appName, version)
}

func SiteDataDir(appName, version string) string {
	root := ProgramData()
	return MakeAppDir(root, appName, version)
}

func SiteConfigDir(appName, version string) string {
	root := ProgramData()
	return MakeAppDir(root, appName, version, "config")
}

func SiteLogDir(appName, version string) string {
	root := ProgramData()
	return MakeAppDir(root, appName, version, "logs")
}
