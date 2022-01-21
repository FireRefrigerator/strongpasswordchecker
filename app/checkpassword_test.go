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

// step 6 将弱密码（长度小于8或大于20，或字符类型少于三种，或存在连续三个相同字符，或有连续三个递增或者递减字符）转化为强密码的最小步数
// less more type count（连续3个重复，或递增，或递减） 四种类型的弱密码过滤器
func TestCheckPaaswordStep(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		wantStep int
	}{
		{
			"pwd #1Zteco is not strong, add 1 char",
			"#1Zteco",
			1,
		},
		{
			"pwd #1Ztecoztecomcnztecomz is not strong, delete 2 char",
			"#1Ztecoztecomcnztecomz",
			2,
		},
		{
			"pwd 13579135 is not strong, replace 2 char",
			"13579135",
			2,
		},
		{
			"pwd 12ztecom is not strong, replace 1 char",
			"12ztecom",
			1,
		},
		{
			"pwd aaaaa1Com is not strong, replace 1 char",
			"aaaaa1Com",
			1,
		},
		{
			"pwd abcde1Com is not strong, replace 1 char",
			"abcde1Com",
			1,
		},
		{
			"pwd edcba1Com is not strong, replace 1 char",
			"edcba1Com",
			1,
		},
		{
			"pwd ztecomc is not strong, add 1 char, replace 1 char",
			"ztecomc",
			2,
		},
		{
			"pwd zzzCom1 is not strong, add 1 char",
			"zzzCom1",
			1,
		},
		{
			"pwd zzzcomc is not strong, add 1 char, replace 1 char",
			"zzzcomc",
			2,
		},
		{
			"pwd ztecomcnztecomcnzteco is not strong, delete 1 replace 2",
			"ztecomcnztecomcnzteco",
			3,
		},
		{
			"pwd Zt1comcnztecomcnzteee is not strong, delete 1",
			"Zt1comcnztecomcnzteee",
			1,
		},
		{
			"pwd ztecomcnztecomcnzteee is not strong, delete 1 replace 2",
			"ztecomcnztecomcnzteee",
			3,
		},
		{
			"pwd AAAAAAAAA123456789AAAA is not strong, delete 2, replace 5 char",
			"AAAAAAAAA123456789AAAA",
			7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			step := CheckPasswordStep(tt.input)
			if step != tt.wantStep {
				t.Errorf("test %s err, want %d actual %d", tt.name, tt.wantStep, step)
			}
		})
	}
}
