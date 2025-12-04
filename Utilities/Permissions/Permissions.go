package Permissions

import "Polybub/Auth/OAuth2"

var MFA_CODE_R = OAuth2.NewPerm("MfaCode", false, true, false, false)
var FOOBAR_CRUD = OAuth2.NewPerm("FooBar", true, true, true, true)
var DASHBOARD_R = OAuth2.NewPerm("Dashboard", true, false, false, false)
