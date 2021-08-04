package helper

import (
	"testing"
)

func TestFormatDate(t *testing.T) {
	type args struct {
		timestamp int64
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "01",
			args: args{
				timestamp: 1628058775,
			},
			want: "2021-08-04",
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				if got := FormatDate(tt.args.timestamp); got != tt.want {
					t.Errorf("FormatDate() = %v, want %v", got, tt.want)
				}
			},
		)
	}
}

func TestFormatTime(t *testing.T) {
	type args struct {
		timestamp int64
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "01",
			args: args{
				timestamp: 1628058775,
			},
			want: "2021-08-04 14:32:55",
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				if got := FormatTime(tt.args.timestamp); got != tt.want {
					t.Errorf("FormatTime() = %v, want %v", got, tt.want)
				}
			},
		)
	}
}

func TestRemoveIllegalCharacters(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "01",
			args: args{
				str: "【你好呀】： / :111",
			},
			want: "【你好呀】111",
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				if got := RemoveIllegalCharacters(tt.args.str); got != tt.want {
					t.Errorf("RemoveIllegalCharacters() = %v, want %v", got, tt.want)
				}
			},
		)
	}
}
