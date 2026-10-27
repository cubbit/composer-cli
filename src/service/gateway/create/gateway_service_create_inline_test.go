package create

import (
	"testing"

	"github.com/cubbit/composer-cli/src/api"
)

func equalSwarmRC(a, b api.SwarmAndRedundancyClassV5) bool {
	return a.SwarmID == b.SwarmID &&
		a.RedundancyClassID == b.RedundancyClassID &&
		a.IsDefault == b.IsDefault
}

func equalSwarmRCSlices(a, b []api.SwarmAndRedundancyClassV5) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if !equalSwarmRC(a[i], b[i]) {
			return false
		}
	}
	return true
}

func TestParseSwarmRCFlags_EmptyInput(t *testing.T) {
	_, err := parseSwarmRCFlags([]string{})
	if err == nil {
		t.Fatal("expected error for empty input, got nil")
	}
	if err.Error() != "at least one --swarm-rc flag is required" {
		t.Fatalf("unexpected error message: %q", err.Error())
	}
}

func TestParseSwarmRCFlags_SingleValid(t *testing.T) {
	result, err := parseSwarmRCFlags([]string{"swarm-1:rc-1:true"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expected := []api.SwarmAndRedundancyClassV5{
		{SwarmID: "swarm-1", RedundancyClassID: "rc-1", IsDefault: true},
	}
	if !equalSwarmRCSlices(result, expected) {
		t.Fatalf("expected %+v, got %+v", expected, result)
	}
}

func TestParseSwarmRCFlags_MultipleValid(t *testing.T) {
	result, err := parseSwarmRCFlags([]string{
		"swarm-1:rc-1:true",
		"swarm-2:rc-2:false",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expected := []api.SwarmAndRedundancyClassV5{
		{SwarmID: "swarm-1", RedundancyClassID: "rc-1", IsDefault: true},
		{SwarmID: "swarm-2", RedundancyClassID: "rc-2", IsDefault: false},
	}
	if !equalSwarmRCSlices(result, expected) {
		t.Fatalf("expected %+v, got %+v", expected, result)
	}
}

func TestParseSwarmRCFlags_DefaultInNonFirstPosition(t *testing.T) {
	result, err := parseSwarmRCFlags([]string{
		"swarm-a:rc-a:false",
		"swarm-b:rc-b:true",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if result[0].IsDefault {
		t.Fatal("expected first entry to not be default")
	}
	if !result[1].IsDefault {
		t.Fatal("expected second entry to be default")
	}
}

func TestParseSwarmRCFlags_WhitespaceTrimmed(t *testing.T) {
	result, err := parseSwarmRCFlags([]string{"  swarm-1  :  rc-1  :  true  "})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if result[0].SwarmID != "swarm-1" {
		t.Fatalf("expected SwarmID 'swarm-1', got %q", result[0].SwarmID)
	}
	if result[0].RedundancyClassID != "rc-1" {
		t.Fatalf("expected RedundancyClassID 'rc-1', got %q", result[0].RedundancyClassID)
	}
	if !result[0].IsDefault {
		t.Fatal("expected IsDefault to be true")
	}
}

func TestParseSwarmRCFlags_InvalidFormat(t *testing.T) {
	_, err := parseSwarmRCFlags([]string{"swarm-1:rc-1"})
	if err == nil {
		t.Fatal("expected error for invalid format, got nil")
	}
}

func TestParseSwarmRCFlags_InvalidFormatExtraParts(t *testing.T) {
	_, err := parseSwarmRCFlags([]string{"a:b:c:d"})
	if err == nil {
		t.Fatal("expected error for extra colon parts, got nil")
	}
}

func TestParseSwarmRCFlags_EmptySwarmID(t *testing.T) {
	_, err := parseSwarmRCFlags([]string{":rc-1:true"})
	if err == nil {
		t.Fatal("expected error for empty swarm ID, got nil")
	}
	if err.Error() != "swarm ID cannot be empty at position 0" {
		t.Fatalf("unexpected error message: %q", err.Error())
	}
}

func TestParseSwarmRCFlags_EmptyRCID(t *testing.T) {
	_, err := parseSwarmRCFlags([]string{"swarm-1::true"})
	if err == nil {
		t.Fatal("expected error for empty RC ID, got nil")
	}
	if err.Error() != "redundancy class ID cannot be empty at position 0" {
		t.Fatalf("unexpected error message: %q", err.Error())
	}
}

func TestParseSwarmRCFlags_MultipleDefaults(t *testing.T) {
	_, err := parseSwarmRCFlags([]string{
		"swarm-1:rc-1:true",
		"swarm-2:rc-2:true",
	})
	if err == nil {
		t.Fatal("expected error for multiple defaults, got nil")
	}
	if err.Error() != "only one swarm-rc pair can be set as default, found multiple at positions" {
		t.Fatalf("unexpected error message: %q", err.Error())
	}
}

func TestParseSwarmRCFlags_NoDefault(t *testing.T) {
	_, err := parseSwarmRCFlags([]string{
		"swarm-1:rc-1:false",
		"swarm-2:rc-2:false",
	})
	if err == nil {
		t.Fatal("expected error when no default is set, got nil")
	}
	if err.Error() != "exactly one --swarm-rc must be marked as default (set is-default to 'true')" {
		t.Fatalf("unexpected error message: %q", err.Error())
	}
}

func TestParseSwarmRCFlags_IsDefaultIsCaseSensitive(t *testing.T) {
	_, err := parseSwarmRCFlags([]string{"swarm-1:rc-1:True"})
	if err == nil {
		t.Fatal("expected error when 'True' (capital T) is used instead of 'true', got nil")
	}
	if err.Error() != "exactly one --swarm-rc must be marked as default (set is-default to 'true')" {
		t.Fatalf("unexpected error message: %q", err.Error())
	}
}
