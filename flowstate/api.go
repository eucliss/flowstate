package flowstate

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"flowstate/flowstate/monitor"

	"errors"

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
	r.Get("/get-node", handleGetNode)
	r.Get("/get-nodes", handleGetNodes)
	r.Post("/node-route-status", handleNodeRouteStatus)

	r.Post("/query", handleQuery)
	r.Post("/add-node", handleAddNode)
	r.Post("/update-node", handleUpdateNode)
	r.Post("/add-edge", handleAddEdge)
	r.Post("/connect-edge", handleConnectEdge)
	r.Post("/delete-edges", handleDeleteEdges)
	r.Post("/delete-node", handleDeleteNode)

}

func ParseNodeFromRequest(req *http.Request) (Node, error) {
	if req.Body == nil {
		return Node{}, errors.New("request body is empty")
	}

	// Read the entire body
	body, err := io.ReadAll(req.Body)
	if err != nil {
		return Node{}, errors.New("Error reading request body: " + err.Error())
	}

	// Log the raw body for debugging
	fmt.Printf("Raw request body: %s\n", string(body))

	// Create a new Node instance
	var node Node
	if err := json.Unmarshal(body, &node); err != nil {
		return Node{}, errors.New("Error parsing JSON: " + err.Error())
	}

	// Log the parsed node
	fmt.Printf("Parsed node: %+v\n", node)

	return node, nil
}

func handleGetStatus(w http.ResponseWriter, req *http.Request) {
	// Set JSON content type
	w.Header().Set("Content-Type", "application/json")
	// Return JSON object
	json.NewEncoder(w).Encode(map[string]string{
		"status": "success",
	})
}

func handleGetNodes(w http.ResponseWriter, req *http.Request) {
	nodes := LoadNodes()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(nodes)
}

func handleAddNode(w http.ResponseWriter, req *http.Request) {
	// Check if the request body is nil
	node, err := ParseNodeFromRequest(req)
	if err != nil {
		http.Error(w, "Error parsing node: "+err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Println("node", node)

	if err := AddNode(node); err != nil {
		http.Error(w, "Error adding node: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Node added",
		"status":  "success",
	})
}

func handleGetNode(w http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	var v map[string]string
	err := decoder.Decode(&v)
	if err != nil {
		panic(err)
	}
	id := v["id"]

	node := GetNode(id)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(node)
}

func handleNodeRouteStatus(w http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	var v map[string]interface{}
	err := decoder.Decode(&v)
	if err != nil {
		panic(err)
	}
	id := v["id"].(string)
	route := v["route"].(string)

	response := QueryRouteStatus(id, route)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func handleUpdateNode(w http.ResponseWriter, req *http.Request) {
	node, err := ParseNodeFromRequest(req)
	if err != nil {
		http.Error(w, "Error parsing node: "+err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Println("node", node)

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

	result := Fs.LogStore.Query(query)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}
