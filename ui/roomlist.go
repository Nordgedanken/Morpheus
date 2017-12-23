package ui

import (
	"github.com/Nordgedanken/Morpheus/matrix"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/uitools"
	"github.com/therecipe/qt/widgets"
)

////////////////////////////////////////////////////
//                                                //
//                                                //
//                                                //
//                                                //
// DON'T TOUCH I HAVE NO IDEA WHY THIS EVEN WORKS //
//                                                //
//                                                //
//                                                //
////////////////////////////////////////////////////

// QRoomVBoxLayoutWithTriggerSlot defines the QVBoxLayout with TriggerMessage slot to add messages to the View
type QRoomVBoxLayoutWithTriggerSlot struct {
	widgets.QVBoxLayout

	_ func(roomID string) `signal:"TriggerRoom"`
	_ func(roomID string) `signal:"ChangeRoom"`
}

// NewRoomList generates a new QRoomVBoxLayoutWithTriggerSlot and adds it to the room scrollArea
func NewRoomList(scrollArea *widgets.QScrollArea) (roomViewLayout *QRoomVBoxLayoutWithTriggerSlot) {
	roomViewLayout = NewQRoomVBoxLayoutWithTriggerSlot2(scrollArea.Widget())

	roomViewLayout.SetSpacing(0)
	roomViewLayout.SetContentsMargins(0, 0, 0, 0)
	scrollArea.Widget().SetContentsMargins(0, 0, 0, 0)
	scrollArea.Widget().SetLayout(roomViewLayout)

	return
}

// NewRoom adds a new room object to the view
func (roomViewLayout *QRoomVBoxLayoutWithTriggerSlot) NewRoom(room *matrix.Room, scrollArea *widgets.QScrollArea, mainUIStruct *MainUI) (err error) {
	roomAvatar, roomAvatarErr := room.GetRoomAvatar()
	if roomAvatarErr != nil {
		err = roomAvatarErr
		return
	}

	var widget = widgets.NewQWidget(nil, 0)

	var loader = uitools.NewQUiLoader(nil)
	var file = core.NewQFile2(":/qml/ui/room.ui")

	file.Open(core.QIODevice__ReadOnly)
	var wrapperWidget = loader.Load(file, widget)
	file.Close()

	roomAvatarQLabel := widgets.NewQLabelFromPointer(widget.FindChild("roomAvatar", core.Qt__FindChildrenRecursively).Pointer())
	roomName := widgets.NewQLabelFromPointer(widget.FindChild("roomName", core.Qt__FindChildrenRecursively).Pointer())
	/*lastMessageContent := widgets.NewQLabelFromPointer(widget.FindChild("lastMessage", core.Qt__FindChildrenRecursively).Pointer())*/

	roomAvatarQLabel.ConnectPaintEvent(func(event *gui.QPaintEvent) {
		painter := gui.NewQPainter()
		painter.SetRenderHint(gui.QPainter__Antialiasing, true)
		r := event.Rect()
		hs := 84 / 2

		brush := gui.NewQBrush()
		brush.SetStyle(core.Qt__SolidPattern)
		brush.SetColor(gui.NewQColor6("white"))

		painter.SetPen3(core.Qt__NoPen)
		painter.SetBrush(brush)
		painter.DrawEllipse5(r.Center(), hs, hs)
	})
	roomAvatarQLabel.SetPixmap(roomAvatar)
	roomName.SetText(room.GetRoomName())

	wrapperWidget.Resize2(scrollArea.Widget().Size().Width(), wrapperWidget.Size().Height())
	widget.Resize2(scrollArea.Widget().Size().Width(), wrapperWidget.Size().Height())

	var filterObject = core.NewQObject(nil)
	filterObject.ConnectEventFilter(func(watched *core.QObject, event *core.QEvent) bool {
		if event.Type() == core.QEvent__MouseButtonPress {
			var mouseEvent = gui.NewQMouseEventFromPointer(event.Pointer())

			if mouseEvent.Button() == core.Qt__LeftButton {
				go roomViewLayout.ChangeRoom(room.RoomID)
				return true
			}

			return false
		}

		return false
	})

	roomViewLayout.SetSpacing(0)
	roomViewLayout.SetContentsMargins(0, 0, 0, 0)

	wrapperWidget.InstallEventFilter(filterObject)

	roomViewLayout.InsertWidget(roomViewLayout.Count()+1, wrapperWidget, 0, 0)

	return
}
