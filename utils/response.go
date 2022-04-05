package utils

import "backend-onboarding/model/dto"

func ResponseError(status string, err interface{}, code int) dto.Result {
	return dto.Result{
		Status:     status,
		StatusCode: code,
		Error:      err,
		Data:       nil,
	}
}

func ResponseSuccess(status string, err interface{}, data interface{}, code int) dto.Result {
	return dto.Result{
		Status:     status,
		StatusCode: code,
		Error:      err,
		Data:       data,
	}
}
