package ptttype

import "testing"

func TestInitConfig(t *testing.T) {
	setupTest()
	defer teardownTest()

	type args struct {
		filename string
	}
	tests := []struct {
		name        string
		args        args
		wantErr     bool
		wantBBSNAME string
	}{
		// TODO: Add test cases.
		{
			args:        args{"./testcase/test.ini"},
			wantBBSNAME: "test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := InitConfig(tt.args.filename); (err != nil) != tt.wantErr {
				t.Errorf("InitConfig() error = %v, wantErr %v", err, tt.wantErr)
			}

			if BBSNAME != tt.wantBBSNAME {
				t.Errorf("BBSNAME: InitConfig(): %v want: %v", BBSNAME, tt.wantBBSNAME)
			}
		})
	}
}
