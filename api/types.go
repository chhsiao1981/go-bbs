package api

type ApiFunc func(params interface{}) (interface{}, error)

type LoginRequiredApiFunc func(userID string, params interface{}) (interface{}, error)
