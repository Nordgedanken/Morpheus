package listLayouts

import (
	"github.com/Nordgedanken/Morpheus/matrix/rooms"
	log "github.com/sirupsen/logrus"
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
func (roomViewLayout *QRoomVBoxLayoutWithTriggerSlot) NewRoom(room *rooms.Room, scrollArea *widgets.QScrollArea) (err error) {
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

	roomAvatarQLabel.ConnectSetPixmap(func(vqp *gui.QPixmap) {
		log.Println("SetPixmapEventRoomAvatar")

		vqp.Scaled2(roomAvatarQLabel.Width(), roomAvatarQLabel.Height(), 0, 0)

		newPixmap := gui.NewQPixmap3(2*roomAvatarQLabel.Width(), 2*roomAvatarQLabel.Height())
		newPixmap.Fill(nil)

		painter := gui.NewQPainter2(newPixmap)

		r := gui.NewQRegion2(roomAvatarQLabel.Width()/2, roomAvatarQLabel.Height()/2, roomAvatarQLabel.Width(), roomAvatarQLabel.Height(), gui.QRegion__Ellipse)

		painter.SetClipRegion(r, 0)

		painter.DrawPixmap10(roomAvatarQLabel.Rect(), vqp)
		newImage := newPixmap.ToImage()
		vqp.FromImage(newImage, 0)
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
	scrollArea.Resize2(wrapperWidget.Size().Width(), scrollArea.Widget().Size().Height())

	return
}
