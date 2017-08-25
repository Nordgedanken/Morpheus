package elements

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
)

func NewRoomList() (roomListView *widgets.QWidget, roomListViewLayout *widgets.QGridLayout) {
	roomListView = widgets.NewQWidget(nil, 0)

	roomListScroll := widgets.NewQScrollArea(nil)
	roomListScroll.SetObjectName("roomListScroll")
	roomListScroll.SetStyleSheet("QAbstractScrollArea#roomListScroll { border: 0px; };")

	roomListViewLayout = widgets.NewQGridLayout(roomListView)
	roomListViewLayout.SetSpacing(0)
	roomListViewLayout.SetContentsMargins(0, 0, 0, 0)

	roomListScroll.SetWidget(roomListView)
	roomListScroll.SetWidgetResizable(true)

	return
}

func NewRoom(name string, roomListViewLayout *widgets.QGridLayout) {
	roomWidget := widgets.NewQWidget(nil, 0)
	roomLayout := widgets.NewQVBoxLayout2(roomWidget)
	room := widgets.NewQLabel2(name, nil, 0)
	roomLayout.AddWidget(room, 0, core.Qt__AlignTop)
	roomListViewLayout.AddWidget(room, 0, 0, core.Qt__AlignTop)

	return
}
