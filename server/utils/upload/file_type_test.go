package upload

import (
	"path"
	"testing"
)

func TestAwsS3_ContentType(t *testing.T) {
	tests := []struct {
		filePath string
		wantErr  string
	}{
		{
			filePath: "C:\\workspace\\1705282371apifox-logo-64 (1).png",
			wantErr:  "image/png",
		}, {
			filePath: "C:\\workspace\\store_sales_order.csv",
			wantErr:  "text/plain; charset=utf-8",
		},
	}
	for _, tt := range tests {
		t.Run(tt.filePath, func(t *testing.T) {
			mineType, err := DetermineByPath(tt.filePath)
			if err != nil {
				t.Errorf("DetermineByPath() failed")
				return
			}

			if mineType != tt.wantErr {
				t.Errorf("DetermineByPath() assert failed, got %v, want %v", mineType, tt.wantErr)
			}
		})
	}
}

func TestJoinPath(t *testing.T) {
	fp := path.Join("http://101.34.6.110:9099/", "/gva-alice3-24a01ad8/", "/1705282371apifox-logo-64 (1).png")
	t.Logf("path: %s", fp)
}
