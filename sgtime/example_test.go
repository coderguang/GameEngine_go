package sgtime

import (
	"fmt"
	"testing"
)

func Test_NextUpdate(t *testing.T) {
	now := New()
	fmt.Println(",now is:", GetTotalSecond(now))
	t.Log("test time ok")
}
