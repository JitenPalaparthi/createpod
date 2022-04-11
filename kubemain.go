package main

import (
	"context"
	"flag"
	"fmt"
	"io/ioutil"
	"path/filepath"

	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/yaml"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

func main() {
	var kubeconfig *string
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()

	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err)
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}

	p1 := &apiv1.Pod{}

	buf, err := ioutil.ReadFile("nginx.yaml")
	if err != nil {
		panic("failed to read the file")
	}
	err = yaml.Unmarshal(buf, p1)
	if err != nil {
		panic(err)
	}
	_, err = clientset.CoreV1().Pods("default").Create(context.TODO(), p1, metav1.CreateOptions{})
	fmt.Println(err)
}
