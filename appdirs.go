package appdirs

import (
	"path"
)

var VERSION = "0.1.0-dev.0"

func MakeAppDir(root, appName, version string, extra ...string) string {

	if len(appName) > 0 {
		root = path.Join(root, appName)
	}

	if len(appName) > 0 && len(version) > 0 {
		root = path.Join(root, version)
	}

	for _, e := range extra {
		root = path.Join(root, e)
	}

	return root
}

type AppDirs struct {
	Name    string
	Version string
}

func NewAppDirs(name, version string) AppDirs {
	return AppDirs{
		Name:    name,
		Version: version,
	}
}

func (d *AppDirs) UserAppDir() string {
	return UserAppDir(d.Name, d.Version)
}

func (d *AppDirs) UserDataDir() string {
	return UserDataDir(d.Name, d.Version)
}

func (d *AppDirs) UserConfigDir() string {
	return UserConfigDir(d.Name, d.Version)
}

func (d *AppDirs) UserCacheDir() string {
	return UserCacheDir(d.Name, d.Version)
}

func (d *AppDirs) UserLogDir() string {
	return UserLogDir(d.Name, d.Version)
}

func (d *AppDirs) SiteAppDir() string {
	return SiteAppDir(d.Name, d.Version)
}

func (d *AppDirs) SiteDataDir() string {
	return SiteDataDir(d.Name, d.Version)
}

func (d *AppDirs) SiteConfigDir() string {
	return SiteConfigDir(d.Name, d.Version)
}

func (d *AppDirs) SiteLogDir() string {
	return SiteLogDir(d.Name, d.Version)
}
