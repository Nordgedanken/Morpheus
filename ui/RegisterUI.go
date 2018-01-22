package ui

import (
	"github.com/Nordgedanken/Morpheus/matrix/globalTypes"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/uitools"
	"github.com/therecipe/qt/widgets"
)

// NewRegUIStruct gives you a RegUI struct with prefilled data
func NewRegUIStruct(windowWidth, windowHeight int, window *widgets.QMainWindow) (regUIStruct *RegUI) {
	configStruct := globalTypes.Config{
		WindowWidth:  windowWidth,
		WindowHeight: windowHeight,
	}
	regUIStruct = &RegUI{
		Config: configStruct,
		window: window,
	}
	return
}

// NewLoginUIStructWithExistingConfig gives you a LoginUI struct with prefilled data and data from a previous Config
func NewRegUIStructWithExistingConfig(configStruct globalTypes.Config, window *widgets.QMainWindow) (regUIStruct *RegUI) {
	regUIStruct = &RegUI{
		Config: configStruct,
		window: window,
	}
	return
}

// GetWidget gives you the widget of the LoginUI struct
func (r *RegUI) GetWidget() (widget *widgets.QWidget) {
	widget = r.widget
	return
}

// NewUI initializes a new login Screen
func (r *RegUI) NewUI() (err error) {
	r.widget = widgets.NewQWidget(nil, 0)

	var loader = uitools.NewQUiLoader(nil)
	var file = core.NewQFile2(":/qml/ui/register.ui")

	file.Open(core.QIODevice__ReadOnly)
	r.RegWidget = loader.Load(file, r.widget)
	file.Close()

	// UsernameInput
	usernameInput := widgets.NewQLineEditFromPointer(r.widget.FindChild("UsernameInput", core.Qt__FindChildrenRecursively).Pointer())

	// PasswordInput
	passwordInput := widgets.NewQLineEditFromPointer(r.widget.FindChild("PasswordInput", core.Qt__FindChildrenRecursively).Pointer())

	// registernButton
	registernButton := widgets.NewQPushButtonFromPointer(r.widget.FindChild("RegisterButton", core.Qt__FindChildrenRecursively).Pointer())

	var layout = widgets.NewQHBoxLayout()
	r.window.SetLayout(layout)
	layout.InsertWidget(0, r.RegWidget, 0, core.Qt__AlignTop|core.Qt__AlignLeft)
	layout.SetSpacing(0)
	layout.SetContentsMargins(0, 0, 0, 0)
	r.widget.SetSizePolicy2(widgets.QSizePolicy__Expanding, widgets.QSizePolicy__Expanding)
	r.RegWidget.SetSizePolicy2(widgets.QSizePolicy__Expanding, widgets.QSizePolicy__Expanding)

	r.widget.ConnectResizeEvent(func(event *gui.QResizeEvent) {
		r.RegWidget.Resize(event.Size())
		event.Accept()
	})

	usernameInput.ConnectTextChanged(func(value string) {
		if usernameInput.StyleSheet() == "border: 1px solid red" {
			usernameInput.SetStyleSheet("")
		}
		r.Username = value
	})

	passwordInput.ConnectTextChanged(func(value string) {
		if passwordInput.StyleSheet() == "border: 1px solid red" {
			passwordInput.SetStyleSheet("")
		}
		r.Password = value
	})

	registernButton.ConnectClicked(func(_ bool) {
		if r.Username != "" && r.Password != "" {
			LoginErr := r.register()
			if LoginErr != nil {
				err = LoginErr
				return
			}
		} else {
			passwordInput.SetStyleSheet("border: 1px solid red")
		}
	})

	usernameInput.ConnectKeyPressEvent(func(ev *gui.QKeyEvent) {
		if int(ev.Key()) == int(core.Qt__Key_Enter) || int(ev.Key()) == int(core.Qt__Key_Return) {
			if r.Password != "" {
				LoginErr := r.register()
				if LoginErr != nil {
					err = LoginErr
					return
				}

				usernameInput.Clear()
				ev.Accept()
			} else {
				passwordInput.SetStyleSheet("border: 1px solid red")
				ev.Ignore()
			}
		} else {
			usernameInput.KeyPressEventDefault(ev)
			ev.Ignore()
		}
	})

	passwordInput.ConnectKeyPressEvent(func(ev *gui.QKeyEvent) {
		if int(ev.Key()) == int(core.Qt__Key_Enter) || int(ev.Key()) == int(core.Qt__Key_Return) {
			if r.Username != "" {
				RegisterErr := r.register()
				if RegisterErr != nil {
					err = RegisterErr
					return
				}

				passwordInput.Clear()
				ev.Accept()
			} else {
				usernameInput.SetStyleSheet("border: 1px solid red")
				ev.Ignore()
			}
		} else {
			passwordInput.KeyPressEventDefault(ev)
			ev.Ignore()
		}
	})

	r.RegWidget.SetWindowTitle("Morpheus - Register")

	return
}

func (r *RegUI) register() error {
	return nil
}
