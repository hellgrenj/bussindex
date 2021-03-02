package system

import (
	"errors"
	"fmt"

	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

// Repository is responsible for storing and retreiving systems
type Repository struct {
	driver neo4j.Driver
}

// NewSystemRepository creates and returns a new SystemRepository
func NewSystemRepository(driver *neo4j.Driver) IRepository {
	return &Repository{driver: *driver}
}

// Save a system node in neo4j
func (r *Repository) Save(system System) (int64, error) {
	session := r.driver.NewSession(neo4j.SessionConfig{})
	defer session.Close()

	persistedSystemID, err := session.WriteTransaction(func(transaction neo4j.Transaction) (interface{}, error) {
		result, err := transaction.Run(
			"CREATE (a:System) SET a.name = $name RETURN id(a)",
			map[string]interface{}{"name": system.Name})
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

	return persistedSystemID.(int64), nil
}

// AddDeveloper adds a relationship to a developer node (Dev working with system)
func (r *Repository) AddDeveloper(systemID int, developerID int) error {
	session := r.driver.NewSession(neo4j.SessionConfig{})
	defer session.Close()

	_, err := session.WriteTransaction(func(transaction neo4j.Transaction) (interface{}, error) {
		result, err := transaction.Run(
			"MATCH (a:System), (b:Developer) WHERE ID(a) = $systemId AND ID(b) = $developerId CREATE (b)-[r:WORKING_ON]->(a)",
			map[string]interface{}{"systemId": systemID, "developerId": developerID})
		if err != nil {
			return nil, err
		}
		if result.Next() {
			return result.Record().Values[0], nil
		}
		return nil, result.Err()
	})
	if err != nil {
		return err
	}

	return nil
}

//
// RemoveDeveloper removes a relationship with a developer node
func (r *Repository) RemoveDeveloper(systemID int, developerID int) error {
	session := r.driver.NewSession(neo4j.SessionConfig{})
	defer session.Close()

	_, err := session.WriteTransaction(func(transaction neo4j.Transaction) (interface{}, error) {
		result, err := transaction.Run(

			"MATCH (s:System)<-[r:WORKING_ON]-(developer) WHERE ID(s) = $systemId AND ID(developer) = $developerId  Delete r",
			map[string]interface{}{"systemId": systemID, "developerId": developerID})
		if err != nil {
			return nil, err
		}
		if result.Next() {
			return result.Record().Values[0], nil
		}
		return nil, result.Err()
	})
	if err != nil {
		return err
	}

	return nil
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

// GetDevIdsWorkingOnSystem takes a system id and returns a list of developer ids working on the system
func (r *Repository) GetDevIdsWorkingOnSystem(systemID int) ([]int64, error) {
	session := r.driver.NewSession(neo4j.SessionConfig{})
	defer session.Close()

	devIds, err := session.ReadTransaction(func(transaction neo4j.Transaction) (interface{}, error) {
		result, err := transaction.Run(
			"MATCH (s:System)<-[:WORKING_ON]-(developer) WHERE ID(s) = $systemId RETURN ID(developer) as id",
			map[string]interface{}{"systemId": systemID})
		if err != nil {
			return nil, err
		}

		var devIds []int64
		for result.Next() {
			fmt.Println(result.Record())
			id, foundID := result.Record().Get("id")
			if !foundID {
				return nil, errors.New("missing id")
			}

			devIds = append(devIds, id.(int64))
		}

		return devIds, result.Err()
	})
	if err != nil {
		return nil, err
	}

	return devIds.([]int64), nil
}

// Get all the system nodes from neo4j
func (r *Repository) Get() ([]System, error) {
	session := r.driver.NewSession(neo4j.SessionConfig{})
	defer session.Close()

	allSystems, err := session.ReadTransaction(func(transaction neo4j.Transaction) (interface{}, error) {
		result, err := transaction.Run(`MATCH (s:System) RETURN ID(s) as id, s.name`, nil)
		if err != nil {
			return nil, err
		}

		var systems []System
		for result.Next() {

			name, foundName := result.Record().Get("s.name")
			if !foundName {
				return nil, errors.New("missing name")
			}

			id, foundID := result.Record().Get("id")
			if !foundID {
				return nil, errors.New("missing id")
			}
			system := &System{ID: id.(int64), Name: name.(string)}
			systems = append(systems, *system)
		}

		return systems, result.Err()
	})
	if err != nil {
		return nil, err
	}

	return allSystems.([]System), nil
}
