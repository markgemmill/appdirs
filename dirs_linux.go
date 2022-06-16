//go:build linux
// +build linux

package appdirs

import (
	"os"
	"path"
)

func GetEnvVar(variable, def string) string {
	v, exists := os.LookupEnv(variable)
	if exists == false {
		return def
	}
	return v
}

func UserAppDir(appName, version string) string {
	home, _ := os.UserHomeDir()

	root := path.Join(home, ".local", "bin")

	return MakeAppDir(root, appName, version)
}

func UserDataDir(appName, version string) string {
	home, _ := os.UserHomeDir()

	root := GetEnvVar("XDG_DATA_HOME", path.Join(home, ".local", "share"))

	return MakeAppDir(root, appName, version)
}

func UserCacheDir(appName, version string) string {
	home, _ := os.UserHomeDir()
	root := GetEnvVar("XDG_CACHE_DIR", path.Join(home, ".cache"))
	return MakeAppDir(root, appName, version)
}

func UserConfigDir(appName, version string) string {
	home, _ := os.UserHomeDir()
	root := GetEnvVar("XDG_CONFIG_HOME", path.Join(home, ".config"))
	return MakeAppDir(root, appName, version)
}

func UserLogDir(appName, version string) string {
	root := UserCacheDir(appName, version)
	return path.Join(root, "log")
}

func SiteAppDir(appName, version string) string {
	root := path.Join("/usr", "bin")
	return MakeAppDir(root, appName, version)
}

func SiteDataDir(appName, version string) string {
	root := path.Join("/usr", "local", "share")
	return MakeAppDir(root, appName, version)
}

func SiteConfigDir(appName, version string) string {
	root := path.Join("/etc", "xdg")
	return MakeAppDir(root, appName, version)
}

func SiteLogDir(appName, version string) string {
	root := path.Join("/var", "log")
	return MakeAppDir(root, appName, version)
}
