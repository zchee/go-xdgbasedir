# go-xdgbasedir

| **linux / darwin**                          | **windows**                                   | **godoc.org**                           |
|:-------------------------------------------:|:---------------------------------------------:|:---------------------------------------:|
| [![circleci.com][circleci-badge]][circleci] | [![appveyor-badge][appveyor-badge]][appveyor] | [![godoc.org][godoc-badge]][godoc]      |
| **coverage**                                | **Release**                                   | **XDG Version**                         |
| [![codecov.io][codecov-badge]][codecov]     | [![Release][release-badge]][release]          | [![XDG version][xdg-badge]][xdg]        |

Package `xdgbasedir` implements a freedesktop XDG Base Directory Specification for Go.

## freedesktop.org XDG Base Directory Specification reference:

  - https://specifications.freedesktop.org/basedir-spec/basedir-spec-0.8.html
  - https://specifications.freedesktop.org/basedir-spec/basedir-spec-latest.html

## Current specification version:
  
  - 0.8

## Specification

| func           | linux, darwin (Mode: `Unix`)  | darwin (Mode: `Native`)         |
|----------------|-------------------------------|---------------------------------|
| `DataHome()`   | `~/.local/share`              | `~/Library/Application Support` |
| `ConfigHome()` | `~/.config`                   | `~/Library/Preferences`         |
| `DataDirs()`   | `/usr/local/share:/usr/share` | `~/Library/Application Support` |
| `ConfigDirs()` | `/etc/xdg`                    | `~/Library/Preferences`         |
| `CacheHome()`  | `~/.cache`                    | `~/Library/Caches`              |
| `RuntimeDir()` | `/run/user/$(id -u)`          | `~/Library/Application Support` |

| func           | windows                               |
|----------------|---------------------------------------|
| `DataHome()`   | `C:\Users\%USER%\AppData\Roaming`     |
| `ConfigHome()` | `C:\Users\%USER%\AppData\Roaming`     |
| `DataDirs()`   | `C:\Users\%USER%\AppData\Roaming`     |
| `ConfigDirs()` | `C:\Users\%USER%\AppData\Roaming`     |
| `CacheHome()`  | `C:\Users\%USER%\AppData\Local\cache` |
| `RuntimeDir()` | `C:\Users\%USER%`                     |

## Note

XDG Base Directory Specification is mainly for GNU/Linux. It does not mention which directory to use with macOS(`darwin`) or `windows`.  
So, We referred to the [qt standard paths document](http://doc.qt.io/qt-5/qstandardpaths.html) for the corresponding directory.

We prepared a `Mode` for users using macOS like Unix. It's `darwin` GOOS specific.  
If it is set to `Unix`, it refers to the same path as linux. If it is set to `Native`, it refers to the [Specification](#specification) path.  
By default, `Unix`.

`Unix`:

```go
// +build darwin

package main

import (
	"fmt"

	"github.com/zchee/go-xdgbasedir"
)

func init() {
	xdgbasedir.Mode = xdgbasedir.Unix // optional, default is Unix
}

func main() {
	fmt.Println(xdgbasedir.DataHome())

	// Output:
	// "/Users/foo/.local/share"
}
```

`Native`:

```go
// +build darwin

package main

import (
	"fmt"

	"github.com/zchee/go-xdgbasedir"
)

func init() {
	xdgbasedir.Mode = xdgbasedir.Native
}

func main() {
	fmt.Println(xdgbasedir.DataHome())

	// Output:
	// "/Users/foo/Library/Application Support"
}
```

## Badge

powered by [shields.io](https://shields.io).

[![XDGBaseDirSpecVersion](https://img.shields.io/badge/%20XDG%20Base%20Dir%20-%200.8-blue.svg?style=for-the-badge)](https://specifications.freedesktop.org/basedir-spec/0.8/)

markdown:

```markdown
[![XDGBaseDirSpecVersion](https://img.shields.io/badge/%20XDG%20Base%20Dir%20-%200.8-blue.svg?style=for-the-badge)](https://specifications.freedesktop.org/basedir-spec/0.8/)
```


<!-- badge links -->
[circleci]: https://circleci.com/gh/zchee/workflows/go-xdgbasedir
[appveyor]: https://ci.appveyor.com/project/zchee/go-xdgbasedir
[godoc]: https://godoc.org/github.com/zchee/go-xdgbasedir
[codecov]: https://codecov.io/gh/zchee/go-xdgbasedir
[release]: https://github.com/zchee/go-xdgbasedir/releases
[xdg]: https://specifications.freedesktop.org/basedir-spec/0.8/

[circleci-badge]: https://img.shields.io/circleci/project/github/zchee/go-xdgbasedir/master.svg?style=for-the-badge&label=CIRCLECI&logo=circleci
[appveyor-badge]: https://img.shields.io/appveyor/ci/zchee/go-xdgbasedir/master.svg?label=AppVeyor&style=for-the-badge&logo=appveyor
[godoc-badge]: https://img.shields.io/badge/godoc-reference-4F73B3.svg?style=for-the-badge&label=GODOC.ORG&logoWidth=25&logo=data%3Aimage%2Fsvg%2Bxml%3Bcharset%3Dutf-8%3Bbase64%2CPHN2ZyB4bWxucz0iaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciIHdpZHRoPSI0MCIgaGVpZ2h0PSI0MCIgdmlld0JveD0iODUgNTUgMTIwIDEyMCI+PHBhdGggZmlsbD0iIzJEQkNBRiIgZD0iTTQwLjIgMTAxLjFjLS40IDAtLjUtLjItLjMtLjVsMi4xLTIuN2MuMi0uMy43LS41IDEuMS0uNWgzNS43Yy40IDAgLjUuMy4zLjZsLTEuNyAyLjZjLS4yLjMtLjcuNi0xIC42bC0zNi4yLS4xek0yNS4xIDExMC4zYy0uNCAwLS41LS4yLS4zLS41bDIuMS0yLjdjLjItLjMuNy0uNSAxLjEtLjVoNDUuNmMuNCAwIC42LjMuNS42bC0uOCAyLjRjLS4xLjQtLjUuNi0uOS42bC00Ny4zLjF6TTQ5LjMgMTE5LjVjLS40IDAtLjUtLjMtLjMtLjZsMS40LTIuNWMuMi0uMy42LS42IDEtLjZoMjBjLjQgMCAuNi4zLjYuN2wtLjIgMi40YzAgLjQtLjQuNy0uNy43bC0yMS44LS4xek0xNTMuMSA5OS4zYy02LjMgMS42LTEwLjYgMi44LTE2LjggNC40LTEuNS40LTEuNi41LTIuOS0xLTEuNS0xLjctMi42LTIuOC00LjctMy44LTYuMy0zLjEtMTIuNC0yLjItMTguMSAxLjUtNi44IDQuNC0xMC4zIDEwLjktMTAuMiAxOSAuMSA4IDUuNiAxNC42IDEzLjUgMTUuNyA2LjguOSAxMi41LTEuNSAxNy02LjYuOS0xLjEgMS43LTIuMyAyLjctMy43aC0xOS4zYy0yLjEgMC0yLjYtMS4zLTEuOS0zIDEuMy0zLjEgMy43LTguMyA1LjEtMTAuOS4zLS42IDEtMS42IDIuNS0xLjZoMzYuNGMtLjIgMi43LS4yIDUuNC0uNiA4LjEtMS4xIDcuMi0zLjggMTMuOC04LjIgMTkuNi03LjIgOS41LTE2LjYgMTUuNC0yOC41IDE3LTkuOCAxLjMtMTguOS0uNi0yNi45LTYuNi03LjQtNS42LTExLjYtMTMtMTIuNy0yMi4yLTEuMy0xMC45IDEuOS0yMC43IDguNS0yOS4zIDcuMS05LjMgMTYuNS0xNS4yIDI4LTE3LjMgOS40LTEuNyAxOC40LS42IDI2LjUgNC45IDUuMyAzLjUgOS4xIDguMyAxMS42IDE0LjEuNi45LjIgMS40LTEgMS43eiIvPjxwYXRoIGZpbGw9IiMyREJDQUYiIGQ9Ik0xODYuMiAxNTQuNmMtOS4xLS4yLTE3LjQtMi44LTI0LjQtOC44LTUuOS01LjEtOS42LTExLjYtMTAuOC0xOS4zLTEuOC0xMS4zIDEuMy0yMS4zIDguMS0zMC4yIDcuMy05LjYgMTYuMS0xNC42IDI4LTE2LjcgMTAuMi0xLjggMTkuOC0uOCAyOC41IDUuMSA3LjkgNS40IDEyLjggMTIuNyAxNC4xIDIyLjMgMS43IDEzLjUtMi4yIDI0LjUtMTEuNSAzMy45LTYuNiA2LjctMTQuNyAxMC45LTI0IDEyLjgtMi43LjUtNS40LjYtOCAuOXptMjMuOC00MC40Yy0uMS0xLjMtLjEtMi4zLS4zLTMuMy0xLjgtOS45LTEwLjktMTUuNS0yMC40LTEzLjMtOS4zIDIuMS0xNS4zIDgtMTcuNSAxNy40LTEuOCA3LjggMiAxNS43IDkuMiAxOC45IDUuNSAyLjQgMTEgMi4xIDE2LjMtLjYgNy45LTQuMSAxMi4yLTEwLjUgMTIuNy0xOS4xeiIvPjwvc3ZnPg==
[codecov-badge]: https://img.shields.io/codecov/c/github/zchee/go-xdgbasedir/master.svg?style=for-the-badge&label=CODECOV&logo=codecov
[release-badge]: https://img.shields.io/github/release/zchee/go-xdgbasedir/all.svg?style=for-the-badge&logo=github
[xdg-badge]: https://img.shields.io/badge/%20XDG%20Base%20Dir%20-%200.8-blue.svg?style=for-the-badge
