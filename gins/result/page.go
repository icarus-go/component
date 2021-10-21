package result

type PageResult struct {
	Page     int         `json:"page"`
	PageSize int         `json:"pageSize"`
	Total    int64       `json:"total"`
	List     interface{} `json:"list"`
}

func NewPageResult(list interface{}, total int64, page int, pageSize int) *PageResult {
	if list == nil || total < 1 {
		return &PageResult{Page: page, PageSize: pageSize, Total: total, List: []struct{}{}}
	}
	return &PageResult{Page: page, PageSize: pageSize, Total: total, List: list}
}
