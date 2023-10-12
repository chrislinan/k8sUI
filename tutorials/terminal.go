package tutorials

import "fyne.io/fyne/v2"
import "github.com/fyne-io/terminal"

func terminalScreen(win fyne.Window) fyne.CanvasObject {
	t := terminal.New()
	go func() {
		_ = t.RunLocalShell()
	}()
	return t
}
