Present Modally
TabBarController
NavigationController
Hamberger drawer
Popovers?

- TabBar
    - Navigation
    - Navigation
    - Navigation (Selected Tab)
        - CurrentPage
    - Navigation
        - Page
        - Page
        - Page
        - Page
    - Navigation
- Presented

UIPageViewController
UISplitViewController
UITableViewController
EKEventViewController
SFSafariViewController
CNContactViewController
MCBrowserViewController
RPPreviewViewController
SKStoreReviewController
UIActivityViewController
GKAchievementViewController
MPMoviePlayerViewController

self.navigationController > This searches for the closest navigationController above us.
Alternately. Pass in the navigation controller.
self.navigation.pop(animated)

self.navigation.ShowDrawer
self.navigation.HideDrawer
self.navigation.SetDrawer

func GetSidebar(v navigator) {
}

sidebar.Get(nav).Show()
sidebar.Get(nav).Hide()

type navigator struct {
}

type sidebarview struct {
}

type tabview struct {
}

type stackview struct {
}


type Screen interface {
    NewView(*view.Context, interface{}) view.View
}

type RootView struct {
    *view.Embed
    nav *Nav
}

func NewRootView(ctx *view.Context, key interface{}) *RootView {
    if v, ok := ctx.Prev(key).(*RootView); ok {
        return v
    }

    // homeView := stacknav.New(nil, nil)
    // homeView.Model.SetScreens([]*stacknav.Screen{
    //  stacknav.Screen{
    //      View:  view1,
    //      Title: "Home",
    //  },
    // })

    // searchView := stacknav.New(nil, nil)
    // searchView.Model.SetScreens([]*stacknav.Screen{
    //  stacknav.Screen{
    //      Screen: view.Screenfunc(func(ctx *view.Context, key interface{}) view.View {
    //          return NewNestedView(ctx, key)
    //      }),
    //      Title: "Search",
    //  },
    // })

    // homeStack := &stacknav.Model{}
    // homeStack.SetScreens([]*view.Screen{
    //  &homeview.Screen{},
    //  // stacknav.Screen{
    //  //  View:  view1,
    //  //  Title: "Home",
    //  // },
    // })

    // searchScreen := &searchview.Screen{}
    // searchScreen.SetQuery("Default Query")

    // searchStack := &stacknav.Screen{}
    // searchStack.SetBackgroundColor()
    // searchStack.SetScreens([]*view.Screen{
    //  searchScreen,
    //  // view.ScreenFunc(func(ctx *view.Context, key interface{}) view.View {
    //  //  return searchpage.New(ctx, key, searchScreen)
    //  // }),
    // })

    homeStack := &stacknav.Model{}
    homeStack.SetScreens([]*view.Screen{
        &homeview.Screen{},
    })

    searchScreen := &searchview.Screen{}
    searchScreen.SetQuery("Default Query")

    searchStack := &stacknav.Screen{}
    searchStack.SetBackgroundColor()
    searchStack.SetScreens([]view.Screen{
        searchScreen,
        
        // view.ScreenFunc(func(ctx *view.Context, key interface{}) view.View {
        //  return searchpage.New(ctx, key, searchScreen)
        // }),
    })

    tab := &tabview.Screen{}
    tab.SetChildScreens([]view.Screen{homeStack, searchStack})

    // captureStack := &stacknav.Model{}
    // captureStack.SetScreens([]*stacknav.Screen{
    //  stacknav.Screen{
    //      View:  view1,
    //      Title: "Camera",
    //  },
    // })

    // likedStack := &stacknav.Model{}
    // likedStack.SetScreens([]*stacknav.Screen{
    //  stacknav.Screen{
    //      View:  view1,
    //      Title: "Camera",
    //  },
    // })

    // profileStack := stacknav.Model{}
    // profileStack.SetScreens([]*stacknav.Screen{
    //  stacknav.Screen{
    //      View:  view1,
    //      Title: "Camera",
    //  },
    // })

    // tab := &tabview.Model{}
    // tab.SetScreens([]*tabview.Screen{
    //  tabview.Screen{
    //      // Screen: nav.homestack, // what if you need to customize more properties?
    //      Screen: view.Screenfunc(func(ctx *view.Context, key interface{}) view.View {
    //          stack := stackview.New(ctx, key)
    //          stack.Model = nav.homeStack
    //          return stack
    //      }),
    //      // View:         stack1,
    //      Title:        title,
    //      Icon:         icon,
    //      SelectedIcon: selIcon,
    //  },
    //  tabview.Screen{
    //      View:         stack2,
    //      Title:        title,
    //      Icon:         icon,
    //      SelectedIcon: selIcon,
    //  },
    // }, tx)

    tab := &tabview.Screen{}
    tab.SetChildScreens([]view.Screen{
        homeStack, searchStack,
        // view.ScreenFunc(func(ctx *view.Context, key interface{}) view.View {
        //  stack := stackview.New(ctx, key)
        //  stack.Model = nav.homeStack
        //  return stack
        // }),
        // view.ScreenFunc(func(ctx *view.Context, key interface{}) view.View {
        //  stack := stackview.New(ctx, key)
        //  stack.Model = nav.searchStack
        //  return stack
        // }),
    })

    return &RootView{
        Embed: view.NewEmbed(ctx.NewId(key)),
    }
}

// type View interface {
//  Build(*Context) *Model
//  Lifecycle(from, to Stage)
//  Id() matcha.Id
//  sync.Locker
//  matcha.Notifier
// }

// // type Nav interface {
// //   Store() store.Store
// // }

type Nav struct {
    store.Store
    RootNav    *tabview.TabView
    HomeNav    *stacknav.StackNav
    SearchNav  *stacknav.StackNav
    CaptureNav *stacknav.StackNav
    LikedNav   *stacknav.StackNav
    ProfileNav *stacknav.StackNav
}

// How does the view get the global object?

// I need to monitor for changes to the navigationitem/tabbaritem/
// Do we want screens to all carry around a Nav object?
// Should all views be able to easily present an overlay? without access to the nav object?
// Navs are container views. That don't/can't rebuild their children.

func (v *TabView) Build(ctx *view.Context) *view.Model {
    l := constraint.New()

    tab := tabnav.New(ctx, 0)
    tab.Store = v.Nav.Tab
    // stack1 := stacknav.New(ctx, 1)
    // stack2 := stacknav.New(ctx, 2)
    // stack3 := stacknav.New(ctx, 3)

    // tx := store.NewWriteTx()
    // defer tx.Commit()

    // v.Nav.SetRootNav(tab.Nav, tx)
    // v.Nav.RootNav().SetScreens([]*tabnav.Screen{
    //  tabnav.Screen{
    //      View:         stack1,
    //      Title:        title,
    //      Icon:         icon,
    //      SelectedIcon: selIcon,
    //  },
    //  tabnav.Screen{
    //      View:         stack2,
    //      Title:        title,
    //      Icon:         icon,
    //      SelectedIcon: selIcon,
    //  },
    // }, tx)

    // v.Nav.SetHomeNav(stack1.Nav(), tx)
    // stack1.Nav().SetScreens([]*stacknav.Screen{
    //  stacknav.Screen{
    //      View:  view1,
    //      Title: "blah",
    //  },
    // }, tx)

    // v.Nav.SetSearchNav(stack2.Nav, tx)
    // v.Nav.SetCaptureNav(stack3.Nav, tx)

    // tab.SetScreens([]*tabnav.Screen{
    //  tabnav.Screen{
    //      Nav: stack1,
    //      // View:  stack1,
    //      // Store: stack1.Store(),
    //      Title: title,
    //  },
    //  tabnav.Screen{
    //      Nav:          stack2,
    //      Icon:         icon,
    //      SelectedIcon: selIcon,
    //      // View:  stack2,
    //      // Store: stack2.Store(),
    //      Title: title,
    //  }}, tx)

    // tab.Store().SetScreens([]*tabnav.Screen{
    //  tabnav.Screen{
    //      Nav: stack1,
    //      // View:  stack1,
    //      // Store: stack1.Store(),
    //      Title: title,
    //  },
    //  tabnav.Screen{
    //      Nav:          stack2,
    //      Icon:         icon,
    //      SelectedIcon: selIcon,
    //      // View:  stack2,
    //      // Store: stack2.Store(),
    //      Title: title,
    //  }}, tx)

    stack1.Store().SetScreens([]*stacknav.Screen{view1, view2, view3}, tx)

    // stack1.Store().SetScreens([]*stacknav.Screen{view1, view2, view3}, tx)

    // stack1.Store().SetScreens([]*stacknav.Screen{
    // stacknav.Screen{
    //  View:  view1,
    //  Title: "blah",
    // },
    // stacknav.Screen{
    //  View:  view2,
    //  Title: "blah2",
    // }}, tx)
    tx.Commit()

    childView.OnClick = func() {
        v.Lock()
        defer v.Unlock()

        tx := store.NewWriteTx()
        defer tx.Commit()
        // v.Nav().Push(NewTableView(nil, nil), tx)
        v.nav.stacknav.Push(v.nav, NewTableView(nil, nil))

        // stack, ok := stacknav.Get(ctx)
        // if !ok {
        //  return
        // }
        // stack.Push(NewTableView(nil, nil), animated, tx)
    }

    // stacknav.Push(nav,

    // view1 := NewNestedView(ctx, 2)

    // stackscreen1 := &stacknav.Screen{}
    // stackscreen1.SetView(view1)
    // stackscreen1.SetTitle("stack title")

    // stack1 := stacknav.New(ctx, 1)
    // stack1.Push(stackscreen1)

    // tab1 := &tabnav.Screen{}
    // tab1.SetView(stack1)
    // tab1.SetTitle("Tab 1")

    // view2 := NewTableView(ctx, "view2")

    // stackscreen2 := &stacknav.Screen{}
    // stackscreen2.SetView(view2)
    // stackscreen2.SetTitle("Table")

    // stack2 := stacknav.New(ctx, "stack2")
    // stack2.Push(stackscreen2)

    // tab2 := &tabnav.Screen{}
    // tab2.SetView(stack2)
    // tab2.SetTitle("Tab 2")

    // tab := tabnav.New(ctx, 100)
    // tab.SetScreens([]*tabnav.Screen{tab1, tab2})
    // l.Add(tab, func(s *constraint.Solver) {
    //  s.WidthEqual(l.Width())
    //  s.HeightEqual(l.Height())
    // })

    return &view.Model{
        Children: []view.View{tab},
        Layouter: l,
        Painter:  &paint.Style{BackgroundColor: colornames.Green},
    }
}