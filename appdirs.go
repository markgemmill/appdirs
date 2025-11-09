package appdirs

import (
	"fmt"
	"path/filepath"
)

func buildAppDir(root, appName, version string, extra ...string) string {
	// this should never panic unless user creates AppDirs directly
	if appName == "" {
		panic("app name is an empty string")
	}

	root = filepath.Join(root, appName)

	if len(appName) > 0 && len(version) > 0 {
		root = filepath.Join(root, version)
	}

	for _, e := range extra {
		root = filepath.Join(root, e)
	}

	return root
}

type AppDirs struct {
	Name    string
	Version string
}

/*
AppDirs generally follows the guidance of it's python namesake. We do not
provide `roaming` (Windows), or `app author` options. Nor do we provide for
multiple directory options on Linux. The main task is to provide a unified
set of standard directory options that related to the execution of a named
application.

Derived from `https://github.com/ActiveState/appdirs` documentation:

Typical user data directories are:

	Mac OS:    ~/Library/Application Support/<AppName>
	Unix:      ~/.local/share/<AppName>    # or in $XDG_DATA_HOME, if defined
	Windows    C:\Users\<username>\AppData\Local\<AppName>

Typical site data directories are:

	Mac OS:    /Library/Application Support/<AppName>
	Unix:      /usr/local/share/<AppName> or /usr/share/<AppName>
	Windowns:  C:\ProgramData\<AppName>

Typical user config directories are:

	Mac OS:    ~/Library/Preferences/<AppName>
	Unix:      ~/.config/<AppName>     # or in $XDG_CONFIG_HOME, if defined
	Windows:   same as user_data_dir

Typical site config directories are:

	Mac OS:    same as site_data_dir
	Unix:      /etc/<AppName>
	Windows:   same as site_data_dir

Typical user cache directories are:

	Mac OS:    ~/Library/Caches/<AppName>
	Unix:      ~/.cache/<AppName> (XDG default)
	Windows:   C:\Users\<username>\AppData\Local\<AppName>\Cache

Typical user state directories are:

	Mac OS:    same as user_data_dir
	Unix:      ~/.local/state/<AppName>   # or in $XDG_STATE_HOME, if defined
	Windows:   same as user_data_dir

Typical user log directories are:

	Mac OS:    ~/Library/Logs/<AppName>
	Unix:      ~/.cache/<AppName>/log  # or under $XDG_CACHE_HOME if defined

Windows:       C:\Users\<username>\AppData\Local\<AppAuthor>\<AppName>\Logs
*/
func NewAppDirs(name, version string) (AppDirs, error) {
	if name == "" {
		return AppDirs{}, fmt.Errorf("app name cannot be an empty string")
	}
	return AppDirs{
		Name:    name,
		Version: version,
	}, nil
}

func (d AppDirs) UserAppDir() string {
	return UserAppDir(d.Name, d.Version)
}

func (d AppDirs) UserDataDir() string {
	return UserDataDir(d.Name, d.Version)
}

func (d AppDirs) UserConfigDir() string {
	return UserConfigDir(d.Name, d.Version)
}

func (d AppDirs) UserCacheDir() string {
	return UserCacheDir(d.Name, d.Version)
}

func (d AppDirs) UserLogDir() string {
	return UserLogDir(d.Name, d.Version)
}

func (d AppDirs) SiteAppDir() string {
	return SiteAppDir(d.Name, d.Version)
}

func (d AppDirs) SiteDataDir() string {
	return SiteDataDir(d.Name, d.Version)
}

func (d AppDirs) SiteConfigDir() string {
	return SiteConfigDir(d.Name, d.Version)
}

func (d AppDirs) SiteLogDir() string {
	return SiteLogDir(d.Name, d.Version)
}
