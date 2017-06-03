Prevent duplicate keys.
* Crash in protobuf if view does not have a layout guide.
* Constraints crash if unconstrained.

// Comm.SyncNotifier SyncNotify() chan struct{}
// Comm.Notifier Notify() chan-> struct{}
// mochi.Id is in the mochi package because view depends on layout. and layout needs to use the Ids

// Why do all views have to expose a Painter object? Is there a better way?
func WithPainter(v view.View, p paint.Painter) view.View {
}
func WithValues(v view.View, vals map[interface{}]interface{}) view.View {
}

High:
* Text Input

Medium:
* Thread & locking. How to prevent deadlocks? Notifiers should be eventually consistent.
* Lock should go bottom to top
* Documentation
* Examples
* Rewrite gomobile

Low:
* Responder Chain
* Rotation
* Accessibility
* Debugging
* Deadlock detection
* LocalStorage / Keychain / UserDefaults
* Cliboard
* Video / Sound / Microphone / Accelerometer
* Cache layout properties

* remove global Middleware list


Done:
* Lifecycle
* Animations
* UIView tree updating
* Switch to protobufs.
* Event Handling / Gestures
* Navigation View Controllers
* UINavigationController
* UITabViewController

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
    
We want to generate a bunch of xcprojects that the user can then add into an xcworkspace.

Accessibility
Touch
Mouse
Keyboard
SupportedRotations