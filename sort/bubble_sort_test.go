package sort

import (
	"fmt"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	a := []int{1,7,9,4,8,90,2,5,7,8,12,45,3,4,6,8,9,45,3}

	type args struct {
		a []int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "normal case",
			args: args{
				a: a,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			BubbleSort(a)
			fmt.Println(a)
		})
	}
}
