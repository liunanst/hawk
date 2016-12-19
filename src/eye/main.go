package main

import (
	"baselib"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"net"
	"os"
	"time"
)

const (
	FRAME_LEN = 18
)

var (
	log       *baselib.Logger
	redispoll redis.Pool
)

func initSvr() {

	err := LoadConfig()
	if err != nil {
		fmt.Printf("init failed , exit!")
		os.Exit(-1)
	}

	config := &baselib.PoolConfig{
		Network:      "tcp",
		Address:      Setting.RedisAddr,
		Passwd:       Setting.RedisPasswd,
		Idle:         5,
		Active:       5,
		ConnTimeout:  5 * time.Second,
		ReadTimeout:  2 * time.Second,
		WriteTimeout: 2 * time.Second,
		IdleTimeout:  10 * time.Second,
	}

	fmt.Printf("Log level %d, log file %s\n", Setting.Log.LogLevel, Setting.Log.LogFile)
	log, _ = baselib.NewLogger(Setting.Log.LogFile, Setting.Log.LogLevel)
	log.Info("server start")

	redispoll = baselib.CreateRedisConnPool(config, log)
	r := redispoll.Get()
	if err := r.Err(); err != nil {
		log.Error("get connection to redis failed,%s", err.Error())
		r.Close()
		return
	}
	r.Close()
	// save pool

	// use pool
	//r = pool.Get()
	//defer r.Close()
	//
	//results, err := redis.Strings(r.Do("zrevrangebyscore", param.zKey, param.max, param.min, "limit", 0, param.count))

}

func main() {
	initSvr()

	l, err := net.Listen("tcp", Setting.LocalAddr)
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}
	defer l.Close()
	fmt.Println("Listening on " + Setting.LocalAddr)

	for {
		// Listen for an incoming connection.
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			os.Exit(1)
		}
		// Handle connections in a new goroutine.
		go handleRequest(conn)
	}
}

func parseFrame(buf []byte) int {
	mac := buf[0:6]
	fmt.Printf("MAC: %x:%x:%x:%x:%x:%x  RSSI %d", mac[0], mac[1], mac[2], mac[3], mac[4], mac[5], buf[7])
	return 0
}

// Handles incoming requests.
func handleRequest(conn net.Conn) {
	var head int
	var tail int

	head = 0
	tail = 0
	buf := make([]byte, 1000)
	for {
		reqLen, err := conn.Read(buf[tail:])
		if err != nil {
			fmt.Println("Error reading:", err.Error())
			conn.Close()
			return
		}
		tail += reqLen
		if tail < FRAME_LEN {
			continue
		}
		for head = 0; head < tail; head += 1 {
			if buf[head] != 0x54 {
				continue
			}
			if head+1 < tail {
				if buf[head+1] != 0x58 {
					continue
				} else {
					break
				}
			}
		}
		for ; head+FRAME_LEN < tail; head += FRAME_LEN {
			ret := parseFrame(buf)
			if ret == -1 {
				conn.Close()
				return
			}
		}

		newBuf := make([]byte, 1000)
		copy(newBuf, buf[head:tail])
		buf = newBuf
		tail = 0
	}
}
