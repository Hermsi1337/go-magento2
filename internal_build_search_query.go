package magento2

import (
	"fmt"
	"net/url"
)

const (
	searchCriteriaFieldTemplate         = "searchCriteria[filter_groups][%d][filters][%d][field]="
	searchCriteriaValueTemplate         = "searchCriteria[filter_groups][%d][filters][%d][value]="
	searchCriteriaConditionTypeTemplate = "searchCriteria[filter_groups][%d][filters][%d][condition_type]="
)

type SearchQueryCriteria struct {
	Fields []FilterFields
}

type FilterFields struct {
	Field         Filter
	Value         Filter
	ConditionType Filter
}

type Filter struct {
	FilterGroups int
	Filters      int
	FilterFor    string
}

type Fields struct {
	Key   string
	Value string
}

func searchCriteriaField(filterGroups, filters int) string {
	return fmt.Sprintf(searchCriteriaFieldTemplate, filterGroups, filters)
}

func searchCriteriaValue(filterGroups, filters int) string {
	return fmt.Sprintf(searchCriteriaValueTemplate, filterGroups, filters)
}

func searchCriteriaConditionType(filterGroups, filters int) string {
	return fmt.Sprintf(searchCriteriaConditionTypeTemplate, filterGroups, filters)
}

// use this method for building simple search queries without any flexibility
func BuildSearchQuery(field, value, conditionType string) string {
	params := url.Values{}
	params.Add(searchCriteriaField(0, 0), field)
	params.Add(searchCriteriaValue(0, 0), value)
	params.Add(searchCriteriaConditionType(0, 0), conditionType)

	return params.Encode()
}

// this method is used to build very flexible search-queries
// for example:
// ?searchCriteria[filter_groups][2][filters][0][field]=increment_id
// &searchCriteria[filter_groups][2][filters][0][value]=XXXXX
// &searchCriteria[filter_groups][2][filters][0][condition_type]=eq
// &fields=items[entity_id]
func BuildFlexibleSearchQuery(criteria []SearchQueryCriteria, additionalQuery ...Fields) string {
	params := url.Values{}
	for i := range criteria {
		for y := range criteria[i].Fields {
			params.Add(
				searchCriteriaField(
					criteria[i].Fields[y].Field.FilterGroups,
					criteria[i].Fields[y].Field.Filters,
				),
				criteria[i].Fields[y].Field.FilterFor,
			)
			params.Add(
				searchCriteriaField(
					criteria[i].Fields[y].Value.FilterGroups,
					criteria[i].Fields[y].Value.Filters,
				),
				criteria[i].Fields[y].Value.FilterFor,
			)
			params.Add(
				searchCriteriaField(
					criteria[i].Fields[y].ConditionType.FilterGroups,
					criteria[i].Fields[y].ConditionType.Filters,
				),
				criteria[i].Fields[y].ConditionType.FilterFor,
			)
		}
	}

	for i := range additionalQuery {
		params.Add(additionalQuery[i].Key, additionalQuery[i].Value)
	}

	return params.Encode()
}
