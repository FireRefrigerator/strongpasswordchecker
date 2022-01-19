package app

import "testing"

func TestCheckPassword(t *testing.T) {
	tests2 := []struct {
		name    string
		input   string
		wantErr bool
	}{
		{
			"pwd cover three increasing Digital is not strong",
			"#1Ztecom123",
			false,
		},
		{
			"pwd cover three decreasing Digital is not strong",
			"#1ztecom321",
			false,
		},
		{
			"pwd cover three increasing LowerCase is not strong",
			"#1ztecomcabc",
			false,
		},
		{
			"pwd cover three decreasing LowerCase is not strong",
			"#1ztecomcncba",
			false,
		},
		{
			"pwd cover three increasing UpperCase is not strong",
			"#1ztecomcnABC",
			false,
		},
		{
			"pwd cover three decreasing UpperCase is not strong",
			"#1ztecomcnCBA",
			false,
		},
	}
	for _, tt := range tests2 {
		t.Run(tt.name, func(t *testing.T) {
			err := CheckPassword(tt.input)
			if err != tt.wantErr {
				t.Errorf("test %s err", tt.name)
			}
		})
	}
}
