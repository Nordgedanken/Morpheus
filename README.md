# Morpheus
[![CircleCI](https://circleci.com/gh/Nordgedanken/Morpheus.svg?style=svg)](https://circleci.com/gh/Nordgedanken/Morpheus)  [![Build status](https://ci.appveyor.com/api/projects/status/a0ke0029ely9w7hu?svg=true)](https://ci.appveyor.com/project/MTRNord/neo)

_CI is known to be broken_

---

A Matrix client written in Go-QT

## Contributing - Matrix Room
Join the Matrix Room at [#Morpheus:matrix.ffslfl.net](https://matrix.to/#/#Morpheus:matrix.ffslfl.net)

## How to build
### Prerequisites
1. https://github.com/therecipe/qt
   * Follow https://github.com/therecipe/qt/wiki/Installation
2. Clone this repo by doing `go get -u github.com/Nordgedanken/Morpheus`

### Build
1. Run `qtdeploy build desktop` inside  `$GOPATH/src/github.com/Nordgedanken/Morpheus`
2. Run the Application from within `deploy/**`


## Versioning
We use [SemVer](http://semver.org/) for versioning. For the versions available, see the [tags on this repository](https://github.com/Nordgedanken/Morpheus/tags).


## License
This project is licensed under the GPLv3 License - see the [LICENSE](LICENSE) file for details

## Acknowledgments
* Inspired by [nheko](http://github.com/mujx/nheko)
