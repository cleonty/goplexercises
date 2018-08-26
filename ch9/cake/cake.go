package cake

type Cake struct{ state string }

func baker(cooked chan<- *Cake) {
	for {
		cake := new(Cake)
		cake.state = "cooked"
		cooked <- cake
	}
}

func icer(iced chan<- *Cake, baked <-chan *Cake) {
	for cake := range baked {
		cake.state = "iced"
		iced <- cake
	}
}

var baked = make(chan *Cake)
var iced = make(chan *Cake)

func init() {
	go baker(baked)
	go icer(iced, baked)
}
