#!/bin/bash
set -ev

#check env
df -h

ls $HOME/*
du -sh $HOME/*

#needed for headless qt installation
export QT_QPA_PLATFORM=minimal

#additional deps for multimedia
sudo apt-get -y install libpulse-dev && sudo apt-get clean

#replace gcc4 with gcc5
echo "deb http://ftp.us.debian.org/debian unstable main contrib non-free" | sudo tee --append /etc/apt/sources.list.d/unstable.list
sudo apt-get update
sudo apt-get install -t unstable gcc-5 g++-5 && sudo apt-get clean
sudo rm -f /etc/apt/sources.list.d/unstable.list
sudo rm -f /usr/bin/gcc; sudo ln -s /usr/bin/gcc-5 /usr/bin/gcc
sudo rm -f /usr/bin/g++; sudo ln -s /usr/bin/g++-5 /usr/bin/g++

#download and install qt
QT=qt-unified-linux-x64-online.run
curl -sL --retry 10 --retry-delay 10 -o /tmp/$QT https://download.qt.io/official_releases/online_installers/$QT
chmod +x /tmp/$QT
/tmp/$QT --script $GOPATH/src/github.com/therecipe/qt/internal/ci/iscript.qs
rm -f /tmp/$QT

#prepare env
sudo chown circleci /usr/local/bin
sudo chown circleci $GOROOT/pkg | true

#check env
df -h

ls $HOME/*
du -sh $HOME/*

exit 0
