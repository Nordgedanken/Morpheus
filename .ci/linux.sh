#!/bin/bash
set -ev

#check env
df -h

ls $HOME/*
du -sh $HOME/*

#needed for headless qt installation
export QT_QPA_PLATFORM=minimal

#additional deps for multimedia
sudo apt-get -y -qq install libpulse-dev && sudo apt-get -qq clean

#replace gcc4 with gcc5
sudo add-apt-repository -y ppa:ubuntu-toolchain-r/test
sudo apt-get -qq update
sudo apt-get -y -qq install gcc-5 g++-5 && sudo apt-get -qq clean
sudo rm -f /usr/bin/gcc; sudo ln -s /usr/bin/gcc-5 /usr/bin/gcc
sudo rm -f /usr/bin/g++; sudo ln -s /usr/bin/g++-5 /usr/bin/g++

#download and install qt
sudo add-apt-repository -y ppa:beineri/opt-qt58-trusty
sudo apt-get -qq update
sudo apt-get -y -qq install qt583d qt58base qt58canvas3d qt58charts-no-lgpl qt58connectivity qt58creator qt58datavis3d-no-lgpl qt58declarative qt58doc qt58gamepad qt58graphicaleffects qt58imageformats qt58location qt58multimedia qt58qbs qt58quickcontrols qt58quickcontrols2 qt58script qt58scxml qt58sensors qt58serialbus qt58serialport qt58svg qt58tools qt58translations qt58virtualkeyboard-no-lgpl qt58webchannel qt58webengine qt58websockets qt58x11extras qt58xmlpatterns qt58speech qt58networkauth-no-lgpl && sudo apt-get -qq clean
QT=qt-opensource-linux-x64-5.8.0.run
curl -sL --retry 10 --retry-delay 10 -o /tmp/$QT https://download.qt.io/official_releases/qt/5.8/5.8.0/$QT
chmod +x /tmp/$QT
/tmp/$QT --script $GOPATH/src/github.com/therecipe/qt/internal/ci/iscript.qs
rm -f /tmp/$QT

#prepare env
sudo chown $USER /usr/local/bin
sudo chown $USER $GOROOT/pkg | true

#check env
df -h

ls $HOME/*
du -sh $HOME/*

exit 0
