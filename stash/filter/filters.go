package filter

import (
	"encoding/json"
	"fmt"

	"github.com/kevwan/go-stash/stash/config"
)

const (
	filterGrok         = "grok"
	filterDrop         = "drop"
	filterRemoveFields = "remove_field"
	filterTransfer     = "transfer"
	filterCopyField    = "copy_field"
	filterReplaceStr   = "replace_str"
	filterMutate       = "mutate"
	opAnd              = "and"
	opOr               = "or"
	typeContains       = "contains"
	typeMatch          = "match"
)

type FilterFunc func(map[string]interface{}) map[string]interface{}

func CreateFilters(p config.Cluster) []FilterFunc {
	var filters []FilterFunc

	for _, f := range p.Filters {
		filters_json, err := json.Marshal(f)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(string(filters_json))
		switch f.Action {
		case filterGrok:
			filters = append(filters, GrokFilter(f.Field, f.Match))
		case filterDrop:
			filters = append(filters, DropFilter(f.Conditions))
		case filterRemoveFields:
			filters = append(filters, RemoveFieldFilter(f.Fields))
		case filterTransfer:
			filters = append(filters, TransferFilter(f.Field, f.Target))
		case filterCopyField:
			filters = append(filters, CopyFiledFilter(f.Field, f.Target))
		case filterReplaceStr:
			filters = append(filters, ReplaceStrFilter(f.Gsub))
		case filterMutate:
			filters = append(filters, MutateFilter(f.Add_fields))
		}

	}

	return filters
}
