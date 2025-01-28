package main

import (
	"net/http"

	"flowstate/flowstate"
)

func main() {

	// logRouter := log_router.NewLogRouter()
	// logRouter.Source = "https://hooks.torq.io/v1/webhooks/7b9ffae7-9867-44a9-8988-e02a7b282bbc/workflows/eb5539bb-0336-412d-b18a-04ceb86099ab/sync"
	// logRouter.Destination = "http://localhost:5080/api/flowstate/test_torq3/_json"
	// logRouter.Router()

	// return
	fs := flowstate.Start()

	http.ListenAndServe(":3000", fs.Router)
}
