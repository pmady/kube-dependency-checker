package version

import (
	"testing"
)

func TestParse(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    *Version
		wantErr bool
	}{
		{
			name:  "simple version",
			input: "1.30",
			want:  &Version{Major: 1, Minor: 30, Patch: 0, Raw: "1.30"},
		},
		{
			name:  "version with v prefix",
			input: "v1.30",
			want:  &Version{Major: 1, Minor: 30, Patch: 0, Raw: "1.30"},
		},
		{
			name:  "full version",
			input: "1.30.5",
			want:  &Version{Major: 1, Minor: 30, Patch: 5, Raw: "1.30.5"},
		},
		{
			name:  "full version with v prefix",
			input: "v1.30.5",
			want:  &Version{Major: 1, Minor: 30, Patch: 5, Raw: "1.30.5"},
		},
		{
			name:    "invalid version",
			input:   "invalid",
			wantErr: true,
		},
		{
			name:    "empty string",
			input:   "",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Parse(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("Parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				return
			}
			if got.Major != tt.want.Major || got.Minor != tt.want.Minor || got.Patch != tt.want.Patch {
				t.Errorf("Parse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVersion_Compare(t *testing.T) {
	tests := []struct {
		name  string
		v1    string
		v2    string
		want  int
	}{
		{"equal versions", "1.30.0", "1.30.0", 0},
		{"v1 newer major", "2.0.0", "1.30.0", 1},
		{"v1 older major", "1.0.0", "2.0.0", -1},
		{"v1 newer minor", "1.31.0", "1.30.0", 1},
		{"v1 older minor", "1.29.0", "1.30.0", -1},
		{"v1 newer patch", "1.30.1", "1.30.0", 1},
		{"v1 older patch", "1.30.0", "1.30.1", -1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v1, _ := Parse(tt.v1)
			v2, _ := Parse(tt.v2)
			if got := v1.Compare(v2); got != tt.want {
				t.Errorf("Compare() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVersion_MinorDiff(t *testing.T) {
	tests := []struct {
		name string
		v1   string
		v2   string
		want int
	}{
		{"same version", "1.30.0", "1.30.0", 0},
		{"one minor diff", "1.31.0", "1.30.0", 1},
		{"two minor diff", "1.32.0", "1.30.0", 2},
		{"three minor diff", "1.33.0", "1.30.0", 3},
		{"negative diff", "1.28.0", "1.30.0", 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v1, _ := Parse(tt.v1)
			v2, _ := Parse(tt.v2)
			if got := v1.MinorDiff(v2); got != tt.want {
				t.Errorf("MinorDiff() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVersion_IsCompatibleWithAPIServer(t *testing.T) {
	tests := []struct {
		name      string
		component string
		apiServer string
		maxSkew   int
		want      bool
	}{
		{"same version", "1.30.0", "1.30.0", 3, true},
		{"one minor older", "1.29.0", "1.30.0", 3, true},
		{"three minor older", "1.27.0", "1.30.0", 3, true},
		{"four minor older", "1.26.0", "1.30.0", 3, false},
		{"newer than api server", "1.31.0", "1.30.0", 3, false},
		{"controller manager one older", "1.29.0", "1.30.0", 1, true},
		{"controller manager two older", "1.28.0", "1.30.0", 1, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			component, _ := Parse(tt.component)
			apiServer, _ := Parse(tt.apiServer)
			if got := component.IsCompatibleWithAPIServer(apiServer, "test", tt.maxSkew); got != tt.want {
				t.Errorf("IsCompatibleWithAPIServer() = %v, want %v", got, tt.want)
			}
		})
	}
}
