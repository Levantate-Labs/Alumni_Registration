package helpers

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	letterBytes = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	digits      = "0123456789"
)

type IDGenerator struct {
	mu sync.Mutex
}

var uid_gen *IDGenerator

func InitUIDGen() {
	uid_gen = &IDGenerator{}
}

func UIDGen() *IDGenerator {
	return uid_gen
}

func (g *IDGenerator) GenerateID(prefix string) string {
	rand.New(rand.NewSource(time.Now().UnixNano()))

	letter := letterBytes[rand.Intn(len(letterBytes))]

	num1 := rand.Intn(10) // Random digit between 0 and 9
	num2 := rand.Intn(10) // Random digit between 0 and 9
	num3 := rand.Intn(9)  // Random digit between 0 and 9

	// Format the string
	randomString := fmt.Sprintf("%s%c%d%d%d", prefix, letter, num1, num2, num3)

	return randomString
}

func GenerateOTP() uint8 {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	min := 100000
	max := 999999
	return uint8(rand.Intn(max-min+1) + min)
}
