package memo

import (
	"testing"

	"github.com/cleonty/gopl/ch9/memotest"
)

var httpGetBody = memotest.HTTPGetBody

func TestMemo(t *testing.T) {
	cancel := make(chan string)
	m := New(httpGetBody, cancel)
	defer m.Close()
	memotest.Sequental(t, m)
}

func TestMemoConcurrent(t *testing.T) {
	cancel := make(chan string)
	m := New(httpGetBody, cancel)
	defer m.Close()
	memotest.Concurrent(t, m)
}

func TestConcurrentWithCancel(t *testing.T) {
	cancel := make(chan string)
	m := New(httpGetBody, cancel)
	defer m.Close()
	memotest.ConcurrentWithCancel(t, m, cancel)
}
