package elements

import (
	"github.com/Nordgedanken/Morpheus/matrix"
	"github.com/therecipe/qt/core"
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

	_ func(roomID string) `slot:"TriggerRoom"`
}

// NewRoomList generates a new QRoomVBoxLayoutWithTriggerSlot and adds it to the room scrollArea
func NewRoomList(scrollArea *widgets.QScrollArea, roomView *widgets.QWidget) (roomViewLayout *QRoomVBoxLayoutWithTriggerSlot) {
	roomViewLayout = NewQRoomVBoxLayoutWithTriggerSlot2(roomView)

	roomViewLayout.SetSpacing(0)
	roomViewLayout.SetContentsMargins(0, 0, 0, 0)
	roomView.SetContentsMargins(0, 0, 0, 0)
	scrollArea.SetWidget(roomView)
	scrollArea.Widget().SetLayout(roomViewLayout)

	return
}

// NewRoom adds a new room object to the view
func (roomViewLayout *QRoomVBoxLayoutWithTriggerSlot) NewRoom(room *matrix.Room, scrollArea *widgets.QScrollArea) (err error) {
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

	roomWidget := widgets.NewQWidgetFromPointer(widget.FindChild("room", core.Qt__FindChildrenRecursively).Pointer())
	roomAvatarQLabel := widgets.NewQLabelFromPointer(widget.FindChild("roomAvatar", core.Qt__FindChildrenRecursively).Pointer())
	roomName := widgets.NewQLabelFromPointer(widget.FindChild("roomName", core.Qt__FindChildrenRecursively).Pointer())
	/*lastMessageContent := widgets.NewQLabelFromPointer(widget.FindChild("lastMessage", core.Qt__FindChildrenRecursively).Pointer())*/

	roomAvatarQLabel.SetPixmap(roomAvatar)
	roomName.SetText(room.GetRoomName())

	/*
		messageContent.SetText(markdownMessage)

		messageContent.SetMinimumWidth(messageContent.LineWidth())

		roomWidget.SetMinimumWidth(messageContent.LineWidth() + 100)
	*/
	roomWidget.Resize(wrapperWidget.Size())

	roomViewLayout.SetSpacing(1)
	roomViewLayout.SetContentsMargins(0, 0, 0, 0)

	roomViewLayout.AddWidget(roomWidget, 0, core.Qt__AlignBottom)
	scrollArea.Widget().SetLayout(roomViewLayout)

	return
}
