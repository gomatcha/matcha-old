// matcha.Id is in the matcha package because view depends on layout. and layout needs to use the Ids
go get github.com/matcha/matcha/...

High:
* We want to generate an bunch of xcprojects that the user can then add into an xcworkspace.
* Switching quickly between navigation item causes loop. 2 quick backs.
* Build objective c files into the framework.
* Collect native resources into assets.
* Build website and documentation.
* How to prevent cycles when sending messages?? We have two trees that need to be kept in sync. The native tree and the go tree.
* Hide StyledText
* Automatically insert copyright notice.

Medium:
* Why doesn't simulator debugger work? Why does simulator crash all the time?
* Rebuild settings app.
* Only send updated views.

Low:
* Support StyledText in textinput and textview.
* Support more flags in matcha command.
* Updating a tabscreen or stackscreen should not trigger a rebuild of its children.
* Webview
* Localization
* Examples. Start rebuild a few apps. Pintrest, Instagram, Settings, Slack
* Constraints should force views onto pixel boundries
* Rotation
* GridView
* Flexbox
* Picker
* ProgressBar
* SegmentedControl
* AlertView
* Modal presentation
* Add preload, and prepreload stages
* Asset catalog
* StackBar height / hidden
* Debug constraints.
* Animations: Spring, Delay, Batch, Reverse, Decay, Repeat, 2d, 3d, Nd
* More Touch Recognizers: Pan, Swipe, Pinch, EdgePan, Rotation

Bugs:
* Crash in protobuf if view does not have a layout guide.
* crash if constraint layouter adds a view that is not in the Children slice
* Constraints crash if unconstrained.
* Auto disable PNGCrush. "Compress PNG Files" and "Remove Text Metadata from PNG Files"
* Should we panic if user tries to unnotify with an unknown comm.Id
* Add attribution link to https://icons8.com

Documentation:
* animate
* cmd
* comm
* docs
* env
* examples
* ios
* keyboard
* layout
    * constraint
    * full
    * table
* paint
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
