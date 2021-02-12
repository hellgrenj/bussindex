package main

import (
	"fmt"
	"log"
	"os"

	"github.com/hellgrenj/bussindex/pkg/db"
	"github.com/hellgrenj/bussindex/pkg/rest"
	"github.com/hellgrenj/bussindex/pkg/system"
)

func main() {
	defer func() { //catch or finally
		if err := recover(); err != nil { //catch
			fmt.Fprintf(os.Stderr, "Exception: %v\n", err)
			os.Exit(1)
		}
	}()

	infoLogger := log.New(os.Stdout, "INFO: ", 3)
	errorLogger := log.New(os.Stderr, "ERROR: ", 3)
	driver, err := db.NewDriver("bolt://bussindexdb-neo4j-community:7687", "neo4j", "mySecretPassword")
	if err != nil {
		panic(err)
	}
	systemRepository := system.NewSystemRepository(driver)
	systemService := system.NewService(systemRepository, infoLogger)
	s := rest.NewServer(systemService, infoLogger, errorLogger)

	s.StartAndListen(8080)
}
