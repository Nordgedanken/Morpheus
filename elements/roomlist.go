package elements

import (
	"github.com/Nordgedanken/Morpheus/matrix"
	"github.com/matrix-org/gomatrix"
	"github.com/rhinoman/go-commonmark"
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

	_ func(messageBody, sender string) `slot:"TriggerMessage"`
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
func (roomViewLayout *QRoomVBoxLayoutWithTriggerSlot) NewRoom(body string, cli *gomatrix.Client, sender string, scrollArea *widgets.QScrollArea) (err error) {
	avatar, AvatarErr := matrix.GetUserAvatar(cli, sender, 84)
	if err != nil {
		err = AvatarErr
		return
	}

	var widget = widgets.NewQWidget(nil, 0)

	var loader = uitools.NewQUiLoader(nil)
	var file = core.NewQFile2(":/qml/ui/message.ui")

	file.Open(core.QIODevice__ReadOnly)
	var wrapperWidget = loader.Load(file, widget)
	file.Close()

	messageWidget := widgets.NewQWidgetFromPointer(widget.FindChild("message", core.Qt__FindChildrenRecursively).Pointer())
	avatarLogo := widgets.NewQLabelFromPointer(widget.FindChild("avatar", core.Qt__FindChildrenRecursively).Pointer())
	messageContent := widgets.NewQLabelFromPointer(widget.FindChild("messageContent", core.Qt__FindChildrenRecursively).Pointer())
	senderContent := widgets.NewQLabelFromPointer(widget.FindChild("sender", core.Qt__FindChildrenRecursively).Pointer())

	markdownMessage := commonmark.Md2Html(body, 0)

	messageContent.SetText(markdownMessage)

	senderDisplayNameResp, _ := cli.GetDisplayName(sender)
	senderDisplayName := senderDisplayNameResp.DisplayName
	senderContent.SetText(senderDisplayName)
	avatarLogo.SetPixmap(avatar)

	messageContent.SetMinimumWidth(messageContent.LineWidth())

	messageWidget.SetMinimumWidth(messageContent.LineWidth() + 100)
	messageWidget.Resize(wrapperWidget.Size())

	roomViewLayout.SetSpacing(1)
	roomViewLayout.SetContentsMargins(0, 0, 0, 0)

	roomViewLayout.AddWidget(messageWidget, 0, core.Qt__AlignBottom)
	scrollArea.Widget().SetLayout(roomViewLayout)

	return
}
