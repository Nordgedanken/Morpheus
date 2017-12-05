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
sudo apt-get update
sudo apt-get -y install qt3d5-dev qtbase5-dev qtconnectivity5-dev qtcreator qtdeclarative5-dev qt5-doc qtlocation5-dev qtmultimedia5-dev qtquickcontrols2-5-dev qtquick1-5-dev qtscript5-dev qtbase5-dev-tools qttools5-dev qttranslations5-l10n qtwebengine5-dev qtxmlpatterns5-dev-tools && sudo apt-get clean

#prepare env
sudo chown $USER /usr/local/bin
sudo chown $USER $GOROOT/pkg | true

#check env
df -h

ls $HOME/*
du -sh $HOME/*

exit 0
