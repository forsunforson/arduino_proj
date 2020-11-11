package arduino

import (
	"reflect"
	"testing"
)

func TestGetArduinos(t *testing.T) {
	tests := []struct {
		name    string
		want    []*Arduino
		wantErr bool
	}{
		{
			name: "case1:normal",
			want: []*Arduino{
				{"usb-Arduino__www.arduino.cc__0043_758343533303519021E0-if00", "/dev/ttyACM0", 9600, 0},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetArduinos()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetArduinos() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetArduinos() = %v, want %v", got, tt.want)
			}
		})
	}
}
