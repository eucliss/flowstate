package flowstate

import (
	"context"
	"encoding/json"
	"flowstate/flowstate/monitor"
	"fmt"
	"strconv"
)

type ComparisonType struct {
	LeftValue  string `json:"leftValue" redis:"leftValue"`
	RightValue string `json:"rightValue" redis:"rightValue"`
	Operator   string `json:"operator" redis:"operator"`
}

type Node struct {
	Id           string          `json:"id" redis:"id"`
	Type         string          `json:"type" redis:"type"`
	X            int             `json:"x" redis:"x"`
	Y            int             `json:"y" redis:"y"`
	Label        string          `json:"label" redis:"label"`
	SQL          string          `json:"sql" redis:"sql"`
	SuccessRoute *ComparisonType `json:"successRoute" redis:"successRoute"`
	FailureRoute *ComparisonType `json:"failureRoute" redis:"failureRoute"`
}

func (c *ComparisonType) MarshalBinary() ([]byte, error) {
	return json.Marshal(c)
}

func (c *ComparisonType) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, c)
}

func convertRedisDataToNode(data map[string]string) Node {
	// Create the node with basic fields
	node := Node{
		Id:    data["id"],
		Type:  data["type"],
		Label: data["label"],
		SQL:   data["sql"],
	}

	// Parse position
	if x, err := strconv.Atoi(data["x"]); err == nil {
		node.X = x
	}
	if y, err := strconv.Atoi(data["y"]); err == nil {
		node.Y = y
	}

	// Parse routes if they exist
	if data["successRoute"] != "" {
		var successRoute ComparisonType
		if err := json.Unmarshal([]byte(data["successRoute"]), &successRoute); err == nil {
			node.SuccessRoute = &successRoute
		}
	}

	if data["failureRoute"] != "" {
		var failureRoute ComparisonType
		if err := json.Unmarshal([]byte(data["failureRoute"]), &failureRoute); err == nil {
			node.FailureRoute = &failureRoute
		}
	}
	return node
}

func GetNode(id string) Node {
	ctx := context.Background()
	result := Fs.DbClient.HGetAll(ctx, "node:"+id)
	node := convertRedisDataToNode(result.Val())
	return node
}

func checkRouteEquality(l string, r string, operator string) bool {
	if operator == "Equals" {
		return l == r
	}
	return l != r
}

func getRouteFromNode(node Node, route string) *ComparisonType {
	if route == "success" {
		return node.SuccessRoute
	}
	return node.FailureRoute
}

func QueryRouteStatus(node_id string, route string) bool {
	node := GetNode(node_id)
	query := node.SQL
	storeQuery := monitor.Query{
		Query:      query,
		Start:      1738040675782000,
		End:        1738041575782000,
		Limit:      1,
		SourceType: "flowstate",
	}
	fmt.Printf("Querying route status for node %s with query %s\n", node_id, query)
	results := Fs.LogStore.Query(storeQuery)
	lastResult := results[0].(map[string]interface{})

	r := getRouteFromNode(node, route)
	leftVal, leftOk := lastResult[r.LeftValue].(string)
	fmt.Println("leftVal ========================", leftVal)
	fmt.Println("rightVal ========================", r.RightValue)
	fmt.Println("operator ========================", r.Operator)

	if !leftOk {
		fmt.Printf("Warning: Could not find values for comparison: left=%v, right=%v\n", r.LeftValue, r.RightValue)
		return false
	}
	return checkRouteEquality(leftVal, r.RightValue, r.Operator)
}

func LoadNodes() []Node {
	ctx := context.Background()
	iter := Fs.DbClient.Scan(ctx, 0, "node:*", 0).Iterator()
	nodes := []Node{}

	for iter.Next(ctx) {
		key := iter.Val()
		var node Node

		// Get all fields for this node
		result := Fs.DbClient.HGetAll(ctx, key)
		data := result.Val()

		// Create the node with basic fields
		node = Node{
			Id:    data["id"],
			Type:  data["type"],
			Label: data["label"],
			SQL:   data["sql"],
		}

		// Parse position
		if x, err := strconv.Atoi(data["x"]); err == nil {
			node.X = x
		}
		if y, err := strconv.Atoi(data["y"]); err == nil {
			node.Y = y
		}

		// Parse routes if they exist
		if data["successRoute"] != "" {
			var successRoute ComparisonType
			if err := json.Unmarshal([]byte(data["successRoute"]), &successRoute); err == nil {
				node.SuccessRoute = &successRoute
			}
		}

		if data["failureRoute"] != "" {
			var failureRoute ComparisonType
			if err := json.Unmarshal([]byte(data["failureRoute"]), &failureRoute); err == nil {
				node.FailureRoute = &failureRoute
			}
		}

		nodes = append(nodes, node)
	}

	if err := iter.Err(); err != nil {
		panic(err)
	}

	fmt.Printf("Loaded nodes: %+v\n", nodes)
	return nodes
}

func AddNode(node Node) error {
	ctx := context.Background()
	err := Fs.DbClient.HSet(ctx, "node:"+node.Id, node).Err()
	if err != nil {
		fmt.Println("error adding node", err)
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
			SQL:   "SELECT * FROM 'test_torq3'",
		},
		{
			Id:    "2",
			Type:  "default",
			X:     100,
			Y:     100,
			Label: "Node 2",
			SQL:   "SELECT * FROM 'test_torq3'",
		},
		{
			Id:    "3",
			Type:  "output",
			X:     400,
			Y:     200,
			Label: "Node 3",
			SQL:   "SELECT * FROM 'test_torq3'",
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
