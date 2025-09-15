package robot

import "testing"

func Test_sort(t *testing.T) {
	type args struct {
		width  int
		height int
		length int
		mass   int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "standard package",
			args: args{
				width:  100,
				height: 10,
				length: 10,
				mass:   10,
			},
			want: STANDARD,
		},
		{
			name: "bulky package due to width",
			args: args{
				width:  150,
				height: 10,
				length: 10,
				mass:   10,
			},
			want: SPECIAL,
		},
		{
			name: "bulky package due to volume",
			args: args{
				width:  100,
				height: 100,
				length: 100,
				mass:   10,
			},
			want: SPECIAL,
		},
		{
			name: "Almost but not bulky package",
			args: args{
				width:  99,
				height: 100,
				length: 100,
				mass:   10,
			},
			want: STANDARD,
		},
		{
			name: "heavy package",
			args: args{
				width:  10,
				height: 10,
				length: 10,
				mass:   20,
			},
			want: SPECIAL,
		},
		{
			name: "bulky and heavy package",
			args: args{
				width:  150,
				height: 10,
				length: 10,
				mass:   20,
			},
			want: REJECTED,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Sort(tt.args.width, tt.args.height, tt.args.length, tt.args.mass); got != tt.want {
				t.Errorf("sort() = %v, want %v", got, tt.want)
			}
		})
	}
}
