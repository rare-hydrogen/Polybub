package Permissions

import "Polybub/Auth/OAuth2"

var FOOBAR_CRUD = OAuth2.NewPerm("FooBar", true, true, true, true)
var DASHBOARD_R = OAuth2.NewPerm("Dashboard", true, false, false, false)
