package Permissions

import "Polybub/Auth/OAuth2"

// Basic system perms allowing users to attempt to authenticate with MFA
var MFA_CODE_R = OAuth2.NewPerm("MfaCode", false, true, false, false)

// Other permissions
var MFA_CODE_CRU = OAuth2.NewPerm("MfaCode", true, true, true, false)
var DASHBOARD_R = OAuth2.NewPerm("Dashboard", true, false, false, false)
var USERS_CRUD = OAuth2.NewPerm("Users", true, true, true, true)
var FOOBARS_CRUD = OAuth2.NewPerm("FooBars", true, true, true, true)
