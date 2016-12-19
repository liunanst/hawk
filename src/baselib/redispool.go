package baselib

import (
	"github.com/garyburd/redigo/redis"
	"time"
)

type PoolConfig struct {
	Network      string
	Address      string
	ConnTimeout  time.Duration
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	Idle         int
	Active       int
	IdleTimeout  time.Duration
	Passwd       string
}

func connect(log *Logger, network, address string, connectTimeout, readTimeout, writeTimeout time.Duration, passwd string) func() (redis.Conn, error) {
	return func() (redis.Conn, error) {
		c, err := redis.DialTimeout(network, address, connectTimeout, readTimeout, writeTimeout)

		if err != nil {
			log.Error("redis Dial failed:%s,%s,%s", network, address, err)
			return nil, err
		}
		//fmt.Printf("redis conn connectTimeout %d, readTimeout %d, writeTimeout %d \n", connectTimeout,readTimeout,writeTimeout)
		if len(passwd) > 0 {
			_, err = c.Do("auth", passwd)
			if err != nil {
				log.Error("redis auth failed:%s,%s,%s", network, address, passwd)
				c.Close()
				return nil, err
			}
		}
		return c, nil
	}

}

//create quote cache connetion pool
func CreateRedisConnPool(config *PoolConfig, log *Logger) redis.Pool {
	pool := redis.Pool{
		Dial:        connect(log, config.Network, config.Address, config.ConnTimeout, config.ReadTimeout, config.WriteTimeout, config.Passwd),
		MaxIdle:     config.Idle,
		MaxActive:   config.Active,
		IdleTimeout: config.IdleTimeout,
	}
	//fmt.Printf("redis pool MaxIdle %d, MaxActive %d, IdleTimeout %d \n", config.Idle,config.Active,config.IdleTimeout)

	return pool
}

func example() {
	config := &PoolConfig{
		Network:      "tcp",
		Address:      "127.0.0.1",
		Passwd:       "passwd",
		Idle:         5,
		Active:       5,
		ConnTimeout:  5 * time.Second,
		ReadTimeout:  2 * time.Second,
		WriteTimeout: 2 * time.Second,
		IdleTimeout:  10 * time.Second,
	}

	log, _ := NewLogger("./demo", 5)
	pool := CreateRedisConnPool(config, log)
	r := pool.Get()
	if err := r.Err(); err != nil {
		//log.Error("get connection from redis %s load pool failed,%s", market.Redishost, err.Error())
		r.Close()
		return
	}
	r.Close()
	// save pool

	// use pool
	r = pool.Get()
	defer r.Close()

	//results, err := redis.Strings(r.Do("zrevrangebyscore", param.zKey, param.max, param.min, "limit", 0, param.count))

}
