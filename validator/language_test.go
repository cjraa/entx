package validator

import "testing"

func TestValidLanguage(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"valid: english",
			args{"en"},
			false,
		},
		{
			"valid: english with locale",
			args{"en-US"},
			false,
		},
		{
			"valid: spanish",
			args{"es"},
			false,
		},
		{
			"invalid: problem",
			args{"problem"},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ValidLanguage(tt.args.s); (err != nil) != tt.wantErr {
				t.Errorf("ValidLanguage() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
