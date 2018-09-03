package memo

import (
	"testing"
	"github.com/cleonty/gopl/ch9/memotest"
)

var httpGetBody = memotest.HTTPGetBody

func TestMemo(t *testing.T) {
	m := New(httpGetBody)
	defer m.Close()
	memotest.Sequental(t, m)
}	

func TestMemoConcurrent(t *testing.T) {
	m := New(httpGetBody)
	defer m.Close()
	memotest.Concurrent(t, m)
}	
