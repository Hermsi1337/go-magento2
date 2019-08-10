package products

type AddProductPayload struct {
	Product     Product `json:"product"`
	SaveOptions bool    `json:"saveOptions"`
}

type MediaGalleryEntries struct {
	ID                  int                    `json:"id"`
	MediaType           string                 `json:"media_type"`
	Label               string                 `json:"label"`
	Position            int                    `json:"position"`
	Disabled            bool                   `json:"disabled"`
	Types               []string               `json:"types"`
	File                string                 `json:"file"`
	Content             Content                `json:"content"`
	ExtensionAttributes map[string]interface{} `json:"extension_attributes"`
}

type TierPrices struct {
	CustomerGroupID     int                    `json:"customer_group_id"`
	Qty                 float64                `json:"qty"`
	Value               float64                `json:"value"`
	ExtensionAttributes map[string]interface{} `json:"extension_attributes"`
}

type ProductLinks struct {
	Sku                 string                 `json:"sku"`
	LinkType            string                 `json:"link_type"`
	LinkedProductSku    string                 `json:"linked_product_sku"`
	LinkedProductType   string                 `json:"linked_product_type"`
	Position            int                    `json:"position"`
	ExtensionAttributes map[string]interface{} `json:"extension_attributes"`
}

type Product struct {
	ID                  int                      `json:"id,omitempty"`
	Sku                 string                   `json:"sku"`
	Name                string                   `json:"name"`
	AttributeSetID      int                      `json:"attribute_set_id"`
	Price               float64                  `json:"price"`
	Status              int                      `json:"status,omitempty"`
	Visibility          int                      `json:"visibility,omitempty"`
	TypeID              string                   `json:"type_id,omitempty"`
	CreatedAt           string                   `json:"created_at,omitempty"`
	UpdatedAt           string                   `json:"updated_at,omitempty"`
	Weight              int                      `json:"weight,omitempty"`
	ExtensionAttributes map[string]interface{}   `json:"extension_attributes,omitempty"`
	ProductLinks        []ProductLinks           `json:"product_links,omitempty"`
	Options             []Options                `json:"options,omitempty"`
	MediaGalleryEntries []MediaGalleryEntries    `json:"media_gallery_entries,omitempty"`
	TierPrices          []TierPrices             `json:"tier_prices,omitempty"`
	CustomAttributes    []map[string]interface{} `json:"custom_attributes,omitempty"`
}

type Content struct {
	Base64EncodedData string `json:"base64_encoded_data"`
	Type              string `json:"type"`
	Name              string `json:"name"`
}

type VideoContent struct {
	MediaType        string `json:"media_type"`
	VideoProvider    string `json:"video_provider"`
	VideoURL         string `json:"video_url"`
	VideoTitle       string `json:"video_title"`
	VideoDescription string `json:"video_description"`
	VideoMetadata    string `json:"video_metadata"`
}

type Values struct {
	Title        string  `json:"title"`
	SortOrder    int     `json:"sort_order"`
	Price        float64 `json:"price"`
	PriceType    string  `json:"price_type"`
	Sku          string  `json:"sku"`
	OptionTypeID int     `json:"option_type_id"`
}

type Options struct {
	ProductSku          string                 `json:"product_sku,omitempty"`
	OptionID            int                    `json:"option_id,omitempty"`
	Title               string                 `json:"title,omitempty"`
	Type                string                 `json:"type,omitempty"`
	SortOrder           int                    `json:"sort_order,omitempty"`
	IsRequired          bool                   `json:"is_required,omitempty"`
	Price               float64                `json:"price,omitempty"`
	PriceType           string                 `json:"price_type,omitempty"`
	Sku                 string                 `json:"sku,omitempty"`
	FileExtension       string                 `json:"file_extension,omitempty"`
	MaxCharacters       int                    `json:"max_characters,omitempty"`
	ImageSizeX          int                    `json:"image_size_x,omitempty"`
	ImageSizeY          int                    `json:"image_size_y,omitempty"`
	Values              []Values               `json:"values,omitempty"`
	ExtensionAttributes map[string]interface{} `json:"extension_attributes,omitempty"`
}
