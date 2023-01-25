package main

import (
	"fmt"
	"github.com/energye/systray"
	"github.com/energye/systray/icon"
	"time"
)

var start func()
var end func()

func main() {
	go func() {
		onExit := func() {
			now := time.Now()
			fmt.Println("Exit at", now.String())
		}

		start, end = systray.RunWithExternalLoop(onReady, onExit) //windows/linux/macos
		start()

	}()
	select {}
}

func MainRun() {
	onExit := func() {
		now := time.Now()
		fmt.Println("Exit at", now.String())
	}

	systray.Run(onReady, onExit)
}

func addQuitItem() {
	mQuit := systray.AddMenuItem("Quit", "Quit the whole app")
	mQuit.Enable()
	mQuit.Click(func() {
		fmt.Println("Requesting quit")
		//systray.Quit()
		//systray.Quit()// macos error
		//end() // macos error
		fmt.Println("Finished quitting")
	})
}

func onReady() {
	fmt.Println("systray.onReady")
	systray.SetTemplateIcon(icon.Data, icon.Data)
	systray.SetTitle("Energy Sys Tray")
	systray.SetTooltip("Energy tooltip")
	systray.SetOnClick(func() {
		fmt.Println("SetOnClick")
	})
	systray.SetOnDClick(func() {
		fmt.Println("SetOnDClick")
	})
	systray.SetOnRClick(func(menu systray.IMenu) {
		menu.ShowMenu()
		fmt.Println("SetOnRClick")
	})
	systray.CreateMenu()
	addQuitItem()
	systray.SetTemplateIcon(icon.Data, icon.Data)
	mChange := systray.AddMenuItem("Change Me", "Change Me")
	mChecked := systray.AddMenuItemCheckbox("Checked", "Check Me", true)
	mEnabled := systray.AddMenuItem("Enabled", "Enabled")
	// Sets the icon of a menu item. Only available on Mac.
	mEnabled.SetTemplateIcon(icon.Data, icon.Data)

	systray.AddMenuItem("Ignored", "Ignored")

	subMenuTop := systray.AddMenuItem("SubMenuTop", "SubMenu Test (top)")
	subMenuMiddle := subMenuTop.AddSubMenuItem("SubMenuMiddle", "SubMenu Test (middle)")
	subMenuBottom := subMenuMiddle.AddSubMenuItemCheckbox("SubMenuBottom - Toggle Panic!", "SubMenu Test (bottom) - Hide/Show Panic!", false)
	subMenuBottom2 := subMenuMiddle.AddSubMenuItem("SubMenuBottom - Panic!", "SubMenu Test (bottom)")
	subMenuBottom2.SetIcon(icon.Data)
	systray.AddSeparator()
	mToggle := systray.AddMenuItem("Toggle", "Toggle some menu items")
	shown := true
	toggle := func() {
		if shown {
			subMenuBottom.Check()
			subMenuBottom2.Hide()
			mEnabled.Hide()
			shown = false
			mEnabled.Disable()
		} else {
			subMenuBottom.Uncheck()
			subMenuBottom2.Show()
			mEnabled.Show()
			mEnabled.Enable()
			shown = true
		}
	}
	mReset := systray.AddMenuItem("Reset", "Reset all items")

	mChange.Click(func() {
		mChange.SetTitle("I've Changed")
	})
	mChecked.Click(func() {
		if mChecked.Checked() {
			mChecked.Uncheck()
			mChecked.SetTitle("Unchecked")
		} else {
			mChecked.Check()
			mChecked.SetTitle("Checked")
		}
	})
	mEnabled.Click(func() {
		mEnabled.SetTitle("Disabled")
		fmt.Println("mEnabled.Disabled()", mEnabled.Disabled())
		mEnabled.Disable()
	})
	subMenuBottom2.Click(func() {
		panic("panic button pressed")
	})
	subMenuBottom.Click(func() {
		toggle()
	})
	mReset.Click(func() {
		systray.ResetMenu()
		addQuitItem()
	})
	mToggle.Click(func() {
		toggle()
	})
}
