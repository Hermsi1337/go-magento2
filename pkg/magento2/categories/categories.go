package categories

import (
	"fmt"
	"github.com/hermsi1337/go-magento2/internal/utils"
	"github.com/hermsi1337/go-magento2/pkg/magento2"
	"github.com/hermsi1337/go-magento2/pkg/magento2/api"
)

type MCategory struct {
	Route     string
	Category  *Category
	ApiClient *api.Client
}

func CreateCategory(c *Category, apiClient *api.Client) (*MCategory, error) {
	mC := &MCategory{
		Category:  &Category{},
		ApiClient: apiClient,
	}
	endpoint := categories
	httpClient := apiClient.HttpClient

	payLoad := createCategoryPayload{
		Category: *c,
	}

	resp, err := httpClient.R().SetBody(payLoad).SetResult(mC.Category).Post(endpoint)
	mC.Route = fmt.Sprintf("%s/%d", categories, mC.Category.ID)

	err = utils.MayReturnErrorForHTTPResponse(err, resp, "create category")
	return mC, err
}

func GetCategoryByName(name string, apiClient *api.Client) (*MCategory, error) {
	mC := &MCategory{
		Category:  &Category{},
		ApiClient: apiClient,
	}
	searchQuery := utils.BuildSearchQuery("name", name, "in")
	endpoint := categoriesList + "?" + searchQuery
	httpClient := apiClient.HttpClient

	response := &categorySearchQueryResponse{}

	resp, err := httpClient.R().SetResult(response).Get(endpoint)
	err = utils.MayReturnErrorForHTTPResponse(err, resp, "get category by name from remote")
	if err != nil {
		return nil, err
	}

	if len(response.Categories) <= 0 {
		return nil, magento2.ErrNotFound
	}

	mC.Category = &response.Categories[0]
	mC.Route = fmt.Sprintf("%s/%d", categories, mC.Category.ID)

	err = utils.MayReturnErrorForHTTPResponse(mC.UpdateCategoryFromRemote(), resp, "get detailed category by name from remote")

	return mC, err
}

func (mC *MCategory) UpdateCategoryFromRemote() error {
	resp, err := mC.ApiClient.HttpClient.R().SetResult(mC.Category).Get(mC.Route)
	return utils.MayReturnErrorForHTTPResponse(err, resp, "get category from remote")
}

func (mC *MCategory) AssignProductByProductLink(pl *ProductLink) error {
	if len(pl.CategoryID) <= 0 {
		pl.CategoryID = fmt.Sprintf("%d", mC.Category.ID)
	}

	httpClient := mC.ApiClient.HttpClient
	endpoint := fmt.Sprintf("%s/%s", mC.Route, categoriesProductsRelative)

	resp, err := httpClient.R().SetBody(pl).Put(endpoint)
	return utils.MayReturnErrorForHTTPResponse(err, resp, "assign product to category")
}
