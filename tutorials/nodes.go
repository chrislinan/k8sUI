package tutorials

import (
	"K8SGUI/k8s"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	v1 "k8s.io/api/core/v1"
	"time"
)

var nodes *v1.NodeList

func showNodes(_ fyne.Window) fyne.CanvasObject {
	var err error
	stack := container.NewStack()

	nodes, err = k8s.ListNode()
	if err != nil {
		panic(err.Error())
	}

	data := make([]string, len(nodes.Items))
	for i := range data {
		data[i] = nodes.Items[i].Name
	}

	icon := widget.NewIcon(nil)

	list := widget.NewList(
		func() int {
			return len(nodes.Items)
		},

		func() fyne.CanvasObject {
			return container.NewHBox(widget.NewIcon(theme.ComputerIcon()), widget.NewLabel("Template Object"))
		},
		func(id widget.ListItemID, item fyne.CanvasObject) {
			hours := time.Now().Sub(nodes.Items[id].ObjectMeta.GetCreationTimestamp().Time).Hours()
			age := ""
			if hours < 24 {
				age = fmt.Sprintf("%.0f hours", time.Now().Sub(nodes.Items[id].ObjectMeta.GetCreationTimestamp().Time).Hours())
			} else {
				age = fmt.Sprintf("%.0f days", time.Now().Sub(nodes.Items[id].ObjectMeta.GetCreationTimestamp().Time).Hours()/24)
			}
			condition := ""
			for _, cond := range nodes.Items[id].Status.Conditions {
				if cond.Status == "True" {
					condition += string(cond.Type) + " "
				}
			}
			content := nodes.Items[id].Name + "\t" + string(nodes.Items[id].Status.Phase) + "\t" +
				nodes.Items[id].Status.NodeInfo.KubeletVersion + "\t" + condition + "\t" +
				age
			item.(*fyne.Container).Objects[1].(*widget.Label).SetText(content)
		},
	)
	list.OnSelected = func(id widget.ListItemID) {
		if nodes != nil {
			result, err := k8s.DescribeNode(nodes.Items[id].Name)
			if err != nil {
				panic(err.Error())
			}
			rich := widget.NewRichTextWithText(result)
			rich.Scroll = container.ScrollBoth
			//data := k8s.BuildNodeTree(n)
			//tree := widget.NewTreeWithStrings(data)
			//tree.OpenAllBranches()
			stack.RemoveAll()
			//stack.Add(tree)
			stack.Add(rich)
			icon.SetResource(theme.ComputerIcon())
		}
	}
	list.OnUnselected = func(id widget.ListItemID) {
		icon.SetResource(nil)
	}

	return container.NewHSplit(list, stack)
}
