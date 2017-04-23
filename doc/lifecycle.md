const (
    Unattached
    Mounted // 
    Visible
    Offscreen
    
    // Dead Stage = iota // nothing
    // BeforeMounted
    // Mounted // in a view somewhere
    // Visible // in a view on screen
    
    //nsapplication
    dead
    active
    launching
    
    //Vue
    BeforeCreated
    Created
    BeforeMount
    Mounted
    BeforeDestroy
    Destroyed
    
    //asyncdisplaykit
    Prepreload
    Preload
    Visible
    
    //
    Dead
    Alive
    Visible
    Focused
    
    // LoadView
    // ViewDidLoad
    // ViewWillAppear
    // ViewDidAppear
    // ViewWillDisappear
    // ViewDidDisappear
    // ViewDidUnload
    
    
    EventLoaded
    StageMounted
    StageVisible
)
