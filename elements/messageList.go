package elements

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/uitools"
	"github.com/therecipe/qt/widgets"
)

type QVBoxLayoutWithTriggerSlot struct {
	widgets.QVBoxLayout

	_ func(messageBody string) `slot:"TriggerMessage"`
}

func NewMessageList() (messageView *widgets.QWidget, messageViewLayout *QVBoxLayoutWithTriggerSlot) {
	messageView = widgets.NewQWidget(nil, 0)

	messageViewLayout = NewQVBoxLayoutWithTriggerSlot2(messageView)

	messageViewLayout.SetSpacing(0)
	messageViewLayout.SetContentsMargins(0, 0, 0, 0)

	return
}

func (messageViewLayout *QVBoxLayoutWithTriggerSlot) NewMessage(messageBody string, widthScrollArea *widgets.QScrollArea, chatWidget *widgets.QWidget) {
	var widget = widgets.NewQWidget(nil, 0)

	var loader = uitools.NewQUiLoader(nil)
	var file = core.NewQFile2(":/qml/ui/message.ui")

	file.Open(core.QIODevice__ReadOnly)
	var wrapperWidget = loader.Load(file, widget)
	file.Close()

	messageWidget := widgets.NewQWidgetFromPointer(widget.FindChild("message", core.Qt__FindChildrenRecursively).Pointer())
	message := widgets.NewQLabelFromPointer(widget.FindChild("messageContent", core.Qt__FindChildrenRecursively).Pointer())

	chatWidget.ConnectResizeEvent(func(event *gui.QResizeEvent) {
		messageWidget.SetMinimumSize2(widthScrollArea.Size().Width(), wrapperWidget.Size().Height())
		message.SetMinimumWidth(widthScrollArea.Size().Width())
	})

	message.SetText(messageBody)
	message.SetMinimumWidth(widthScrollArea.Size().Width())

	messageWidget.SetMinimumSize2(widthScrollArea.Size().Width(), wrapperWidget.Size().Height())
	messageWidget.Resize(wrapperWidget.Size())

	messageViewLayout.SetSpacing(0)

	messageViewLayout.AddWidget(messageWidget, 0, core.Qt__AlignBottom)

	return
}
