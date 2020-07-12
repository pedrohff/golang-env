package pkg

import (
	"os"
	"testing"
)

func TestFormatFilePathForWorkingDir(t *testing.T) {
	wd, _ := os.Getwd()
	fileNameSuccess := "asdfasdf.txt"
	fileNameError := "asdfasdf"
	type args struct {
		filename string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{"success", args{filename: fileNameSuccess}, wd + fileNameSuccess, false},
		{"error", args{filename: fileNameError}, "", true},
		{"error1", args{filename: fileNameError}, "", true},
		{"error1", args{filename: fileNameError}, "", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := FormatFilePathForWorkingDir(tt.args.filename)
			if (err != nil) != tt.wantErr {
				t.Errorf("FormatFilePathForWorkingDir() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("FormatFilePathForWorkingDir() got = %v, want %v", got, tt.want)
			}
		})
	}
}
