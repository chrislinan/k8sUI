package tutorials

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func showWorkloads(_ fyne.Window) fyne.CanvasObject {
	return container.NewVBox(
		widget.NewLabelWithStyle("func Length() int", fyne.TextAlignLeading, fyne.TextStyle{Monospace: true}),
		widget.NewLabelWithStyle("func CreateItem() fyne.CanvasObject", fyne.TextAlignLeading, fyne.TextStyle{Monospace: true}),
		widget.NewLabelWithStyle("func UpdateItem(ListItemID, fyne.CanvasObject)", fyne.TextAlignLeading, fyne.TextStyle{Monospace: true}),
		widget.NewLabelWithStyle("func OnSelected(ListItemID)", fyne.TextAlignLeading, fyne.TextStyle{Monospace: true}),
		widget.NewLabelWithStyle("func OnUnselected(ListItemID)", fyne.TextAlignLeading, fyne.TextStyle{Monospace: true}))
}
