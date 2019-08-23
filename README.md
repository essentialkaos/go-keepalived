<p align="center"><a href="#readme"><img src="https://gh.kaos.st/go-keepalived.svg"/></a></p>

<p align="center">
  <a href="https://godoc.org/pkg.re/essentialkaos/keepalived.v1"><img src="https://godoc.org/pkg.re/essentialkaos/keepalived.v1?status.svg"></a>
  <a href="https://goreportcard.com/report/github.com/essentialkaos/keepalived"><img src="https://goreportcard.com/badge/github.com/essentialkaos/keepalived"></a>
  <a href="https://travis-ci.org/essentialkaos/keepalived"><img src="https://travis-ci.org/essentialkaos/keepalived.svg"></a>
  <a href='https://coveralls.io/github/essentialkaos/keepalived?branch=master'><img src='https://coveralls.io/repos/github/essentialkaos/keepalived/badge.svg?branch=master' alt='Coverage Status' /></a>
</p>

<p align="center"><a href="#installation">Installation</a> • <a href="#build-status">Build Status</a> • <a href="#license">License</a></p>

<br/>

`keepalived` is a very simple Go package for reading virtual IP info from keepalived configuration file.

### Installation

Before the initial install allows git to use redirects for [pkg.re](https://github.com/essentialkaos/pkgre) service (_reason why you should do this described [here](https://github.com/essentialkaos/pkgre#git-support)_):

```
git config --global http.https://pkg.re.followRedirects true
```

Make sure you have a working Go 1.11+ workspace (_[instructions](https://golang.org/doc/install)_), then:

```
go get pkg.re/essentialkaos/keepalived.v1
```

For update to the latest stable release, do:

```
go get -u pkg.re/essentialkaos/keepalived.v1
```

### Build Status

| Branch | Status |
|--------|--------|
| `master` | [![Build Status](https://travis-ci.org/essentialkaos/keepalived.svg?branch=master)](https://travis-ci.org/essentialkaos/keepalived) |
| `develop` | [![Build Status](https://travis-ci.org/essentialkaos/keepalived.svg?branch=develop)](https://travis-ci.org/essentialkaos/keepalived) |

### License

[EKOL](https://essentialkaos.com/ekol)

<p align="center"><a href="https://essentialkaos.com"><img src="https://gh.kaos.st/ekgh.svg"/></a></p>
