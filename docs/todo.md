// Comm.SyncNotifier SyncNotify() chan struct{}
// Comm.Notifier Notify() chan-> struct{}
// mochi.Id is in the mochi package because view depends on layout. and layout needs to use the Ids
// Remove view.model.children, Just get the list of child views from the layouter.

High:
* Text Input / Keyboard
* Thread & locking. Switch to syncnotifier for all things views.
* Native images.
* Examples. Start rebuild a few apps. Pintrest, Instagram, Settings, Slack

Medium:
* Documentation
* Rewrite gomobile. We want to generate a bunch of xcprojects that the user can then add into an xcworkspace.
* Cache layout properties
* Images/ data blobs

Low:
* Animations
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
* Modal presentation

Bugs:
* Prevent duplicate keys.
* Crash in protobuf if view does not have a layout guide.
* Constraints crash if unconstrained.

Done:
* Lifecycle
* UIView tree updating
* Switch to protobufs.
* Event Handling / Gestures
* Navigation View Controllers
* UINavigationController
* UITabViewController
* Routing / Navs