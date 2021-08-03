downloader
============

[![Build Status](https://github.com/jormin/downloaderer/workflows/test/badge.svg?branch=master)](https://github.com/jormin/downloaderer/actions?query=workflow%3Atest)
[![Codecov](https://codecov.io/gh/jormin/downloader/branch/master/graph/badge.svg)](https://codecov.io/gh/jormin/downloader)
[![GoDoc](https://godoc.org/github.com/jormin/downloaderer?status.svg)](http://godoc.org/github.com/jormin/downloaderer)
[![Go Report Card](https://goreportcard.com/badge/github.com/jormin/downloaderer)](https://goreportcard.com/report/github.com/jormin/downloaderer)

This is a tool to download video from third-paty video sites such as bilibili, aiyiqi, youku etc. Only support public free sources, no cracking of vip resources.

Support
-----

- [x] bilibili
- [ ] aiqiyi
- [ ] youku
- [ ] tudou
- [ ] youtube

Usage
-----

```
# clone source code
git clone https://github.com/jormin/downloaderer.git

# download module
go mod download

# install
go install
```

Command
-----

```shell
NAME:
   downloader - This is a tool to download video from third-paty video sites such as bilibili, aiyiqi, youku etc. Only support public free sources, no cracking of vip resources.

USAGE:
   downloader [global options] command [command options] [arguments...]

VERSION:
   v0.0.1

DESCRIPTION:
   A simple tool to manage your todo list

COMMANDS:
   bilibili  download video from bilibili (https://www.bilibili.com/)
   help, h   Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help (default: false)
   --version, -v  print the version (default: false)
```

Example
-----

##### download from bilibili

- save in default directory

```shell script
downloader bilibili BV1Zi4y1x7Q2
```

- save in specified directory with tag `-d`

```shell
downloader bilibili -d ~/Desktop BV1Zi4y1x7Q2
```

License
-------

under the [MIT](./LICENSE) License