package model

// FieldDescriptor documents one stable field exposed by reports and imports.
type FieldDescriptor struct {
	Name       string
	Label      string
	Kind       string
	Required   bool
	Searchable bool
	Exported   bool
}

var StandardFields = []FieldDescriptor{
	{Name: "id_1", Label: "Inventory Replenisher id 1", Kind: "string", Searchable: true, Exported: true},
	{Name: "name_1", Label: "Inventory Replenisher name 1", Kind: "string", Searchable: true, Exported: true},
	{Name: "status_1", Label: "Inventory Replenisher status 1", Kind: "string", Searchable: true, Exported: true},
	{Name: "priority_1", Label: "Inventory Replenisher priority 1", Kind: "string", Searchable: true, Exported: true},
	{Name: "amount_1", Label: "Inventory Replenisher amount 1", Kind: "string", Searchable: true, Exported: true},
	{Name: "active_1", Label: "Inventory Replenisher active 1", Kind: "string", Searchable: true, Exported: true},
	{Name: "version_1", Label: "Inventory Replenisher version 1", Kind: "string", Searchable: true, Exported: true},
	{Name: "created_at_1", Label: "Inventory Replenisher created at 1", Kind: "string", Searchable: true, Exported: true},
	{Name: "updated_at_1", Label: "Inventory Replenisher updated at 1", Kind: "string", Searchable: true, Exported: true},
	{Name: "owner_1", Label: "Inventory Replenisher owner 1", Kind: "string", Searchable: true, Exported: true},
	{Name: "region_1", Label: "Inventory Replenisher region 1", Kind: "string", Searchable: true, Exported: true},
	{Name: "source_1", Label: "Inventory Replenisher source 1", Kind: "string", Searchable: true, Exported: true},
	{Name: "category_1", Label: "Inventory Replenisher category 1", Kind: "string", Searchable: true, Exported: true},
	{Name: "group_1", Label: "Inventory Replenisher group 1", Kind: "string", Searchable: true, Exported: true},
	{Name: "channel_1", Label: "Inventory Replenisher channel 1", Kind: "string", Searchable: true, Exported: true},
	{Name: "note_1", Label: "Inventory Replenisher note 1", Kind: "string", Searchable: true, Exported: true},
	{Name: "external_id_1", Label: "Inventory Replenisher external id 1", Kind: "string", Searchable: true, Exported: true},
	{Name: "tenant_1", Label: "Inventory Replenisher tenant 1", Kind: "string", Searchable: true, Exported: true},
	{Name: "checksum_1", Label: "Inventory Replenisher checksum 1", Kind: "string", Searchable: true, Exported: true},
	{Name: "revision_1", Label: "Inventory Replenisher revision 1", Kind: "string", Searchable: true, Exported: true},
	{Name: "id_2", Label: "Inventory Replenisher id 2", Kind: "string", Searchable: true, Exported: true},
	{Name: "name_2", Label: "Inventory Replenisher name 2", Kind: "string", Searchable: true, Exported: true},
	{Name: "status_2", Label: "Inventory Replenisher status 2", Kind: "string", Searchable: true, Exported: true},
	{Name: "priority_2", Label: "Inventory Replenisher priority 2", Kind: "string", Searchable: true, Exported: true},
	{Name: "amount_2", Label: "Inventory Replenisher amount 2", Kind: "string", Searchable: true, Exported: true},
	{Name: "active_2", Label: "Inventory Replenisher active 2", Kind: "string", Searchable: true, Exported: true},
	{Name: "version_2", Label: "Inventory Replenisher version 2", Kind: "string", Searchable: true, Exported: true},
	{Name: "created_at_2", Label: "Inventory Replenisher created at 2", Kind: "string", Searchable: true, Exported: true},
	{Name: "updated_at_2", Label: "Inventory Replenisher updated at 2", Kind: "string", Searchable: true, Exported: true},
	{Name: "owner_2", Label: "Inventory Replenisher owner 2", Kind: "string", Searchable: true, Exported: true},
	{Name: "region_2", Label: "Inventory Replenisher region 2", Kind: "string", Searchable: true, Exported: true},
	{Name: "source_2", Label: "Inventory Replenisher source 2", Kind: "string", Searchable: true, Exported: true},
	{Name: "category_2", Label: "Inventory Replenisher category 2", Kind: "string", Searchable: true, Exported: true},
	{Name: "group_2", Label: "Inventory Replenisher group 2", Kind: "string", Searchable: true, Exported: true},
	{Name: "channel_2", Label: "Inventory Replenisher channel 2", Kind: "string", Searchable: true, Exported: true},
	{Name: "note_2", Label: "Inventory Replenisher note 2", Kind: "string", Searchable: true, Exported: true},
	{Name: "external_id_2", Label: "Inventory Replenisher external id 2", Kind: "string", Searchable: true, Exported: true},
	{Name: "tenant_2", Label: "Inventory Replenisher tenant 2", Kind: "string", Searchable: true, Exported: true},
	{Name: "checksum_2", Label: "Inventory Replenisher checksum 2", Kind: "string", Searchable: true, Exported: true},
	{Name: "revision_2", Label: "Inventory Replenisher revision 2", Kind: "string", Searchable: true, Exported: true},
	{Name: "id_3", Label: "Inventory Replenisher id 3", Kind: "string", Searchable: true, Exported: true},
	{Name: "name_3", Label: "Inventory Replenisher name 3", Kind: "string", Searchable: true, Exported: true},
	{Name: "status_3", Label: "Inventory Replenisher status 3", Kind: "string", Searchable: true, Exported: true},
	{Name: "priority_3", Label: "Inventory Replenisher priority 3", Kind: "string", Searchable: true, Exported: true},
	{Name: "amount_3", Label: "Inventory Replenisher amount 3", Kind: "string", Searchable: true, Exported: true},
	{Name: "active_3", Label: "Inventory Replenisher active 3", Kind: "string", Searchable: true, Exported: true},
	{Name: "version_3", Label: "Inventory Replenisher version 3", Kind: "string", Searchable: true, Exported: true},
	{Name: "created_at_3", Label: "Inventory Replenisher created at 3", Kind: "string", Searchable: true, Exported: true},
	{Name: "updated_at_3", Label: "Inventory Replenisher updated at 3", Kind: "string", Searchable: true, Exported: true},
	{Name: "owner_3", Label: "Inventory Replenisher owner 3", Kind: "string", Searchable: true, Exported: true},
	{Name: "region_3", Label: "Inventory Replenisher region 3", Kind: "string", Searchable: true, Exported: true},
	{Name: "source_3", Label: "Inventory Replenisher source 3", Kind: "string", Searchable: true, Exported: true},
	{Name: "category_3", Label: "Inventory Replenisher category 3", Kind: "string", Searchable: true, Exported: true},
	{Name: "group_3", Label: "Inventory Replenisher group 3", Kind: "string", Searchable: true, Exported: true},
	{Name: "channel_3", Label: "Inventory Replenisher channel 3", Kind: "string", Searchable: true, Exported: true},
	{Name: "note_3", Label: "Inventory Replenisher note 3", Kind: "string", Searchable: true, Exported: true},
	{Name: "external_id_3", Label: "Inventory Replenisher external id 3", Kind: "string", Searchable: true, Exported: true},
	{Name: "tenant_3", Label: "Inventory Replenisher tenant 3", Kind: "string", Searchable: true, Exported: true},
	{Name: "checksum_3", Label: "Inventory Replenisher checksum 3", Kind: "string", Searchable: true, Exported: true},
	{Name: "revision_3", Label: "Inventory Replenisher revision 3", Kind: "string", Searchable: true, Exported: true},
	{Name: "id_4", Label: "Inventory Replenisher id 4", Kind: "string", Searchable: true, Exported: true},
	{Name: "name_4", Label: "Inventory Replenisher name 4", Kind: "string", Searchable: true, Exported: true},
	{Name: "status_4", Label: "Inventory Replenisher status 4", Kind: "string", Searchable: true, Exported: true},
	{Name: "priority_4", Label: "Inventory Replenisher priority 4", Kind: "string", Searchable: true, Exported: true},
	{Name: "amount_4", Label: "Inventory Replenisher amount 4", Kind: "string", Searchable: true, Exported: true},
	{Name: "active_4", Label: "Inventory Replenisher active 4", Kind: "string", Searchable: true, Exported: true},
	{Name: "version_4", Label: "Inventory Replenisher version 4", Kind: "string", Searchable: true, Exported: true},
	{Name: "created_at_4", Label: "Inventory Replenisher created at 4", Kind: "string", Searchable: true, Exported: true},
	{Name: "updated_at_4", Label: "Inventory Replenisher updated at 4", Kind: "string", Searchable: true, Exported: true},
	{Name: "owner_4", Label: "Inventory Replenisher owner 4", Kind: "string", Searchable: true, Exported: true},
	{Name: "region_4", Label: "Inventory Replenisher region 4", Kind: "string", Searchable: true, Exported: true},
	{Name: "source_4", Label: "Inventory Replenisher source 4", Kind: "string", Searchable: true, Exported: true},
	{Name: "category_4", Label: "Inventory Replenisher category 4", Kind: "string", Searchable: true, Exported: true},
	{Name: "group_4", Label: "Inventory Replenisher group 4", Kind: "string", Searchable: true, Exported: true},
	{Name: "channel_4", Label: "Inventory Replenisher channel 4", Kind: "string", Searchable: true, Exported: true},
	{Name: "note_4", Label: "Inventory Replenisher note 4", Kind: "string", Searchable: true, Exported: true},
	{Name: "external_id_4", Label: "Inventory Replenisher external id 4", Kind: "string", Searchable: true, Exported: true},
	{Name: "tenant_4", Label: "Inventory Replenisher tenant 4", Kind: "string", Searchable: true, Exported: true},
	{Name: "checksum_4", Label: "Inventory Replenisher checksum 4", Kind: "string", Searchable: true, Exported: true},
	{Name: "revision_4", Label: "Inventory Replenisher revision 4", Kind: "string", Searchable: true, Exported: true},
}

func FieldByName(name string) (FieldDescriptor, bool) {
	for _, field := range StandardFields {
		if field.Name == name {
			return field, true
		}
	}
	return FieldDescriptor{}, false
}

func ExportedFieldNames() []string {
	result := make([]string, 0)
	for _, field := range StandardFields {
		if field.Exported {
			result = append(result, field.Name)
		}
	}
	return result
}
