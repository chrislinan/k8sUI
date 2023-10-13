package k8s

import (
	"context"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/kubectl/pkg/describe"
)

func ListPod(namespace string) (*v1.PodList, error) {
	// list pods
	if namespace == "" {
		namespace = "default"
	}
	return Clientset.CoreV1().Pods(namespace).List(context.TODO(), metav1.ListOptions{})
}

func DescribePod(namespace, name string) (string, error) {
	d := describe.PodDescriber{
		Interface: Clientset,
	}
	result, err := d.Describe(namespace, name, describe.DescriberSettings{})
	if err != nil {
		return "", err
	}
	return result, nil
}
