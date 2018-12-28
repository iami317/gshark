package models

import (
	"time"
	"github.com/neal1991/gshark/vars"
	"fmt"
	"github.com/neal1991/gshark/util/common"
)

// AppSearchResult represents a single search result for app market search
type AppSearchResult struct {
	Id          int64
	Name        *string `json:"name,omitempty"`
	Description *string
	Market      *string `json:"market,omitempty"`
	Developer   *string
	Version     *string
	DeployDate  *string
	AppUrl      *string
	Status      int
	CreatedTime time.Time
	UpdatedTime time.Time
}

func (r *AppSearchResult) Insert() (int64, error) {
	return Engine.Insert(r)
}

func (r *AppSearchResult) Exist() (bool, error) {
	return Engine.Table("app_search_result").Where("name=? and market=?",
		r.Name, r.Market).Exist()
}

func ListAppSearchResultByPage(page int, status int) ([]AppSearchResult, int, int) {
	results := make([]AppSearchResult, 0)
	totalPages, err := Engine.Table("app_search_result").Where("status=?", status).Count()
	var pages int

	page, pages = common.GetPageAndPagesByTotalPages(page, int(totalPages))

	err = Engine.Where("status=?", status).
		Limit(vars.PAGE_SIZE, (page-1)*vars.PAGE_SIZE).Desc("id").Find(&results)

	if err != nil {
		fmt.Errorf("search failed:%s", err)
	}

	return results, pages, int(totalPages)
}