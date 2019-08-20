package utils

import (
	"net/url"
)

const (
	searchCriteriaField         = "searchCriteria[filter_groups][0][filters][0][field]="
	searchCriteriaValue         = "searchCriteria[filter_groups][0][filters][0][value]="
	searchCriteriaConditionType = "searchCriteria[filter_groups][0][filters][0][condition_type]="
)

func BuildSearchQuery(field, value, conditionType string) string {
	params := url.Values{}
	params.Add(searchCriteriaField, field)
	params.Add(searchCriteriaValue, value)
	params.Add(searchCriteriaConditionType, conditionType)

	return params.Encode()
}
