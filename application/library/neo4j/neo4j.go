package neo4j

import neo4j "github.com/johnnadratowski/golang-neo4j-bolt-driver"

var (
	err  error
	pool neo4j.DriverPool
)

//Init 初始化neo4j
func Init(host, port, username, password string, maxpool int) error {
	if pool, err = neo4j.NewDriverPool("bolt://"+username+":"+password+"@"+host+":"+port, maxpool); err != nil {
		return err
	}
	return nil
}

//OpenClient get客户端
func OpenClient() (neo4j.Conn, error) {
	return pool.OpenPool()
}
