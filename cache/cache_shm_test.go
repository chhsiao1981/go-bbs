package cache

import (
	"testing"

	"github.com/PichuChen/go-bbs/ptttype"
)

func TestAttachSHM(t *testing.T) {
	setupTest()
	defer teardownTest()

	_ = NewSHM(TestShmKey, ptttype.USE_HUGETLB, true)
	defer CloseSHM()

	tests := []struct {
		name    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := AttachSHM(); (err != nil) != tt.wantErr {
				t.Errorf("AttachSHM() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}

}
