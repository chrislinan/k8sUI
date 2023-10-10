package tutorials

import (
	"K8SGUI/k8s"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"time"
)

func showNodes(_ fyne.Window) fyne.CanvasObject {
	nodes, err := k8s.ListNode()
	if err != nil {
		panic(err.Error())
	}

	data := make([]string, len(nodes.Items))
	for i := range data {
		data[i] = nodes.Items[i].Name
	}

	icon := widget.NewIcon(nil)
	label := widget.NewLabel("Select An Item From The List")
	hbox := container.NewHBox(label)

	list := widget.NewList(
		func() int {
			return len(nodes.Items)
		},

		func() fyne.CanvasObject {
			return container.NewHBox(widget.NewIcon(theme.ComputerIcon()), widget.NewLabel("Template Object"))
		},
		func(id widget.ListItemID, item fyne.CanvasObject) {
			age := time.Now().Sub(nodes.Items[id].ObjectMeta.GetCreationTimestamp().Time).Hours() / 24
			condition := ""
			for _, cond := range nodes.Items[id].Status.Conditions {
				if cond.Status == "True" {
					condition += string(cond.Type) + " "
				}
			}
			content := nodes.Items[id].Name + "\t" + string(nodes.Items[id].Status.Phase) + "\t" +
				nodes.Items[id].Status.NodeInfo.KubeletVersion + "\t" + condition + "\t" +
				fmt.Sprintf("%.0f days", age)
			item.(*fyne.Container).Objects[1].(*widget.Label).SetText(content)
		},
	)
	list.OnSelected = func(id widget.ListItemID) {
		label.SetText(data[id])
		icon.SetResource(theme.ComputerIcon())
	}
	list.OnUnselected = func(id widget.ListItemID) {
		label.SetText("Select An Item From The List")
		icon.SetResource(nil)
	}
	return container.NewHSplit(list, container.NewCenter(hbox))
}
