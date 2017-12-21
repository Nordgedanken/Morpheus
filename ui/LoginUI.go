package ui

import (
	"sync"

	"github.com/Nordgedanken/Morpheus/matrix"
	"github.com/Nordgedanken/Morpheus/matrix/scalar"
	"github.com/matrix-org/gomatrix"
	log "github.com/sirupsen/logrus"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/uitools"
	"github.com/therecipe/qt/widgets"
)

// NewLoginUIStruct gives you a LoginUI struct with prefilled data
func NewLoginUIStruct(windowWidth, windowHeight int, window *widgets.QMainWindow) (loginUIStruct LoginUI) {
	configStruct := config{
		windowWidth:  windowWidth,
		windowHeight: windowHeight,
	}
	loginUIStruct = LoginUI{
		config: configStruct,
		window: window,
	}
	return
}

// NewLoginUIStructWithExistingConfig gives you a LoginUI struct with prefilled data and data from a previous Config
func NewLoginUIStructWithExistingConfig(configStruct config, window *widgets.QMainWindow) (loginUIStruct LoginUI) {
	loginUIStruct = LoginUI{
		config: configStruct,
		window: window,
	}
	return
}

// GetWidget gives you the widget of the LoginUI struct
func (l *LoginUI) GetWidget() (widget *widgets.QWidget) {
	widget = l.widget
	return
}

// NewUI initializes a new login Screen
func (l *LoginUI) NewUI() (err error) {
	l.widget = widgets.NewQWidget(nil, 0)

	var loader = uitools.NewQUiLoader(nil)
	var file = core.NewQFile2(":/qml/ui/login.ui")

	file.Open(core.QIODevice__ReadOnly)
	l.LoginWidget = loader.Load(file, l.widget)
	file.Close()

	// UsernameInput
	usernameInput := widgets.NewQLineEditFromPointer(l.widget.FindChild("UsernameInput", core.Qt__FindChildrenRecursively).Pointer())

	// PasswordInput
	passwordInput := widgets.NewQLineEditFromPointer(l.widget.FindChild("PasswordInput", core.Qt__FindChildrenRecursively).Pointer())

	// loginButton
	loginButton := widgets.NewQPushButtonFromPointer(l.widget.FindChild("LoginButton", core.Qt__FindChildrenRecursively).Pointer())

	var layout = widgets.NewQHBoxLayout()
	l.window.SetLayout(layout)
	layout.InsertWidget(0, l.LoginWidget, 0, core.Qt__AlignTop|core.Qt__AlignLeft)
	layout.SetSpacing(0)
	layout.SetContentsMargins(0, 0, 0, 0)
	l.widget.SetSizePolicy2(widgets.QSizePolicy__Expanding, widgets.QSizePolicy__Expanding)
	l.LoginWidget.SetSizePolicy2(widgets.QSizePolicy__Expanding, widgets.QSizePolicy__Expanding)

	l.widget.ConnectResizeEvent(func(event *gui.QResizeEvent) {
		l.LoginWidget.Resize(event.Size())
		event.Accept()
	})

	usernameInput.ConnectTextChanged(func(value string) {
		if usernameInput.StyleSheet() == "border: 1px solid red" {
			usernameInput.SetStyleSheet("")
		}
		l.username = value
	})

	passwordInput.ConnectTextChanged(func(value string) {
		if passwordInput.StyleSheet() == "border: 1px solid red" {
			passwordInput.SetStyleSheet("")
		}
		l.password = value
	})

	loginButton.ConnectClicked(func(_ bool) {
		if l.username != "" {
			if l.password != "" {
				LoginErr := l.login()
				if LoginErr != nil {
					err = LoginErr
					return
				}
			} else {
				passwordInput.SetStyleSheet("border: 1px solid red")
			}
		} else {
			usernameInput.SetStyleSheet("border: 1px solid red")
		}
	})

	usernameInput.ConnectKeyPressEvent(func(ev *gui.QKeyEvent) {
		if int(ev.Key()) == int(core.Qt__Key_Enter) || int(ev.Key()) == int(core.Qt__Key_Return) {
			if l.password != "" {
				LoginErr := l.login()
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
			if l.username != "" {
				LoginErr := l.login()
				if LoginErr != nil {
					err = LoginErr
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

	l.LoginWidget.SetWindowTitle("Morpheus - Login")

	return
}

func (l *LoginUI) login() (err error) {
	//TODO register enter and show loader or so

	var wg sync.WaitGroup

	if l.username != "" && l.password != "" {
		log.Infoln("Starting Login Sequenze in background")
		results := make(chan *gomatrix.Client)

		wg.Add(1)
		go matrix.DoLogin(l.username, l.password, "", "", "", results, &wg)

		go func() {
			wg.Wait()      // wait for each execTask to return
			close(results) // then close the results channel
		}()

		//Show MainUI
		for result := range results {
			//TODO Don't switch screen on wrong login data.
			l.cli = result
			MainUIStruct := NewMainUIStructWithExistingConfig(l.config, l.window)
			mainUIErr := MainUIStruct.NewUI()
			if mainUIErr != nil {
				err = mainUIErr
				return
			}

			scalar.ReqAndSaveScalarToken(MainUIStruct.GetCli())

			l.window.SetCentralWidget(MainUIStruct.GetWidget())
			l.window.Resize(l.widget.Size())
		}
	} else {
		log.Warningln("Username and/or password is empty. Do Nothing.")
	}
	return
}
