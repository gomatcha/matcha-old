package env

import (
	"github.com/overcyn/matchabridge"
)

func AssetsDir() (string, error) {
	return matchabridge.Bridge().Call("assetsDir").ToString(), nil
}
