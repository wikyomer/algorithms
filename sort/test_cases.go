package sort

type args struct {
	array []int
}

var TestCases = []struct {
	name string
	args args
	want []int
}{
	{name: "Unsorted array", args: args{[]int{7, 2, 9, 12, 11}}, want: []int{2, 7, 9, 11, 12}},
	{name: "Identical elements", args: args{[]int{7, 7, 7, 7, 7}}, want: []int{7, 7, 7, 7, 7}},
	{name: "Sorted array", args: args{[]int{1, 2, 9, 12, 21}}, want: []int{1, 2, 9, 12, 21}},
	{name: "Empty array", args: args{[]int{}}, want: []int{}},
}
