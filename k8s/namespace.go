package k8s

import (
	"context"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func ListNameSpaces() (*v1.NamespaceList, error) {
	return Clientset.CoreV1().Namespaces().List(context.TODO(), metav1.ListOptions{})
}
