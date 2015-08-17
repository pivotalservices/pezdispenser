package vcloud_client

import "errors"

const (
	//VCloudTokenHeaderName - response header name for the auth token
	VCloudTokenHeaderName = "X-Vcloud-Authorization"
	//AuthSuccessStatusCode - status code expected for a successful auth call to the vcd api
	AuthSuccessStatusCode = 200
	//VAppTemplateName - the name of the vapp template to seed our apps from by defualt
	VAppTemplateName = "vSphere6-base-pcfaas-0.9.2"
)

var (
	//ErrAuthFailure - error message returned for authentication responses not having a success statuscode
	ErrAuthFailure = errors.New("status code failure on authentication call to api")
	//ErrNoTokenToApply - error response if decorator can not find a token to decorate with
	ErrNoTokenToApply = errors.New("no token to decorate the given request with")
)
