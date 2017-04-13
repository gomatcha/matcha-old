# Bridge

Mochi need some way of connecting the iOS/Android app with Go code. 

## Mochi

This is our fork of Gomobile to better understand it.

    go install golang.org/x/mobile/cmd/gomobile
    gomobile init
    gomobile bind -target=ios golang.org/x/mobile/example/bind/hello

	-n      print the commands but do not run them.
    -x      print the commands
    -v      print the names of packages as they are compiled.
    -work   print the name of the temporary work directory and do not delete it when exiting.
    -o      output
/var/folders/c3/8xb38pfj2kvg1c87nfdq8h140000gn/T/gomobile-work-618387728


## Gomobile

Go provides some basic tools for binding in [Go Mobile](golang.org/x/mobile). There are two methods of using `gomobile` to build a Go backed application. 

* Building an all Go mobile application.
* Calling into Go code through generated bindings from Java and Objective-C.

Mochi will be focusing on the second option, as developers will likely need to fallback into "native" code for various reasons. Integration into preexisting codebases, using non-go libraries, or needing deeper control than Mochi provides. To build bindings, `gomobile` takes a specified package and generates a framework with translated ObjC headers that can be added to XCode project. The exported APIs are limited to a subset of types that includes:

* Signed integer and floating point types.
* String and boolean types.
* Byte slice types. Note that byte slices are passed by reference, and support mutation.
* Any function type all of whose parameters and results have supported types. Functions must return either no results, one result, or two results where the type of the second is the built-in 'error' type.
* Any interface type, all of whose exported methods have supported function types.
* Any struct type, all of whose exported methods have supported function types and all of whose exported fields have supported types.

Note that maps and slices are unexported. And any types that are declared outside of the package are not included.

Mobile Wiki
https://github.com/golang/go/wiki/Mobile

Expose Java to Gomobile
https://github.com/golang/go/issues/16876

Expose ObjC to Gomobile
https://github.com/golang/go/issues/17102

Binding Go and Java
https://docs.google.com/document/d/1y9hStonl9wpj-5VM-xWrSTuEJFUAxGOXOhxvAs7GZHE/edit

## JavaScriptCore

JavaScriptCore is a JavaScript engine for Apple devices introduced in iOS7. It includes bindings for Swift and ObjectiveC, and is used for many projects including ReactNative and AppStore/iTunesStore. JavaScript is a weakly typed language, so it manages to translate well into Objective C. The type conversions are as follows.

| ObjC Type                        | JavaScript Type | Notes                                                             |
| -------------------------------- | --------------- | ----------------------------------------------------------------- |
| nil                              | undefined       |                                                                   |
| NSNull                           | null            |                                                                   |
| NSString                         | String          |                                                                   |
| NSNumber                         | Number, Boolean |                                                                   |
| NSDictionary                     | Object          | Recursive                                                         |
| NSArray                          | Array           | Recursive                                                         |
| NSDate                           | Date            |                                                                   |
| Object, Class                    | Object          | Object properties and methods are not automatically exposed to JS |
| NSRange, CGRect, CGPoint, CGSize | Object          |                                                                   |
| Block                            | Function        | Blocks can convert to JS functions but not the other way around   |

Object methods and properties can be exposed to the JavaScript runtime using the JSExport protocol. Each class must create a new protocol that conforms to JSExport, and place the desired methods within it.

```
@protocol MyPointExports <JSExport>
@property double x;
@property double y;
- (NSString *)description;
- (instancetype)initWithX:(double)x y:(double)y;
+ (MyPoint *)makePointWithX:(double)x y:(double)y;
@end
@interface MyPoint : NSObject <MyPointExports>
- (void)myPrivateMethod;  // Not in the MyPointExports protocol, so not visible to JavaScript code.
@end
@implementation MyPoint
// ...
@end
```

This is a fairly clean approach for flagging functions but it is unclear how it would translate to Go.

## CGo 

...

