package tutorials

import (
	"K8SGUI/k8s"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	v1 "k8s.io/api/core/v1"
	"strings"
	"time"
)

func showPods(_ fyne.Window) fyne.CanvasObject {
	var err error
	var selectedNs, podName string
	podList := make([]v1.Pod, 0)
	stack := container.NewStack()

	ns, err := k8s.ListNameSpaces()
	if err != nil {
		panic(err.Error())
	}
	nsList := make([]string, 0)
	for _, n := range ns.Items {
		nsList = append(nsList, n.Name)
	}
	namespace := widget.NewSelect(nsList, func(s string) {
		selectedNs = s
	})
	namespace.SetSelected("default")
	search := widget.NewEntry()
	search.SetPlaceHolder("input pod name")
	search.TextStyle = fyne.TextStyle{Monospace: true}
	search.OnChanged = func(s string) {
		podName = search.Text
	}
	podList = getPodList(selectedNs, podName)
	icon := widget.NewIcon(nil)

	list := widget.NewList(
		func() int {
			return len(podList)
		},
		func() fyne.CanvasObject {
			return container.NewHBox(widget.NewIcon(theme.ComputerIcon()), widget.NewLabel("Template Object"))
		},
		func(id widget.ListItemID, item fyne.CanvasObject) {
			age := time.Now().Sub(podList[id].ObjectMeta.GetCreationTimestamp().Time).Hours() / 24
			content := podList[id].Name + "\t" + string(podList[id].Status.Phase) + "\t" + fmt.Sprintf("%.0f days", age)
			item.(*fyne.Container).Objects[1].(*widget.Label).SetText(content)
		},
	)
	list.OnSelected = func(id widget.ListItemID) {
		if podList != nil {
			//n, err := k8s.DescribeNode(pods.Items[id].Name)
			//if err != nil {
			//	panic(err.Error())
			//}
			//data := k8s.BuildNodeTree(n)
			//tree := widget.NewTreeWithStrings(data)
			//tree.OpenAllBranches()
			//stack.Objects = nil
			//stack.Add(tree)
			icon.SetResource(theme.ComputerIcon())
		}
	}
	list.OnUnselected = func(id widget.ListItemID) {
		icon.SetResource(nil)
	}
	load := widget.NewButton("Load", func() {
		podList = getPodList(selectedNs, podName)
		list.Refresh()
	})
	split := container.NewVSplit(container.NewHBox(widget.NewLabel("Namespace:"), namespace, widget.NewLabel("Search:"), search, load), container.NewHSplit(list, stack))
	split.Offset = 0.03
	return split
}

func getPodList(selectedNs, podName string) []v1.Pod {
	var podList []v1.Pod
	pods, err := k8s.ListPod(selectedNs)
	if err != nil {
		panic(err.Error())
	}
	if podName != "" && pods != nil {
		for _, p := range pods.Items {
			if strings.Contains(p.Name, podName) {
				podList = append(podList, p)
			}
		}
	} else {
		podList = pods.Items
	}
	return podList
}
