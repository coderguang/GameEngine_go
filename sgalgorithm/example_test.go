package sgalgorithm

import (
	"log"
	"math"
	"strconv"
	"testing"

	"github.com/coderguang/GameEngine_go/sglog"
	"github.com/coderguang/GameEngine_go/sgserver"
)

func TestPermutation(t *testing.T) {
	sgserver.StartLogServer("debug", "../../log/", log.LstdFlags, true)

	src := []string{"1", "2", "3", "4"}
	result := []string{}

	GenPermutation(src, 3, &result)

	sglog.Info("result:", result)

	if len(result) != 64 {
		t.Error("result error")
	}

	sgserver.StopLogServer()

	t.Log("test permutation ok")

}

func BenchmarkPermutation(b *testing.B) {
	sgserver.StartLogServer("debug", "../../log/", log.LstdFlags, true)

	src := []string{}
	for i := 0; i < 100; i++ {
		src = append(src, strconv.Itoa(i))
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		if i == 0 {
			continue
		}

		num := i
		if num > 3 {
			num = 3
		}
		result := []string{}
		GenPermutation(src, num, &result)
		resultNum := math.Pow(float64(len(src)), float64(num))
		if int(resultNum) != len(result) {
			sglog.Error("data error,num:", num, "should:", resultNum, ",real:", len(result))
		}
	}

	sgserver.StopLogServer()

}
