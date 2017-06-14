<!-- type Notifier interface {
    int64 Notify() func()
    Unnotify(int64)
} -->
// mochi.Id is in the mochi package because view depends on layout. and layout needs to use the Ids

High:
* Text Input / Keyboard
* Rebuild settings app.

Medium:
* Thread & locking. Switch to closures for notifiers
* Documentation
* Rewrite gomobile. We want to generate a bunch of xcprojects that the user can then add into an xcworkspace.
* Cache layout properties
* faster transferring of Images/ data blobs
* Collect native resources into assets.

Low:
* Examples. Start rebuild a few apps. Pintrest, Instagram, Settings, Slack
* Constraints should force views onto pixel boundries
* Animations
* uitabbaritem and uinavigationitem should observer their children, so they can update their button/bar
* Responder Chain
* Rotation
* Accessibility
* Debugging
* Asset catalog
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
* More Touch Recognizers: Pan, Swipe, Pinch, EdgePan, Rotation

Bugs:
* Crash in protobuf if view does not have a layout guide.
* crash if constraint layouter adds a view that is not in the Children slice
* Constraints crash if unconstrained.
* Auto disable PNGCrush. "Compress PNG Files" and "Remove Text Metadata from PNG Files"

Done:
* Lifecycle
* UIView tree updating
* Switch to protobufs.
* Event Handling / Gestures
* Navigation View Controllers
* UINavigationController
* UITabViewController
* Routing / Navs
