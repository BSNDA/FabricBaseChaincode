package models

// Page Info
type PageInfo struct {
	PageIndex int    `json:"pageIndex"` // PageIndex
	PageSize  int    `json:"pageSize"`  // PageSize
	Bookmark  string `json:"bookmark"`  // bookmark
}

type PageListResult struct {
	TotalCount int32       `json:"totalCount"` // totalCount
	List       interface{} `json:"list"`       // result data
}
