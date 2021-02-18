package system

import (
	"errors"

	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

// Repository is responsible for storing and retreiving systems
type Repository struct {
	driver neo4j.Driver
}

// NewSystemRepository creates and returns a new SystemRepository
func NewSystemRepository(driver *neo4j.Driver) DBRepository {
	return &Repository{driver: *driver}
}

// Save a system node in neo4j
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

// Delete a system node in neo4j
func (r *Repository) Delete(id int) error {

	session := r.driver.NewSession(neo4j.SessionConfig{})
	defer session.Close()

	_, err := session.WriteTransaction(func(transaction neo4j.Transaction) (interface{}, error) {
		_, err := transaction.Run(
			"MATCH (n) where ID(n)=$id DETACH DELETE n",
			map[string]interface{}{"id": id})
		if err != nil {
			return nil, err
		}
		return nil, nil
	})
	if err != nil {
		return err
	}

	return nil
}

// Get all the system nodes from neo4j
func (r *Repository) Get() ([]System, error) {
	session := r.driver.NewSession(neo4j.SessionConfig{})
	defer session.Close()

	allSystems, err := session.ReadTransaction(func(transaction neo4j.Transaction) (interface{}, error) {
		result, err := transaction.Run(`MATCH (s:System) RETURN ID(s) as id, s.description`, nil)
		if err != nil {
			return nil, err
		}

		var systems []System
		for result.Next() {

			desc, foundDescripton := result.Record().Get("s.description")
			if !foundDescripton {
				return nil, errors.New("missing description")
			}

			id, foundID := result.Record().Get("id")
			if !foundID {
				return nil, errors.New("missing id")
			}
			system := &System{ID: id.(int64), Description: desc.(string)}
			systems = append(systems, *system)
		}

		return systems, result.Err()
	})
	if err != nil {
		return nil, err
	}

	return allSystems.([]System), nil
}
