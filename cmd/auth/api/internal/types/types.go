// Code generated by goctl. DO NOT EDIT.
package types

type CommonResp struct {
	Success bool   `json:"success"`
	Detail  string `json:"detail,omitempty"`
}

type AuthReq struct {
	AccessToken string `header:"Access-Token"`
}

type AuthResp struct {
	CommonResp
}
