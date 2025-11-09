# appdirs

A cross-platform library for generating common application directory paths.

## Installation

```bash
go get github.com/markgemmill/appdirs
```

## Usage

### Basic Example

```go
package main

import (
    "fmt"
    "github.com/markgemmill/appdirs"
)

func main() {
    // Create an AppDirs instance for your application
    dirs, err := appdirs.NewAppDirs("myapp", "1.0")

    if err != nil {
        panic("this is only going to happen if app name is blank")
    }

    // Get user-level directories
    fmt.Println("User Data:", dirs.UserDataDir())
    fmt.Println("User Config:", dirs.UserConfigDir())
    fmt.Println("User Cache:", dirs.UserCacheDir())
    fmt.Println("User Logs:", dirs.UserLogDir())

    // Get system-level directories
    fmt.Println("Site Data:", dirs.SiteDataDir())
    fmt.Println("Site Config:", dirs.SiteConfigDir())
}
```

### Example Output

**macOS:**

```text
User Data:   /Users/username/Library/Application Support/myapp/1.0
User Config: /Users/username/Library/Preferences/myapp/1.0
User Cache:  /Users/username/Library/Caches/myapp/1.0
User Logs:   /Users/username/Library/Logs/myapp/1.0
Site Data:   /Library/Application Support/myapp/1.0
Site Config: /Library/Preferences/myapp/1.0
```

**Linux:**

```text
User Data:   /home/username/.local/share/myapp/1.0
User Config: /home/username/.config/myapp/1.0
User Cache:  /home/username/.cache/myapp/1.0
User Logs:   /home/username/.cache/myapp/1.0/log
Site Data:   /usr/local/share/myapp/1.0
Site Config: /etc/xdg/myapp/1.0
```

**Windows:**

```text
User Data:   C:\Users\username\AppData\Local\myapp\1.0
User Config: C:\Users\username\AppData\Local\myapp\1.0
User Cache:  C:\Users\username\AppData\Local\myapp\1.0\cache
User Logs:   C:\Users\username\AppData\Local\myapp\1.0\logs
Site Data:   C:\ProgramData\myapp\1.0
Site Config: C:\ProgramData\myapp\1.0\config
```

### Note on Directory Creation

This library only returns directory paths; it does not create them.
