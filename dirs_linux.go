//go:build linux
// +build linux

package appdirs

import (
	"os"
	"path/filepath"
)

func GetEnvVar(variable, def string) string {
	v, exists := os.LookupEnv(variable)
	if !exists {
		return def
	}
	return v
}

func UserAppDir(appName, version string) string {
	home, _ := os.UserHomeDir()

	root := filepath.Join(home, ".local", "bin")

	return buildAppDir(root, appName, version)
}

func UserDataDir(appName, version string) string {
	home, _ := os.UserHomeDir()

	root := GetEnvVar("XDG_DATA_HOME", filepath.Join(home, ".local", "share"))

	return buildAppDir(root, appName, version)
}

func UserCacheDir(appName, version string) string {
	// TODO
	// Uses XDG_CACHE_DIR instead of XDG_CACHE_HOME:
	// root := GetEnvVar("XDG_CACHE_DIR", filepath.Join(home, ".cache"))  // WRONG
	home, _ := os.UserHomeDir()
	root := GetEnvVar("XDG_CACHE_DIR", filepath.Join(home, ".cache"))
	return buildAppDir(root, appName, version)
}

func UserConfigDir(appName, version string) string {
	home, _ := os.UserHomeDir()
	root := GetEnvVar("XDG_CONFIG_HOME", filepath.Join(home, ".config"))
	return buildAppDir(root, appName, version)
}

func UserLogDir(appName, version string) string {
	root := UserCacheDir(appName, version)
	return filepath.Join(root, "log")
}

func SiteAppDir(appName, version string) string {
	root := filepath.Join("/usr", "bin")
	return buildAppDir(root, appName, version)
}

func SiteDataDir(appName, version string) string {
	root := filepath.Join("/usr", "local", "share")
	return buildAppDir(root, appName, version)
}

func SiteConfigDir(appName, version string) string {
	root := filepath.Join("/etc", "xdg")
	return buildAppDir(root, appName, version)
}

func SiteLogDir(appName, version string) string {
	root := filepath.Join("/var", "log")
	return buildAppDir(root, appName, version)
}
