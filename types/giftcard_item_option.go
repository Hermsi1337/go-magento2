package types

type GiftcardItemOption struct {
	GiftcardAmount         string                 `json:"giftcard_amount,omitempty"`
	CustomGiftcardAmount   int                    `json:"custom_giftcard_amount,omitempty"`
	GiftcardSenderName     string                 `json:"giftcard_sender_name,omitempty"`
	GiftcardRecipientName  string                 `json:"giftcard_recipient_name,omitempty"`
	GiftcardSenderEmail    string                 `json:"giftcard_sender_email,omitempty"`
	GiftcardRecipientEmail string                 `json:"giftcard_recipient_email,omitempty"`
	GiftcardMessage        string                 `json:"giftcard_message,omitempty"`
	ExtensionAttributes    map[string]interface{} `json:"extension_attributes,omitempty"`
}
