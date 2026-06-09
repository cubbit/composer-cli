package service

import (
	"testing"
)

func TestValidateAAG(t *testing.T) {
	tests := []struct {
		name     string
		aag      int
		minDisks int
		expected string
	}{
		{name: "valid minimum", aag: 1, minDisks: 10, expected: ""},
		{name: "valid maximum", aag: 10, minDisks: 10, expected: ""},
		{name: "valid middle", aag: 5, minDisks: 10, expected: ""},
		{name: "zero aag", aag: 0, minDisks: 10, expected: "AAG must be at least 1"},
		{name: "negative aag", aag: -1, minDisks: 10, expected: "AAG must be at least 1"},
		{name: "exceeds max disks", aag: 11, minDisks: 10, expected: "AAG cannot exceed minimum disks per node (10)"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ValidateAAG(tt.aag, tt.minDisks)
			if result != tt.expected {
				t.Errorf("ValidateAAG(%d, %d) = %q, want %q", tt.aag, tt.minDisks, result, tt.expected)
			}
		})
	}
}

func TestValidateOuterK(t *testing.T) {
	tests := []struct {
		name         string
		outerK       int
		numLocations int
		expected     string
	}{
		{name: "valid zero", outerK: 0, numLocations: 3, expected: ""},
		{name: "valid one", outerK: 1, numLocations: 3, expected: ""},
		{name: "valid max", outerK: 2, numLocations: 3, expected: ""},
		{name: "negative outerK", outerK: -1, numLocations: 3, expected: "Outer K must be at least 0"},
		{name: "equals locations", outerK: 3, numLocations: 3, expected: "Outer K must be less than number of locations (3)"},
		{name: "exceeds locations", outerK: 5, numLocations: 3, expected: "Outer K must be less than number of locations (3)"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ValidateOuterK(tt.outerK, tt.numLocations)
			if result != tt.expected {
				t.Errorf("ValidateOuterK(%d, %d) = %q, want %q", tt.outerK, tt.numLocations, result, tt.expected)
			}
		})
	}
}

func TestValidateLocalNK(t *testing.T) {
	tests := []struct {
		name                string
		localNK             int
		aag                 int
		minNodesPerLocation int
		minDisksAcrossLoc   int
		expected            string
	}{
		{name: "valid minimum", localNK: 1, aag: 1, minNodesPerLocation: 1, minDisksAcrossLoc: 10, expected: ""},
		{name: "valid divisible", localNK: 6, aag: 3, minNodesPerLocation: 2, minDisksAcrossLoc: 10, expected: ""},
		{name: "zero localNK", localNK: 0, aag: 1, minNodesPerLocation: 1, minDisksAcrossLoc: 10, expected: "Local N+K must be at least 1"},
		{name: "exceeds min disks", localNK: 15, aag: 1, minNodesPerLocation: 1, minDisksAcrossLoc: 10, expected: "Local N+K cannot exceed minimum disks across locations (10)"},
		{name: "not divisible by aag", localNK: 5, aag: 3, minNodesPerLocation: 2, minDisksAcrossLoc: 10, expected: "Local N+K must be divisible by AAG (3)"},
		{name: "exceeds aag x nodes", localNK: 10, aag: 3, minNodesPerLocation: 2, minDisksAcrossLoc: 10, expected: "Local N+K (10) exceeds AAG × min nodes per location (3 × 2 = 6)"},
		{name: "aag is zero", localNK: 5, aag: 0, minNodesPerLocation: 1, minDisksAcrossLoc: 10, expected: ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ValidateLocalNK(tt.localNK, tt.aag, tt.minNodesPerLocation, tt.minDisksAcrossLoc)
			if result != tt.expected {
				t.Errorf("ValidateLocalNK(%d, %d, %d, %d) = %q, want %q", tt.localNK, tt.aag, tt.minNodesPerLocation, tt.minDisksAcrossLoc, result, tt.expected)
			}
		})
	}
}

func TestValidateLocalK(t *testing.T) {
	tests := []struct {
		name     string
		localK   int
		localNK  int
		expected string
	}{
		{name: "valid zero", localK: 0, localNK: 5, expected: ""},
		{name: "valid max", localK: 4, localNK: 5, expected: ""},
		{name: "valid middle", localK: 2, localNK: 5, expected: ""},
		{name: "negative localK", localK: -1, localNK: 5, expected: "Local K must be at least 0"},
		{name: "equals localNK", localK: 5, localNK: 5, expected: "Local K must be less than Local N+K (5)"},
		{name: "exceeds localNK", localK: 10, localNK: 5, expected: "Local K must be less than Local N+K (5)"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ValidateLocalK(tt.localK, tt.localNK)
			if result != tt.expected {
				t.Errorf("ValidateLocalK(%d, %d) = %q, want %q", tt.localK, tt.localNK, result, tt.expected)
			}
		})
	}
}

func TestValidateRCName(t *testing.T) {
	tests := []struct {
		name          string
		rcName        string
		existingNames []string
		expected      string
	}{
		{name: "valid new name", rcName: "my-rc", existingNames: []string{}, expected: ""},
		{name: "valid among existing", rcName: "my-rc", existingNames: []string{"other-rc"}, expected: ""},
		{name: "empty name", rcName: "", existingNames: []string{}, expected: "RC name is required"},
		{name: "duplicate name", rcName: "my-rc", existingNames: []string{"my-rc"}, expected: "RC name 'my-rc' is already used"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ValidateRCName(tt.rcName, tt.existingNames)
			if result != tt.expected {
				t.Errorf("ValidateRCName(%q, %v) = %q, want %q", tt.rcName, tt.existingNames, result, tt.expected)
			}
		})
	}
}

func TestParseInt(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{input: "", expected: 0},
		{input: "0", expected: 0},
		{input: "1", expected: 1},
		{input: "42", expected: 42},
		{input: "999", expected: 999},
		{input: "abc", expected: -1},
		{input: "12a34", expected: -1},
	}
	for _, tt := range tests {
		t.Run("parseInt_"+tt.input, func(t *testing.T) {
			result := parseInt(tt.input)
			if result != tt.expected {
				t.Errorf("parseInt(%q) = %d, want %d", tt.input, result, tt.expected)
			}
		})
	}
}
