package flowstate

import (
	"github.com/go-chi/chi/v5"
	redis "github.com/redis/go-redis/v9"
)

var Fs *Flowstate

type Flowstate struct {
	Nodes    []Node
	Edges    []Edge
	DbClient *redis.Client
	Router   *chi.Mux
}

func Start() *Flowstate {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // No password set
		DB:       0,  // Use default DB
		Protocol: 2,  // Connection protocol
	})
	router := InitializeAPI()
	Fs = &Flowstate{
		DbClient: client,
		Router:   router,
	}
	Fs.LoadState()
	return Fs
}

func (f *Flowstate) LoadState() {
	f.Nodes = LoadNodes()
	f.Edges = LoadEdges()
}

func (f *Flowstate) Reset() {
	f.DeleteAll()
	f.InitializeDummyGraph()
}

func (f *Flowstate) InitializeDummyGraph() {
	AddStartingNodes()
	AddStartingEdges()
}

func (f *Flowstate) DeleteAll() {
	DeleteAllNodes()
	DeleteAllEdges()
}
