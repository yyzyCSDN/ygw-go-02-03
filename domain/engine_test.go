package domain

import (
	"encoding/json"
	"testing"
)

func TestInventoryBookLifecycle(t *testing.T) {
	engine := NewInventoryBook()
	if err := engine.ReleaseStock(); err != nil {
		t.Fatal(err)
	}
	if err := engine.RegisterSKU(InventoryBookRecord{ID: "primary", Quantity: 4, Labels: map[string]string{"zone": "a"}}); err != nil {
		t.Fatal(err)
	}
	if err := engine.ReceiveStock("primary", 3); err != nil {
		t.Fatal(err)
	}
	if err := engine.ReserveStock("primary", 2); err != nil {
		t.Fatal(err)
	}
	if got := engine.CountAvailable(); got != 5 {
		t.Fatalf("count = %d; want 5", got)
	}
	if err := engine.ScheduleReorder(); err != nil {
		t.Fatal(err)
	}
}

func TestInventoryBookPrioritiesAndExport(t *testing.T) {
	engine := NewInventoryBook()
	_ = engine.RegisterSKU(InventoryBookRecord{ID: "low", Quantity: 1})
	_ = engine.RegisterSKU(InventoryBookRecord{ID: "high", Quantity: 2})
	if err := engine.ForecastDemand("high", 9); err != nil {
		t.Fatal(err)
	}
	values := engine.List()
	if len(values) != 2 || values[0].ID != "high" {
		t.Fatalf("unexpected order: %#v", values)
	}
	values[0].Labels = map[string]string{"changed": "yes"}
	data, err := engine.ExportStock()
	if err != nil || !json.Valid(data) {
		t.Fatalf("invalid export: %s, %v", data, err)
	}
}

func TestInventoryBookRejectsInvalidOperations(t *testing.T) {
	engine := NewInventoryBook()
	if err := engine.RegisterSKU(InventoryBookRecord{}); err == nil {
		t.Fatal("expected blank id error")
	}
	if err := engine.RegisterSKU(InventoryBookRecord{ID: "one", Quantity: 1}); err != nil {
		t.Fatal(err)
	}
	if err := engine.RegisterSKU(InventoryBookRecord{ID: "one"}); err == nil {
		t.Fatal("expected duplicate error")
	}
	if err := engine.ReserveStock("one", 2); err == nil {
		t.Fatal("expected insufficient quantity error")
	}
	if err := engine.ForecastDemand("missing", 1); err == nil {
		t.Fatal("expected missing record error")
	}
}
