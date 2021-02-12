package system

import "github.com/neo4j/neo4j-go-driver/v4/neo4j"

// Repository is responsible for storing and retreiving systems
type Repository struct {
	driver neo4j.Driver
}

// NewSystemRepository creates and returns a new SystemRepository
func NewSystemRepository(driver *neo4j.Driver) DBRepository {
	return &Repository{driver: *driver}
}

// Save a system in neo4j
func (r *Repository) Save(system System) (int64, error) {
	session := r.driver.NewSession(neo4j.SessionConfig{})
	defer session.Close()

	persistedSystem, err := session.WriteTransaction(func(transaction neo4j.Transaction) (interface{}, error) {
		result, err := transaction.Run(
			"CREATE (a:System) SET a.description = $description RETURN id(a)",
			map[string]interface{}{"description": system.Description})
		if err != nil {
			return nil, err
		}

		if result.Next() {
			return result.Record().Values[0], nil
		}

		return nil, result.Err()
	})
	if err != nil {
		return 0, err
	}

	return persistedSystem.(int64), nil
}
