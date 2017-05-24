Prevent duplicate keys.
Install Protobuf

// Comm.SyncNotifier SyncNotify() chan struct{}
// Comm.Notifier Notify() chan-> struct{}

High:
* Event Handling / Gestures
* Navigation View Controllers
* Switch to protobufs.

Medium:
* Thread & locking. How to prevent deadlocks?
* Lock should go bottom to top

Low:
* Responder Chain
* Rotation
* Accessibility
* Debugging
* LocalStorage / Keychain / UserDefaults
* Cliboard
* Video / Sound / Microphone / Accelerometer

Done:
* Lifecycle
* Animations
* UIView tree updating

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