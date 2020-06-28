package magento2

import (
	"fmt"
)

type MCategory struct {
	Route     string
	Category  *Category
	Products  *[]ProductLink
	APIClient *Client
}

func CreateCategory(c *Category, apiClient *Client) (*MCategory, error) {
	mC := &MCategory{
		Category:  &Category{},
		Products:  &[]ProductLink{},
		APIClient: apiClient,
	}
	endpoint := categories
	httpClient := apiClient.HTTPClient

	payLoad := createCategoryPayload{
		Category: *c,
	}

	resp, err := httpClient.R().SetBody(payLoad).SetResult(mC.Category).Post(endpoint)
	mC.Route = fmt.Sprintf("%s/%d", categories, mC.Category.ID)

	err = mayReturnErrorForHTTPResponse(err, resp, "create category")
	return mC, err
}

func GetCategoryByName(name string, apiClient *Client) (*MCategory, error) {
	mC := &MCategory{
		Category:  &Category{},
		Products:  &[]ProductLink{},
		APIClient: apiClient,
	}
	searchQuery := BuildSearchQuery("name", name, "in")
	endpoint := categoriesList + "?" + searchQuery
	httpClient := apiClient.HTTPClient

	response := &categorySearchQueryResponse{}

	resp, err := httpClient.R().SetResult(response).Get(endpoint)
	err = mayReturnErrorForHTTPResponse(err, resp, "get category by name from remote")
	if err != nil {
		return nil, err
	}

	if len(response.Categories) == 0 {
		return nil, ErrNotFound
	}

	mC.Category = &response.Categories[0]
	mC.Route = fmt.Sprintf("%s/%d", categories, mC.Category.ID)

	err = mayReturnErrorForHTTPResponse(mC.UpdateCategoryFromRemote(), resp, "get detailed category by name from remote")

	return mC, err
}

func (mC *MCategory) UpdateCategoryFromRemote() error {
	resp, err := mC.APIClient.HTTPClient.R().SetResult(mC.Category).Get(mC.Route)
	err = mayReturnErrorForHTTPResponse(err, resp, "get category from remote")
	if err != nil {
		return err
	}

	return mC.UpdateCategoryProductsFromRemote()
}

func (mC *MCategory) UpdateCategoryProductsFromRemote() error {
	resp, err := mC.APIClient.HTTPClient.R().SetResult(mC.Products).Get(fmt.Sprintf("%s/%s", mC.Route, categoriesProductsRelative))
	return mayReturnErrorForHTTPResponse(err, resp, "get category products from remote")
}

func (mC *MCategory) AssignProductByProductLink(pl *ProductLink) error {
	if pl.CategoryID == "" {
		pl.CategoryID = fmt.Sprintf("%d", mC.Category.ID)
	}

	httpClient := mC.APIClient.HTTPClient
	endpoint := fmt.Sprintf("%s/%s", mC.Route, categoriesProductsRelative)

	payLoad := assignProductPayload{ProductLink: *pl}

	resp, err := httpClient.R().SetBody(payLoad).Put(endpoint)
	err = mayReturnErrorForHTTPResponse(err, resp, "assign product to category")

	if err == nil {
		*mC.Products = append(*mC.Products, *pl)
	}

	return err
}
