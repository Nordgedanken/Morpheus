package rooms

/*
#cgo CFLAGS: -fno-keep-inline-dllexport -O2 -Wextra -Wall -W -DUNICODE -D_UNICODE -DQT_NO_DEBUG -DQT_WIDGETS_LIB -DQT_GUI_LIB -DQT_CORE_LIB -DQT_NEEDS_QMAIN
#cgo CXXFLAGS: -fno-keep-inline-dllexport -O2 -std=gnu++11 -Wextra -Wall -W -fexceptions -mthreads -DUNICODE -D_UNICODE -DQT_NO_DEBUG -DQT_WIDGETS_LIB -DQT_GUI_LIB -DQT_CORE_LIB -DQT_NEEDS_QMAIN
#cgo CXXFLAGS: -I../../matrix -I. -IC:/Qt/Qt5.9.1/5.9.3/mingw53_32/include -IC:/Qt/Qt5.9.1/5.9.3/mingw53_32/include/QtWidgets -IC:/Qt/Qt5.9.1/5.9.3/mingw53_32/include/QtGui -IC:/Qt/Qt5.9.1/5.9.3/mingw53_32/include/QtANGLE -IC:/Qt/Qt5.9.1/5.9.3/mingw53_32/include/QtCore -Irelease -IC:/Qt/Qt5.9.1/5.9.3/mingw53_32/mkspecs/win32-g++
#cgo LDFLAGS:        -Wl,-s -Wl,-subsystem,windows -mthreads
#cgo LDFLAGS:        -lmingw32 -LC:/Qt/Qt5.9.1/5.9.3/mingw53_32/lib C:/Qt/Qt5.9.1/5.9.3/mingw53_32/lib/libqtmain.a -LC:/utils/my_sql/my_sql/lib -LC:/utils/postgresql/pgsql/lib -lshell32 C:/Qt/Qt5.9.1/5.9.3/mingw53_32/lib/libQt5Widgets.a C:/Qt/Qt5.9.1/5.9.3/mingw53_32/lib/libQt5Gui.a C:/Qt/Qt5.9.1/5.9.3/mingw53_32/lib/libQt5Core.a
#cgo LDFLAGS: -Wl,--allow-multiple-definition
*/
import "C"
