package models

type Query struct {
	QueryLanguage   string   `json:"query_language"`
	QueryString     string   `json:"query_string"`
	ResultLanguages []string `json:"result_languages"`
	Limit           int    `json:"limit"`
	Bdats           []string `json:"bdats"`
	Tables          []string `json:"tables"`
}
