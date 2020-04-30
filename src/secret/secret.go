package secret

import "github.com/gbrlsnchs/jwt/v3"

var AppKey = jwt.NewHS256([]byte("HI_THIS_IS TEST"))

