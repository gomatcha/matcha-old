// matcha.Id is in the matcha package because view depends on layout. and layout needs to use the Ids
go get github.com/matcha/matcha/...

High:
* Switching quickly between navigation item causes loop. 2 quick backs.
* How to prevent cycles when sending messages?? We have two trees that need to be kept in sync. The native tree and the go tree.
* Build website and documentation. Make sure go get "gomatcha.io/..." works.

Medium:
* Add a Notifyall, and notifyBetween to store.
* Rebuild settings app, Todo App.
* Only send updated views.
* Make sure development team isn't packaged with sample app.

Low:
* Have flag that generates a new xcodeproj for easy setup.
* Add tests around core functionality. Store, etc.
* Automatically insert copyright notice.
* Collect native resources into assets.
* StyledText
* Support more flags in matcha command.
* Updating a tabscreen or stackscreen should not trigger a rebuild of its children.
* Webview
* Localization
* Examples. Start rebuild a few apps. Pintrest, Instagram, Settings, Slack
* Constraints should force views onto pixel boundries
* GridView
* Flexbox
* Picker
* TextField
* SegmentedControl
* AlertView
* Modal presentation
* Add preload, and prepreload stages
* Asset catalog
* StackBar height / hidden, color
* Debug constraints.
* Animations: Spring, Delay, Batch, Reverse, Decay, Repeat, 2d, 3d, Nd
* Rotation
* More Touch Recognizers: Pan, Swipe, Pinch, EdgePan, Rotation
* Table scroll position?
* Table ScrollBehaviors, Table Direction
* Custom painters.
* View 3d transforms.

Bugs:
* Crash in protobuf if view does not have a layout guide.
* crash if constraint layouter adds a view that is not in the Children slice
* Constraints crash if unconstrained.
* Auto disable PNGCrush. "Compress PNG Files" and "Remove Text Metadata from PNG Files"
* Should we panic if user tries to unnotify with an unknown comm.Id
* Add attribution link to https://icons8.com

Documentation:
* Build process.
* animate
* cmd
* comm
* store
* docs
* env
* examples
* layout
* text
* touch
* view 
    * basicview
    * button
    * imageview
    * scrollview
    * slider
    * stackscreen
    * switchview
    * tabscrceen
    * textinput
    * textview
    * urlimageview
    * progressview

Pro:
* Debugging
* Deadlock detection
* LocalStorage / Keychain / UserDefaults
* Cliboard
* Notifications
* Video / Sound / Microphone / Accelerometer
* ActionSheet
* CameraView
* MapView
* GPS
* Accessibility
