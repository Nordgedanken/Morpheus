#!/bin/bash
appname=`basename $0 | sed s,\.sh$,,`

export LD_LIBRARY_PATH=$dirname/lib
export QT_PLUGIN_PATH=$dirname/plugins
export QML_IMPORT_PATH=$dirname/qml
export QML2_IMPORT_PATH=$dirname/qml
/usr/bin/$appname "$@"