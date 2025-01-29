package monitor

type Monitor interface {
	Query(query Query) []map[string]interface{}
}

type Query struct {
	Query      string
	Limit      int
	Start      int
	End        int
	SourceType string
}
