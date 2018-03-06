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

// RoomList defines the TriggerRoom and ChangeRoom method to add messages to the View
type RoomList struct {
	widgets.QVBoxLayout
	_         func(roomID string) `slot:"triggerRoom"`
	_         func(roomID string) `slot:"changeRoom"`
	RoomCount int64
}

type RoomThread struct {
	core.QThread
	RoomFunc func()
}

func (rt *RoomThread) Run() {
	log.Infoln("Run!")
	rt.RoomFunc()
}

// Init generates a new RoomList and adds it to the room scrollArea
func (r *RoomList) Init(scrollArea *widgets.QScrollArea) {
	r.SetSpacing(0)
	r.SetContentsMargins(0, 0, 0, 0)
	scrollArea.Widget().SetContentsMargins(0, 0, 0, 0)
	scrollArea.Widget().SetLayout(r)

	return
}

// NewRoom adds a new room object to the view
func (r *RoomList) NewRoom(room *rooms.Room, scrollArea *widgets.QScrollArea) (err error) {
	var widget = widgets.NewQWidget(nil, 0)

	var loader = uitools.NewQUiLoader(nil)
	var file = core.NewQFile2(":/qml/ui/room.ui")

	file.Open(core.QIODevice__ReadOnly)
	var wrapperWidget = loader.Load(file, widget)
	file.Close()

	roomAvatarQLabel := widgets.NewQLabelFromPointer(widget.FindChild("roomAvatar", core.Qt__FindChildrenRecursively).Pointer())
	roomName := widgets.NewQLabelFromPointer(widget.FindChild("roomName", core.Qt__FindChildrenRecursively).Pointer())
	/*lastMessageContent := widgets.NewQLabelFromPointer(widget.FindChild("lastMessage", core.Qt__FindChildrenRecursively).Pointer())*/

	roomName.SetText(room.GetRoomName())

	wrapperWidget.Resize2(scrollArea.Widget().Size().Width(), wrapperWidget.Size().Height())
	widget.Resize2(scrollArea.Widget().Size().Width(), wrapperWidget.Size().Height())

	var filterObject = core.NewQObject(nil)
	filterObject.ConnectEventFilter(func(watched *core.QObject, event *core.QEvent) bool {
		if event.Type() == core.QEvent__MouseButtonPress {
			var mouseEvent = gui.NewQMouseEventFromPointer(event.Pointer())

			if mouseEvent.Button() == core.Qt__LeftButton {
				r.ChangeRoom(room.RoomID)
				return true
			}

			return false
		}

		return false
	})

	r.SetSpacing(0)
	r.SetContentsMargins(0, 0, 0, 0)

	wrapperWidget.InstallEventFilter(filterObject)

	r.InsertWidget(r.Count()+1, wrapperWidget, 0, 0)
	scrollArea.SetWidgetResizable(true)
	scrollArea.Resize2(wrapperWidget.Size().Width(), scrollArea.Widget().Size().Height())
	scrollArea.Widget().Resize2(wrapperWidget.Size().Width(), scrollArea.Widget().Size().Height())

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

	room.ConnectSetAvatar(func(IMGdata []byte) {
		avatar := gui.NewQPixmap()

		str := string(IMGdata[:])
		avatar.LoadFromData(str, uint(len(str)), "", 0)

		roomAvatarQLabel.SetPixmap(avatar)

		return
	})

	go room.GetRoomAvatar()

	return
}
