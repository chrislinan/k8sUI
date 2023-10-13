package tutorials

import (
	"K8SGUI/k8s"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
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
			return container.NewHBox(widget.NewIcon(theme.DocumentIcon()), widget.NewLabel("Template Object"))
		},
		func(id widget.ListItemID, item fyne.CanvasObject) {
			hours := time.Now().Sub(podList[id].ObjectMeta.GetCreationTimestamp().Time).Hours()
			age := ""
			if hours < 24 {
				age = fmt.Sprintf("%.0f hours", time.Now().Sub(podList[id].ObjectMeta.GetCreationTimestamp().Time).Hours())
			} else {
				age = fmt.Sprintf("%.0f days", time.Now().Sub(podList[id].ObjectMeta.GetCreationTimestamp().Time).Hours()/24)
			}
			content := podList[id].Name + "\t" + string(podList[id].Status.Phase) + "\t" + age
			item.(*fyne.Container).Objects[1].(*widget.Label).SetText(content)
		},
	)
	list.OnSelected = func(id widget.ListItemID) {
		if podList != nil {
			result, err := k8s.DescribePod(selectedNs, podList[id].Name)
			if err != nil {
			}
			rich := widget.NewRichTextWithText(result)
			rich.Scroll = container.ScrollBoth
			stack.RemoveAll()
			stack.Add(rich)
			icon.SetResource(theme.DocumentIcon())
		}
	}
	list.OnUnselected = func(id widget.ListItemID) {
		icon.SetResource(nil)
	}
	load := widget.NewButton("Load", func() {
		podList = getPodList(selectedNs, podName)
		list.Refresh()
	})
	form := widget.NewForm()
	form.Append("Namespace:", namespace)
	form.Append("Search:", search)

	center := container.NewGridWithColumns(5,
		container.NewCenter(widget.NewLabel("Namespace:")),
		namespace,
		container.NewCenter(widget.NewLabel("Search:")),
		search,
		load,
	)
	topContainer := container.NewGridWithColumns(2, center, layout.NewSpacer())
	split := container.NewBorder(topContainer, nil, nil, nil, container.NewHSplit(list, stack))
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
