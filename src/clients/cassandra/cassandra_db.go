package cassandra

import (
	"fmt"

	"github.com/gocql/gocql"
)

var (
	cluster *gocql.ClusterConfig
)

func init() {
	// connect to the Cassandra cluster
	cluster = gocql.NewCluster("127.0.0.1")
	cluster.Keyspace = "oauth"
	cluster.Consistency = gocql.Quorum

	session, err := cluster.CreateSession()
	if err != nil {
		panic(err)
	}
	defer session.Close()

	fmt.Println("Cassandra connection successfully created!")
}

func GetSession() (*gocql.Session, error) {
	return cluster.CreateSession()
}
