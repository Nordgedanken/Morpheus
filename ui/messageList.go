package ui

import (
	"fmt"
	"time"

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

	_ func(messageBody, sender string, timestamp int64) `slot:"TriggerMessage"`
}

// NewMessageList generates a new QVBoxLayoutWithTriggerSlot and adds it to the message scrollArea
func NewMessageList(scrollArea *widgets.QScrollArea, messageView *widgets.QWidget) (messageViewLayout *QVBoxLayoutWithTriggerSlot) {
	messageViewLayout = NewQVBoxLayoutWithTriggerSlot2(messageView)

	messageViewLayout.SetSpacing(0)
	messageViewLayout.SetContentsMargins(0, 0, 0, 0)
	messageView.SetContentsMargins(0, 0, 0, 0)
	scrollArea.SetWidget(messageView)
	scrollArea.SetAlignment(core.Qt__AlignLeading | core.Qt__AlignLeft | core.Qt__AlignVCenter)
	scrollArea.Widget().SetLayout(messageViewLayout)

	return
}

// NewMessage adds a new message object to the view
func (messageViewLayout *QVBoxLayoutWithTriggerSlot) NewMessage(body string, cli *gomatrix.Client, sender string, timestamp int64, scrollArea *widgets.QScrollArea) (err error) {
	avatar, AvatarErr := matrix.GetUserAvatar(cli, sender, 61)
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

	timestampFormat := time.Unix(0, int64(timestamp)*int64(time.Millisecond))
	timestampString := timestampFormat.Format("15:04:05 - Mon 2.01.2006")

	messageWidget := widgets.NewQWidgetFromPointer(widget.FindChild("message", core.Qt__FindChildrenRecursively).Pointer())
	avatarLogo := widgets.NewQLabelFromPointer(widget.FindChild("avatar", core.Qt__FindChildrenRecursively).Pointer())
	messageContent := widgets.NewQLabelFromPointer(widget.FindChild("messageContent", core.Qt__FindChildrenRecursively).Pointer())
	timestampContent := widgets.NewQLabelFromPointer(widget.FindChild("timestamp", core.Qt__FindChildrenRecursively).Pointer())
	senderContent := widgets.NewQLabelFromPointer(widget.FindChild("sender", core.Qt__FindChildrenRecursively).Pointer())

	markdownMessage := commonmark.Md2Html(body, 0)

	messageContent.SetText(markdownMessage)

	senderDisplayNameResp, _ := cli.GetDisplayName(sender)
	var senderDisplayName string
	if senderDisplayNameResp == nil {
		senderDisplayName = sender
	} else if senderDisplayNameResp.DisplayName == "" {
		senderDisplayName = sender
	} else {
		senderDisplayName = senderDisplayNameResp.DisplayName
	}
	fmt.Println(senderDisplayName)
	senderContent.SetText(senderDisplayName)
	timestampContent.SetText(timestampString)
	avatarLogo.SetPixmap(avatar)

	var lineLength int
	lineLength = messageContent.FontMetrics().Width(messageContent.Text(), -1) - 87
	fmt.Println(scrollArea.Widget().Size().Width())
	fmt.Println("lineLengthO: ", lineLength)
	if lineLength >= scrollArea.Widget().Size().Width() {
		lineLength = scrollArea.Widget().Size().Width() - 20 - 87
	}
	fmt.Println("lineLengthN: ", lineLength)

	messageContent.SetMinimumWidth(lineLength + 10)

	messageWidget.SetMinimumWidth(lineLength)
	messageWidget.Resize2(lineLength, wrapperWidget.Size().Height())

	messageViewLayout.SetSpacing(1)
	messageViewLayout.SetContentsMargins(0, 0, 0, 0)

	messageViewLayout.InsertWidget(0, messageWidget, 0, core.Qt__AlignBottom)

	return
}

// NewOwnMessage adds a new message object to the view
func (messageViewLayout *QVBoxLayoutWithTriggerSlot) NewOwnMessage(body string, cli *gomatrix.Client, sender string, timestamp int64, scrollArea *widgets.QScrollArea) (err error) {
	avatar, AvatarErr := matrix.GetUserAvatar(cli, sender, 61)
	if AvatarErr != nil {
		err = AvatarErr
		return
	}

	var widget = widgets.NewQWidget(nil, 0)

	var loader = uitools.NewQUiLoader(nil)
	var file = core.NewQFile2(":/qml/ui/ownmessage.ui")

	file.Open(core.QIODevice__ReadOnly)
	var wrapperWidget = loader.Load(file, widget)
	file.Close()

	timestampFormat := time.Unix(0, int64(timestamp)*int64(time.Millisecond))
	timestampString := timestampFormat.Format("15:04:05 - Mon 2.01.2006")

	messageWidget := widgets.NewQWidgetFromPointer(widget.FindChild("message", core.Qt__FindChildrenRecursively).Pointer())
	avatarLogo := widgets.NewQLabelFromPointer(widget.FindChild("avatar", core.Qt__FindChildrenRecursively).Pointer())
	messageContent := widgets.NewQLabelFromPointer(widget.FindChild("messageContent", core.Qt__FindChildrenRecursively).Pointer())
	timestampContent := widgets.NewQLabelFromPointer(widget.FindChild("timestamp", core.Qt__FindChildrenRecursively).Pointer())
	senderContent := widgets.NewQLabelFromPointer(widget.FindChild("sender", core.Qt__FindChildrenRecursively).Pointer())

	markdownMessage := commonmark.Md2Html(body, 0)

	messageContent.SetText(markdownMessage)

	senderDisplayNameResp, _ := cli.GetDisplayName(sender)
	var senderDisplayName string
	if senderDisplayNameResp == nil {
		senderDisplayName = sender
	} else if senderDisplayNameResp.DisplayName == "" {
		senderDisplayName = sender
	} else {
		senderDisplayName = senderDisplayNameResp.DisplayName
	}
	fmt.Println(senderDisplayName)
	senderContent.SetText(senderDisplayName)
	timestampContent.SetText(timestampString)
	avatarLogo.SetPixmap(avatar)

	messageContent.SetMinimumWidth(messageContent.LineWidth())

	messageWidget.SetMinimumWidth(messageContent.LineWidth() + 100)
	messageWidget.Resize(wrapperWidget.Size())

	messageViewLayout.SetSpacing(0)
	messageViewLayout.SetContentsMargins(0, 0, 0, 0)

	messageViewLayout.InsertWidget(0, messageWidget, 0, core.Qt__AlignBottom)

	return
}
