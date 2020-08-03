package magento2

type Order struct {
	AdjustmentNegative                      float64         `json:"adjustment_negative,omitempty"`
	AdjustmentPositive                      float64         `json:"adjustment_positive,omitempty"`
	AppliedRuleIds                          string          `json:"applied_rule_ids,omitempty"`
	BaseAdjustmentNegative                  float64         `json:"base_adjustment_negative,omitempty"`
	BaseAdjustmentPositive                  float64         `json:"base_adjustment_positive,omitempty"`
	BaseCurrencyCode                        string          `json:"base_currency_code,omitempty"`
	BaseDiscountAmount                      float64         `json:"base_discount_amount,omitempty"`
	BaseDiscountCanceled                    float64         `json:"base_discount_canceled,omitempty"`
	BaseDiscountInvoiced                    float64         `json:"base_discount_invoiced,omitempty"`
	BaseDiscountRefunded                    float64         `json:"base_discount_refunded,omitempty"`
	BaseGrandTotal                          float64         `json:"base_grand_total,omitempty"`
	BaseDiscountTaxCompensationAmount       float64         `json:"base_discount_tax_compensation_amount,omitempty"`
	BaseDiscountTaxCompensationInvoiced     float64         `json:"base_discount_tax_compensation_invoiced,omitempty"`
	BaseDiscountTaxCompensationRefunded     float64         `json:"base_discount_tax_compensation_refunded,omitempty"`
	BaseShippingAmount                      float64         `json:"base_shipping_amount,omitempty"`
	BaseShippingCanceled                    float64         `json:"base_shipping_canceled,omitempty"`
	BaseShippingDiscountAmount              float64         `json:"base_shipping_discount_amount,omitempty"`
	BaseShippingDiscountTaxCompensationAmnt float64         `json:"base_shipping_discount_tax_compensation_amnt,omitempty"`
	BaseShippingInclTax                     float64         `json:"base_shipping_incl_tax,omitempty"`
	BaseShippingInvoiced                    float64         `json:"base_shipping_invoiced,omitempty"`
	BaseShippingRefunded                    float64         `json:"base_shipping_refunded,omitempty"`
	BaseShippingTaxAmount                   float64         `json:"base_shipping_tax_amount,omitempty"`
	BaseShippingTaxRefunded                 float64         `json:"base_shipping_tax_refunded,omitempty"`
	BaseSubtotal                            float64         `json:"base_subtotal,omitempty"`
	BaseSubtotalCanceled                    float64         `json:"base_subtotal_canceled,omitempty"`
	BaseSubtotalInclTax                     float64         `json:"base_subtotal_incl_tax,omitempty"`
	BaseSubtotalInvoiced                    float64         `json:"base_subtotal_invoiced,omitempty"`
	BaseSubtotalRefunded                    float64         `json:"base_subtotal_refunded,omitempty"`
	BaseTaxAmount                           float64         `json:"base_tax_amount,omitempty"`
	BaseTaxCanceled                         float64         `json:"base_tax_canceled,omitempty"`
	BaseTaxInvoiced                         float64         `json:"base_tax_invoiced,omitempty"`
	BaseTaxRefunded                         float64         `json:"base_tax_refunded,omitempty"`
	BaseTotalCanceled                       float64         `json:"base_total_canceled,omitempty"`
	BaseTotalDue                            float64         `json:"base_total_due,omitempty"`
	BaseTotalInvoiced                       float64         `json:"base_total_invoiced,omitempty"`
	BaseTotalInvoicedCost                   float64         `json:"base_total_invoiced_cost,omitempty"`
	BaseTotalOfflineRefunded                float64         `json:"base_total_offline_refunded,omitempty"`
	BaseTotalOnlineRefunded                 float64         `json:"base_total_online_refunded,omitempty"`
	BaseTotalPaid                           float64         `json:"base_total_paid,omitempty"`
	BaseTotalQtyOrdered                     float64         `json:"base_total_qty_ordered,omitempty"`
	BaseTotalRefunded                       float64         `json:"base_total_refunded,omitempty"`
	BaseToGlobalRate                        float64         `json:"base_to_global_rate,omitempty"`
	BaseToOrderRate                         float64         `json:"base_to_order_rate,omitempty"`
	BillingAddressID                        float64         `json:"billing_address_id,omitempty"`
	CanShipPartially                        float64         `json:"can_ship_partially,omitempty"`
	CanShipPartiallyItem                    float64         `json:"can_ship_partially_item,omitempty"`
	CouponCode                              string          `json:"coupon_code,omitempty"`
	CreatedAt                               string          `json:"created_at,omitempty"`
	CustomerDob                             string          `json:"customer_dob,omitempty"`
	CustomerEmail                           string          `json:"customer_email,omitempty"`
	CustomerFirstname                       string          `json:"customer_firstname,omitempty"`
	CustomerGender                          float64         `json:"customer_gender,omitempty"`
	CustomerGroupID                         float64         `json:"customer_group_id,omitempty"`
	CustomerID                              float64         `json:"customer_id,omitempty"`
	CustomerIsGuest                         float64         `json:"customer_is_guest,omitempty"`
	CustomerLastname                        string          `json:"customer_lastname,omitempty"`
	CustomerMiddlename                      string          `json:"customer_middlename,omitempty"`
	CustomerNote                            string          `json:"customer_note,omitempty"`
	CustomerNoteNotify                      float64         `json:"customer_note_notify,omitempty"`
	CustomerPrefix                          string          `json:"customer_prefix,omitempty"`
	CustomerSuffix                          string          `json:"customer_suffix,omitempty"`
	CustomerTaxvat                          string          `json:"customer_taxvat,omitempty"`
	DiscountAmount                          float64         `json:"discount_amount,omitempty"`
	DiscountCanceled                        float64         `json:"discount_canceled,omitempty"`
	DiscountDescription                     string          `json:"discount_description,omitempty"`
	DiscountInvoiced                        float64         `json:"discount_invoiced,omitempty"`
	DiscountRefunded                        float64         `json:"discount_refunded,omitempty"`
	EditIncrement                           float64         `json:"edit_increment,omitempty"`
	EmailSent                               float64         `json:"email_sent,omitempty"`
	EntityID                                int             `json:"entity_id,omitempty"`
	ExtCustomerID                           string          `json:"ext_customer_id,omitempty"`
	ExtOrderID                              string          `json:"ext_order_id,omitempty"`
	ForcedShipmentWithInvoice               float64         `json:"forced_shipment_with_invoice,omitempty"`
	GlobalCurrencyCode                      string          `json:"global_currency_code,omitempty"`
	GrandTotal                              float64         `json:"grand_total,omitempty"`
	DiscountTaxCompensationAmount           float64         `json:"discount_tax_compensation_amount,omitempty"`
	DiscountTaxCompensationInvoiced         float64         `json:"discount_tax_compensation_invoiced,omitempty"`
	DiscountTaxCompensationRefunded         float64         `json:"discount_tax_compensation_refunded,omitempty"`
	HoldBeforeState                         string          `json:"hold_before_state,omitempty"`
	HoldBeforeStatus                        string          `json:"hold_before_status,omitempty"`
	IncrementID                             string          `json:"increment_id,omitempty"`
	IsVirtual                               float64         `json:"is_virtual,omitempty"`
	OrderCurrencyCode                       string          `json:"order_currency_code,omitempty"`
	OriginalIncrementID                     string          `json:"original_increment_id,omitempty"`
	PaymentAuthorizationAmount              float64         `json:"payment_authorization_amount,omitempty"`
	PaymentAuthExpiration                   float64         `json:"payment_auth_expiration,omitempty"`
	ProtectCode                             string          `json:"protect_code,omitempty"`
	QuoteAddressID                          float64         `json:"quote_address_id,omitempty"`
	QuoteID                                 float64         `json:"quote_id,omitempty"`
	RelationChildID                         string          `json:"relation_child_id,omitempty"`
	RelationChildRealID                     string          `json:"relation_child_real_id,omitempty"`
	RelationParentID                        string          `json:"relation_parent_id,omitempty"`
	RelationParentRealID                    string          `json:"relation_parent_real_id,omitempty"`
	RemoteIP                                string          `json:"remote_ip,omitempty"`
	ShippingAmount                          float64         `json:"shipping_amount,omitempty"`
	ShippingCanceled                        float64         `json:"shipping_canceled,omitempty"`
	ShippingDescription                     string          `json:"shipping_description,omitempty"`
	ShippingDiscountAmount                  float64         `json:"shipping_discount_amount,omitempty"`
	ShippingDiscountTaxCompensationAmount   float64         `json:"shipping_discount_tax_compensation_amount,omitempty"`
	ShippingInclTax                         float64         `json:"shipping_incl_tax,omitempty"`
	ShippingInvoiced                        float64         `json:"shipping_invoiced,omitempty"`
	ShippingRefunded                        float64         `json:"shipping_refunded,omitempty"`
	ShippingTaxAmount                       float64         `json:"shipping_tax_amount,omitempty"`
	ShippingTaxRefunded                     float64         `json:"shipping_tax_refunded,omitempty"`
	State                                   string          `json:"state,omitempty"`
	Status                                  string          `json:"status,omitempty"`
	StoreCurrencyCode                       string          `json:"store_currency_code,omitempty"`
	StoreID                                 float64         `json:"store_id,omitempty"`
	StoreName                               string          `json:"store_name,omitempty"`
	StoreToBaseRate                         float64         `json:"store_to_base_rate,omitempty"`
	StoreToOrderRate                        float64         `json:"store_to_order_rate,omitempty"`
	Subtotal                                float64         `json:"subtotal,omitempty"`
	SubtotalCanceled                        float64         `json:"subtotal_canceled,omitempty"`
	SubtotalInclTax                         float64         `json:"subtotal_incl_tax,omitempty"`
	SubtotalInvoiced                        float64         `json:"subtotal_invoiced,omitempty"`
	SubtotalRefunded                        float64         `json:"subtotal_refunded,omitempty"`
	TaxAmount                               float64         `json:"tax_amount,omitempty"`
	TaxCanceled                             float64         `json:"tax_canceled,omitempty"`
	TaxInvoiced                             float64         `json:"tax_invoiced,omitempty"`
	TaxRefunded                             float64         `json:"tax_refunded,omitempty"`
	TotalCanceled                           float64         `json:"total_canceled,omitempty"`
	TotalDue                                float64         `json:"total_due,omitempty"`
	TotalInvoiced                           float64         `json:"total_invoiced,omitempty"`
	TotalItemCount                          float64         `json:"total_item_count,omitempty"`
	TotalOfflineRefunded                    float64         `json:"total_offline_refunded,omitempty"`
	TotalOnlineRefunded                     float64         `json:"total_online_refunded,omitempty"`
	TotalPaid                               float64         `json:"total_paid,omitempty"`
	TotalQtyOrdered                         float64         `json:"total_qty_ordered,omitempty"`
	TotalRefunded                           float64         `json:"total_refunded,omitempty"`
	UpdatedAt                               string          `json:"updated_at,omitempty"`
	Weight                                  float64         `json:"weight,omitempty"`
	XForwardedFor                           string          `json:"x_forwarded_for,omitempty"`
	Items                                   []Item          `json:"items,omitempty"`
	BillingAddress                          *BillingAddress `json:"billing_address,omitempty"`
	Payment                                 *struct {
		AccountStatus             string   `json:"account_status,omitempty"`
		AdditionalData            string   `json:"additional_data,omitempty"`
		AdditionalInformation     []string `json:"additional_information,omitempty"`
		AddressStatus             string   `json:"address_status,omitempty"`
		AmountAuthorized          float64  `json:"amount_authorized,omitempty"`
		AmountCanceled            float64  `json:"amount_canceled,omitempty"`
		AmountOrdered             float64  `json:"amount_ordered,omitempty"`
		AmountPaid                float64  `json:"amount_paid,omitempty"`
		AmountRefunded            float64  `json:"amount_refunded,omitempty"`
		AnetTransMethod           string   `json:"anet_trans_method,omitempty"`
		BaseAmountAuthorized      float64  `json:"base_amount_authorized,omitempty"`
		BaseAmountCanceled        float64  `json:"base_amount_canceled,omitempty"`
		BaseAmountOrdered         float64  `json:"base_amount_ordered,omitempty"`
		BaseAmountPaid            float64  `json:"base_amount_paid,omitempty"`
		BaseAmountPaidOnline      float64  `json:"base_amount_paid_online,omitempty"`
		BaseAmountRefunded        float64  `json:"base_amount_refunded,omitempty"`
		BaseAmountRefundedOnline  float64  `json:"base_amount_refunded_online,omitempty"`
		BaseShippingAmount        float64  `json:"base_shipping_amount,omitempty"`
		BaseShippingCaptured      float64  `json:"base_shipping_captured,omitempty"`
		BaseShippingRefunded      float64  `json:"base_shipping_refunded,omitempty"`
		CcApproval                string   `json:"cc_approval,omitempty"`
		CcAvsStatus               string   `json:"cc_avs_status,omitempty"`
		CcCidStatus               string   `json:"cc_cid_status,omitempty"`
		CcDebugRequestBody        string   `json:"cc_debug_request_body,omitempty"`
		CcDebugResponseBody       string   `json:"cc_debug_response_body,omitempty"`
		CcDebugResponseSerialized string   `json:"cc_debug_response_serialized,omitempty"`
		CcExpMonth                string   `json:"cc_exp_month,omitempty"`
		CcExpYear                 string   `json:"cc_exp_year,omitempty"`
		CcLast4                   string   `json:"cc_last4,omitempty"`
		CcNumberEnc               string   `json:"cc_number_enc,omitempty"`
		CcOwner                   string   `json:"cc_owner,omitempty"`
		CcSecureVerify            string   `json:"cc_secure_verify,omitempty"`
		CcSsIssue                 string   `json:"cc_ss_issue,omitempty"`
		CcSsStartMonth            string   `json:"cc_ss_start_month,omitempty"`
		CcSsStartYear             string   `json:"cc_ss_start_year,omitempty"`
		CcStatus                  string   `json:"cc_status,omitempty"`
		CcStatusDescription       string   `json:"cc_status_description,omitempty"`
		CcTransID                 string   `json:"cc_trans_id,omitempty"`
		CcType                    string   `json:"cc_type,omitempty"`
		EcheckAccountName         string   `json:"echeck_account_name,omitempty"`
		EcheckAccountType         string   `json:"echeck_account_type,omitempty"`
		EcheckBankName            string   `json:"echeck_bank_name,omitempty"`
		EcheckRoutingNumber       string   `json:"echeck_routing_number,omitempty"`
		EcheckType                string   `json:"echeck_type,omitempty"`
		EntityID                  int      `json:"entity_id,omitempty"`
		LastTransID               string   `json:"last_trans_id,omitempty"`
		Method                    string   `json:"method,omitempty"`
		ParentID                  float64  `json:"parent_id,omitempty"`
		PoNumber                  string   `json:"po_number,omitempty"`
		ProtectionEligibility     string   `json:"protection_eligibility,omitempty"`
		QuotePaymentID            float64  `json:"quote_payment_id,omitempty"`
		ShippingAmount            float64  `json:"shipping_amount,omitempty"`
		ShippingCaptured          float64  `json:"shipping_captured,omitempty"`
		ShippingRefunded          float64  `json:"shipping_refunded,omitempty"`
		ExtensionAttributes       *struct {
			VaultPaymentToken *struct {
				EntityID          int     `json:"entity_id,omitempty"`
				CustomerID        float64 `json:"customer_id,omitempty"`
				PublicHash        string  `json:"public_hash,omitempty"`
				PaymentMethodCode string  `json:"payment_method_code,omitempty"`
				Type              string  `json:"type,omitempty"`
				CreatedAt         string  `json:"created_at,omitempty"`
				ExpiresAt         string  `json:"expires_at,omitempty"`
				GatewayToken      string  `json:"gateway_token,omitempty"`
				TokenDetails      string  `json:"token_details,omitempty"`
				IsActive          bool    `json:"is_active,omitempty"`
				IsVisible         bool    `json:"is_visible,omitempty"`
			} `json:"vault_payment_token,omitempty"`
		} `json:"extension_attributes,omitempty"`
	} `json:"payment,omitempty"`
	StatusHistories     []StatusHistory `json:"status_histories,omitempty"`
	ExtensionAttributes *struct {
		ShippingAssignments []struct {
			Shipping *struct {
				Address *ShippingAddress `json:"address,omitempty"`
				Method  string           `json:"method,omitempty"`
				Total   *struct {
					BaseShippingAmount                      float64 `json:"base_shipping_amount,omitempty"`
					BaseShippingCanceled                    float64 `json:"base_shipping_canceled,omitempty"`
					BaseShippingDiscountAmount              float64 `json:"base_shipping_discount_amount,omitempty"`
					BaseShippingDiscountTaxCompensationAmnt float64 `json:"base_shipping_discount_tax_compensation_amnt,omitempty"`
					BaseShippingInclTax                     float64 `json:"base_shipping_incl_tax,omitempty"`
					BaseShippingInvoiced                    float64 `json:"base_shipping_invoiced,omitempty"`
					BaseShippingRefunded                    float64 `json:"base_shipping_refunded,omitempty"`
					BaseShippingTaxAmount                   float64 `json:"base_shipping_tax_amount,omitempty"`
					BaseShippingTaxRefunded                 float64 `json:"base_shipping_tax_refunded,omitempty"`
					ShippingAmount                          float64 `json:"shipping_amount,omitempty"`
					ShippingCanceled                        float64 `json:"shipping_canceled,omitempty"`
					ShippingDiscountAmount                  float64 `json:"shipping_discount_amount,omitempty"`
					ShippingDiscountTaxCompensationAmount   float64 `json:"shipping_discount_tax_compensation_amount,omitempty"`
					ShippingInclTax                         float64 `json:"shipping_incl_tax,omitempty"`
					ShippingInvoiced                        float64 `json:"shipping_invoiced,omitempty"`
					ShippingRefunded                        float64 `json:"shipping_refunded,omitempty"`
					ShippingTaxAmount                       float64 `json:"shipping_tax_amount,omitempty"`
					ShippingTaxRefunded                     float64 `json:"shipping_tax_refunded,omitempty"`
					ExtensionAttributes                     struct {
					} `json:"extension_attributes,omitempty"`
				} `json:"total,omitempty"`
				ExtensionAttributes *struct {
					ExtOrderID         string `json:"ext_order_id,omitempty"`
					ShippingExperience *struct {
						Label string  `json:"label,omitempty"`
						Code  string  `json:"code,omitempty"`
						Cost  float64 `json:"cost,omitempty"`
					} `json:"shipping_experience,omitempty"`
					CollectionPofloat64 *struct {
						RecipientAddressID    float64  `json:"recipient_address_id,omitempty"`
						CollectionPofloat64ID string   `json:"collection_pofloat64_id,omitempty"`
						Name                  string   `json:"name,omitempty"`
						Country               string   `json:"country,omitempty"`
						Region                string   `json:"region,omitempty"`
						Postcode              string   `json:"postcode,omitempty"`
						City                  string   `json:"city,omitempty"`
						Street                []string `json:"street,omitempty"`
					} `json:"collection_pofloat64,omitempty"`
				} `json:"extension_attributes,omitempty"`
			} `json:"shipping,omitempty"`
			Items               []Item  `json:"items,omitempty"`
			StockID             float64 `json:"stock_id,omitempty"`
			ExtensionAttributes *struct {
			} `json:"extension_attributes,omitempty"`
		} `json:"shipping_assignments,omitempty"`
		PaymentAdditionalInfo []struct {
			Key   string `json:"key,omitempty"`
			Value string `json:"value,omitempty"`
		} `json:"payment_additional_info,omitempty"`
		CompanyOrderAttributes *struct {
			OrderID             float64 `json:"order_id,omitempty"`
			CompanyID           float64 `json:"company_id,omitempty"`
			CompanyName         string  `json:"company_name,omitempty"`
			ExtensionAttributes *struct {
			} `json:"extension_attributes,omitempty"`
		} `json:"company_order_attributes,omitempty"`
		AppliedTaxes []struct {
			Code                string  `json:"code,omitempty"`
			Title               string  `json:"title,omitempty"`
			Percent             float64 `json:"percent,omitempty"`
			Amount              float64 `json:"amount,omitempty"`
			BaseAmount          float64 `json:"base_amount,omitempty"`
			ExtensionAttributes *struct {
				Rates []struct {
					Code                string  `json:"code,omitempty"`
					Title               string  `json:"title,omitempty"`
					Percent             float64 `json:"percent,omitempty"`
					ExtensionAttributes *struct {
					} `json:"extension_attributes,omitempty"`
				} `json:"rates,omitempty"`
			} `json:"extension_attributes,omitempty"`
		} `json:"applied_taxes,omitempty"`
		ItemAppliedTaxes []struct {
			Type             string  `json:"type,omitempty"`
			ItemID           float64 `json:"item_id,omitempty"`
			AssociatedItemID float64 `json:"associated_item_id,omitempty"`
			AppliedTaxes     []struct {
				Code                string  `json:"code,omitempty"`
				Title               string  `json:"title,omitempty"`
				Percent             float64 `json:"percent,omitempty"`
				Amount              float64 `json:"amount,omitempty"`
				BaseAmount          float64 `json:"base_amount,omitempty"`
				ExtensionAttributes *struct {
					Rates []struct {
						Code                string  `json:"code,omitempty"`
						Title               string  `json:"title,omitempty"`
						Percent             float64 `json:"percent,omitempty"`
						ExtensionAttributes *struct {
						} `json:"extension_attributes,omitempty"`
					} `json:"rates,omitempty"`
				} `json:"extension_attributes,omitempty"`
			} `json:"applied_taxes,omitempty"`
			ExtensionAttributes *struct {
			} `json:"extension_attributes,omitempty"`
		} `json:"item_applied_taxes,omitempty"`
		ConvertingFromQuote              bool    `json:"converting_from_quote,omitempty"`
		BaseCustomerBalanceAmount        float64 `json:"base_customer_balance_amount,omitempty"`
		CustomerBalanceAmount            float64 `json:"customer_balance_amount,omitempty"`
		BaseCustomerBalanceInvoiced      float64 `json:"base_customer_balance_invoiced,omitempty"`
		CustomerBalanceInvoiced          float64 `json:"customer_balance_invoiced,omitempty"`
		BaseCustomerBalanceRefunded      float64 `json:"base_customer_balance_refunded,omitempty"`
		CustomerBalanceRefunded          float64 `json:"customer_balance_refunded,omitempty"`
		BaseCustomerBalanceTotalRefunded float64 `json:"base_customer_balance_total_refunded,omitempty"`
		CustomerBalanceTotalRefunded     float64 `json:"customer_balance_total_refunded,omitempty"`
		GiftCards                        []struct {
			ID         float64 `json:"id,omitempty"`
			Code       string  `json:"code,omitempty"`
			Amount     float64 `json:"amount,omitempty"`
			BaseAmount float64 `json:"base_amount,omitempty"`
		} `json:"gift_cards,omitempty"`
		BaseGiftCardsAmount   float64 `json:"base_gift_cards_amount,omitempty"`
		GiftCardsAmount       float64 `json:"gift_cards_amount,omitempty"`
		BaseGiftCardsInvoiced float64 `json:"base_gift_cards_invoiced,omitempty"`
		GiftCardsInvoiced     float64 `json:"gift_cards_invoiced,omitempty"`
		BaseGiftCardsRefunded float64 `json:"base_gift_cards_refunded,omitempty"`
		GiftCardsRefunded     float64 `json:"gift_cards_refunded,omitempty"`
		GiftMessage           *struct {
			GiftMessageID       float64 `json:"gift_message_id,omitempty"`
			CustomerID          float64 `json:"customer_id,omitempty"`
			Sender              string  `json:"sender,omitempty"`
			Recipient           string  `json:"recipient,omitempty"`
			Message             string  `json:"message,omitempty"`
			ExtensionAttributes *struct {
				EntityID                   string  `json:"entity_id,omitempty"`
				EntityType                 string  `json:"entity_type,omitempty"`
				WrappingID                 float64 `json:"wrapping_id,omitempty"`
				WrappingAllowGiftReceipt   bool    `json:"wrapping_allow_gift_receipt,omitempty"`
				WrappingAddPrfloat64edCard bool    `json:"wrapping_add_prfloat64ed_card,omitempty"`
			} `json:"extension_attributes,omitempty"`
		} `json:"gift_message,omitempty"`
		GwID                     string  `json:"gw_id,omitempty"`
		GwAllowGiftReceipt       string  `json:"gw_allow_gift_receipt,omitempty"`
		GwAddCard                string  `json:"gw_add_card,omitempty"`
		GwBasePrice              string  `json:"gw_base_price,omitempty"`
		GwPrice                  string  `json:"gw_price,omitempty"`
		GwItemsBasePrice         string  `json:"gw_items_base_price,omitempty"`
		GwItemsPrice             string  `json:"gw_items_price,omitempty"`
		GwCardBasePrice          string  `json:"gw_card_base_price,omitempty"`
		GwCardPrice              string  `json:"gw_card_price,omitempty"`
		GwBaseTaxAmount          string  `json:"gw_base_tax_amount,omitempty"`
		GwTaxAmount              string  `json:"gw_tax_amount,omitempty"`
		GwItemsBaseTaxAmount     string  `json:"gw_items_base_tax_amount,omitempty"`
		GwItemsTaxAmount         string  `json:"gw_items_tax_amount,omitempty"`
		GwCardBaseTaxAmount      string  `json:"gw_card_base_tax_amount,omitempty"`
		GwCardTaxAmount          string  `json:"gw_card_tax_amount,omitempty"`
		GwBasePriceInclTax       string  `json:"gw_base_price_incl_tax,omitempty"`
		GwPriceInclTax           string  `json:"gw_price_incl_tax,omitempty"`
		GwItemsBasePriceInclTax  string  `json:"gw_items_base_price_incl_tax,omitempty"`
		GwItemsPriceInclTax      string  `json:"gw_items_price_incl_tax,omitempty"`
		GwCardBasePriceInclTax   string  `json:"gw_card_base_price_incl_tax,omitempty"`
		GwCardPriceInclTax       string  `json:"gw_card_price_incl_tax,omitempty"`
		GwBasePriceInvoiced      string  `json:"gw_base_price_invoiced,omitempty"`
		GwPriceInvoiced          string  `json:"gw_price_invoiced,omitempty"`
		GwItemsBasePriceInvoiced string  `json:"gw_items_base_price_invoiced,omitempty"`
		GwItemsPriceInvoiced     string  `json:"gw_items_price_invoiced,omitempty"`
		GwCardBasePriceInvoiced  string  `json:"gw_card_base_price_invoiced,omitempty"`
		GwCardPriceInvoiced      string  `json:"gw_card_price_invoiced,omitempty"`
		GwBaseTaxAmountInvoiced  string  `json:"gw_base_tax_amount_invoiced,omitempty"`
		GwTaxAmountInvoiced      string  `json:"gw_tax_amount_invoiced,omitempty"`
		GwItemsBaseTaxInvoiced   string  `json:"gw_items_base_tax_invoiced,omitempty"`
		GwItemsTaxInvoiced       string  `json:"gw_items_tax_invoiced,omitempty"`
		GwCardBaseTaxInvoiced    string  `json:"gw_card_base_tax_invoiced,omitempty"`
		GwCardTaxInvoiced        string  `json:"gw_card_tax_invoiced,omitempty"`
		GwBasePriceRefunded      string  `json:"gw_base_price_refunded,omitempty"`
		GwPriceRefunded          string  `json:"gw_price_refunded,omitempty"`
		GwItemsBasePriceRefunded string  `json:"gw_items_base_price_refunded,omitempty"`
		GwItemsPriceRefunded     string  `json:"gw_items_price_refunded,omitempty"`
		GwCardBasePriceRefunded  string  `json:"gw_card_base_price_refunded,omitempty"`
		GwCardPriceRefunded      string  `json:"gw_card_price_refunded,omitempty"`
		GwBaseTaxAmountRefunded  string  `json:"gw_base_tax_amount_refunded,omitempty"`
		GwTaxAmountRefunded      string  `json:"gw_tax_amount_refunded,omitempty"`
		GwItemsBaseTaxRefunded   string  `json:"gw_items_base_tax_refunded,omitempty"`
		GwItemsTaxRefunded       string  `json:"gw_items_tax_refunded,omitempty"`
		GwCardBaseTaxRefunded    string  `json:"gw_card_base_tax_refunded,omitempty"`
		GwCardTaxRefunded        string  `json:"gw_card_tax_refunded,omitempty"`
		RewardPofloat64sBalance  float64 `json:"reward_pofloat64s_balance,omitempty"`
		RewardCurrencyAmount     float64 `json:"reward_currency_amount,omitempty"`
		BaseRewardCurrencyAmount float64 `json:"base_reward_currency_amount,omitempty"`
		AmazonOrderReferenceID   *struct {
			AmazonOrderReferenceID string  `json:"amazon_order_reference_id,omitempty"`
			OrderID                float64 `json:"order_id,omitempty"`
		} `json:"amazon_order_reference_id,omitempty"`
	} `json:"extension_attributes,omitempty"`
}

type OrdersProductOption struct {
	ExtensionAttributes *struct {
		CustomOptions []struct {
			OptionID            string `json:"option_id,omitempty"`
			OptionValue         string `json:"option_value,omitempty"`
			ExtensionAttributes *struct {
				FileInfo *struct {
					Base64EncodedData string `json:"base64_encoded_data,omitempty"`
					Type              string `json:"type,omitempty"`
					Name              string `json:"name,omitempty"`
				} `json:"file_info,omitempty"`
			} `json:"extension_attributes,omitempty"`
		} `json:"custom_options,omitempty"`
		BundleOptions []struct {
			OptionID            float64   `json:"option_id,omitempty"`
			OptionQty           float64   `json:"option_qty,omitempty"`
			OptionSelections    []float64 `json:"option_selections,omitempty"`
			ExtensionAttributes *struct {
			} `json:"extension_attributes,omitempty"`
		} `json:"bundle_options,omitempty"`
		ConfigurableItemOptions []struct {
			OptionID            string  `json:"option_id,omitempty"`
			OptionValue         float64 `json:"option_value,omitempty"`
			ExtensionAttributes *struct {
			} `json:"extension_attributes,omitempty"`
		} `json:"configurable_item_options,omitempty"`
		DownloadableOption *struct {
			DownloadableLinks []float64 `json:"downloadable_links,omitempty"`
		} `json:"downloadable_option,omitempty"`
		GiftcardItemOption *struct {
			GiftcardAmount         string  `json:"giftcard_amount,omitempty"`
			CustomGiftcardAmount   float64 `json:"custom_giftcard_amount,omitempty"`
			GiftcardSenderName     string  `json:"giftcard_sender_name,omitempty"`
			GiftcardRecipientName  string  `json:"giftcard_recipient_name,omitempty"`
			GiftcardSenderEmail    string  `json:"giftcard_sender_email,omitempty"`
			GiftcardRecipientEmail string  `json:"giftcard_recipient_email,omitempty"`
			GiftcardMessage        string  `json:"giftcard_message,omitempty"`
			ExtensionAttributes    *struct {
			} `json:"extension_attributes,omitempty"`
		} `json:"giftcard_item_option,omitempty"`
	} `json:"extension_attributes,omitempty"`
}

type StatusHistory struct {
	Comment             string  `json:"comment,omitempty"`
	CreatedAt           string  `json:"created_at,omitempty"`
	EntityID            int     `json:"entity_id,omitempty"`
	EntityName          string  `json:"entity_name,omitempty"`
	IsCustomerNotified  float64 `json:"is_customer_notified,omitempty"`
	IsVisibleOnFront    float64 `json:"is_visible_on_front,omitempty"`
	ParentID            float64 `json:"parent_id,omitempty"`
	Status              string  `json:"status,omitempty"`
	ExtensionAttributes *struct {
	} `json:"extension_attributes,omitempty"`
}
