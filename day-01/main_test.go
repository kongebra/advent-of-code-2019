package main

import "testing"

func TestFuelCounterUpper(t *testing.T) {
	type args struct {
		mass float64
	}

	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "one",
			args: args{
				mass: 12,
			},
			want: 2,
		},
		{
			name: "two",
			args: args{
				mass: 14,
			},
			want: 2,
		},
		{
			name: "three",
			args: args{
				mass: 1969,
			},
			want: 654,
		},
		{
			name: "four",
			args: args{
				mass: 100756,
			},
			want: 33583,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FuelCounterUpper(tt.args.mass); got != tt.want {
				t.Errorf("FuelCounterUpper() = %v, want %v", got, tt.want)
			}
		})
	}
}
