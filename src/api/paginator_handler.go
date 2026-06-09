package api

import (
	"strconv"

	request_utils "github.com/cubbit/composer-cli/src/request_utils"
)

type PaginatorHandlerOptions struct {
	PageSize       int
	PageNumer      int
	NextPageNumber *int
	FetchAllPages  bool
}

func NewDefaultPaginatorHandlerOptions() *PaginatorHandlerOptions {
	pageSize := 1000
	pageNumber := 1
	nextPageNumber := 1
	fetchAllPages := false

	return &PaginatorHandlerOptions{
		PageSize:       pageSize,
		PageNumer:      pageNumber,
		NextPageNumber: &nextPageNumber,
		FetchAllPages:  fetchAllPages,
	}
}

func (o *PaginatorHandlerOptions) GetNextPageNumber() *int {
	return o.NextPageNumber
}

func (o *PaginatorHandlerOptions) SetNextPage(nextPage *int) {
	o.NextPageNumber = nextPage
}

type PaginatorHandlerOptionsFunc = func(*PaginatorHandlerOptions) error

func WithPageSize(pageSize int) PaginatorHandlerOptionsFunc {
	return func(options *PaginatorHandlerOptions) error {
		options.PageSize = pageSize
		return nil
	}
}

func WithPageNumber(pageNumber int) PaginatorHandlerOptionsFunc {
	return func(options *PaginatorHandlerOptions) error {
		options.PageNumer = pageNumber
		options.NextPageNumber = &pageNumber
		return nil
	}
}

func WithFetchAllPages(fetchAll bool) PaginatorHandlerOptionsFunc {
	return func(options *PaginatorHandlerOptions) error {
		options.FetchAllPages = fetchAll
		return nil
	}
}

func PaginatorHandler[T any](
	requestFunc request_utils.BoundRequest,
	urlBuilder URLBuilder,
	opts ...PaginatorHandlerOptionsFunc,
) ([]T, error) {
	options := NewDefaultPaginatorHandlerOptions()

	for _, opt := range opts {
		if err := opt(options); err != nil {
			return nil, err
		}
	}

	response := make([]T, 0)

	for options.GetNextPageNumber() != nil {
		pageString := strconv.Itoa(*options.GetNextPageNumber())
		itemsString := strconv.Itoa(options.PageSize)

		url := urlBuilder.
			QueryParam("page", pageString).
			QueryParam("items", itemsString).
			Build()

		var pageResponse GenericPaginatedResponse[T]
		if err := requestFunc.Do(
			url,
			ExtractGenericModel(&pageResponse),
		); err != nil {
			return nil, err
		}

		response = append(response, pageResponse.Data...)

		if !options.FetchAllPages {
			break
		}

		options.SetNextPage(pageResponse.NextPage)
	}

	return response, nil
}
