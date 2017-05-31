Prevent duplicate keys.
Root background color isn't set.

// Comm.SyncNotifier SyncNotify() chan struct{}
// Comm.Notifier Notify() chan-> struct{}

High:
* Navigation View Controllers

Medium:
* Thread & locking. How to prevent deadlocks? Notifiers should be eventually consistent.
* Lock should go bottom to top

Low:
* Responder Chain
* Rotation
* Accessibility
* Debugging
* LocalStorage / Keychain / UserDefaults
* Cliboard
* Video / Sound / Microphone / Accelerometer
* Cache layout properties

Done:
* Lifecycle
* Animations
* UIView tree updating
* Switch to protobufs.
* Event Handling / Gestures

* GridView
* UINavigationController
* UITabViewController
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


    <!-- return &view.Model{
        Children: []View{chl},
        Layouter: l,
        Painter:  &paint.Style{BackgroundColor: colornames.Green},
        Values: map[interface{}]interface{}{
            touch.Key(): []touch.Recognizer{tap},
        },
    } -->
    
We want to generate a bunch of xcprojects that the user can then add into an xcworkspace.

Accessibility
Touch
Mouse
Keyboard
SupportedRotations