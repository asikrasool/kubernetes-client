package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	rest "k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

var config *rest.Config
var clientSet *kubernetes.Clientset

// export USE_KUBECONFIG=true to use kubeconfig file from local. it is not needed when deployed to cluster as pod
func main() {
	useKubeConfig := os.Getenv("USE_KUBECONFIG")
	kubeConfigFilePath := os.Getenv("KUBECONFIG")

	if len(useKubeConfig) == 0 {
		// default to service account in cluster token
		c, err := rest.InClusterConfig()
		if err != nil {
			panic(err.Error())
		}
		config = c
	} else {
		//load from a kube config
		var kubeconfig string

		if kubeConfigFilePath == "" {
			if home := homedir.HomeDir(); home != "" {
				kubeconfig = filepath.Join(home, ".kube", "config")
			}
		} else {
			kubeconfig = kubeConfigFilePath
		}

		fmt.Println("kubeconfig: " + kubeconfig)
		// create the config object from kubeconfig
		c, err := clientcmd.BuildConfigFromFlags("", kubeconfig)

		if err != nil {
			panic(err.Error())
		}
		config = c

	}

	// kubeconfig := flag.String("kubeconfig", "/Users/asikrasool/.kube/config", "location to your kubeconfig file")
	// config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)

	// if err != nil {
	// 	panic(err.Error())
	// }
	// fmt.Println(config)

	// create clientset (set of muliple clients) for each Group (e.g. Core),
	// the Version (V1) of Group and Kind (e.g. Pods) so GVK.
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	// executes LIST of pods from all namespaces
	pods, err := clientset.CoreV1().Pods("").List(context.Background(), metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("List of pods")
	for _, pod := range pods.Items {
		fmt.Println(pod.Name, "\n")
	}
	// executes LIST of Deployments from all namespaces
	deployments, err := clientset.AppsV1().Deployments("").List(context.Background(), metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("List of Deployments")
	for _, d := range deployments.Items {
		fmt.Println(d.Name, "\n")
	}

	// TO check the full deployment details using Marshal
	// x := map[string]interface{}{"a": 1, "b": 2}
	x := deployments
	b, err := json.MarshalIndent(x, "", "  ")
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Print(string(b))
}
