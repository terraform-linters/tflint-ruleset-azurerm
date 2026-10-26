package resourceproviders

import (
	"slices"
	"strings"
	"testing"
)

func TestCanonical(t *testing.T) {
	cases := map[string]string{
		"Microsoft.Dashboard": "Microsoft.Dashboard",
		"microsoft.dashboard": "Microsoft.Dashboard",
		"Microsoft.Insights":  "microsoft.insights",
		"Github.Network":      "GitHub.Network",
		"Microsoft.ScVmm":     "Microsoft.SCVMM",
	}
	for input, want := range cases {
		if got, ok := Canonical(input); !ok || got != want {
			t.Errorf("Canonical(%q) = %q, %t, want %q", input, got, ok, want)
		}
	}
	if _, ok := Canonical("Microsoft.Dashboardd"); ok {
		t.Error("expected Microsoft.Dashboardd to be unknown")
	}
}

// TestCasingConflicts ensures that a namespace with different casings in the generated data has an entry in
// azureCasing, so that the casing the Azure API uses is chosen deliberately.
func TestCasingConflicts(t *testing.T) {
	casings := map[string]map[string]struct{}{}
	add := func(namespaces ...string) {
		for _, namespace := range namespaces {
			key := strings.ToLower(namespace)
			if casings[key] == nil {
				casings[key] = map[string]struct{}{}
			}
			casings[key][namespace] = struct{}{}
		}
	}
	for _, namespaces := range Registrations {
		add(namespaces...)
	}
	for _, namespaces := range Resources {
		add(namespaces...)
	}
	add(RegisteredByDefault...)

	for key, variants := range casings {
		overridden := slices.ContainsFunc(azureCasing, func(namespace string) bool { return strings.EqualFold(namespace, key) })
		if len(variants) > 1 && !overridden {
			t.Errorf("namespace %q has multiple casings %v, add the casing the Azure API uses to azureCasing", key, variants)
		}
	}

	// The rule compares RegisteredByDefault case-sensitively with canonical namespaces.
	for _, namespace := range RegisteredByDefault {
		if got, _ := Canonical(namespace); got != namespace {
			t.Errorf("RegisteredByDefault contains %q, use the casing the Azure API uses: %q", namespace, got)
		}
	}
}

func TestRegistrations(t *testing.T) {
	for _, mode := range []string{"none", "core", "extended", "all", "legacy"} {
		if _, ok := Registrations[mode]; !ok {
			t.Errorf("expected registration set for mode %q", mode)
		}
	}
}

func TestResources(t *testing.T) {
	if len(Resources) == 0 {
		t.Fatal("expected generated resources")
	}
	for resourceType, namespaces := range Resources {
		if len(namespaces) == 0 {
			t.Errorf("expected namespaces for %s", resourceType)
		}
	}
}
