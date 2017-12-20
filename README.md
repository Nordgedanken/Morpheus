# Morpheus
[![CircleCI branch](https://img.shields.io/circleci/project/github/Nordgedanken/Morpheus/master.svg)](https://circleci.com/gh/Nordgedanken/Morpheus)

[![Github All Releases](https://img.shields.io/github/downloads/Nordgedanken/Morpheus/total.svg)]()
---

A Matrix client written in Go-QT

## Contributing - Matrix Room
Join the Matrix Room at [#Morpheus:matrix.ffslfl.net](https://matrix.to/#/#Morpheus:matrix.ffslfl.net)
Read the [Contribution Guideline](CONTRIBUTING.md)

## How to build
### Prerequisites
1. https://github.com/therecipe/qt
   * Follow https://github.com/therecipe/qt/wiki/Installation
2. Clone this repo by doing `go get -u github.com/Nordgedanken/Morpheus`

### Build
1. Run `qtdeploy build desktop` inside  `$GOPATH/src/github.com/Nordgedanken/Morpheus`
2. Run the Application from within `deploy/**`

## How to Build the Windows Installer
1. Get a Windows PC
2. Follow setup on https://github.com/mh-cbon/go-msi
3. Download the Morpheus.exe from the latest CI Build for your tag into the Source Folder or the build result (it needs to be static linked)
4. Run `go-msi make --msi Morpheus.msi --version 0.0.1 --src scripts/templates`

## Versioning
We use [SemVer](http://semver.org/) for versioning. For the versions available, see the [tags on this repository](https://github.com/Nordgedanken/Morpheus/tags).


## License
This project is licensed under the GPLv3 License - see the [LICENSE](LICENSE) file for details

## Acknowledgments
* Inspired by [nheko](http://github.com/mujx/nheko)
