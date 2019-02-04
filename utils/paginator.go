package utils

import (
	"math"

	"github.com/jinzhu/gorm"
)

const (
	Page    = 1
	PerPage = 25
)

type Paginator struct {
	DB             *gorm.DB
	PerPage        int
	Representation interface{}
}

type Data struct {
	TotalRecords int         `json:"total_records"`
	Records      interface{} `json:"records"`
	CurrentPage  int         `json:"current_page"`
	TotalPages   int         `json:"total_pages"`
}

func (p *Paginator) validatePages(page int, TotalPages int) int {
	if page > 0 && page <= TotalPages {
		return page
	} else if page > TotalPages {
		return TotalPages
	}
	return 1
}

func (p *Paginator) validatePerPage() {
	if p.PerPage < 1 && p.PerPage > 100 {
		p.PerPage = PerPage

	}
}

func (p *Paginator) calculateOffSet(page int) int {
	pageAux := page - 1
	return pageAux * p.PerPage
}

func (p *Paginator) Paginate(dataSource interface{}, page int) *Data {
	db := p.DB

	var count int
	countRecords(db, dataSource, &count)

	p.validatePerPage()
	totalPages := getTotalPages(p.PerPage, count)
	page = p.validatePages(page, totalPages)

	var output Data
	output.TotalRecords = count
	output.CurrentPage = page
	output.TotalPages = totalPages

	offset := p.calculateOffSet(page)
	db = db.Limit(p.PerPage).Offset(offset)

	if p.Representation != nil {
		db.Scan(p.Representation)
		output.Records = p.Representation
	} else {
		db.Find(dataSource)
		output.Records = dataSource
	}

	return &output
}

func countRecords(db *gorm.DB, countDataSource interface{}, count *int) {
	if countDataSource != nil {
		db = db.Model(countDataSource)
	}
	db.Count(count)
}

func getTotalPages(perPage int, totalRecords int) int {
	totalPages := (float64(totalRecords) / float64(perPage))
	return int(math.Ceil(totalPages))
}
