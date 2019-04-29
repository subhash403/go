package main

import (
    "fmt"
    "time"

     "github.com/google/glog"

     "k8s.io/api/core/v1"
     "k8s.io/apimachinery/pkg/fields"
     "k8s.io/client-go/kubernetes"
     "k8s.io/client-go/tools/cache"
     "k8s.io/client-go/tools/clientcmd"
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

    watchlist := cache.NewListWatchFromClient(
        clientset.CoreV1().RESTClient(),
        string(v1.ResourceServices),
        v1.NamespaceAll,
        fields.Everything(),
    )
    _, controller := cache.NewInformer( // also take a look at NewSharedIndexInformer
        watchlist,
        &v1.Service{},
        0, //Duration is int64
        cache.ResourceEventHandlerFuncs{
            AddFunc: func(obj interface{}) {
                fmt.Printf("Pod added: %s \n", obj)
            },
            DeleteFunc: func(obj interface{}) {
                fmt.Printf("Pod  deleted: %s \n", obj)
            },
            UpdateFunc: func(oldObj, newObj interface{}) {
                fmt.Printf("Pod changed \n")
            },
    })

         // I found it in k8s scheduler module. Maybe it's help if you interested in.
     // serviceInformer := cache.NewSharedIndexInformer(watchlist, &v1.Service{},0, cache.Indexers{
     //     cache.NamespaceIndex: cache.MetaNamespaceIndexFunc,
     // })
     // go serviceInformer.Run(stop)
    stop := make(chan struct{})
    defer close(stop)
    go controller.Run(stop)
    for {
        time.Sleep(time.Second)
    }
}
