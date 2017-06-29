package env // import "gomatcha.io/matcha/env"

import (
	"gomatcha.io/bridge"
)

func AssetsDir() (string, error) {
	return bridge.Bridge().Call("assetsDir").ToString(), nil
}
