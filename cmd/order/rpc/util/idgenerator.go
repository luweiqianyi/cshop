package util

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var gIDGeneratorOnce sync.Once
var gIDGenerator IDGenerator

func OrderIdGenerator() IDGenerator {
	gIDGeneratorOnce.Do(func() {
		gIDGenerator = IDGenerator{}
	})
	return gIDGenerator
}

type IDGenerator struct {
}

func (g IDGenerator) OrderID(businessCategory string) string {
	return fmt.Sprintf(
		"%s-%s%s",
		businessCategory,
		time.Now().Format("20060102_15:04:05_07:00.000"),
		Krand(8, 0))
}

func Krand(size int, kind int) string {
	ikind, kinds, result := kind, [][]int{[]int{10, 48}, []int{26, 97}, []int{26, 65}}, make([]byte, size)
	is_all := kind > 2 || kind < 0
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		if is_all { // random ikind
			ikind = rand.Intn(3)
		}
		scope, base := kinds[ikind][0], kinds[ikind][1]
		result[i] = uint8(base + rand.Intn(scope))
	}
	return string(result)
}
