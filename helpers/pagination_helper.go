package helpers

import "test-echo/dto"

type PaginationHelper struct{}

func (h *PaginationHelper) PaginationOptions(input *dto.PaginationInput) *dto.PaginationOptions {
	var resp dto.PaginationOptions
	if input != nil {
		if input.Limit <= 0 {
			input.Limit = 20
		}
		if input.Page <= 0 {
			input.Page = 1
		}

		resp.Limit = input.Limit
		resp.Offset = (input.Page - 1) * input.Limit
	}

	return &resp
}
