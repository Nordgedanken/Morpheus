package main

import (
	"fmt"
	"log"
	"sync"

	"github.com/Nordgedanken/Neo/matrix"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/uitools"
	"github.com/therecipe/qt/widgets"
)

var username string
var password string

//NewLoginUI initializes the login Screen
func NewLoginUI(windowWidth, windowHeight int) *widgets.QWidget {
	widget := widgets.NewQWidget(nil, 0)
	topLayout := widgets.NewQVBoxLayout()

	formWidget := widgets.NewQWidget(nil, 0)
	formWrapper := widgets.NewQHBoxLayout()
	formWidget.SetMinimumSize2(350, 200)

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

	// UsernameInput-Label
	// usernameLabel := widgets.NewQLabel(nil, 0)
	// usernameLabel.SetText("Username: ")
	// usernameLabel.SetBuddy(usernameInput)
	// layout.AddWidget(usernameLabel, 0, 0)

	// PasswordInput-Label
	// passwordLabel := widgets.NewQLabel(nil, 0)
	// passwordLabel.SetText("Password: ")
	// passwordLabel.SetBuddy(passwordInput)
	// layout.AddWidget(passwordLabel, 0, 0)

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
		username = value
	})

	passwordInput.ConnectTextChanged(func(value string) {
		password = value
	})

	loginButton.ConnectClicked(func(checked bool) {
		//TODO register enter and show loader or so
		if username != "" && password != "" {
			localLog.Println("Starting Login Sequenze in background")
			var wg sync.WaitGroup
			results := make(chan *matrix.Client)

			wg.Add(1)
			go func(username, password string, localLog *log.Logger, results chan<- *matrix.Client) {
				defer wg.Done()
				cli, err := matrix.LoginUser(username, password)
				if err != nil {
					localLog.Println(err)
				}

				results <- cli

			}(username, password, localLog, results)

			go func() {
				wg.Wait()      // wait for each execTask to return
				close(results) // then close the results channel
			}()

			//Show MainUI
			for result := range results {
				//TODO Don't switch screen on wrong login data.
				MainUI := NewMainUI(windowWidth, windowHeight, result)
				window.SetCentralWidget(MainUI)
			}
		} else {
			localLog.Println("Username and/or password is empty. Do Nothing.")
		}
	})

	widget.SetWindowTitle("Neo - Login")

	return widget
}

//NewMainUI initializes the login Screen
func NewMainUI(windowWidth, windowHeight int, cli *matrix.Client) *widgets.QWidget {
	widget := widgets.NewQWidget(nil, 0)

	loader := uitools.NewQUiLoader(nil)
	file := core.NewQFile2(":/qml/Main.ui")

	file.Open(core.QIODevice__ReadOnly)
	mainWidget := loader.Load(file, widget)
	file.Close()

	var (
		usernameLabel = widgets.NewQLabelFromPointer(widget.FindChild("UsernameLabel", core.Qt__FindChildrenRecursively).Pointer())
		mxidLabel     = widgets.NewQLabelFromPointer(widget.FindChild("MXIDLabel", core.Qt__FindChildrenRecursively).Pointer())
		avatarLogo    = widgets.NewQLabelFromPointer(widget.FindChild("AvatarLabel", core.Qt__FindChildrenRecursively).Pointer())
	)

	// Set MXID Label
	mxidLabel.SetText(fmt.Sprint(username))

	// Set Dispalyname Label
	displayName, displayNameErr := cli.GetUserDisplayName(username)
	if displayNameErr != nil {
		localLog.Println(displayNameErr)
	}
	usernameLabel.SetText(fmt.Sprint(displayName))

	// Set Avatar
	avatarLogo.SetAlignment(core.Qt__AlignBottom | core.Qt__AlignRight)
	avatarLogo.SetPixmap(cli.GetOwnUserAvatar())

	layout := widgets.NewQVBoxLayout()
	layout.AddWidget(mainWidget, 0, 0)
	widget.SetLayout(layout)

	return widget
}
