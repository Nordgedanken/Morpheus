package elements

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
)

type QGridLayoutWithTriggerSlot struct {
	widgets.QGridLayout

	_ func(messageBody string) `slot:"TriggerMessage"`
}

func NewMessageList() (messageView *widgets.QWidget, messageViewLayout *QGridLayoutWithTriggerSlot) {
	messageView = widgets.NewQWidget(nil, 0)

	messageScroll := widgets.NewQScrollArea(nil)
	messageScroll.SetObjectName("messageScroll")
	messageScroll.SetStyleSheet("QScrollArea#messageScroll { border: 0px; };")
	messageScroll.SetStyleSheet("QScrollArea#messageScroll { border: 0px; };")

	messageViewLayout = NewQGridLayoutWithTriggerSlot(messageView)

	//messageViewLayout = widgets.NewQGridLayout(messageView)
	messageViewLayout.SetSpacing(0)
	messageViewLayout.SetContentsMargins(0, 0, 0, 0)

	messageScroll.SetWidget(messageView)
	messageScroll.SetWidgetResizable(true)

	return
}

func (messageViewLayout *QGridLayoutWithTriggerSlot) NewMessage(messageBody string) {
	mesageWidget := widgets.NewQWidget(nil, 0)
	mesageWidget.SetStyleSheet("QLabel { border: 0px; }; QWidget { border: 0px; };")

	messageLayout := widgets.NewQVBoxLayout2(mesageWidget)

	message := widgets.NewQLabel2(messageBody, nil, 0)

	message.SetLayout(messageLayout)

	messageViewLayout.AddWidget(message, 0, 0, core.Qt__AlignTop)

	return
}
