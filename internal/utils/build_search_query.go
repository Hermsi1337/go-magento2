package utils

import (
	"fmt"
)

const (
	searchCriteriaField         = "searchCriteria[filter_groups][0][filters][0][field]="
	searchCriteriaValue         = "searchCriteria[filter_groups][0][filters][0][value]="
	searchCriteriaConditionType = "searchCriteria[filter_groups][0][filters][0][condition_type]="
)

func BuildSearchQuery(field, value, conditionType string) string {
	return fmt.Sprintf(searchCriteriaField + field + "&" + searchCriteriaValue + value + "&" + searchCriteriaConditionType + conditionType)
}
