package main

import (
	"fmt"
	"github.com/markgemmill/appdirs"
)

func main() {
	appdir := appdirs.NewAppDirs("", "")
	fmt.Printf("\n--- user raw directories ---\n")
	fmt.Printf("User Apps:    %s\n", appdir.UserAppDir())
	fmt.Printf("User Data:    %s\n", appdir.UserDataDir())
	fmt.Printf("User Config:  %s\n", appdir.UserConfigDir())
	fmt.Printf("User Cache:   %s\n", appdir.UserCacheDir())
	fmt.Printf("User Logs:    %s\n", appdir.UserLogDir())

	appdir = appdirs.NewAppDirs("myapp", "1.0")
	fmt.Printf("\n--- user directories with app name and version ---\n")
	fmt.Printf("User Apps:    %s\n", appdir.UserAppDir())
	fmt.Printf("User Data:    %s\n", appdir.UserDataDir())
	fmt.Printf("User Config:  %s\n", appdir.UserConfigDir())
	fmt.Printf("User Cache:   %s\n", appdir.UserCacheDir())
	fmt.Printf("User Logs:    %s\n", appdir.UserLogDir())

	appdir = appdirs.NewAppDirs("", "")
	fmt.Printf("\n--- site raw directories ---\n")
	fmt.Printf("Site Apps:    %s\n", appdir.SiteAppDir())
	fmt.Printf("Site Data:    %s\n", appdir.SiteDataDir())
	fmt.Printf("Site Config:  %s\n", appdir.SiteConfigDir())
	fmt.Printf("Site Logs:    %s\n", appdir.SiteLogDir())

	appdir = appdirs.NewAppDirs("myapp", "1.0")
	fmt.Printf("\n--- site directories with app name and version ---\n")
	fmt.Printf("Site Apps:    %s\n", appdir.SiteAppDir())
	fmt.Printf("Site Data:    %s\n", appdir.SiteDataDir())
	fmt.Printf("Site Config:  %s\n", appdir.SiteConfigDir())
	fmt.Printf("Site Logs:    %s\n", appdir.SiteLogDir())

}
