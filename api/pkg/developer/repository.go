package developer

import (
	"errors"
	"time"

	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

// Repository is responsible for storing and retreiving systems
type Repository struct {
	driver neo4j.Driver
}

// NewDeveloperRepository creates and returns a new DeveloperRepository
func NewDeveloperRepository(driver *neo4j.Driver) IRepository {
	return &Repository{driver: *driver}
}

// Save a system node in neo4j
func (r *Repository) Save(developer Developer) (int64, error) {
	session := r.driver.NewSession(neo4j.SessionConfig{})
	defer session.Close()

	persistedDeveloperID, err := session.WriteTransaction(func(transaction neo4j.Transaction) (interface{}, error) {
		result, err := transaction.Run(
			"CREATE (d:Developer) SET d.name = $name, d.dateOfEmployment = $doe RETURN id(d)",
			map[string]interface{}{"name": developer.Name, "doe": developer.DateOfEmployment.UTC()})
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

	return persistedDeveloperID.(int64), nil
}

// Delete a system node in neo4j
func (r *Repository) Delete(id int) error {

	session := r.driver.NewSession(neo4j.SessionConfig{})
	defer session.Close()

	_, err := session.WriteTransaction(func(transaction neo4j.Transaction) (interface{}, error) {
		_, err := transaction.Run(
			"MATCH (d) where ID(d)=$id DETACH DELETE d",
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
func (r *Repository) Get() ([]Developer, error) {
	session := r.driver.NewSession(neo4j.SessionConfig{})
	defer session.Close()

	allDevelopers, err := session.ReadTransaction(func(transaction neo4j.Transaction) (interface{}, error) {
		result, err := transaction.Run(`MATCH (d:Developer) RETURN ID(d) as id, d.name, d.dateOfEmployment`, nil)
		if err != nil {
			return nil, err
		}

		var developers []Developer
		for result.Next() {

			name, foundName := result.Record().Get("d.name")
			if !foundName {
				return nil, errors.New("missing name")
			}
			dateOfEmployment, foundDateOfEmployment := result.Record().Get("d.dateOfEmployment")
			if !foundDateOfEmployment {
				return nil, errors.New("missing dateOfEmployment")
			}

			id, foundID := result.Record().Get("id")
			if !foundID {
				return nil, errors.New("missing id")
			}
			developer := &Developer{ID: id.(int64), Name: name.(string), DateOfEmployment: dateOfEmployment.(time.Time)}
			developers = append(developers, *developer)
		}

		return developers, result.Err()
	})
	if err != nil {
		return nil, err
	}

	return allDevelopers.([]Developer), nil
}
