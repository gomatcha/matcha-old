High:
* Switching quickly between navigation item causes loop. 2 quick backs.
* How to prevent cycles when sending messages?? We have two trees that need to be kept in sync. The native tree and the go tree.
* Build website and documentation. Make sure go get "gomatcha.io/..." works.

Medium:
* Rebuild settings app, Todo App.
* Only send updated views.
* Make sure development team isn't packaged with sample app.
* Table scroll position?
* User registered views.

Low:
* Have matcha flag that generates a new xcodeproj for easy setup.
* Add tests around core functionality. Store, etc.
* Updating a tabscreen should not trigger a rebuild of its children.
* Webview
* Examples. Start rebuild a few apps. Pintrest, Instagram, Settings, Slack
* Constraints should force views onto pixel boundries
* Flexbox
* Picker
* TextField
* SegmentedControl
* AlertView
* Modal presentation
* Asset catalog
* StackBar height / hidden, color
* Rotation
* More Touch Recognizers: Pan, Swipe, Pinch, EdgePan, Rotation
* Table ScrollBehaviors, Table Direction
* Custom painters.
* Compile a list of things that should be easy to do and implement them. Button activation cancelled by vertical scrolling but not horizontal, Pinch to zoom, Highlighting a view and dragging outside of it and back in., Horizontal swipe on tableview to show delete button, Touch driven animations. AKA swipe back to navigate.

Very Low:
* Automatically insert copyright notice.
* StyledText
* Text selection.
* Localization
* View 3d transforms.
* GridView
* Add preload, and prepreload stages
* Debug constraints.
* Collect native resources into assets.
* Animations: Spring, Delay, Batch, Reverse, Decay, Repeat

Bugs:
* Crash in protobuf if view does not have a layout guide.
* Constraints crash if unconstrained.
* Auto disable PNGCrush. "Compress PNG Files" and "Remove Text Metadata from PNG Files"
* Should we panic if user tries to unnotify with an unknown comm.Id

Documentation:
* Build process.
* cmd
* comm
* store
* docs
* env
* examples
* layout
* view 
    * scrollview
    * slider
    * stackscreen
    * switchview
    * tabscrceen
    * textinput
    * textview
    * urlimageview
    * progressview
    * segmentview

Pro:
* Debugging
* LocalStorage / Keychain / UserDefaults
* Cliboard
* Notifications
* Video / Sound / Microphone / Accelerometer
* ActionSheet
* CameraView
* MapView
* GPS
* Accessibility
