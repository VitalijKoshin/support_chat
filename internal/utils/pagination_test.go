package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetTotalPages(t *testing.T) {
	type args struct {
		totalCount int
		pageSize   int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Normal 0", args{totalCount: 0, pageSize: 1}, 0},
		{"Normal 1", args{totalCount: 1, pageSize: 2}, 1},
		{"Normal 10", args{totalCount: 10, pageSize: 3}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GetTotalPages(tt.args.totalCount, tt.args.pageSize)
			assert.Equal(t, got, tt.want)
		})
	}
}

func TestGetHasMore(t *testing.T) {
	type args struct {
		currentPage int
		totalCount  int
		pageSize    int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"Normal 0", args{currentPage: 0, totalCount: 0, pageSize: 0}, false},
		{"Normal 1", args{currentPage: 1, totalCount: 1, pageSize: 2}, false},
		{"Normal 10", args{currentPage: 1, totalCount: 10, pageSize: 3}, true},
		{"Normal 10", args{currentPage: 9, totalCount: 10, pageSize: 3}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GetHasMore(tt.args.currentPage, tt.args.totalCount, tt.args.pageSize)
			assert.Equal(t, got, tt.want)
		})
	}
}

func TestPaginationQuery_GetSize(t *testing.T) {
	type fields struct {
		Size    int
		Page    int
		OrderBy string
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{"Normal 0", fields{Size: 0, Page: 0, OrderBy: ""}, 0},
		{"Normal 1", fields{Size: 1, Page: 0, OrderBy: ""}, 1},
		{"Normal 10", fields{Size: 10, Page: 0, OrderBy: ""}, 10},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &PaginationQuery{
				Size:    tt.fields.Size,
				Page:    tt.fields.Page,
				OrderBy: tt.fields.OrderBy,
			}
			got := q.GetSize()
			assert.Equal(t, got, tt.want)
		})
	}
}

func TestPaginationQuery_GetPage(t *testing.T) {
	type fields struct {
		Size    int
		Page    int
		OrderBy string
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{"Normal 0", fields{Size: 0, Page: 0, OrderBy: ""}, 0},
		{"Normal 1", fields{Size: 1, Page: 1, OrderBy: ""}, 1},
		{"Normal 10", fields{Size: 10, Page: 10, OrderBy: ""}, 10},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &PaginationQuery{
				Size:    tt.fields.Size,
				Page:    tt.fields.Page,
				OrderBy: tt.fields.OrderBy,
			}
			got := q.GetPage()
			assert.Equal(t, got, tt.want)
		})
	}
}

func TestPaginationQuery_GetOffset(t *testing.T) {
	type fields struct {
		Size    int
		Page    int
		OrderBy string
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{"Normal 0", fields{Size: 0, Page: 0, OrderBy: ""}, 0},
		{"Normal 1", fields{Size: 1, Page: 1, OrderBy: ""}, 0},
		{"Normal 2", fields{Size: 1, Page: 2, OrderBy: ""}, 1},
		{"Normal 10", fields{Size: 10, Page: 10, OrderBy: ""}, 90},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &PaginationQuery{
				Size:    tt.fields.Size,
				Page:    tt.fields.Page,
				OrderBy: tt.fields.OrderBy,
			}
			got := q.GetOffset()
			assert.Equal(t, got, tt.want)
		})
	}
}

func TestPaginationQuery_GetLimit(t *testing.T) {
	type fields struct {
		Size    int
		Page    int
		OrderBy string
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{"Normal 0", fields{Size: 0, Page: 0, OrderBy: ""}, 0},
		{"Normal 1", fields{Size: 1, Page: 1, OrderBy: ""}, 1},
		{"Normal 2", fields{Size: 1, Page: 2, OrderBy: ""}, 1},
		{"Normal 10", fields{Size: 10, Page: 10, OrderBy: ""}, 10},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &PaginationQuery{
				Size:    tt.fields.Size,
				Page:    tt.fields.Page,
				OrderBy: tt.fields.OrderBy,
			}
			got := q.GetLimit()
			assert.Equal(t, got, tt.want)
		})
	}
}

func TestPaginationQuery_SetSize(t *testing.T) {
	type fields struct {
		Size    int
		Page    int
		OrderBy string
	}
	type args struct {
		sizeQuery string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{"Normal", fields{Size: 0, Page: 0, OrderBy: ""}, args{sizeQuery: "2"}, false},
		{"Normal 10", fields{Size: 10, Page: 1, OrderBy: ""}, args{sizeQuery: "10"}, false},
		{"Normal -1", fields{Size: 10, Page: 1, OrderBy: ""}, args{sizeQuery: "-1"}, false},
		{"Normal empty str", fields{Size: 10, Page: 1, OrderBy: ""}, args{sizeQuery: ""}, false},
		{"Error str", fields{Size: 1, Page: 1, OrderBy: ""}, args{sizeQuery: "wevw"}, true},
		{"Error --1", fields{Size: 10, Page: 1, OrderBy: ""}, args{sizeQuery: "--1"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &PaginationQuery{
				Size:    tt.fields.Size,
				Page:    tt.fields.Page,
				OrderBy: tt.fields.OrderBy,
			}
			err := q.SetSize(tt.args.sizeQuery)
			isErr := (err != nil)
			assert.Equal(t, isErr, tt.wantErr)
		})
	}
}

func TestPaginationQuery_SetPage(t *testing.T) {
	type fields struct {
		Size    int
		Page    int
		OrderBy string
	}
	type args struct {
		pageQuery string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{"Normal", fields{Size: 0, Page: 0, OrderBy: ""}, args{pageQuery: "2"}, false},
		{"Normal 10", fields{Size: 10, Page: 1, OrderBy: ""}, args{pageQuery: "10"}, false},
		{"Normal -1", fields{Size: 10, Page: 1, OrderBy: ""}, args{pageQuery: "-1"}, false},
		{"Normal empty str", fields{Size: 10, Page: 1, OrderBy: ""}, args{pageQuery: ""}, false},
		{"Error str", fields{Size: 1, Page: 1, OrderBy: ""}, args{pageQuery: "wevw"}, true},
		{"Error --1", fields{Size: 10, Page: 1, OrderBy: ""}, args{pageQuery: "--1"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &PaginationQuery{
				Size:    tt.fields.Size,
				Page:    tt.fields.Page,
				OrderBy: tt.fields.OrderBy,
			}
			err := q.SetPage(tt.args.pageQuery)
			isErr := (err != nil)
			assert.Equal(t, isErr, tt.wantErr)
		})
	}
}

func TestPaginationQuery_SetOrderBy(t *testing.T) {
	type fields struct {
		Size    int
		Page    int
		OrderBy string
	}
	type args struct {
		orderByQuery string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{"Normal", fields{Size: 0, Page: 0, OrderBy: ""}, args{orderByQuery: "orderBy"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &PaginationQuery{
				Size:    tt.fields.Size,
				Page:    tt.fields.Page,
				OrderBy: tt.fields.OrderBy,
			}
			q.SetOrderBy(tt.args.orderByQuery)
			assert.Equal(t, q.OrderBy, tt.args.orderByQuery)
		})
	}
}

func TestGetPaginationFromCtx(t *testing.T) {
	type args struct {
		page    string
		size    string
		orderBy string
	}
	tests := []struct {
		name    string
		args    args
		want    *PaginationQuery
		wantErr bool
	}{
		{"page 0", args{page: "0", size: "0", orderBy: ""}, &PaginationQuery{Size: 2, Page: 0, OrderBy: ""}, false},
		{"page 1", args{page: "1", size: "0", orderBy: ""}, &PaginationQuery{Size: 2, Page: 1, OrderBy: ""}, false},
		{"page 1", args{page: "2", size: "0", orderBy: ""}, &PaginationQuery{Size: 2, Page: 2, OrderBy: ""}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetPaginationFromCtx(tt.args.page, tt.args.size, tt.args.orderBy)
			assert.Equal(t, got, tt.want)
			assert.Equal(t, (err != nil), tt.wantErr)
		})
	}
}
