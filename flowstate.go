package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"context"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	redis "github.com/redis/go-redis/v9"
)

type FlowState struct {
	Nodes []Node
	Edges []Edge
}

type Node struct {
	Id    string `redis:"id"`
	Type  string `redis:"type"`
	X     int    `redis:"x"`
	Y     int    `redis:"y"`
	Label string `redis:"label"`
}

type Edge struct {
	Id       string `redis:"id"`
	Source   string `redis:"source"`
	Target   string `redis:"target"`
	Animated bool   `redis:"animated"`
}

func addStartingEdges(client *redis.Client) {
	edges := []Edge{
		{
			Id:     "1->2",
			Source: "1",
			Target: "2",
		},
		{
			Id:       "2->3",
			Source:   "2",
			Target:   "3",
			Animated: true,
		},
	}

	ctx := context.Background()
	for _, edge := range edges {
		err := client.HSet(ctx, "edge:"+edge.Id, edge).Err()
		if err != nil {
			panic(err)
		}
	}
}

func addStartingNodes(client *redis.Client) {
	nodes := []Node{
		{
			Id:    "1",
			Type:  "input",
			X:     250,
			Y:     5,
			Label: "Node 1",
		},
		{
			Id:    "2",
			Type:  "default",
			X:     100,
			Y:     100,
			Label: "Node 2",
		},
		{
			Id:    "3",
			Type:  "output",
			X:     400,
			Y:     200,
			Label: "Node 3",
		},
	}

	ctx := context.Background()
	for _, node := range nodes {
		err := client.HSet(ctx, "node:"+node.Id, node).Err()
		if err != nil {
			panic(err)
		}
	}

}

func handleAddNode(client *redis.Client, req *http.Request) {
	ctx := context.Background()
	fmt.Println("Got add node request")
	fmt.Println(req)
	fmt.Println("--------------------------------")
	decoder := json.NewDecoder(req.Body)
	var node Node
	err := decoder.Decode(&node)
	if err != nil {
		panic(err)
	}

	fmt.Println("node: ", node)
	fmt.Println("--------------------------------")
	fmt.Println("client: ", client)
	fmt.Println("--------------------------------")
	node.Id = "node:" + node.Id
	err = client.HSet(ctx, node.Id, node).Err()
	if err != nil {
		panic(err)
	}
}

func handleAddEdge(client *redis.Client, req *http.Request) {
	ctx := context.Background()
	fmt.Println("Got add edge request")
	fmt.Println(req)
	fmt.Println("--------------------------------")
	decoder := json.NewDecoder(req.Body)
	var edge Edge
	err := decoder.Decode(&edge)
	if err != nil {
		panic(err)
	}

	fmt.Println("edge: ", edge)
	fmt.Println("--------------------------------")
	fmt.Println("client: ", client)
	fmt.Println("--------------------------------")
	edge.Id = "edge:" + edge.Id
	err = client.HSet(ctx, edge.Id, edge).Err()
	if err != nil {
		panic(err)
	}
}

func loadNodes(client *redis.Client) []Node {
	ctx := context.Background()
	iter := client.Scan(ctx, 0, "node:*", 0).Iterator()
	nodes := []Node{}

	for iter.Next(ctx) {
		key := iter.Val()
		var node Node
		// Get all fields for this node

		if err := client.HGetAll(ctx, key).Scan(&node); err != nil {
			panic(err)
		}
		fmt.Printf("Key: %s, Value: %v\n", key, node)
		nodes = append(nodes, node)
	}

	if err := iter.Err(); err != nil {
		panic(err)
	}

	return nodes
}

func loadEdges(client *redis.Client) []Edge {
	ctx := context.Background()
	iter := client.Scan(ctx, 0, "edge:*", 0).Iterator()
	fmt.Println("iter: ", iter)
	edges := []Edge{}

	for iter.Next(ctx) {
		key := iter.Val()
		var edge Edge
		// Get all fields for this node

		if err := client.HGetAll(ctx, key).Scan(&edge); err != nil {
			panic(err)
		}
		fmt.Printf("Key: %s, Value: %v\n", key, edge)
		edges = append(edges, edge)
	}

	if err := iter.Err(); err != nil {
		panic(err)
	}

	return edges
}

func loadFlowState(client *redis.Client) FlowState {
	nodes := loadNodes(client)
	edges := loadEdges(client)

	return FlowState{
		Nodes: nodes,
		Edges: edges,
	}
}

func removeNode(client *redis.Client, id string) {
	ctx := context.Background()
	res, err := client.Del(ctx, "node:"+id).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("res: ", res)

}

func main() {

	// Start Redis Client

	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // No password set
		DB:       0,  // Use default DB
		Protocol: 2,  // Connection protocol
	})
	flowState := loadFlowState(client)
	fmt.Println("flowState: ", flowState)

	// addStartingNodes(client)

	// Start Router for API
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
	status := "success"
	r.Get("/status", func(w http.ResponseWriter, r *http.Request) {
		if status == "success" {
			status = "failure"
		} else {
			status = "success"
		}
		// Set JSON content type
		w.Header().Set("Content-Type", "application/json")
		// Return JSON object
		json.NewEncoder(w).Encode(map[string]string{
			"status": status,
		})
	})
	r.Post("/add-node", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Got add node request")
		fmt.Println(r)
		fmt.Println("--------------------------------")
		// Add a node to the flow state
		// Get the node from the request body
		// Add the node to the flow state
		// Return the flow state
		handleAddNode(client, r)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"message": "Node added",
			"status":  "success",
		})
	})

	r.Get("/load-flow-state", func(w http.ResponseWriter, r *http.Request) {
		flowState := loadFlowState(client)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"message": "Flow state loaded",
			"status":  "success",
			"nodes":   flowState.Nodes,
			"edges":   flowState.Edges,
		})
	})
	http.ListenAndServe(":3000", r)
}
