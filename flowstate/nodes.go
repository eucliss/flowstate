package flowstate

import (
	"context"
)

type Node struct {
	Id    string `redis:"id"`
	Type  string `redis:"type"`
	X     int    `redis:"x"`
	Y     int    `redis:"y"`
	Label string `redis:"label"`
}

func LoadNodes() []Node {
	ctx := context.Background()
	iter := Fs.DbClient.Scan(ctx, 0, "node:*", 0).Iterator()
	nodes := []Node{}

	for iter.Next(ctx) {
		key := iter.Val()
		var node Node
		// Get all fields for this node

		if err := Fs.DbClient.HGetAll(ctx, key).Scan(&node); err != nil {
			panic(err)
		}
		nodes = append(nodes, node)
	}

	if err := iter.Err(); err != nil {
		panic(err)
	}

	return nodes
}

func AddNode(node Node) error {
	ctx := context.Background()
	err := Fs.DbClient.HSet(ctx, "node:"+node.Id, node).Err()
	if err != nil {
		return err
	}
	return nil
}

func UpdateNode(node Node) error {
	ctx := context.Background()
	err := Fs.DbClient.HSet(ctx, "node:"+node.Id, node).Err()
	if err != nil {
		return err
	}
	return nil
}

func DeleteNode(id string) error {
	ctx := context.Background()
	err := Fs.DbClient.Del(ctx, "node:"+id).Err()
	if err != nil {
		return err
	}
	return nil
}

func AddStartingNodes() {
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
		err := Fs.DbClient.HSet(ctx, "node:"+node.Id, node).Err()
		if err != nil {
			panic(err)
		}
	}

}

func DeleteAllNodes() {
	ctx := context.Background()
	iter := Fs.DbClient.Scan(ctx, 0, "node:*", 0).Iterator()

	for iter.Next(ctx) {
		key := iter.Val()
		Fs.DbClient.Del(ctx, key).Result()
	}

	if err := iter.Err(); err != nil {
		panic(err)
	}
}
