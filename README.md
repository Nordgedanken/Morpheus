# Morpheus
[![CircleCI](https://circleci.com/gh/Nordgedanken/Morpheus.svg?style=svg)](https://circleci.com/gh/Nordgedanken/Morpheus)  [![Build status](https://ci.appveyor.com/api/projects/status/a0ke0029ely9w7hu?svg=true)](https://ci.appveyor.com/project/MTRNord/neo)

A Matrix client written in Go-QT

# How to build

1. Follow https://github.com/therecipe/qt#installation
2. Run `go get -u github.com/Nordgedanken/Morpheus` to clone this Repo
3. Run `qtdeploy build desktop .` inside  `$GOPATH/src/github.com/Nordgedanken/Morpheus`
4. Run the Application from within `deploy/windows`
