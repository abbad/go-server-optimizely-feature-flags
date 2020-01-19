package main

import (
	"os"
	"log"
	"net/http"
	"pkg/featureflags"

)

func main() {
	client := &featureflags.OptiService{Client: featureflags.GetClient(os.Getenv("OPTMIZILY_SDK_KEY"))}
	http.HandleFunc("/api/feature-flags", client.GetEnabledFeatures)
	log.Println("running server on 0.0.0.0:10000 \n")
	log.Fatal(http.ListenAndServe("0.0.0.0:10000", nil))
}
