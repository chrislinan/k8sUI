package k8s

import (
	"context"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/kubectl/pkg/describe"
)

func ListNode() (*v1.NodeList, error) {
	return Clientset.CoreV1().Nodes().List(context.TODO(), metav1.ListOptions{})
}

func DescribeNode(name string) (string, error) {
	d := describe.NodeDescriber{
		Interface: Clientset,
	}
	result, err := d.Describe("all", name, describe.DescriberSettings{})
	return result, err
}

func BuildNodeTree(n *v1.Node) map[string][]string {
	annotation := make([]string, 0)
	for key, val := range n.Annotations {
		annotation = append(annotation, key+"="+val)
	}
	labels := make([]string, 0)
	for key, val := range n.Labels {
		labels = append(labels, key+"="+val)
	}
	conditions := make([]string, 0)
	for _, c := range n.Status.Conditions {
		conditions = append(conditions, string(c.Type)+"="+string(c.Status))
	}
	images := make([]string, 0)
	for _, i := range n.Status.Images {
		images = append(images, i.Names...)
	}
	volumeAttached := make([]string, 0)
	for _, v := range n.Status.VolumesAttached {
		volumeAttached = append(volumeAttached, string(v.Name))
	}
	podCidr := make([]string, 0)
	for _, v := range n.Spec.PodCIDRs {
		podCidr = append(podCidr, v)
	}
	taints := make([]string, 0)
	for _, t := range n.Spec.Taints {
		taints = append(taints, string(t.Key)+"="+t.Value)
	}
	data := map[string][]string{
		"":                {"Status", "Spec", "Labels", "Annotations", "Creation"},
		"Name":            {n.Name},
		"Creation":        {n.CreationTimestamp.String()},
		"version":         {n.Status.NodeInfo.KubeletVersion},
		"conditions":      conditions,
		"Annotations":     annotation,
		"Labels":          labels,
		"Status":          {"conditions", "image", "version", "volumesAttached"},
		"volumesAttached": volumeAttached,
		"image":           {n.Status.NodeInfo.OSImage},
		"Spec":            {"podCidr", "ProviderID", "Taints"},
		"podCidr":         podCidr,
		"ProviderID":      {n.Spec.ProviderID},
		"Taints":          taints,
	}
	return data
}
