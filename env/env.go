package env

import (
	"gomatcha.io/bridge"
)

func AssetsDir() (string, error) {
	return matchabridge.Bridge().Call("assetsDir").ToString(), nil
}
