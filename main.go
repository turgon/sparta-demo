// File: main.go
package main

import (
	"fmt"
	"net/http"

	sparta "github.com/mweagle/Sparta"
)

// Standard AWS λ function
func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

func main() {
	var lambdaFunctions []*sparta.LambdaAWSInfo
	helloWorldFn := sparta.HandleAWSLambda("Hello World",
		http.HandlerFunc(helloWorld),
		sparta.IAMRoleDefinition{})
	lambdaFunctions = append(lambdaFunctions, helloWorldFn)
	sparta.Main("MyHelloWorldStack",
		"Simple Sparta application that demonstrates core functionality",
		lambdaFunctions,
		nil,
		nil)
}
