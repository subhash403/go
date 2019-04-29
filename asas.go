package main
import (
	"fmt"
	"k8s.io/api/core/v1"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/fake"
	//"testing"
)
func main() {
	client := fake.NewSimpleClientset(&v1.Pod{
		TypeMeta: metaV1.TypeMeta{
			Kind:       "Pod",
			APIVersion: "v1",
		},
		ObjectMeta: metaV1.ObjectMeta{
			Name:      "kubecon-demo-*",
			Namespace: "default",
			Labels: map[string]string{
				"tag": "",
			},
		},
		Status: v1.PodStatus{
			Phase: v1.PodRunning,
			PodIP: "172.17.0.4",
		},
	})

	podList, _ := client.CoreV1().Pods("default").List(metaV1.ListOptions)
            if err != nil {
		fmt.Printf("unexpected error: %v", err)
	}
	if len(pods.Items) != 0 {
		fmt.Errorf("expected no pods, got %#v", pods)
	}

	// get a validation error
	pod := &v1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			GenerateName: "test",
			Namespace:    "default",
		},
		Spec: v1.PodSpec{
			Containers: []v1.Container{
				{
					Name: "test",
				},
			},
		},
	}

}
