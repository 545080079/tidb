package commontest

import (
	"github.com/stretchr/testify/require"
	"go.uber.org/goleak"
	"math/rand"
	"strconv"
	"testing"
)

func TestMain(m *testing.M) {
	goleak.VerifyTestMain(m)
}

func TestStrLenOfUint64Fast(t *testing.T) {
	for i := 0; i < 1000000; i++ {
		num := rand.Uint64()
		expected := len(strconv.FormatUint(num, 10))
		actual := StrLenOfUint64Fast(num)
		require.Equal(t, expected, actual)
	}
}

func StrLenOfUint64Fast(num uint64) int {
	return len(strconv.FormatUint(num, 10))
}
