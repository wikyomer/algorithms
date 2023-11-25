package sort

import (
	"reflect"
	"testing"
)

func TestMergeSort(t *testing.T) {
	tests := TestCases
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MergeSort(tt.args.array); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MergeSort() = %v, want %v", got, tt.want)
			}
		})
	}
}
