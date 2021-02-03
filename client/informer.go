package client

import (
	"time"

	core_v1 "k8s.io/api/core/v1"
	"k8s.io/client-go/informers"
	informer_corev1 "k8s.io/client-go/informers/core/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
)

var (
	sharedInformerFactory informers.SharedInformerFactory
	podInformer           cache.SharedIndexInformer
	nodeInformer          cache.SharedIndexInformer
)

// InitInformer init informer
func InitInformer(clientset *kubernetes.Clientset) {
	sharedInformerFactory = informers.NewSharedInformerFactory(clientset, 0)
	stopCh := make(chan struct{})

	podInformer = sharedInformerFactory.InformerFor(&core_v1.Pod{}, NewPodInformer)
	nodeInformer = sharedInformerFactory.InformerFor(&core_v1.Node{}, NewNodeInformer)
	sharedInformerFactory.Start(stopCh)
	sharedInformerFactory.WaitForCacheSync(stopCh)
}

// NewPodInformer new pod informer
func NewPodInformer(client kubernetes.Interface, syncTime time.Duration) cache.SharedIndexInformer {
	return informer_corev1.NewPodInformer(client, "default",syncTime, cache.Indexers{})
}

// NewNodeInformer new node informer
func NewNodeInformer(client kubernetes.Interface, syncTime time.Duration) cache.SharedIndexInformer {
	return informer_corev1.NewNodeInformer(client, syncTime, cache.Indexers{})
}
