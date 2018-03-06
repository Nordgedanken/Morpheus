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

// MessageList defines the TriggerMessage method to add messages to the View
type MessageList struct {
	widgets.QVBoxLayout
	_            func(messageID string) `slot:"triggerMessage"`
	MessageCount int64
}

// Init generates a new widgets.QVBoxLayout and adds it to the message scrollArea
func (m *MessageList) Init(scrollArea *widgets.QScrollArea) {
	m.SetSpacing(0)
	m.SetContentsMargins(0, 0, 0, 0)
	scrollArea.Widget().SetContentsMargins(0, 0, 0, 0)
	scrollArea.Widget().SetLayout(m)
	return
}

// NewMessage adds a new message object to the view
func (m *MessageList) NewMessage(message *messages.Message, scrollArea *widgets.QScrollArea, own bool, height, width int) (err error) {
	barAtBottom := false
	bar := scrollArea.VerticalScrollBar()
	log.Println("BarVal: ", bar.Value())
	log.Println("BarMin: ", bar.Minimum())
	log.Println("BarMax: ", bar.Maximum())
	if bar.Value() == bar.Minimum() {
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

	m.InsertWidget(m.Count()+1, wrapperWidget, 0, 0)
	scrollArea.SetWidgetResizable(true)
	scrollArea.Resize2(wrapperWidget.Size().Width(), scrollArea.Widget().Size().Height())
	scrollArea.Widget().Resize2(wrapperWidget.Size().Width(), scrollArea.Widget().Size().Height())

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

		return
	})
	//TODO: Debug width and height
	log.Println("Width: ", width)
	log.Println("Height: ", height)
	var aWidth int
	if (width + 10) > scrollArea.Widget().Width() {
		aWidth = scrollArea.Widget().Width()
	} else {
		aWidth = width
	}
	log.Println("aWidth: ", aWidth)

	messageContent.SetMinimumWidth(aWidth + 10)

	messageWidget.SetMinimumWidth(aWidth)
	messageWidget.Resize2(aWidth, height+10)

	go message.GetUserAvatar()

	log.Println("BarVal: ", bar.Value())
	log.Println("BarMin: ", bar.Minimum())
	log.Println("BarMax: ", bar.Maximum())
	if barAtBottom {
		bar.Update()
		bar.SetValue(bar.Minimum())
	}

	return
}
