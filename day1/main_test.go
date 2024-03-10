package main

import (
	"testing"
)

func Test_first(t *testing.T) {
	type args struct {
		line       string
		comparator []string
		currIdx    int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test case 1: target exists in comparator",
			want: 10,
			args: args{
				line:       "12345",
				comparator: []string{"1", "2", "3", "4", "5"},
				currIdx:    0,
			},
		},
		{
			name: "Test case 2: target does not exist in comparator",
			want: 0,
			args: args{
				line:       "12345",
				comparator: []string{"6", "7", "8", "9"},
				currIdx:    0,
			},
		},
		{
			name: "Test case 3: target is not a valid integer",
			want: 0,
			args: args{
				line:       "12345",
				comparator: []string{"a", "b", "c"},
				currIdx:    0,
			},
		},
		{
			name: "Test case 4: curr idx is greater than the length of line",
			want: 0,
			args: args{
				line:       "12345",
				comparator: []string{"1", "2", "3", "4", "5"},
				currIdx:    11,
			},
		},
		{
			name: "Test case 5: empty line",
			want: 0,
			args: args{
				line:       "",
				comparator: []string{"1", "2", "3", "4", "5"},
				currIdx:    0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := first(tt.args.line, tt.args.comparator, tt.args.currIdx); got != tt.want {
				t.Errorf("first() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_last(t *testing.T) {
	type args struct {
		line       string
		comparator []string
		currIdx    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test case 1: target exists in comparator",
			want: 5,
			args: args{
				line:       "12345",
				comparator: []string{"1", "2", "3", "4", "5"},
				currIdx:    4,
			},
		},
		{
			name: "Test case 2: target does not exists in comparator",
			want: 0,
			args: args{
				line:       "onetwothreefourfive",
				comparator: []string{"1", "2", "3", "4", "5"},
				currIdx:    4,
			},
		},
		{
			name: "Test case 3: target is not a valid integer",
			want: 0,
			args: args{
				line:       "12345",
				comparator: []string{"a", "b", "c"},
				currIdx:    2,
			},
		},
		{
			name: "Test case 4: currIdx is greater than the size of the line",
			want: 0,
			args: args{
				line:       "12345",
				comparator: []string{"1", "2", "3", "4", "5"},
				currIdx:    5,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := last(tt.args.line, tt.args.comparator, tt.args.currIdx); got != tt.want {
				t.Errorf("last() = %v, want %v", got, tt.want)
			}
		})
	}
}
