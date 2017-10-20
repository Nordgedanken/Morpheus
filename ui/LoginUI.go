package ui

import (
	"sync"

	"github.com/Nordgedanken/Morpheus/matrix"
	"github.com/Nordgedanken/Morpheus/util"
	"github.com/matrix-org/gomatrix"
	"github.com/therecipe/qt/core"
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

// InitLogger adds a new logger to the LoginUI struct
func (l *LoginUI) InitLogger() error {
	localLog := util.Logger()
	localLog, _, err := util.StartFileLog(localLog)
	if err != nil {
		return err
	}
	l.localLog = localLog
	return nil
}

// NewUI initializes a new login Screen
func (l *LoginUI) NewUI() (err error) {
	widget := widgets.NewQWidget(nil, 0)
	widget.SetObjectName("LoginWrapper")
	widget.SetStyleSheet("QWidget#LoginWrapper { border: 0px; };")
	topLayout := widgets.NewQVBoxLayout()

	formWidget := widgets.NewQWidget(nil, 0)
	formWrapper := widgets.NewQHBoxLayout()

	formLayout := widgets.NewQVBoxLayout()
	formLayout.SetSpacing(20)
	formLayout.SetContentsMargins(0, 0, 0, 30)
	formWidget.SetLayout(formLayout)

	formWrapper.AddStretch(1)
	formWrapper.AddWidget(formWidget, 0, 0)
	formWrapper.AddStretch(1)

	// UsernameInput
	usernameInput := widgets.NewQLineEdit(nil)
	usernameInput.SetPlaceholderText("Insert MXID")

	usernameLayout := widgets.NewQHBoxLayout()
	usernameLayout.AddWidget(usernameInput, 0, core.Qt__AlignVCenter)

	// PasswordInput
	passwordInput := widgets.NewQLineEdit(nil)
	passwordInput.SetPlaceholderText("Insert password")
	passwordInput.SetEchoMode(widgets.QLineEdit__Password)

	passwordLayout := widgets.NewQHBoxLayout()
	passwordLayout.AddWidget(passwordInput, 0, core.Qt__AlignVCenter)

	formLayout.AddLayout(usernameLayout, 0)
	formLayout.AddLayout(passwordLayout, 0)

	// loginButton
	buttonLayout := widgets.NewQHBoxLayout()
	buttonLayout.SetSpacing(0)
	buttonLayout.SetContentsMargins(0, 0, 0, 30)

	loginButton := widgets.NewQPushButton2("LOGIN", nil)
	loginButton.SetMinimumSize2(350, 65)

	buttonLayout.AddStretch(1)
	buttonLayout.AddWidget(loginButton, 0, 0)
	buttonLayout.AddStretch(1)

	topLayout.AddStretch(1)
	topLayout.AddLayout(formWrapper, 0)
	topLayout.AddStretch(1)
	topLayout.AddLayout(buttonLayout, 0)
	topLayout.AddStretch(1)

	widget.SetLayout(topLayout)

	usernameInput.ConnectTextChanged(func(value string) {
		l.username = value
	})

	passwordInput.ConnectTextChanged(func(value string) {
		l.password = value
	})

	loginButton.ConnectClicked(func(_ bool) {
		LoginErr := l.login()
		if err != nil {
			err = LoginErr
			return
		}
	})

	widget.SetWindowTitle("Morpheus - Login")

	l.widget = widget
	return
}

func (l *LoginUI) login() (err error) {
	//TODO register enter and show loader or so

	var wg sync.WaitGroup

	if l.username != "" && l.password != "" {
		l.localLog.Println("Starting Login Sequenze in background")
		results := make(chan *gomatrix.Client)

		wg.Add(1)
		go matrix.DoLogin(l.username, l.password, "", "", "", l.localLog, results, &wg)

		go func() {
			wg.Wait()      // wait for each execTask to return
			close(results) // then close the results channel
		}()

		//Show MainUI
		for result := range results {
			//TODO Don't switch screen on wrong login data.
			l.cli = result
			MainUIStruct := NewMainUIStructWithExistingConfig(l.config, l.window)
			MainUILoggerInitErr := MainUIStruct.InitLogger()
			if MainUILoggerInitErr != nil {
				err = MainUILoggerInitErr
				return
			}
			mainUIErr := MainUIStruct.NewUI()
			if mainUIErr != nil {
				err = mainUIErr
				return
			}
			l.window.SetCentralWidget(MainUIStruct.GetWidget())
			l.window.Resize(l.widget.Size())
		}
	} else {
		l.localLog.Println("Username and/or password is empty. Do Nothing.")
	}
	return
}
