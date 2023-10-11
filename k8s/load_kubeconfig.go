package k8s

import (
	"fyne.io/fyne/v2"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"log"
)

func LoadKubeconfig(f fyne.URIReadCloser) {
	if f == nil {
		log.Println("Cancelled")
		return
	}
	defer func(f fyne.URIReadCloser) {
		err := f.Close()
		if err != nil {
			log.Println("Error closing file", err)
		}
	}(f)

	log.Println("Loaded KUBECONFIG from...", f.URI().Path())

	config, err := clientcmd.BuildConfigFromFlags("", f.URI().Path())
	if err != nil {
		panic(err.Error())
	}
	Clientset, err = kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
}
