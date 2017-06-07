// Comm.SyncNotifier SyncNotify() chan struct{}
// Comm.Notifier Notify() chan-> struct{}
// mochi.Id is in the mochi package because view depends on layout. and layout needs to use the Ids

// Why do all views have to expose a Painter object? Is there a better way?
func WithPainter(v view.View, p paint.Painter) view.View {
}
func WithValues(v view.View, vals map[interface{}]interface{}) view.View {
}

High:
* Text Input / Keyboard
* Routing / Navs
* Thread & locking. Switch to syncnotifier for all things views.

Medium:
* Documentation
* Examples. Start rebuild a few apps. Instagram, Settings, Slack
* Rewrite gomobile. We want to generate a bunch of xcprojects that the user can then add into an xcworkspace.
* Cache layout properties
* Images/ data blobs

Low:
* Middleware for uitabbaritem and uinavigationitem
* Responder Chain
* Rotation
* Accessibility
* Debugging
* Deadlock detection
* LocalStorage / Keychain / UserDefaults
* Cliboard
* Video / Sound / Microphone / Accelerometer
* remove global Middleware list
* GridView
* Picker
* ProgressBar
* SegmentedControl
* Slider
* TextInput
* Webview
* AlertView
* ActionSheet
* CameraView
* MapView

Bugs:
* Prevent duplicate keys.
* Crash in protobuf if view does not have a layout guide.
* Constraints crash if unconstrained.

Done:
* Lifecycle
* Animations
* UIView tree updating
* Switch to protobufs.
* Event Handling / Gestures
* Navigation View Controllers
* UINavigationController
* UITabViewController