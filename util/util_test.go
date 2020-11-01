package util

import "testing"

func TestIsValidEmail(t *testing.T) {
	type args struct {
		email string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "valid email format",
			args: args{
				email: "abcd@gmail.com",
			},
			want: true,
		},
		{
			name: "Invalid email format",
			args: args{
				email: "invalidEmail",
			},
			want: false,
		},
		{
			name: "empty email",
			args: args{},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsValidEmail(tt.args.email); got != tt.want {
				t.Errorf("IsValidEmail() = %v, want %v", got, tt.want)
			}
		})
	}
}
