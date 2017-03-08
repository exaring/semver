# semver-cli
A command line tool for bumping versions written in [golang](https://github.com/golang/go). Based on [blang/semver](https://github.com/blang/semver)

## Install
`go get github.com/exaring/semver`

Make sure $GOPATH/bin is on your PATH.

## Use
```bash
> semver
Usage: semver [OPTIONS] <version>
OPTIONS:
   -major, -minor, -patch   increase version part
   -build build-name        include additional build name (e.g. alpha)
```

## Example
```bash
> semver -minor 1.0.0
1.1.0

>semver -minor  1.3.19
1.4.0

>semver -minor -build=alpha 1.0.0
1.2.0+alpha
```
