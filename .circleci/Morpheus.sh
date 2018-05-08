#!/bin/bash
appname=`basename $0 | sed s,\.sh$,,`

export LD_LIBRARY_PATH=/usr/local/lib
export QT_PLUGIN_PATH=/usr/local/Morpheus/plugins
export QML_IMPORT_PATH=/usr/local/Morpheus/qml
export QML2_IMPORT_PATH=/usr/local/Morpheus/qml
/usr/bin/$appname "$@"