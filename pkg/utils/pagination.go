package utils

import (
	"errors"
	"math"
	"strconv"
	"strings"
)

const (
	defaultItemsPerPage = 25
)

// ResponseWithPagination - Get response data with pagination object
type ResponseWithPagination struct {
	Response         interface{}
	PaginationObject Pagination
}

// Pagination - json is optimized for VuetifyJS
type Pagination struct {
	Page         int64 `json:"page" jsonList:"page"`                 // current page
	ItemsPerPage int64 `json:"itemsPerPage" jsonList:"itemsPerPage"` // Count of items per page
	PageStart    int64 `json:"pageStart" jsonList:"pageStart"`       // First page
	PageStop     int64 `json:"pageStop" jsonList:"pageStop"`         // Last page
	PageCount    int64 `json:"pageCount" jsonList:"pageCount"`       // Pages count
	PageNext     int64 `json:"pageNext" jsonList:"pageNext"`         // Next page
	PagePrev     int64 `json:"pagePrev" jsonList:"pagePrev"`         // Previous page
	ItemsCount   int64 `json:"totalItems" jsonList:"totalItems"`     // Count of all items in dataset
}

// QueryOptions - object
type QueryOptions struct {
	SortBy    []string
	SortDesc  []bool
	GroupBy   []string
	GroupDesc []bool
	MultiSort bool
	MustSort  bool
}

// SetValuesFromPage - set all values from current page and items per page AND recount it
func (me *Pagination) SetValuesFromPage(currentPage int64, itemsPerPage int64, itemsCount int64) {
	me.Page = currentPage
	me.ItemsPerPage = itemsPerPage
	me.ItemsCount = itemsCount

	me.RecountValuesFromPage()
}

// RecountValuesFromPage - method recount values from page, itemsPerPage and itemsCount
func (me *Pagination) RecountValuesFromPage() (err error) {
	if me.ItemsPerPage < 1 {
		me.ItemsPerPage = defaultItemsPerPage
	}

	// nastavit kdyz jsou hodnoty nulove!
	if me.ItemsCount < 1 {
		me.PageStart = 0
		me.PageStop = 0
		me.PageCount = 0
		me.PageNext = 0
		me.PagePrev = 0

		return errors.New("ItemsCount must be greater than zero")
	}

	// number of pages
	me.PageCount = int64(math.Round(float64(me.ItemsCount / me.ItemsPerPage)))
	if (me.ItemsCount % me.ItemsPerPage) > 0 {
		me.PageCount = me.PageCount + 1
	}

	me.PageStart = 1
	me.PageStop = me.PageCount

	// correct bad assigned current page
	if !(me.Page > 0 && me.Page <= me.PageCount) {
		me.Page = me.PageStart
		println("proslo podminkou")
	}

	// initialy set
	me.PageNext = me.PageStop

	if me.Page < me.PageCount {
		me.PageNext = me.Page + 1
	}

	// initialy set
	me.PagePrev = me.PageStart

	// set if it's possible
	if me.Page > me.PageStart {
		me.PagePrev = me.Page - 1
	}
	return nil
}

// GetLimits - Vraci limity pro databazi
func (me *Pagination) GetLimits(currentPage int64) (offset int64, limit int64, err error) {

	// test if it's recounted first
	if me.ItemsPerPage < 1 && currentPage < 1 {
		return 0, 0, errors.New("First you must call RecountValuesFromPage")
	}

	// test limits
	if currentPage > me.PageCount {
		return 0, 0, errors.New("current page out of limits")
	}

	limit = me.ItemsPerPage
	offset = (currentPage - 1) * me.ItemsPerPage

	return offset, limit, nil
}

func GetCurrentPage(pageStr string) (currentPage int, err error) {
	pageID := strings.ReplaceAll(pageStr, "/", "")
	if len(pageID) == 0 {
		return 1, nil
	}

	id, err := strconv.Atoi(pageID)
	if err != nil {
		return 0, err
	}
	return id, nil
}
