package main

import (
	"github.com/rhinoman/go-commonmark"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/uitools"
	"github.com/therecipe/qt/widgets"
)

func AddMessage(body string, avatar *gui.QPixmap) *widgets.QWidget {
	var widget = widgets.NewQWidget(nil, 0)

	var loader = uitools.NewQUiLoader(nil)
	var file = core.NewQFile2(":/qml/ui/message.ui")

	file.Open(core.QIODevice__ReadOnly)
	var messageWidget = loader.Load(file, widget)
	file.Close()

	var layout = widgets.NewQHBoxLayout()
	layout.AddWidget(messageWidget, 1, core.Qt__AlignTop|core.Qt__AlignLeft)
	widget.SetLayout(layout)
	layout.SetSpacing(0)
	layout.SetContentsMargins(0, 0, 0, 0)

	avatarLogo := widgets.NewQLabelFromPointer(widget.FindChild("avatar", core.Qt__FindChildrenRecursively).Pointer())
	messageContent := widgets.NewQLabelFromPointer(widget.FindChild("messageContent", core.Qt__FindChildrenRecursively).Pointer())

	mardownMessage := commonmark.Md2Html(body, 0)

	messageContent.SetText(mardownMessage)
	avatarLogo.SetPixmap(avatar)

	return widget
}
