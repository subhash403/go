package main

import (
    "fmt"
    //"time"

     "github.com/google/glog"

     //"k8s.io/api/core/v1"
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

     pods, err := clientset.CoreV1().Pods("default").List(metaV1.ListOptions{})
     if err != nil {
	// handle error
     }
     for _, pod := range pods.Items {
	fmt.Println(pod.Name,pod.Status,pod.Status.PodIP)
     }

    // I found it in k8s scheduler module. Maybe it's help if you interested in.
    //serviceInformer := cache.NewSharedIndexInformer(watchlist, &v1.Service{},0, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc,
    //})
    //go serviceInformer.Run(stop)
    //stop := make(chan struct{})
    //defer close(stop)
    //go watch.Run(stop)
    //for {
    //    time.Sleep(time.Second)
    //}
}
