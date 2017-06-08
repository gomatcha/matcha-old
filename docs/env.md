NSBundle *bundle = [NSBundle bundleWithURL:[[NSBundle mainBundle] URLForResource:@"YourBundle" withExtension:@"bundle"]];
NSString *imagePath = [bundle pathForResource:@"imageInBundle" ofType:@"png"];
UIImage *image = [UIImage imageWithContentsOfFile:imagePath];


UIDevice 
// OS Version
// OS Name
// Device Name
// Device size (iPad, iPhone)

UIApplication
// Window
// application state
// background state
// app badge
// open url 
// local notifications
// status bar style
// status bar orientation
// Application path? bundle path?

UIScreen
* mainScreen
* Coordinate space
* bounds
* scale
* Brightness

UIWindow
....

NSBundle
// 

Android assets vs resources.