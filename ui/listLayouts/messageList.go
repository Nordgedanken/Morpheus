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

// messageList defines the TriggerMessage method to add messages to the View
type MessageList struct {
	MessageViewLayout   *widgets.QVBoxLayout
	ScrollArea          *widgets.QScrollArea
	triggerMessageFuncs []func(message *messages.Message)
}

func NewMessageList() *MessageList {
	return &MessageList{}
}

func (m *MessageList) ConnectTriggerMessage(f func(message *messages.Message)) {
	m.triggerMessageFuncs = append(m.triggerMessageFuncs, f)
	return
}

func (m *MessageList) TriggerMessage(message *messages.Message) {
	for _, f := range m.triggerMessageFuncs {
		f(message)
	}
	return
}

// InitMessageListLayout generates a new widgets.QVBoxLayout and adds it to the message scrollArea
func (m *MessageList) InitMessageListLayout() {
	log.Println(m.ScrollArea.Widget())
	messageViewLayout := widgets.NewQVBoxLayout()
	log.Println(messageViewLayout)

	messageViewLayout.SetSpacing(0)
	messageViewLayout.SetContentsMargins(15, 0, 15, 15)
	m.ScrollArea.Widget().SetContentsMargins(0, 0, 0, 0)
	m.ScrollArea.SetAlignment(core.Qt__AlignLeading | core.Qt__AlignLeft | core.Qt__AlignVCenter)
	m.ScrollArea.Widget().SetLayout(messageViewLayout)

	m.MessageViewLayout = messageViewLayout

	return
}

// NewMessage adds a new message object to the view
func (m *MessageList) NewMessage(message *messages.Message, own bool) (err error) {
	barAtBottom := false
	bar := m.ScrollArea.VerticalScrollBar()
	if bar.Value() == bar.Maximum() {
		barAtBottom = true
	}

	var widget = widgets.NewQWidget(m.ScrollArea, 0)

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

	avatarLogo.ConnectSetPixmap(func(vqp *gui.QPixmap) {
		log.Println("SetPixmapEventRoomAvatar")

		vqp.Scaled2(avatarLogo.Width(), avatarLogo.Height(), 0, 0)

		newPixmap := gui.NewQPixmap3(2*avatarLogo.Width(), 2*avatarLogo.Height())
		newPixmap.Fill(nil)

		painter := gui.NewQPainter2(newPixmap)

		r := gui.NewQRegion2(avatarLogo.Width()/2, avatarLogo.Height()/2, avatarLogo.Width(), avatarLogo.Height(), gui.QRegion__Ellipse)

		painter.SetClipRegion(r, 0)

		painter.DrawPixmap10(avatarLogo.Rect(), vqp)
		newImage := newPixmap.ToImage()
		vqp.FromImage(newImage, 0)
	})

	message.ConnectSetAvatar(func(IMGdata []byte) {
		avatar := gui.NewQPixmap()

		str := string(IMGdata[:])

		avatar.LoadFromData(string(str[:]), uint(len(str)), "", 0)

		avatarLogo.SetPixmap(avatar)
	})

	var lineLength int
	lineLength = messageContent.FontMetrics().Width(messageContent.Text(), -1) - 87
	if lineLength >= m.ScrollArea.Widget().Size().Width() {
		lineLength = m.ScrollArea.Widget().Size().Width() - 20 - 87
	}

	messageContent.SetMinimumWidth(lineLength + 10)

	messageWidget.SetMinimumWidth(lineLength)
	messageWidget.Resize2(lineLength, wrapperWidget.Size().Height())

	m.MessageViewLayout.SetSpacing(1)
	m.MessageViewLayout.SetContentsMargins(0, 0, 0, 0)

	log.Println(m.MessageViewLayout)
	m.MessageViewLayout.InsertWidget(m.MessageViewLayout.Count()+1, messageWidget, 0, core.Qt__AlignBottom)

	if barAtBottom {
		bar.SetValue(bar.Maximum())
	}

	go message.GetUserAvatar()

	return
}
