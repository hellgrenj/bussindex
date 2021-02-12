package db

import (
	"fmt"

	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

// NewDriver initiliazes and returns a new Neo4j driver
func NewDriver(uri string, username string, password string) (*neo4j.Driver, error) {
	driver, err := neo4j.NewDriver(uri, neo4j.BasicAuth(username, password, ""))
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return &driver, nil
}
