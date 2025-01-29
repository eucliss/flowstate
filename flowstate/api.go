package flowstate

import (
	"encoding/json"
	"net/http"

	"flowstate/flowstate/monitor"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func InitializeAPI() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	// Add CORS middleware
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"}, // Or specify origins like "http://localhost:5173"
		AllowedMethods:   []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Content-Type"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))
	AddRoutes(r)
	return r
}

func AddRoutes(r *chi.Mux) {
	r.Get("/status", handleGetStatus)
	r.Get("/load-flow-state", handleLoadFlowState)
	r.Get("/reset", handleReset)

	r.Post("/query", handleQuery)
	r.Post("/add-node", handleAddNode)
	r.Post("/update-node", handleUpdateNode)
	r.Post("/add-edge", handleAddEdge)
	r.Post("/connect-edge", handleConnectEdge)
	r.Post("/delete-edges", handleDeleteEdges)
	r.Post("/delete-node", handleDeleteNode)
}

func handleGetStatus(w http.ResponseWriter, req *http.Request) {
	// Set JSON content type
	w.Header().Set("Content-Type", "application/json")
	// Return JSON object
	json.NewEncoder(w).Encode(map[string]string{
		"status": "success",
	})
}

func handleAddNode(w http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)

	var node Node
	err := decoder.Decode(&node)
	if err != nil {
		panic(err)
	}

	err = AddNode(node)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Node added",
		"status":  "success",
	})

}

func handleUpdateNode(w http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	var node Node
	err := decoder.Decode(&node)
	if err != nil {
		panic(err)
	}

	err = UpdateNode(node)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Node updated",
		"status":  "success",
	})
}

func handleAddEdge(w http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	var edge Edge
	err := decoder.Decode(&edge)
	if err != nil {
		panic(err)
	}

	err = AddEdge(edge)
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Edge added",
		"status":  "success",
	})
}

func handleConnectEdge(w http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	var edge Edge
	err := decoder.Decode(&edge)
	if err != nil {
		panic(err)
	}

	err = ConnectEdge(edge)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Edge connected",
		"status":  "success",
	})
}

func handleDeleteEdges(w http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	var ids []string
	err := decoder.Decode(&ids)
	if err != nil {
		panic(err)
	}

	err = DeleteEdges(ids)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Edges deleted",
		"status":  "success",
	})
}

func handleDeleteNode(w http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	var id string
	err := decoder.Decode(&id)
	if err != nil {
		panic(err)
	}

	err = DeleteNode(id)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Node deleted",
		"status":  "success",
	})
}

func handleLoadFlowState(w http.ResponseWriter, req *http.Request) {
	Fs.LoadState()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Flow state loaded",
		"status":  "success",
		"nodes":   Fs.Nodes,
		"edges":   Fs.Edges,
	})
}

func handleReset(w http.ResponseWriter, req *http.Request) {
	Fs.Reset()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Flow state reset",
		"status":  "success",
	})
}

func handleQuery(w http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	var query monitor.Query
	err := decoder.Decode(&query)
	if err != nil {
		panic(err)
	}

	openObserve := monitor.OpenObserve{
		URL: "http://localhost:5080",
	}
	result := openObserve.Query(query)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}
