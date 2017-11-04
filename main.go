// File: main.go
package main

import (
	"fmt"
	"net/http"

	"encoding/json"

	"github.com/Sirupsen/logrus"

	sparta "github.com/mweagle/Sparta"
)

// Standard AWS λ function
func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

func echoS3Event(w http.ResponseWriter, r *http.Request) {
	logger, _ := r.Context().Value(sparta.ContextKeyLogger).(*logrus.Logger)
	lambdaContext, _ := r.Context().Value(sparta.ContextKeyLambdaContext).(*sparta.LambdaContext)

	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()
	var jsonMessage json.RawMessage
	err := decoder.Decode(&jsonMessage)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	logger.WithFields(logrus.Fields{
		"RequestID": lambdaContext.AWSRequestID,
		"Event":     string(jsonMessage),
	}).Info("Request received")

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonMessage)
}

func main() {
	var lambdaFunctions []*sparta.LambdaAWSInfo

	lambdaFn := sparta.HandleAWSLambda(sparta.LambdaName(echoS3Event),
		http.HandlerFunc(echoS3Event),
		sparta.IAMRoleDefinition{})

	lambdaFunctions = append(lambdaFunctions, lambdaFn)

	stage := sparta.NewStage("prod")
	apiGateway := sparta.NewAPIGateway("MySpartaAPI", stage)

	apiGatewayResource, _ := apiGateway.NewResource("/hello/world/test", lambdaFn)
	apiGatewayResource.NewMethod("GET", http.StatusOK)

	sparta.Main("SpartaApplication",
		"Simple Sparta application",
		lambdaFunctions,
		apiGateway,
		nil)

}
