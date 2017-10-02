package main

import (
	"log"

	//"k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/client-go/rest"
	//"k8s.io/client-go/tools/cache"
	//"k8s.io/client-go/util/workqueue"

	//"github.com/joshvanl/k8s-simple-api/pkg/apis/simple/v1alpha1"
	"github.com/joshvanl/k8s-simple-api/pkg/client"
	//factory "github.com/munnerz/k8s-api-pager-demo/pkg/informers/externalversions"
)

var (
	// apiserverURL is the URL of the API server to connect to
	apiserverURL = "http://127.0.0.1:8001"

	// cl is a Kubernetes API client for our custom resource definition type
	cl client.Interface
)

func main() {

	cl, err := client.NewForConfig(&rest.Config{
		Host: apiserverURL,
	})
	if err != nil {
		log.Fatal("failed to create api client: %v", err)
	}

	log.Printf("Created Kubernets client.")
	log.Print("%s", cl)

}
