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
	RoomViewLayout   *widgets.QVBoxLayout
	ScrollArea       *widgets.QScrollArea
	triggerRoomFuncs []func(roomID string)
	changeRoomFuncs  []func(roomID string)
}

func NewRoomList() *RoomList {
	return &RoomList{}
}

func (r *RoomList) ConnectTriggerRoom(f func(roomID string)) {
	r.triggerRoomFuncs = append(r.triggerRoomFuncs, f)
	return
}

func (r *RoomList) TriggerRoom(roomID string) {
	log.Println(r.triggerRoomFuncs)
	for _, f := range r.triggerRoomFuncs {
		log.Println("Trigger Room")
		f(roomID)
	}
	return
}

func (r *RoomList) ConnectChangeRoom(f func(roomID string)) {
	r.changeRoomFuncs = append(r.changeRoomFuncs, f)
	return
}

func (r *RoomList) ChangeRoom(roomID string) {
	for _, f := range r.changeRoomFuncs {
		f(roomID)
	}
	return
}

// InitRoomListLayout generates a new QRoomVBoxLayoutWithTriggerSlot and adds it to the room scrollArea
func (r *RoomList) InitRoomListLayout() {
	roomViewLayout := widgets.NewQVBoxLayout2(r.ScrollArea.Widget())

	roomViewLayout.SetSpacing(0)
	roomViewLayout.SetContentsMargins(0, 0, 0, 0)
	r.ScrollArea.Widget().SetContentsMargins(0, 0, 0, 0)
	r.ScrollArea.Widget().SetLayout(roomViewLayout)

	r.RoomViewLayout = roomViewLayout

	return
}

// NewRoom adds a new room object to the view
func (r *RoomList) NewRoom(room *rooms.Room) (err error) {
	log.Println("New Room")
	var widget = widgets.NewQWidget(r.ScrollArea, 0)

	var loader = uitools.NewQUiLoader(nil)
	var file = core.NewQFile2(":/qml/ui/room.ui")

	file.Open(core.QIODevice__ReadOnly)
	var wrapperWidget = loader.Load(file, widget)
	file.Close()

	roomAvatarQLabel := widgets.NewQLabelFromPointer(widget.FindChild("roomAvatar", core.Qt__FindChildrenRecursively).Pointer())
	roomName := widgets.NewQLabelFromPointer(widget.FindChild("roomName", core.Qt__FindChildrenRecursively).Pointer())
	/*lastMessageContent := widgets.NewQLabelFromPointer(widget.FindChild("lastMessage", core.Qt__FindChildrenRecursively).Pointer())*/

	roomName.SetText(room.GetRoomName())

	wrapperWidget.Resize2(r.ScrollArea.Widget().Size().Width(), wrapperWidget.Size().Height())
	widget.Resize2(r.ScrollArea.Widget().Size().Width(), wrapperWidget.Size().Height())

	var filterObject = core.NewQObject(nil)
	filterObject.ConnectEventFilter(func(watched *core.QObject, event *core.QEvent) bool {
		if event.Type() == core.QEvent__MouseButtonPress {
			var mouseEvent = gui.NewQMouseEventFromPointer(event.Pointer())

			if mouseEvent.Button() == core.Qt__LeftButton {
				go r.ChangeRoom(room.RoomID)
				return true
			}

			return false
		}

		return false
	})

	r.RoomViewLayout.SetSpacing(0)
	r.RoomViewLayout.SetContentsMargins(0, 0, 0, 0)

	wrapperWidget.InstallEventFilter(filterObject)

	log.Println("r.RoomViewLayout")
	log.Println(r.RoomViewLayout)
	r.RoomViewLayout.InsertWidget(r.RoomViewLayout.Count()+1, wrapperWidget, 0, 0)
	r.ScrollArea.SetWidgetResizable(true)
	r.ScrollArea.Resize2(wrapperWidget.Size().Width(), r.ScrollArea.Widget().Size().Height())
	r.ScrollArea.Widget().Resize2(wrapperWidget.Size().Width(), r.ScrollArea.Widget().Size().Height())

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
