package config

import (
	"reflect"
	"testing"
)

var mockFilePathVar string

func TestInit(t *testing.T) {
	tests := []struct {
		name     string
		wantCfg  Config
		wantErr  bool
		setup    func()
		teardown func()
	}{
		{
			name: "valid no error",
			wantCfg: Config{
				Port: PortConfig{
					Main: 111,
				},
				Worker: WorkerConfig{
					Default: 1,
				},
			},
			wantErr: false,
			setup: func() {
				mockFilePath("mock_file/mock_valid_configurations.json")
			},
			teardown: func() {
				resetFilePath()
			},
		},
		{
			name:    "error - misconfigured file content",
			wantCfg: Config{},
			wantErr: true,
			setup: func() {
				mockFilePath("mock_file/mock_fail_configurations.json")
			},
			teardown: func() {
				resetFilePath()
			},
		},
		{
			name:    "error - invalid file path",
			wantCfg: Config{},
			wantErr: true,
			setup: func() {
				mockFilePath("invalid_file_path")
			},
			teardown: func() {
				resetFilePath()
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.setup != nil {
				tt.setup()
			}
			gotCfg, err := Init()
			if (err != nil) != tt.wantErr {
				t.Errorf("Init() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotCfg, tt.wantCfg) {
				t.Errorf("Init() = %v, want %v", gotCfg, tt.wantCfg)
			}

			if tt.teardown != nil {
				tt.teardown()
			}
		})
	}
}

func mockFilePath(path string) {
	mockFilePathVar = filePath
	filePath = path
}

func resetFilePath() {
	filePath = mockFilePathVar
}
