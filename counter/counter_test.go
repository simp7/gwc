package counter

import (
	"testing"
)

func Test_byteCounter_Count(t *testing.T) {
	type args struct {
		text []byte
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "normal", args: args{text: []byte("hello world")}, want: 11},
		{name: "empty", args: args{text: []byte("")}, want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &byteCounter{}
			if got := b.Count(tt.args.text); got != tt.want {
				t.Errorf("Count() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_characterCounter_Count(t *testing.T) {
	type args struct {
		text []byte
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "normal", args: args{text: []byte("hello world")}, want: 11},
		{name: "hangul", args: args{text: []byte("안녕 세상아")}, want: 6},
		{name: "mixed", args: args{text: []byte("Hello, 世界")}, want: 9},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &characterCounter{}
			if got := c.Count(tt.args.text); got != tt.want {
				t.Errorf("Count() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_lineCounter_Count(t *testing.T) {
	type args struct {
		text []byte
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "normal", args: args{text: []byte("hello world")}, want: 0},
		{name: "manyLines", args: args{text: []byte("hel\nlo,\n world")}, want: 2},
		{name: "empty", args: args{text: []byte("")}, want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &lineCounter{}
			if got := l.Count(tt.args.text); got != tt.want {
				t.Errorf("Count() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_wordCounter_Count(t *testing.T) {
	type args struct {
		text []byte
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "space", args: args{text: []byte("hello world")}, want: 2},
		{name: "tab", args: args{text: []byte("hello\tworld")}, want: 2},
		{name: "lineBreak", args: args{text: []byte("hello\nworld")}, want: 2},
		{name: "justSpace", args: args{text: []byte("   \n\t  \t \n")}, want: 0},
		{name: "manySpaces", args: args{text: []byte("\n\t\n  H \n E\tL L  O\t\n, World")}, want: 7},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &wordCounter{}
			if got := w.Count(tt.args.text); got != tt.want {
				t.Errorf("Count() = %v, want %v", got, tt.want)
			}
		})
	}
}
