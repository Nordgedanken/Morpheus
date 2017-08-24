package main

import (
	"log"
	"sync"

	"github.com/Nordgedanken/Neo/matrix"
	"github.com/therecipe/qt/core"
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
				MainUI := newMainUI(windowWidth, windowHeight, result)
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
func newMainUI(windowWidth, windowHeight int, cli *matrix.Client) *widgets.QWidget {
	widget := widgets.NewQWidget(nil, 0)
	topLayout := widgets.NewQVBoxLayout2(widget)

	// var (
	// 	usernameLabel = widgets.NewQLabelFromPointer(widget.FindChild("UsernameLabel", core.Qt__FindChildrenRecursively).Pointer())
	// 	mxidLabel     = widgets.NewQLabelFromPointer(widget.FindChild("MXIDLabel", core.Qt__FindChildrenRecursively).Pointer())
	// 	avatarLogo    = widgets.NewQLabelFromPointer(widget.FindChild("AvatarLabel", core.Qt__FindChildrenRecursively).Pointer())
	// )
	//
	// // Set MXID Label
	// mxidLabel.SetText(fmt.Sprint(username))
	//
	// // Set Dispalyname Label
	// displayName, displayNameErr := cli.GetUserDisplayName(username)
	// if displayNameErr != nil {
	// 	localLog.Println(displayNameErr)
	// }
	// usernameLabel.SetText(fmt.Sprint(displayName))
	//
	// // Set Avatar
	// avatarLogo.SetAlignment(core.Qt__AlignBottom | core.Qt__AlignRight)
	// avatarLogo.SetPixmap(cli.GetOwnUserAvatar())

	// Wrapper
	wrapperWidget := widgets.NewQGroupBox2("", nil)
	wrapperLayout := widgets.NewQGridLayout2()

	// Roomlist
	roomListView := widgets.NewQWidget(nil, 0)
	roomListView.SetMinimumHeight(windowHeight)
	roomListScroll := widgets.NewQScrollArea(nil)
	roomListViewLayout := widgets.NewQGridLayout(roomListView)
	roomListScroll.SetWidget(roomListView)
	roomListScroll.SetWidgetResizable(true)

	// Fake Room
	RoomWidget := widgets.NewQWidget(nil, 0)
	roomLayout := widgets.NewQVBoxLayout2(RoomWidget)
	room := widgets.NewQLabel2("test", nil, 0)
	roomLayout.AddWidget(room, 0, core.Qt__AlignTop)
	roomListViewLayout.AddWidget(room, 0, 0, core.Qt__AlignTop)

	wrapperLayout.AddWidget(roomListScroll, 0, 0, 0)

	// Message View
	messageView := widgets.NewQWidget(nil, 0)
	messageView.SetMinimumHeight(windowHeight)
	messageScroll := widgets.NewQScrollArea(nil)
	messageViewLayout := widgets.NewQGridLayout(messageView)
	messageScroll.SetWidget(messageView)
	messageScroll.SetWidgetResizable(true)

	// Fake Message
	mesageWidget := widgets.NewQWidget(nil, 0)
	messageLayout := widgets.NewQVBoxLayout2(mesageWidget)
	message := widgets.NewQLabel2("test", nil, 0)
	messageLayout.AddWidget(message, 0, core.Qt__AlignTop)
	messageViewLayout.AddWidget(message, 0, 0, core.Qt__AlignTop)

	wrapperLayout.AddWidget(messageScroll, 0, 1, 0)

	// logoutButton //TODO move it to correct place
	buttonWidget := widgets.NewQWidget(nil, 0)
	buttonLayout := widgets.NewQVBoxLayout()

	logoutButton := widgets.NewQPushButton2("LOGOUT", nil)
	logoutButton.SetMinimumSize2(350, 65)

	buttonLayout.AddStretch(1)
	buttonLayout.AddWidget(logoutButton, 0, 0)
	buttonLayout.AddStretch(1)

	buttonWidget.SetLayout(buttonLayout)
	wrapperLayout.AddWidget(buttonWidget, 1, 0, core.Qt__AlignBottom)

	logoutButton.ConnectClicked(func(checked bool) {
		//TODO register enter and show loader or so
		localLog.Println("Starting Logout Sequenze in background")
		var wg sync.WaitGroup
		results := make(chan bool)

		wg.Add(1)
		go func(cli *matrix.Client, localLog *log.Logger, results chan<- bool) {
			defer wg.Done()
			_, err := cli.Logout()
			if err != nil {
				localLog.Println(err)
				results <- false
			} else {
				cli.ClearCredentials()
				results <- true
			}
		}(cli, localLog, results)

		go func() {
			wg.Wait()      // wait for each execTask to return
			close(results) // then close the results channel
		}()

		//Show MainUI
		for result := range results {
			if result {
				LoginUI := NewLoginUI(windowWidth, windowHeight)
				window.SetCentralWidget(LoginUI)
			}
		}
	})

	wrapperLayout.SetColumnMinimumWidth(0, windowWidth/3)
	wrapperLayout.SetColumnMinimumWidth(1, (windowWidth/3)*2)
	wrapperLayout.SetRowMinimumHeight(0, windowHeight)
	wrapperWidget.SetLayout(wrapperLayout)
	topLayout.AddWidget(wrapperWidget, 1, core.Qt__AlignVCenter)
	widget.SetLayout(topLayout)

	return widget
}
