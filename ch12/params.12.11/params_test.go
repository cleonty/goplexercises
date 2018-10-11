package params

import "testing"

type data struct {
	Labels     []string `http:"l"`
	MaxResults int      `http:"max"`
	Exact      bool     `http:"x"`
}

func TestPack(t *testing.T) {
	type args struct {
		ptr interface{}
	}
	d := data{
		Labels:     []string{"l1", "l2", "l3"},
		MaxResults: 5,
		Exact:      true,
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"test 1",
			args{&d},
			"l=l1&l=l2&l=l3&max=5&x=true",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Pack(tt.args.ptr); got != tt.want {
				t.Errorf("Pack() = %v, want %v", got, tt.want)
			}
		})
	}
}
