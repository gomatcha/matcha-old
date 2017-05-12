package store

import (
	"github.com/overcyn/mochi"
)

func Notifier(f func(tx *Tx)) mochi.Notifier {
	return nil
}

// func InterfaceNotifier(f func(tx *Tx) interface{}) mochi.InterfaceNotifier {
// 	return nil
// }

// func BoolNotifier(f func(tx *Tx) bool) mochi.BoolNotifier {
// 	return nil
// }

// func IntNotifier(f func(tx *Tx) int) mochi.IntNotifier {
// 	return nil
// }

// func UintNotifier(f func(tx *Tx) uint) mochi.UintNotifier {
// 	return nil
// }

// func Int64Notifier(f func(tx *Tx) int64) mochi.Int64Notifier {
// 	return nil
// }

// func Uint64Notifier(f func(tx *Tx) uint64) mochi.Uint64Notifier {
// 	return nil
// }

// func Float64Notifier(f func(tx *Tx) float64) mochi.Float64Notifier {
// 	return nil
// }

// func StringNotifier(f func(tx *Tx) string) mochi.StringNotifier {
//  return nil
// }

// func BytesNotifier(f func(tx *Tx) []byte) mochi.ByteNotifier {
// 	return nil
// }
