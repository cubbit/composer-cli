package service

import (
	"bytes"
	"strings"
	"testing"
	"time"

	"github.com/cubbit/composer-cli/src/api"
	"github.com/spf13/cobra"
)

func TestPrintDomainList_Human(t *testing.T) {
	cmd := &cobra.Command{}
	cmd.Flags().Bool("quiet", false, "quiet mode")
	cmd.Flags().Bool("no-headers", false, "no headers")
	cmd.Flags().Set("quiet", "false")
	cmd.Flags().Set("no-headers", "false")

	var out bytes.Buffer
	cmd.SetOut(&out)

	createdAt, err := time.Parse("2006-01-02 15:04:05", "2024-01-15 10:30:00")
	if err != nil {
		t.Fatal(err)
	}

	domains := []api.DomainDTO{
		{
			ID:         "domain-001",
			DomainName: "verified-domain.com",
			CreatedAt:  createdAt,
			Challenge:  "challenge-abc",
			VerifiedAt: &createdAt,
		},
		{
			ID:         "domain-002",
			DomainName: "unverified-domain.com",
			CreatedAt:  createdAt,
			Challenge:  "challenge-def",
		},
	}

	err = PrintDomainList(cmd, domains)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expected := strings.TrimSpace(`
╭───────────────────────┬────────────┬──────────┬─────────────────────┬───────────────╮
│ Domain Name           │ ID         │ Verified │ Created On          │ Challenge     │
├───────────────────────┼────────────┼──────────┼─────────────────────┼───────────────┤
│ verified-domain.com   │ domain-001 │ Yes      │ 2024-01-15 10:30:00 │ challenge-abc │
│ unverified-domain.com │ domain-002 │ No       │ 2024-01-15 10:30:00 │ challenge-def │
╰───────────────────────┴────────────┴──────────┴─────────────────────┴───────────────╯
`)

	if strings.TrimSpace(out.String()) != expected {
		t.Fatalf("Expected output:\n%s\n\nactual:\n%s", expected, strings.TrimSpace(out.String()))
	}
}

func TestPrintDomainList_Human_Empty(t *testing.T) {
	cmd := &cobra.Command{}
	cmd.Flags().Bool("quiet", false, "quiet mode")
	cmd.Flags().Bool("no-headers", false, "no headers")
	cmd.Flags().Set("quiet", "false")
	cmd.Flags().Set("no-headers", "false")

	var out bytes.Buffer
	cmd.SetOut(&out)

	err := PrintDomainList(cmd, []api.DomainDTO{})
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	output := out.String()
	if !strings.Contains(output, "No domains found") {
		t.Fatalf("Expected empty state message, got %q", output)
	}
}

func TestPrintDomainList_NoHeaders(t *testing.T) {
	cmd := &cobra.Command{}
	cmd.Flags().Bool("quiet", false, "quiet mode")
	cmd.Flags().Bool("no-headers", false, "no headers")
	cmd.Flags().Set("quiet", "false")
	cmd.Flags().Set("no-headers", "true")

	var out bytes.Buffer
	cmd.SetOut(&out)

	createdAt, err := time.Parse("2006-01-02 15:04:05", "2024-01-15 10:30:00")
	if err != nil {
		t.Fatal(err)
	}

	domains := []api.DomainDTO{
		{
			ID:         "domain-001",
			DomainName: "no-headers.com",
			CreatedAt:  createdAt,
			Challenge:  "challenge-noh",
		},
	}

	err = PrintDomainList(cmd, domains)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	output := out.String()
	if strings.Contains(output, "Domain Name") || strings.Contains(output, "ID") {
		t.Fatalf("Expected output without headers, got %q", output)
	}
	if !strings.Contains(output, "no-headers.com") {
		t.Fatalf("Expected output to contain domain name, got %q", output)
	}
}

func TestPrintDomainDetails_Human_Verified(t *testing.T) {
	cmd := &cobra.Command{}
	cmd.Flags().Bool("quiet", false, "quiet mode")
	cmd.Flags().Set("quiet", "false")

	var out bytes.Buffer
	cmd.SetOut(&out)

	createdAt, err := time.Parse("2006-01-02 15:04:05", "2024-01-15 10:30:00")
	if err != nil {
		t.Fatal(err)
	}
	verifiedAt, err := time.Parse("2006-01-02 15:04:05", "2024-02-20 14:00:00")
	if err != nil {
		t.Fatal(err)
	}

	domain := api.DomainDTO{
		ID:             "domain-verified",
		DomainName:     "verified-domain.com",
		CreatedAt:      createdAt,
		VerifiedAt:     &verifiedAt,
		Challenge:      "verify-challenge-token",
		OrganizationID: "org-123",
		IsShared:       true,
	}

	err = PrintDomainDetails(cmd, domain)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expected := strings.TrimSpace(`
Domain: verified-domain.com
ID: domain-verified
Organization: org-123
Challenge: verify-challenge-token
Verified: Yes
Verified At: 2024-02-20 14:00:00
Deleted: No
Shared: Yes
Created At: 2024-01-15 10:30:00
`)

	if strings.TrimSpace(out.String()) != expected {
		t.Fatalf("Expected output:\n%s\n\nactual:\n%s", expected, strings.TrimSpace(out.String()))
	}
}

func TestPrintDomainDetails_Human_NotVerified(t *testing.T) {
	cmd := &cobra.Command{}
	cmd.Flags().Bool("quiet", false, "quiet mode")
	cmd.Flags().Set("quiet", "false")

	var out bytes.Buffer
	cmd.SetOut(&out)

	createdAt, err := time.Parse("2006-01-02 15:04:05", "2024-03-10 08:00:00")
	if err != nil {
		t.Fatal(err)
	}

	domain := api.DomainDTO{
		ID:             "domain-unverified",
		DomainName:     "unverified-domain.com",
		CreatedAt:      createdAt,
		Challenge:      "unverified-challenge",
		OrganizationID: "org-456",
		IsShared:       false,
	}

	err = PrintDomainDetails(cmd, domain)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expected := strings.TrimSpace(`
Domain: unverified-domain.com
ID: domain-unverified
Organization: org-456
Challenge: unverified-challenge
Verified: No
Verified At: N/A
Deleted: No
Shared: No
Created At: 2024-03-10 08:00:00
`)

	if strings.TrimSpace(out.String()) != expected {
		t.Fatalf("Expected output:\n%s\n\nactual:\n%s", expected, strings.TrimSpace(out.String()))
	}
}

func TestPrintDomainDetails_Human_Deleted(t *testing.T) {
	cmd := &cobra.Command{}
	cmd.Flags().Bool("quiet", false, "quiet mode")
	cmd.Flags().Set("quiet", "false")

	var out bytes.Buffer
	cmd.SetOut(&out)

	createdAt, err := time.Parse("2006-01-02 15:04:05", "2024-01-15 10:30:00")
	if err != nil {
		t.Fatal(err)
	}
	deletedAt, err := time.Parse("2006-01-02 15:04:05", "2024-03-01 12:00:00")
	if err != nil {
		t.Fatal(err)
	}

	domain := api.DomainDTO{
		ID:             "domain-deleted",
		DomainName:     "deleted-domain.com",
		CreatedAt:      createdAt,
		DeletedAt:      &deletedAt,
		Challenge:      "deleted-challenge",
		OrganizationID: "org-789",
		IsShared:       false,
	}

	err = PrintDomainDetails(cmd, domain)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expected := strings.TrimSpace(`
Domain: deleted-domain.com
ID: domain-deleted
Organization: org-789
Challenge: deleted-challenge
Verified: No
Verified At: N/A
Deleted: Yes
Shared: No
Created At: 2024-01-15 10:30:00
`)

	if strings.TrimSpace(out.String()) != expected {
		t.Fatalf("Expected output:\n%s\n\nactual:\n%s", expected, strings.TrimSpace(out.String()))
	}
}
