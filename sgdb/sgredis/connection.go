package sgredis

import (
	"github.com/gomodule/redigo/redis"
)

func NewConnection(host string, port string) (redis.Conn, error) {
	conn, err := redis.Dial("tcp", host+":"+port)

	if err != nil {
		return nil, err
	}
	return conn, nil
}
