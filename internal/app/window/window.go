type Window { // direct manage of Window, opening, closing, event
	handle *glfw.Window // native window, handle is the arrow?
	app AppInterface // app interface
	// mousecallbacks and stuff
}


type AppInterface { // dependency injection - window is calling app's methods
	Update{}
	Render{}
	ShouldClose{} bool
}