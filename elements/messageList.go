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

// QVBoxLayoutWithTriggerSlot defines the QVBoxLayout with TriggerMessage slot to add messages to the View
type QVBoxLayoutWithTriggerSlot struct {
	widgets.QVBoxLayout

	_ func(messageBody, sender string) `slot:"TriggerMessage"`
}

// NewMessageList generates a new QVBoxLayoutWithTriggerSlot and adds it to the message scrollArea
func NewMessageList(scrollArea *widgets.QScrollArea, messageView *widgets.QWidget) (messageViewLayout *QVBoxLayoutWithTriggerSlot) {
	messageViewLayout = NewQVBoxLayoutWithTriggerSlot2(messageView)

	messageViewLayout.SetSpacing(0)
	messageViewLayout.SetContentsMargins(0, 0, 0, 0)
	messageView.SetContentsMargins(0, 0, 0, 0)
	scrollArea.SetWidget(messageView)
	scrollArea.Widget().SetLayout(messageViewLayout)

	return
}

// NewMessage adds a new message object to the view
func (messageViewLayout *QVBoxLayoutWithTriggerSlot) NewMessage(body string, cli *gomatrix.Client, sender string, scrollArea *widgets.QScrollArea) {
	avatar := matrix.GetUserAvatar(cli, sender)

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

	messageViewLayout.SetSpacing(1)
	messageViewLayout.SetContentsMargins(0, 0, 0, 0)

	messageViewLayout.AddWidget(messageWidget, 0, core.Qt__AlignBottom)
	scrollArea.Widget().SetLayout(messageViewLayout)

	return
}
