#!/bin/bash
set -ev

#check env
df -h

ls $HOME/*
du -sh $HOME/*

#needed for headless qt installation
export QT_QPA_PLATFORM=minimal

#additional deps for multimedia
sudo apt-get -y install libpulse-dev software-properties-common && sudo apt-get clean

#replace gcc4 with gcc5
sudo apt-get update

#download and install qt
sudo add-apt-repository -y ppa:beineri/opt-qt58-trusty
sudo apt-get update
sudo apt-get -y install qt583d qt58base qt58canvas3d qt58charts-no-lgpl qt58connectivity qt58creator qt58datavis3d-no-lgpl qt58declarative qt58doc qt58gamepad qt58graphicaleffects qt58imageformats qt58location qt58multimedia qt58qbs qt58quickcontrols qt58quickcontrols2 qt58script qt58scxml qt58sensors qt58serialbus qt58serialport qt58svg qt58tools qt58translations qt58virtualkeyboard-no-lgpl qt58webchannel qt58webengine qt58websockets qt58x11extras qt58xmlpatterns qt58speech qt58networkauth-no-lgpl && sudo apt-get clean

#prepare env
sudo chown $USER /usr/local/bin
sudo chown $USER $GOROOT/pkg | true

#check env
df -h

ls $HOME/*
du -sh $HOME/*

exit 0
