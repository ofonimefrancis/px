package driver

import (
	"fmt"

	"github.com/labstack/gommon/log"
	"github.com/ofonimefrancis/pixels/pkg/datastore"
	"github.com/ofonimefrancis/pixels/pkg/datastore/mongodb"
)

type Driver struct {
}

func NewConnectionDriver() *Driver {
	return &Driver{}
}

func (c *Driver) Connect(dbDriver, dbConnectionUrl, dbName string, timeout int) datastore.UserRepository {
	switch dbDriver {
	case "mongo":
		if dbConnectionUrl == "" {
			dbConnectionUrl = "mongodb://localhost:27017"
		}

		if dbName == "" {
			dbName = "test"
		}

		if timeout == 0 {
			timeout = 10
		}

		uri := fmt.Sprintf("%s/%s", dbConnectionUrl, dbName)
		ds, err := mongodb.NewUserRepository(uri, dbName, timeout)
		if err != nil {
			log.Fatal(err)
		}

		return ds
	}
	return nil
}
