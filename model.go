package main

// Data holds the information about OBD reading
// split by key and value
type Data struct {
	Key   string
	Value string
}

// Store is our interface on where to put the data
type Store interface {
	Open() error
	Insert(map[string]interface{}) error
	Close() error
}
