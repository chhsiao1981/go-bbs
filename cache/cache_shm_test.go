package cache

import (
	"testing"

	"github.com/PichuChen/go-bbs/ptttype"
	"github.com/PichuChen/go-bbs/types"
	log "github.com/sirupsen/logrus"
)

func TestAttachSHM(t *testing.T) {
	setupTest()
	defer teardownTest()

	log.Infof("TestAttachSHM: to NewSHM: shm_key: %v USE_HUGETLB: %v", ptttype.SHM_KEY, ptttype.USE_HUGETLB)

	err := NewSHM(types.Key_t(ptttype.SHM_KEY), ptttype.USE_HUGETLB, true)
	if err != nil {
		return
	}
	defer CloseSHM()

	log.Infof("TestAttachSHM: after NewSHM")

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
