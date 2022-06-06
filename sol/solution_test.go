package sol

import (
	"reflect"
	"testing"
)

func BenchmarkTest(b *testing.B) {
	nums := []int{1, 2, 3}
	for idx := 0; idx < b.N; idx++ {
		subsets(nums)
	}
}
func Test_subsets(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "nums = [1,2,3]",
			args: args{nums: []int{1, 2, 3}},
			want: [][]int{
				{1, 2, 3}, {1, 2}, {1, 3}, {1}, {2, 3}, {2}, {3}, {}},
		},
		{
			name: "nums = [0]",
			args: args{nums: []int{0}},
			want: [][]int{{0}, {}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := subsets(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("subsets() = %v, want %v", got, tt.want)
			}
		})
	}
}
