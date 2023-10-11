package k8s

import (
	"context"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func ListPod(namespace string) (*v1.PodList, error) {
	// list pods
	if namespace == "" {
		namespace = "default"
	}
	return Clientset.CoreV1().Pods(namespace).List(context.TODO(), metav1.ListOptions{})
}
