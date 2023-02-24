package main

import (
	"reflect"
	"testing"
)

// Unit test fucntion
func Test_main(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "",
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			main()
		})
	}
}

func Test_removeSuccess(t *testing.T) {
	type args struct {
		slice []int
		s     int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "test-index-0",
			args: args{
				slice: []int{1, 2, 3},
				s:     0,
			},
			want: []int{2, 3},
		},
		{
			name: "test-index-1",
			args: args{
				slice: []int{1, 2, 3},
				s:     1,
			},
			want: []int{1, 3},
		},
		{
			name: "test-index-2",
			args: args{
				slice: []int{1, 2, 3},
				s:     2,
			},
			want: []int{1, 2},
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := remove(tt.args.slice, tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("remove() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_removePlayerSuccess(t *testing.T) {
	type args struct {
		slice []Player
		s     int
	}
	tests := []struct {
		name string
		args args
		want []Player
	}{
		{
			name: "test-index-0",
			args: args{
				slice: []Player{
					{
						Dice:      []int{},
						TotalDice: 0,
						Score:     0,
						Name:      0,
					},
					{
						Dice:      []int{},
						TotalDice: 1,
						Score:     1,
						Name:      1,
					},
				},
				s: 0,
			},
			want: []Player{
				{
					Dice:      []int{},
					TotalDice: 1,
					Score:     1,
					Name:      1,
				}},
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removePlayer(tt.args.slice, tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("removePlayer() = %v, want %v", got, tt.want)
			}
		})
	}
}
