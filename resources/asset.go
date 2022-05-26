package resources

type Asset struct {
	Symbol string  `json:"symbol"`
	Quotes []Quote `json:"quotes"`
}
