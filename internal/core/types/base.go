package types

type BaseResponse struct {
	Status int
}

type CheckResult struct {
	Name   string
	Passed bool
}

type CheckFunc func(lang string) CheckResult
