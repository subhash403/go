package main

import (
    "fmt"
    //"time"
    //"log"
    "github.com/google/glog"
     apiv1 "k8s.io/api/core/v1" 
     metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
     //"k8s.io/apimachinery/pkg/fields"
     "k8s.io/client-go/kubernetes"
     //"k8s.io/client-go/tools/cache"
     "k8s.io/client-go/tools/clientcmd"
     //"k8s.io/client-go/pkg/watch"
)

//var (
//    kubeconfig = flag.String("kubeconfig", "./config", "/home/easyway/.kube/config")
//)

func main() {
    config, err := clientcmd.BuildConfigFromFlags("","/home/easyway/.kube/config")
    if err != nil {
        glog.Errorln(err)
    } 
    
    
    clientset, err := kubernetes.NewForConfig(config)
    if err != nil {
        glog.Errorln(err)
    }


    podsonly, _ := clientset.CoreV1().Pods("").List(metaV1.ListOptions{})
    for _, pod := range podsonly.Items {
	    fmt.Println(pod.Name, pod.Status)
    }


    fmt.Printf("\n")    

//List the deployments under default namespace 
    deploymentsClient := clientset.AppsV1().Deployments(apiv1.NamespaceDefault)
    fmt.Printf("Listing deployments in namespace %q:\n", apiv1.NamespaceDefault)
	list, err := deploymentsClient.List(metaV1.ListOptions{})
	if err != nil {
		panic(err)
	}
	for _, d := range list.Items {
		fmt.Printf(" * %s (%d replicas)\n", d.Name, *d.Spec.Replicas)
	}
        
	for i, pod := range podsonly.Items {
		fmt.Printf("[%d] %s\n", i, pod.GetName())
	}

        fmt.Printf("\n")	
	
//List the service 
        services, err := clientset.CoreV1().Services("").List(metaV1.ListOptions{})
        if err != nil {
           fmt.Printf("Get service from kubernetes cluster error:%v", err)
           return
        }
	fmt.Printf("Listing services in namespace %q:\n", apiv1.NamespaceDefault)
	for _, service := range services.Items {
                fmt.Printf(" * %s\n", service.Name)
        }


}


