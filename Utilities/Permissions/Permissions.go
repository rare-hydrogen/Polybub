package Permissions

import "Polybub/Auth/OAuth2"

var MFA_CODE_CRU = OAuth2.NewPerm("MfaCode", true, true, true, false)
var FOOBAR_CRUD = OAuth2.NewPerm("FooBar", true, true, true, true)
var DASHBOARD_R = OAuth2.NewPerm("Dashboard", true, false, false, false)
