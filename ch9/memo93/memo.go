package memo

type entry struct {
	res   result
	ready chan struct{}
}

type Memo struct {
	requests chan request
	cancel   <-chan string
}

type Func func(key string) (interface{}, error)
type FuncWithCancel struct {
	fn     Func
	cancel <-chan string
}

type result struct {
	value interface{}
	err   error
}

type request struct {
	key      string
	response chan<- result
}

func New(f Func, cancel <-chan string) *Memo {
	memo := &Memo{make(chan request), cancel}
	go memo.server(f)
	return memo
}

func (memo *Memo) Get(key string) (interface{}, error) {
	response := make(chan result)
	memo.requests <- request{key, response}
	res := <-response
	return res.value, res.err
}

func (memo *Memo) Close() { close(memo.requests) }

func (memo *Memo) server(f Func) {
	cache := make(map[string]*entry)
	for {
		select {
		case req := <-memo.requests:
			e := cache[req.key]
			if e == nil {
				e = &entry{ready: make(chan struct{})}
				cache[req.key] = e
				go e.call(f, req.key)
			}
			go e.deliver(req.response)
		case key := <-memo.cancel:
			delete(cache, key)
		}
	}
}

func (e *entry) call(f Func, key string) {
	e.res.value, e.res.err = f(key)
	close(e.ready)
}

func (e *entry) deliver(response chan<- result) {
	<-e.ready
	response <- e.res
}
