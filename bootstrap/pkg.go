package bootstrap

import (
	"database/sql"
	"fmt"
)

// DBConnect connects you to Postgresql based on Config.
func DBConnect(c *Config) (*sql.DB, error) {
	info := fmt.Sprintf("host=%v port=%v user=%v password=%v dbname=%v sslmode=disable",
		c.DBHost, c.DBPort, c.DBUser, c.DBPassword, c.DBName)

	db, err := sql.Open("postgres", info)
	if err != nil {
		return nil, err
	}

	return db, db.Ping()
}

//func NewRedisClient(p string) (*redis.Client, error) {
//	client := redis.NewClient(&redis.Options{
//		Addr: fmt.Sprintf("localhost:%v", p),
//	})
//
//	return client, client.Ping(context.Background()).Err()
//}
