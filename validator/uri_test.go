package validator

import "testing"

func Test_httpURI(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"valid: http",
			args{
				"http://www.github.com",
			},
			false,
		},
		{
			"valid: https",
			args{
				"https://www.github.com",
			},
			false,
		},
		{
			"valid: https with uppercase",
			args{
				"HTTPS://www.github.com/?test=1",
			},
			false,
		},
		{
			"valid: http with query",
			args{
				"https://www.github.com/?test=1",
			},
			false,
		},
		{
			"invalid: ssh",
			args{
				"ssh://www.github.com/?test=1",
			},
			true,
		},
		{
			"invalid: string",
			args{
				"github.com",
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := HttpURI(tt.args.s); (err != nil) != tt.wantErr {
				t.Errorf("uriValidator() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
