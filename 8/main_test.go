package main

import (
	"reflect"
	"testing"
)

func Test_median(t *testing.T) {
	type args struct {
		array []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			"three",
			args{
				array: []string{"раз", "два", "три"},
			},
			[]string{"два"},
		},
		{
			"four",
			args{
				array: []string{"раз", "два", "три", "четыре"},
			},
			[]string{"два", "три"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := median(tt.args.array); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("median() = %v, want %v", got, tt.want)
			}
		})
	}
}
