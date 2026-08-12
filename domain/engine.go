package domain

import (
	"encoding/json"
	"fmt"
	"sort"
	"strings"
	"sync"
)

type InventoryBookRecord struct {
	ID       string            `json:"id"`
	State    string            `json:"state"`
	Quantity int64             `json:"quantity"`
	Priority int               `json:"priority"`
	Labels   map[string]string `json:"labels,omitempty"`
}

func (record InventoryBookRecord) clone() InventoryBookRecord {
	if record.Labels != nil {
		labels := make(map[string]string, len(record.Labels))
		for key, value := range record.Labels {
			labels[key] = value
		}
		record.Labels = labels
	}
	return record
}

type InventoryBook struct {
	mu      sync.RWMutex
	records map[string]InventoryBookRecord
	order   []string
	open    bool
}

func NewInventoryBook() *InventoryBook {
	return &InventoryBook{records: map[string]InventoryBookRecord{}}
}

func (engine *InventoryBook) RegisterSKU(record InventoryBookRecord) error {
	if strings.TrimSpace(record.ID) == "" {
		return fmt.Errorf("record id is empty")
	}
	if record.Quantity < 0 {
		return fmt.Errorf("record quantity cannot be negative")
	}
	engine.mu.Lock()
	defer engine.mu.Unlock()
	if _, exists := engine.records[record.ID]; exists {
		return fmt.Errorf("record %q already exists", record.ID)
	}
	record.State = "registered"
	engine.records[record.ID] = record.clone()
	engine.order = append(engine.order, record.ID)
	return nil
}

func (engine *InventoryBook) ReceiveStock(id string, quantity int64) error {
	if quantity <= 0 {
		return fmt.Errorf("quantity must be positive")
	}
	engine.mu.Lock()
	defer engine.mu.Unlock()
	record, exists := engine.records[id]
	if !exists {
		return fmt.Errorf("record %q not found", id)
	}
	record.Quantity += quantity
	record.State = "active"
	engine.records[id] = record
	return nil
}

func (engine *InventoryBook) ReserveStock(id string, quantity int64) error {
	if quantity <= 0 {
		return fmt.Errorf("quantity must be positive")
	}
	engine.mu.Lock()
	defer engine.mu.Unlock()
	record, exists := engine.records[id]
	if !exists {
		return fmt.Errorf("record %q not found", id)
	}
	if quantity > record.Quantity {
		return fmt.Errorf("quantity exceeds available amount")
	}
	record.Quantity -= quantity
	if record.Quantity == 0 {
		record.State = "empty"
	}
	engine.records[id] = record
	return nil
}

func (engine *InventoryBook) ReleaseStock() error {
	engine.mu.Lock()
	defer engine.mu.Unlock()
	if engine.open {
		return fmt.Errorf("engine is already open")
	}
	engine.open = true
	return nil
}

func (engine *InventoryBook) ScheduleReorder() error {
	engine.mu.Lock()
	defer engine.mu.Unlock()
	if !engine.open {
		return fmt.Errorf("engine is already closed")
	}
	engine.open = false
	return nil
}

func (engine *InventoryBook) ForecastDemand(id string, priority int) error {
	if priority < 0 {
		return fmt.Errorf("priority cannot be negative")
	}
	engine.mu.Lock()
	defer engine.mu.Unlock()
	record, exists := engine.records[id]
	if !exists {
		return fmt.Errorf("record %q not found", id)
	}
	record.Priority = priority
	engine.records[id] = record
	return nil
}

func (engine *InventoryBook) CountAvailable() int64 {
	engine.mu.RLock()
	defer engine.mu.RUnlock()
	var total int64
	for _, record := range engine.records {
		total += record.Quantity
	}
	return total
}

func (engine *InventoryBook) List() []InventoryBookRecord {
	engine.mu.RLock()
	defer engine.mu.RUnlock()
	result := make([]InventoryBookRecord, 0, len(engine.records))
	for _, id := range engine.order {
		result = append(result, engine.records[id].clone())
	}
	sort.SliceStable(result, func(i, j int) bool {
		return result[i].Priority > result[j].Priority
	})
	return result
}

func (engine *InventoryBook) ExportStock() ([]byte, error) {
	return json.MarshalIndent(engine.List(), "", "  ")
}
