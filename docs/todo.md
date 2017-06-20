// matcha.Id is in the matcha package because view depends on layout. and layout needs to use the Ids

High:
* Rewrite gomobile. We want to generate a bunch of xcprojects that the user can then add into an xcworkspace.
* UINavigation item button. uitabbar item image
* uitabbaritem and uinavigationitem should observer their children, so they can update their button/bar
* UINavigationItem cycle.
* How to handle image scale


Medium:
* Rebuild settings app.
* Documentation
* Fix middleware.
* Cache layout properties
* Faster transferring of Images/ data blobs
* Collect native resources into assets.
* How to prevent cycles when sending messages?? We have two trees that need to be kept in sync. The native tree and the go tree.
* Animations
* remove global Middleware list
* Webview
* Define constraints
* Make painters easier to set.

Low:
* Localization
* Examples. Start rebuild a few apps. Pintrest, Instagram, Settings, Slack
* Constraints should force views onto pixel boundries
* Rotation
* Accessibility
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