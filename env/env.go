package env

import (
	"github.com/overcyn/mochibridge"
)

func AssetsDir() (string, error) {
	return mochibridge.Bridge().Call("assetsDir").ToString(), nil
}
