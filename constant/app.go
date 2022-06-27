package constant

import "os"

var dbSecret = os.Getenv("DB_SECRET")
var SECRET_JWT = dbSecret
