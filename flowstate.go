package main

import (
	"net/http"

	"flowstate/flowstate"
)

func main() {

	// logRouter := flowstate.NewLogRouter()
	// logRouter.Source = "https://hooks.torq.io/v1/webhooks/7b9ffae7-9867-44a9-8988-e02a7b282bbc/workflows/eb5539bb-0336-412d-b18a-04ceb86099ab/sync"
	// logRouter.Destination = "http://localhost:5080/api/flowstate/test_torq3/_json"
	// logRouter.Router()

	// return

	// openObserve := monitor.OpenObserve{
	// 	URL: "http://localhost:5080",
	// }

	// result := openObserve.Query(monitor.Query{
	// 	Query:      "SELECT * FROM 'test_torq3'",
	// 	Limit:      2,
	// 	Start:      1738040675782000, // 10000 seconds ago
	// 	End:        1738041575782000,
	// 	SourceType: "flowstate",
	// })

	fs := flowstate.Start()
	// fs.Reset()

	http.ListenAndServe(":3000", fs.Router)
}
