package listLayouts

import (
	"time"

	"github.com/Nordgedanken/Morpheus/matrix/messages"
	"github.com/rhinoman/go-commonmark"
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

// QVBoxLayoutWithTriggerSlot defines the QVBoxLayout with TriggerMessage slot to add messages to the View
type QVBoxLayoutWithTriggerSlot struct {
	widgets.QVBoxLayout

	_ func(message *messages.Message) `signal:"TriggerMessage"`
}

// NewMessageList generates a new QVBoxLayoutWithTriggerSlot and adds it to the message scrollArea
func NewMessageList(scrollArea *widgets.QScrollArea) (messageViewLayout *QVBoxLayoutWithTriggerSlot) {
	messageViewLayout = NewQVBoxLayoutWithTriggerSlot2(scrollArea.Widget())

	messageViewLayout.SetSpacing(0)
	messageViewLayout.AddStretch(1)
	messageViewLayout.SetContentsMargins(15, 0, 15, 15)
	scrollArea.Widget().SetContentsMargins(0, 0, 0, 0)
	scrollArea.SetAlignment(core.Qt__AlignLeading | core.Qt__AlignLeft | core.Qt__AlignVCenter)
	scrollArea.Widget().SetLayout(messageViewLayout)

	return
}

// NewMessage adds a new message object to the view
func (messageViewLayout *QVBoxLayoutWithTriggerSlot) NewMessage(message *messages.Message, scrollArea *widgets.QScrollArea, own bool) (err error) {
	barAtBottom := false
	bar := scrollArea.VerticalScrollBar()
	if bar.Value() == bar.Maximum() {
		barAtBottom = true
	}

	var widget = widgets.NewQWidget(nil, 0)

	var loader = uitools.NewQUiLoader(nil)
	var file *core.QFile
	if own {
		file = core.NewQFile2(":/qml/ui/ownmessage.ui")
	} else {
		file = core.NewQFile2(":/qml/ui/message.ui")
	}

	file.Open(core.QIODevice__ReadOnly)
	var wrapperWidget = loader.Load(file, widget)
	file.Close()

	timestampFormat := time.Unix(0, int64(message.Timestamp)*int64(time.Millisecond))
	timestampString := timestampFormat.Format("15:04:05 - Mon 2.01.2006")

	messageWidget := widgets.NewQWidgetFromPointer(widget.FindChild("message", core.Qt__FindChildrenRecursively).Pointer())
	avatarLogo := widgets.NewQLabelFromPointer(widget.FindChild("avatar", core.Qt__FindChildrenRecursively).Pointer())
	messageContent := widgets.NewQLabelFromPointer(widget.FindChild("messageContent", core.Qt__FindChildrenRecursively).Pointer())
	timestampContent := widgets.NewQLabelFromPointer(widget.FindChild("timestamp", core.Qt__FindChildrenRecursively).Pointer())
	senderContent := widgets.NewQLabelFromPointer(widget.FindChild("sender", core.Qt__FindChildrenRecursively).Pointer())

	markdownMessage := commonmark.Md2Html(message.Message, 0)

	messageContent.SetText(markdownMessage)

	senderDisplayNameResp, _ := message.Cli.GetDisplayName(message.Author)
	var senderDisplayName string
	if senderDisplayNameResp == nil {
		senderDisplayName = message.Author
	} else if senderDisplayNameResp.DisplayName == "" {
		senderDisplayName = message.Author
	} else {
		senderDisplayName = senderDisplayNameResp.DisplayName
	}
	senderContent.SetText(senderDisplayName)
	timestampContent.SetText(timestampString)
	message.ConnectSetAvatar(func(avatar *gui.QPixmap) {
		log.Println("Set Avatar")
		avatarNew := gui.NewQPixmap()
		avatarLogo.ConnectPaintEvent(func(event *gui.QPaintEvent) {
			log.Println("PaintEventAvatar")
			painter := gui.NewQPainter2(avatarLogo)

			aWidth := 61 / 2
			aHeight := 61 / 2

			painter.DrawEllipse3(aWidth, aHeight, 61.0, 61.0)
			painter.DrawPixmap2(avatarLogo.Rect(), avatarNew, avatar.Rect())
			//avatarLogo.Update()
			avatarLogo.SetPixmap(avatarNew)
		})

		avatarLogo.SetPixmap(avatar)
	})

	var lineLength int
	lineLength = messageContent.FontMetrics().Width(messageContent.Text(), -1) - 87
	if lineLength >= scrollArea.Widget().Size().Width() {
		lineLength = scrollArea.Widget().Size().Width() - 20 - 87
	}

	messageContent.SetMinimumWidth(lineLength + 10)

	messageWidget.SetMinimumWidth(lineLength)
	messageWidget.Resize2(lineLength, wrapperWidget.Size().Height())

	messageViewLayout.SetSpacing(1)
	messageViewLayout.SetContentsMargins(0, 0, 0, 0)

	messageViewLayout.InsertWidget(messageViewLayout.Count()+1, messageWidget, 0, core.Qt__AlignBottom)

	if barAtBottom {
		bar.SetValue(bar.Maximum())
	}

	go message.GetUserAvatar()

	return
}
