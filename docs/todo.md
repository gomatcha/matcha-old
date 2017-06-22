// matcha.Id is in the matcha package because view depends on layout. and layout needs to use the Ids

High:
* Rewrite gomobile. We want to generate a bunch of xcprojects that the user can then add into an xcworkspace.f
* Switching quickly between navigation item causes loop. 2 quick backs.

Medium:
* Rebuild settings app.
* Documentation.
* Fix middleware.
* Cache layout properties.
* Only send updated views.
* Collect native resources into assets.
* How to prevent cycles when sending messages?? We have two trees that need to be kept in sync. The native tree and the go tree.
* Animations
* remove global Middleware list
* Add PaintStyle back as a property for all views. so as to make painters easier to set.
* Build website and documentation.

Low:
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
* Slider
* AlertView
* Modal presentation
* Add preload, and prepreload stages
* Support StyledText in textinput and textview.
* Asset catalog
* StackBar height / hidden
* Debug constraints.

Bugs:
* Crash in protobuf if view does not have a layout guide.
* crash if constraint layouter adds a view that is not in the Children slice
* Constraints crash if unconstrained.
* Auto disable PNGCrush. "Compress PNG Files" and "Remove Text Metadata from PNG Files"
* Should we panic if user tries to unnotify with an unknown comm.Id

* Add attribution link to https://icons8.com


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
* More Touch Recognizers: Pan, Swipe, Pinch, EdgePan, Rotation
* GPS
* Accessibility
