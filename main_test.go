package main

import "testing"

func Test_kelem(t *testing.T) {
	s1 := []int{1, 3, 5, 7, 9}
	s2 := []int{2, 4, 6}
	type args struct {
		s1 []int
		s2 []int
		k  int
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{name: "k is zero", args: args{s1: s1, s2: s2, k: 0}, wantErr: true},
		{name: "k is too big", args: args{s1: s1, s2: s2, k: len(s1) + len(s2) + 1}, wantErr: true},
		{name: "k is the first", args: args{s1: s1, s2: s2, k: 1}, want: 1, wantErr: false},
		{name: "k is the last", args: args{s1: s1, s2: s2, k: len(s1) + len(s2)}, want: 9, wantErr: false},
		{name: "k is the middle", args: args{s1: s1, s2: s2, k: 4}, want: 4, wantErr: false},
		{name: "k is the last of s1", args: args{s1: s1, s2: s2, k: 8}, want: 9, wantErr: false},
		{name: "k is the last of s2", args: args{s1: s1, s2: s2, k: 6}, want: 6, wantErr: false},
		{name: "k is the first of s2", args: args{s1: s1, s2: s2, k: 2}, want: 2, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := kelem(tt.args.s1, tt.args.s2, tt.args.k)
			if (err != nil) != tt.wantErr {
				t.Errorf("kelem() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("kelem() = %v, want %v", got, tt.want)
			}
		})
	}
}
