package flowstate

import (
	"context"
	"fmt"
)

type Edge struct {
	Id           string `redis:"id"`
	Source       string `redis:"source"`
	Target       string `redis:"target"`
	Animated     bool   `redis:"animated"`
	SourceHandle string `redis:"sourceHandle"`
	TargetHandle string `redis:"targetHandle"`
	Type         string `redis:"type"`
}

func AddEdge(edge Edge) error {
	ctx := context.Background()
	fmt.Println("Adding edge: ", edge)
	err := Fs.DbClient.HSet(ctx, "edge:"+edge.Id, edge).Err()
	if err != nil {
		return err
	}
	return nil
}

func ConnectEdge(edge Edge) error {
	ctx := context.Background()
	fmt.Println("Connecting edge: ", edge)
	err := Fs.DbClient.HSet(ctx, "edge:"+edge.Id, edge).Err()
	if err != nil {
		return err
	}
	return nil
}

func DeleteEdges(ids []string) error {
	ctx := context.Background()
	for _, id := range ids {
		Fs.DbClient.Del(ctx, "edge:"+id).Result()
	}
	return nil
}

func LoadEdges() []Edge {
	ctx := context.Background()
	iter := Fs.DbClient.Scan(ctx, 0, "edge:*", 0).Iterator()
	edges := []Edge{}

	for iter.Next(ctx) {
		key := iter.Val()
		var edge Edge
		// Get all fields for this node

		if err := Fs.DbClient.HGetAll(ctx, key).Scan(&edge); err != nil {
			panic(err)
		}
		edges = append(edges, edge)
	}

	if err := iter.Err(); err != nil {
		panic(err)
	}

	fmt.Println("Loaded edges: ", edges)
	return edges
}

func AddStartingEdges() {
	edges := []Edge{
		{
			Id:           "1->2",
			Source:       "1",
			Target:       "2",
			SourceHandle: "bottom",
			TargetHandle: "top",
			Animated:     true,
		},
		{
			Id:           "2->3",
			Source:       "2",
			Target:       "3",
			SourceHandle: "top",
			TargetHandle: "bottom",
			Animated:     true,
		},
	}

	ctx := context.Background()
	for _, edge := range edges {
		err := Fs.DbClient.HSet(ctx, "edge:"+edge.Id, edge).Err()
		if err != nil {
			panic(err)
		}
	}
}

func DeleteAllEdges() {
	ctx := context.Background()
	iter := Fs.DbClient.Scan(ctx, 0, "edge:*", 0).Iterator()

	for iter.Next(ctx) {
		key := iter.Val()
		Fs.DbClient.Del(ctx, key).Result()
	}

	if err := iter.Err(); err != nil {
		panic(err)
	}
}
