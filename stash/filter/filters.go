package filter

import (
	"encoding/json"
	"fmt"

	"github.com/kevwan/go-stash/stash/config"
)

const (
	filterDrop         = "drop"
	filterRemoveFields = "remove_field"
	filterTransfer     = "transfer"
	opAnd              = "and"
	opOr               = "or"
	typeContains       = "contains"
	typeMatch          = "match"
)

type FilterFunc func(map[string]interface{}) map[string]interface{}

func CreateFilters(p config.Cluster) []FilterFunc {
	var filters []FilterFunc

	for _, f := range p.Filters {
		json_str, err := json.Marshal(f)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(string(json_str))
		switch f.Action {
		case filterDrop:
			filters = append(filters, DropFilter(f.Conditions))
		case filterRemoveFields:
			filters = append(filters, RemoveFieldFilter(f.Fields))
		case filterTransfer:
			filters = append(filters, TransferFilter(f.Field, f.Target))
		}
	}

	return filters
}
