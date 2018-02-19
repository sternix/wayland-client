// package xdg acts as a client for the xdg_shell wayland protocol.

// generated by wl-scanner
// https://github.com/dkolbly/wl-scanner
// from: https://raw.githubusercontent.com/wayland-project/wayland-protocols/master/stable/xdg-shell/xdg-shell.xml
// on 2018-02-19 10:39:29 -0600
package xdg

import (
	"github.com/dkolbly/wl"
	"sync"
)

type WmBasePingEvent struct {
	Serial uint32
}

func (p *WmBase) AddPingHandler(h wl.Handler) {
	if h != nil {
		p.mu.Lock()
		p.pingHandlers = append(p.pingHandlers, h)
		p.mu.Unlock()
	}
}

func (p *WmBase) RemovePingHandler(h wl.Handler) {
	p.mu.Lock()
	defer p.mu.Unlock()

	for i, e := range p.pingHandlers {
		if e == h {
			p.pingHandlers = append(p.pingHandlers[:i], p.pingHandlers[i+1:]...)
			break
		}
	}
}

func (p *WmBase) Dispatch(event *wl.Event) {
	switch event.Opcode {
	case 0:
		if len(p.pingHandlers) > 0 {
			ev := WmBasePingEvent{}
			ev.Serial = event.Uint32()
			p.mu.RLock()
			for _, h := range p.pingHandlers {
				h.Handle(ev)
			}
			p.mu.RUnlock()
		}
	}
}

type WmBase struct {
	wl.BaseProxy
	mu           sync.RWMutex
	pingHandlers []wl.Handler
}

func NewWmBase(ctx *wl.Context) *WmBase {
	ret := new(WmBase)
	ctx.Register(ret)
	return ret
}

// Destroy will destroy xdg_wm_base.
//
//
// Destroy this xdg_wm_base object.
//
// Destroying a bound xdg_wm_base object while there are surfaces
// still alive created by this xdg_wm_base object instance is illegal
// and will result in a protocol error.
//
func (p *WmBase) Destroy() error {
	return p.Context().SendRequest(p, 0)
}

// CreatePositioner will create a positioner object.
//
//
// Create a positioner object. A positioner object is used to position
// surfaces relative to some parent surface. See the interface description
// and xdg_surface.get_popup for details.
//
func (p *WmBase) CreatePositioner() (*Positioner, error) {
	ret := NewPositioner(p.Context())
	return ret, p.Context().SendRequest(p, 1, wl.Proxy(ret))
}

// GetXdgSurface will create a shell surface from a surface.
//
//
// This creates an xdg_surface for the given surface. While xdg_surface
// itself is not a role, the corresponding surface may only be assigned
// a role extending xdg_surface, such as xdg_toplevel or xdg_popup.
//
// This creates an xdg_surface for the given surface. An xdg_surface is
// used as basis to define a role to a given surface, such as xdg_toplevel
// or xdg_popup. It also manages functionality shared between xdg_surface
// based surface roles.
//
// See the documentation of xdg_surface for more details about what an
// xdg_surface is and how it is used.
//
func (p *WmBase) GetXdgSurface(surface *wl.Surface) (*Surface, error) {
	ret := NewSurface(p.Context())
	return ret, p.Context().SendRequest(p, 2, wl.Proxy(ret), surface)
}

// Pong will respond to a ping event.
//
//
// A client must respond to a ping event with a pong request or
// the client may be deemed unresponsive. See xdg_wm_base.ping.
//
func (p *WmBase) Pong(serial uint32) error {
	return p.Context().SendRequest(p, 3, serial)
}

const (
	WmBaseErrorRole                = 0
	WmBaseErrorDefunctSurfaces     = 1
	WmBaseErrorNotTheTopmostPopup  = 2
	WmBaseErrorInvalidPopupParent  = 3
	WmBaseErrorInvalidSurfaceState = 4
	WmBaseErrorInvalidPositioner   = 5
)

type Positioner struct {
	wl.BaseProxy
}

func NewPositioner(ctx *wl.Context) *Positioner {
	ret := new(Positioner)
	ctx.Register(ret)
	return ret
}

// Destroy will destroy the xdg_positioner object.
//
//
// Notify the compositor that the xdg_positioner will no longer be used.
//
func (p *Positioner) Destroy() error {
	return p.Context().SendRequest(p, 0)
}

// SetSize will set the size of the to-be positioned rectangle.
//
//
// Set the size of the surface that is to be positioned with the positioner
// object. The size is in surface-local coordinates and corresponds to the
// window geometry. See xdg_surface.set_window_geometry.
//
// If a zero or negative size is set the invalid_input error is raised.
//
func (p *Positioner) SetSize(width int32, height int32) error {
	return p.Context().SendRequest(p, 1, width, height)
}

// SetAnchorRect will set the anchor rectangle within the parent surface.
//
//
// Specify the anchor rectangle within the parent surface that the child
// surface will be placed relative to. The rectangle is relative to the
// window geometry as defined by xdg_surface.set_window_geometry of the
// parent surface.
//
// When the xdg_positioner object is used to position a child surface, the
// anchor rectangle may not extend outside the window geometry of the
// positioned child's parent surface.
//
// If a negative size is set the invalid_input error is raised.
//
func (p *Positioner) SetAnchorRect(x int32, y int32, width int32, height int32) error {
	return p.Context().SendRequest(p, 2, x, y, width, height)
}

// SetAnchor will set anchor rectangle anchor.
//
//
// Defines the anchor point for the anchor rectangle. The specified anchor
// is used derive an anchor point that the child surface will be
// positioned relative to. If a corner anchor is set (e.g. 'top_left' or
// 'bottom_right'), the anchor point will be at the specified corner;
// otherwise, the derived anchor point will be centered on the specified
// edge, or in the center of the anchor rectangle if no edge is specified.
//
func (p *Positioner) SetAnchor(anchor uint32) error {
	return p.Context().SendRequest(p, 3, anchor)
}

// SetGravity will set child surface gravity.
//
//
// Defines in what direction a surface should be positioned, relative to
// the anchor point of the parent surface. If a corner gravity is
// specified (e.g. 'bottom_right' or 'top_left'), then the child surface
// will be placed towards the specified gravity; otherwise, the child
// surface will be centered over the anchor point on any axis that had no
// gravity specified.
//
func (p *Positioner) SetGravity(gravity uint32) error {
	return p.Context().SendRequest(p, 4, gravity)
}

// SetConstraintAdjustment will set the adjustment to be done when constrained.
//
//
// Specify how the window should be positioned if the originally intended
// position caused the surface to be constrained, meaning at least
// partially outside positioning boundaries set by the compositor. The
// adjustment is set by constructing a bitmask describing the adjustment to
// be made when the surface is constrained on that axis.
//
// If no bit for one axis is set, the compositor will assume that the child
// surface should not change its position on that axis when constrained.
//
// If more than one bit for one axis is set, the order of how adjustments
// are applied is specified in the corresponding adjustment descriptions.
//
// The default adjustment is none.
//
func (p *Positioner) SetConstraintAdjustment(constraint_adjustment uint32) error {
	return p.Context().SendRequest(p, 5, constraint_adjustment)
}

// SetOffset will set surface position offset.
//
//
// Specify the surface position offset relative to the position of the
// anchor on the anchor rectangle and the anchor on the surface. For
// example if the anchor of the anchor rectangle is at (x, y), the surface
// has the gravity bottom|right, and the offset is (ox, oy), the calculated
// surface position will be (x + ox, y + oy). The offset position of the
// surface is the one used for constraint testing. See
// set_constraint_adjustment.
//
// An example use case is placing a popup menu on top of a user interface
// element, while aligning the user interface element of the parent surface
// with some user interface element placed somewhere in the popup surface.
//
func (p *Positioner) SetOffset(x int32, y int32) error {
	return p.Context().SendRequest(p, 6, x, y)
}

const (
	PositionerErrorInvalidInput = 0
)

const (
	PositionerAnchorNone        = 0
	PositionerAnchorTop         = 1
	PositionerAnchorBottom      = 2
	PositionerAnchorLeft        = 3
	PositionerAnchorRight       = 4
	PositionerAnchorTopLeft     = 5
	PositionerAnchorBottomLeft  = 6
	PositionerAnchorTopRight    = 7
	PositionerAnchorBottomRight = 8
)

const (
	PositionerGravityNone        = 0
	PositionerGravityTop         = 1
	PositionerGravityBottom      = 2
	PositionerGravityLeft        = 3
	PositionerGravityRight       = 4
	PositionerGravityTopLeft     = 5
	PositionerGravityBottomLeft  = 6
	PositionerGravityTopRight    = 7
	PositionerGravityBottomRight = 8
)

const (
	PositionerConstraintAdjustmentNone    = 0
	PositionerConstraintAdjustmentSlideX  = 1
	PositionerConstraintAdjustmentSlideY  = 2
	PositionerConstraintAdjustmentFlipX   = 4
	PositionerConstraintAdjustmentFlipY   = 8
	PositionerConstraintAdjustmentResizeX = 16
	PositionerConstraintAdjustmentResizeY = 32
)

type SurfaceConfigureEvent struct {
	Serial uint32
}

func (p *Surface) AddConfigureHandler(h wl.Handler) {
	if h != nil {
		p.mu.Lock()
		p.configureHandlers = append(p.configureHandlers, h)
		p.mu.Unlock()
	}
}

func (p *Surface) RemoveConfigureHandler(h wl.Handler) {
	p.mu.Lock()
	defer p.mu.Unlock()

	for i, e := range p.configureHandlers {
		if e == h {
			p.configureHandlers = append(p.configureHandlers[:i], p.configureHandlers[i+1:]...)
			break
		}
	}
}

func (p *Surface) Dispatch(event *wl.Event) {
	switch event.Opcode {
	case 0:
		if len(p.configureHandlers) > 0 {
			ev := SurfaceConfigureEvent{}
			ev.Serial = event.Uint32()
			p.mu.RLock()
			for _, h := range p.configureHandlers {
				h.Handle(ev)
			}
			p.mu.RUnlock()
		}
	}
}

type Surface struct {
	wl.BaseProxy
	mu                sync.RWMutex
	configureHandlers []wl.Handler
}

func NewSurface(ctx *wl.Context) *Surface {
	ret := new(Surface)
	ctx.Register(ret)
	return ret
}

// Destroy will destroy the xdg_surface.
//
//
// Destroy the xdg_surface object. An xdg_surface must only be destroyed
// after its role object has been destroyed.
//
func (p *Surface) Destroy() error {
	return p.Context().SendRequest(p, 0)
}

// GetToplevel will assign the xdg_toplevel surface role.
//
//
// This creates an xdg_toplevel object for the given xdg_surface and gives
// the associated wl_surface the xdg_toplevel role.
//
// See the documentation of xdg_toplevel for more details about what an
// xdg_toplevel is and how it is used.
//
func (p *Surface) GetToplevel() (*Toplevel, error) {
	ret := NewToplevel(p.Context())
	return ret, p.Context().SendRequest(p, 1, wl.Proxy(ret))
}

// GetPopup will assign the xdg_popup surface role.
//
//
// This creates an xdg_popup object for the given xdg_surface and gives
// the associated wl_surface the xdg_popup role.
//
// If null is passed as a parent, a parent surface must be specified using
// some other protocol, before committing the initial state.
//
// See the documentation of xdg_popup for more details about what an
// xdg_popup is and how it is used.
//
func (p *Surface) GetPopup(parent *Surface, positioner *Positioner) (*Popup, error) {
	ret := NewPopup(p.Context())
	return ret, p.Context().SendRequest(p, 2, wl.Proxy(ret), parent, positioner)
}

// SetWindowGeometry will set the new window geometry.
//
//
// The window geometry of a surface is its "visible bounds" from the
// user's perspective. Client-side decorations often have invisible
// portions like drop-shadows which should be ignored for the
// purposes of aligning, placing and constraining windows.
//
// The window geometry is double buffered, and will be applied at the
// time wl_surface.commit of the corresponding wl_surface is called.
//
// When maintaining a position, the compositor should treat the (x, y)
// coordinate of the window geometry as the top left corner of the window.
// A client changing the (x, y) window geometry coordinate should in
// general not alter the position of the window.
//
// Once the window geometry of the surface is set, it is not possible to
// unset it, and it will remain the same until set_window_geometry is
// called again, even if a new subsurface or buffer is attached.
//
// If never set, the value is the full bounds of the surface,
// including any subsurfaces. This updates dynamically on every
// commit. This unset is meant for extremely simple clients.
//
// The arguments are given in the surface-local coordinate space of
// the wl_surface associated with this xdg_surface.
//
// The width and height must be greater than zero. Setting an invalid size
// will raise an error. When applied, the effective window geometry will be
// the set window geometry clamped to the bounding rectangle of the
// combined geometry of the surface of the xdg_surface and the associated
// subsurfaces.
//
func (p *Surface) SetWindowGeometry(x int32, y int32, width int32, height int32) error {
	return p.Context().SendRequest(p, 3, x, y, width, height)
}

// AckConfigure will ack a configure event.
//
//
// When a configure event is received, if a client commits the
// surface in response to the configure event, then the client
// must make an ack_configure request sometime before the commit
// request, passing along the serial of the configure event.
//
// For instance, for toplevel surfaces the compositor might use this
// information to move a surface to the top left only when the client has
// drawn itself for the maximized or fullscreen state.
//
// If the client receives multiple configure events before it
// can respond to one, it only has to ack the last configure event.
//
// A client is not required to commit immediately after sending
// an ack_configure request - it may even ack_configure several times
// before its next surface commit.
//
// A client may send multiple ack_configure requests before committing, but
// only the last request sent before a commit indicates which configure
// event the client really is responding to.
//
func (p *Surface) AckConfigure(serial uint32) error {
	return p.Context().SendRequest(p, 4, serial)
}

const (
	SurfaceErrorNotConstructed     = 1
	SurfaceErrorAlreadyConstructed = 2
	SurfaceErrorUnconfiguredBuffer = 3
)

type ToplevelConfigureEvent struct {
	Width  int32
	Height int32
	States []int32
}

func (p *Toplevel) AddConfigureHandler(h wl.Handler) {
	if h != nil {
		p.mu.Lock()
		p.configureHandlers = append(p.configureHandlers, h)
		p.mu.Unlock()
	}
}

func (p *Toplevel) RemoveConfigureHandler(h wl.Handler) {
	p.mu.Lock()
	defer p.mu.Unlock()

	for i, e := range p.configureHandlers {
		if e == h {
			p.configureHandlers = append(p.configureHandlers[:i], p.configureHandlers[i+1:]...)
			break
		}
	}
}

type ToplevelCloseEvent struct {
}

func (p *Toplevel) AddCloseHandler(h wl.Handler) {
	if h != nil {
		p.mu.Lock()
		p.closeHandlers = append(p.closeHandlers, h)
		p.mu.Unlock()
	}
}

func (p *Toplevel) RemoveCloseHandler(h wl.Handler) {
	p.mu.Lock()
	defer p.mu.Unlock()

	for i, e := range p.closeHandlers {
		if e == h {
			p.closeHandlers = append(p.closeHandlers[:i], p.closeHandlers[i+1:]...)
			break
		}
	}
}

func (p *Toplevel) Dispatch(event *wl.Event) {
	switch event.Opcode {
	case 0:
		if len(p.configureHandlers) > 0 {
			ev := ToplevelConfigureEvent{}
			ev.Width = event.Int32()
			ev.Height = event.Int32()
			ev.States = event.Array()
			p.mu.RLock()
			for _, h := range p.configureHandlers {
				h.Handle(ev)
			}
			p.mu.RUnlock()
		}
	case 1:
		if len(p.closeHandlers) > 0 {
			ev := ToplevelCloseEvent{}
			p.mu.RLock()
			for _, h := range p.closeHandlers {
				h.Handle(ev)
			}
			p.mu.RUnlock()
		}
	}
}

type Toplevel struct {
	wl.BaseProxy
	mu                sync.RWMutex
	configureHandlers []wl.Handler
	closeHandlers     []wl.Handler
}

func NewToplevel(ctx *wl.Context) *Toplevel {
	ret := new(Toplevel)
	ctx.Register(ret)
	return ret
}

// Destroy will destroy the xdg_toplevel.
//
//
// This request destroys the role surface and unmaps the surface;
// see "Unmapping" behavior in interface section for details.
//
func (p *Toplevel) Destroy() error {
	return p.Context().SendRequest(p, 0)
}

// SetParent will set the parent of this surface.
//
//
// Set the "parent" of this surface. This surface should be stacked
// above the parent surface and all other ancestor surfaces.
//
// Parent windows should be set on dialogs, toolboxes, or other
// "auxiliary" surfaces, so that the parent is raised when the dialog
// is raised.
//
// Setting a null parent for a child window removes any parent-child
// relationship for the child. Setting a null parent for a window which
// currently has no parent is a no-op.
//
// If the parent is unmapped then its children are managed as
// though the parent of the now-unmapped parent has become the
// parent of this surface. If no parent exists for the now-unmapped
// parent then the children are managed as though they have no
// parent surface.
//
func (p *Toplevel) SetParent(parent *Toplevel) error {
	return p.Context().SendRequest(p, 1, parent)
}

// SetTitle will set surface title.
//
//
// Set a short title for the surface.
//
// This string may be used to identify the surface in a task bar,
// window list, or other user interface elements provided by the
// compositor.
//
// The string must be encoded in UTF-8.
//
func (p *Toplevel) SetTitle(title string) error {
	return p.Context().SendRequest(p, 2, title)
}

// SetAppId will set application ID.
//
//
// Set an application identifier for the surface.
//
// The app ID identifies the general class of applications to which
// the surface belongs. The compositor can use this to group multiple
// surfaces together, or to determine how to launch a new application.
//
// For D-Bus activatable applications, the app ID is used as the D-Bus
// service name.
//
// The compositor shell will try to group application surfaces together
// by their app ID. As a best practice, it is suggested to select app
// ID's that match the basename of the application's .desktop file.
// For example, "org.freedesktop.FooViewer" where the .desktop file is
// "org.freedesktop.FooViewer.desktop".
//
// See the desktop-entry specification [0] for more details on
// application identifiers and how they relate to well-known D-Bus
// names and .desktop files.
//
// [0] http://standards.freedesktop.org/desktop-entry-spec/
//
func (p *Toplevel) SetAppId(app_id string) error {
	return p.Context().SendRequest(p, 3, app_id)
}

// ShowWindowMenu will show the window menu.
//
//
// Clients implementing client-side decorations might want to show
// a context menu when right-clicking on the decorations, giving the
// user a menu that they can use to maximize or minimize the window.
//
// This request asks the compositor to pop up such a window menu at
// the given position, relative to the local surface coordinates of
// the parent surface. There are no guarantees as to what menu items
// the window menu contains.
//
// This request must be used in response to some sort of user action
// like a button press, key press, or touch down event.
//
func (p *Toplevel) ShowWindowMenu(seat *wl.Seat, serial uint32, x int32, y int32) error {
	return p.Context().SendRequest(p, 4, seat, serial, x, y)
}

// Move will start an interactive move.
//
//
// Start an interactive, user-driven move of the surface.
//
// This request must be used in response to some sort of user action
// like a button press, key press, or touch down event. The passed
// serial is used to determine the type of interactive move (touch,
// pointer, etc).
//
// The server may ignore move requests depending on the state of
// the surface (e.g. fullscreen or maximized), or if the passed serial
// is no longer valid.
//
// If triggered, the surface will lose the focus of the device
// (wl_pointer, wl_touch, etc) used for the move. It is up to the
// compositor to visually indicate that the move is taking place, such as
// updating a pointer cursor, during the move. There is no guarantee
// that the device focus will return when the move is completed.
//
func (p *Toplevel) Move(seat *wl.Seat, serial uint32) error {
	return p.Context().SendRequest(p, 5, seat, serial)
}

// Resize will start an interactive resize.
//
//
// Start a user-driven, interactive resize of the surface.
//
// This request must be used in response to some sort of user action
// like a button press, key press, or touch down event. The passed
// serial is used to determine the type of interactive resize (touch,
// pointer, etc).
//
// The server may ignore resize requests depending on the state of
// the surface (e.g. fullscreen or maximized).
//
// If triggered, the client will receive configure events with the
// "resize" state enum value and the expected sizes. See the "resize"
// enum value for more details about what is required. The client
// must also acknowledge configure events using "ack_configure". After
// the resize is completed, the client will receive another "configure"
// event without the resize state.
//
// If triggered, the surface also will lose the focus of the device
// (wl_pointer, wl_touch, etc) used for the resize. It is up to the
// compositor to visually indicate that the resize is taking place,
// such as updating a pointer cursor, during the resize. There is no
// guarantee that the device focus will return when the resize is
// completed.
//
// The edges parameter specifies how the surface should be resized,
// and is one of the values of the resize_edge enum. The compositor
// may use this information to update the surface position for
// example when dragging the top left corner. The compositor may also
// use this information to adapt its behavior, e.g. choose an
// appropriate cursor image.
//
func (p *Toplevel) Resize(seat *wl.Seat, serial uint32, edges uint32) error {
	return p.Context().SendRequest(p, 6, seat, serial, edges)
}

// SetMaxSize will set the maximum size.
//
//
// Set a maximum size for the window.
//
// The client can specify a maximum size so that the compositor does
// not try to configure the window beyond this size.
//
// The width and height arguments are in window geometry coordinates.
// See xdg_surface.set_window_geometry.
//
// Values set in this way are double-buffered. They will get applied
// on the next commit.
//
// The compositor can use this information to allow or disallow
// different states like maximize or fullscreen and draw accurate
// animations.
//
// Similarly, a tiling window manager may use this information to
// place and resize client windows in a more effective way.
//
// The client should not rely on the compositor to obey the maximum
// size. The compositor may decide to ignore the values set by the
// client and request a larger size.
//
// If never set, or a value of zero in the request, means that the
// client has no expected maximum size in the given dimension.
// As a result, a client wishing to reset the maximum size
// to an unspecified state can use zero for width and height in the
// request.
//
// Requesting a maximum size to be smaller than the minimum size of
// a surface is illegal and will result in a protocol error.
//
// The width and height must be greater than or equal to zero. Using
// strictly negative values for width and height will result in a
// protocol error.
//
func (p *Toplevel) SetMaxSize(width int32, height int32) error {
	return p.Context().SendRequest(p, 7, width, height)
}

// SetMinSize will set the minimum size.
//
//
// Set a minimum size for the window.
//
// The client can specify a minimum size so that the compositor does
// not try to configure the window below this size.
//
// The width and height arguments are in window geometry coordinates.
// See xdg_surface.set_window_geometry.
//
// Values set in this way are double-buffered. They will get applied
// on the next commit.
//
// The compositor can use this information to allow or disallow
// different states like maximize or fullscreen and draw accurate
// animations.
//
// Similarly, a tiling window manager may use this information to
// place and resize client windows in a more effective way.
//
// The client should not rely on the compositor to obey the minimum
// size. The compositor may decide to ignore the values set by the
// client and request a smaller size.
//
// If never set, or a value of zero in the request, means that the
// client has no expected minimum size in the given dimension.
// As a result, a client wishing to reset the minimum size
// to an unspecified state can use zero for width and height in the
// request.
//
// Requesting a minimum size to be larger than the maximum size of
// a surface is illegal and will result in a protocol error.
//
// The width and height must be greater than or equal to zero. Using
// strictly negative values for width and height will result in a
// protocol error.
//
func (p *Toplevel) SetMinSize(width int32, height int32) error {
	return p.Context().SendRequest(p, 8, width, height)
}

// SetMaximized will maximize the window.
//
//
// Maximize the surface.
//
// After requesting that the surface should be maximized, the compositor
// will respond by emitting a configure event with the "maximized" state
// and the required window geometry. The client should then update its
// content, drawing it in a maximized state, i.e. without shadow or other
// decoration outside of the window geometry. The client must also
// acknowledge the configure when committing the new content (see
// ack_configure).
//
// It is up to the compositor to decide how and where to maximize the
// surface, for example which output and what region of the screen should
// be used.
//
// If the surface was already maximized, the compositor will still emit
// a configure event with the "maximized" state.
//
// If the surface is in a fullscreen state, this request has no direct
// effect. It will alter the state the surface is returned to when
// unmaximized if not overridden by the compositor.
//
func (p *Toplevel) SetMaximized() error {
	return p.Context().SendRequest(p, 9)
}

// UnsetMaximized will unmaximize the window.
//
//
// Unmaximize the surface.
//
// After requesting that the surface should be unmaximized, the compositor
// will respond by emitting a configure event without the "maximized"
// state. If available, the compositor will include the window geometry
// dimensions the window had prior to being maximized in the configure
// event. The client must then update its content, drawing it in a
// regular state, i.e. potentially with shadow, etc. The client must also
// acknowledge the configure when committing the new content (see
// ack_configure).
//
// It is up to the compositor to position the surface after it was
// unmaximized; usually the position the surface had before maximizing, if
// applicable.
//
// If the surface was already not maximized, the compositor will still
// emit a configure event without the "maximized" state.
//
// If the surface is in a fullscreen state, this request has no direct
// effect. It will alter the state the surface is returned to when
// unmaximized if not overridden by the compositor.
//
func (p *Toplevel) UnsetMaximized() error {
	return p.Context().SendRequest(p, 10)
}

// SetFullscreen will set the window as fullscreen on an output.
//
//
// Make the surface fullscreen.
//
// After requesting that the surface should be fullscreened, the
// compositor will respond by emitting a configure event with the
// "fullscreen" state and the fullscreen window geometry. The client must
// also acknowledge the configure when committing the new content (see
// ack_configure).
//
// The output passed by the request indicates the client's preference as
// to which display it should be set fullscreen on. If this value is NULL,
// it's up to the compositor to choose which display will be used to map
// this surface.
//
// If the surface doesn't cover the whole output, the compositor will
// position the surface in the center of the output and compensate with
// with border fill covering the rest of the output. The content of the
// border fill is undefined, but should be assumed to be in some way that
// attempts to blend into the surrounding area (e.g. solid black).
//
// If the fullscreened surface is not opaque, the compositor must make
// sure that other screen content not part of the same surface tree (made
// up of subsurfaces, popups or similarly coupled surfaces) are not
// visible below the fullscreened surface.
//
func (p *Toplevel) SetFullscreen(output *wl.Output) error {
	return p.Context().SendRequest(p, 11, output)
}

// UnsetFullscreen will unset the window as fullscreen.
//
//
// Make the surface no longer fullscreen.
//
// After requesting that the surface should be unfullscreened, the
// compositor will respond by emitting a configure event without the
// "fullscreen" state.
//
// Making a surface unfullscreen sets states for the surface based on the following:
// * the state(s) it may have had before becoming fullscreen
// * any state(s) decided by the compositor
// * any state(s) requested by the client while the surface was fullscreen
//
// The compositor may include the previous window geometry dimensions in
// the configure event, if applicable.
//
// The client must also acknowledge the configure when committing the new
// content (see ack_configure).
//
func (p *Toplevel) UnsetFullscreen() error {
	return p.Context().SendRequest(p, 12)
}

// SetMinimized will set the window as minimized.
//
//
// Request that the compositor minimize your surface. There is no
// way to know if the surface is currently minimized, nor is there
// any way to unset minimization on this surface.
//
// If you are looking to throttle redrawing when minimized, please
// instead use the wl_surface.frame event for this, as this will
// also work with live previews on windows in Alt-Tab, Expose or
// similar compositor features.
//
func (p *Toplevel) SetMinimized() error {
	return p.Context().SendRequest(p, 13)
}

const (
	ToplevelResizeEdgeNone        = 0
	ToplevelResizeEdgeTop         = 1
	ToplevelResizeEdgeBottom      = 2
	ToplevelResizeEdgeLeft        = 4
	ToplevelResizeEdgeTopLeft     = 5
	ToplevelResizeEdgeBottomLeft  = 6
	ToplevelResizeEdgeRight       = 8
	ToplevelResizeEdgeTopRight    = 9
	ToplevelResizeEdgeBottomRight = 10
)

const (
	ToplevelStateMaximized  = 1
	ToplevelStateFullscreen = 2
	ToplevelStateResizing   = 3
	ToplevelStateActivated  = 4
)

type PopupConfigureEvent struct {
	X      int32
	Y      int32
	Width  int32
	Height int32
}

func (p *Popup) AddConfigureHandler(h wl.Handler) {
	if h != nil {
		p.mu.Lock()
		p.configureHandlers = append(p.configureHandlers, h)
		p.mu.Unlock()
	}
}

func (p *Popup) RemoveConfigureHandler(h wl.Handler) {
	p.mu.Lock()
	defer p.mu.Unlock()

	for i, e := range p.configureHandlers {
		if e == h {
			p.configureHandlers = append(p.configureHandlers[:i], p.configureHandlers[i+1:]...)
			break
		}
	}
}

type PopupPopupDoneEvent struct {
}

func (p *Popup) AddPopupDoneHandler(h wl.Handler) {
	if h != nil {
		p.mu.Lock()
		p.popupDoneHandlers = append(p.popupDoneHandlers, h)
		p.mu.Unlock()
	}
}

func (p *Popup) RemovePopupDoneHandler(h wl.Handler) {
	p.mu.Lock()
	defer p.mu.Unlock()

	for i, e := range p.popupDoneHandlers {
		if e == h {
			p.popupDoneHandlers = append(p.popupDoneHandlers[:i], p.popupDoneHandlers[i+1:]...)
			break
		}
	}
}

func (p *Popup) Dispatch(event *wl.Event) {
	switch event.Opcode {
	case 0:
		if len(p.configureHandlers) > 0 {
			ev := PopupConfigureEvent{}
			ev.X = event.Int32()
			ev.Y = event.Int32()
			ev.Width = event.Int32()
			ev.Height = event.Int32()
			p.mu.RLock()
			for _, h := range p.configureHandlers {
				h.Handle(ev)
			}
			p.mu.RUnlock()
		}
	case 1:
		if len(p.popupDoneHandlers) > 0 {
			ev := PopupPopupDoneEvent{}
			p.mu.RLock()
			for _, h := range p.popupDoneHandlers {
				h.Handle(ev)
			}
			p.mu.RUnlock()
		}
	}
}

type Popup struct {
	wl.BaseProxy
	mu                sync.RWMutex
	configureHandlers []wl.Handler
	popupDoneHandlers []wl.Handler
}

func NewPopup(ctx *wl.Context) *Popup {
	ret := new(Popup)
	ctx.Register(ret)
	return ret
}

// Destroy will remove xdg_popup interface.
//
//
// This destroys the popup. Explicitly destroying the xdg_popup
// object will also dismiss the popup, and unmap the surface.
//
// If this xdg_popup is not the "topmost" popup, a protocol error
// will be sent.
//
func (p *Popup) Destroy() error {
	return p.Context().SendRequest(p, 0)
}

// Grab will make the popup take an explicit grab.
//
//
// This request makes the created popup take an explicit grab. An explicit
// grab will be dismissed when the user dismisses the popup, or when the
// client destroys the xdg_popup. This can be done by the user clicking
// outside the surface, using the keyboard, or even locking the screen
// through closing the lid or a timeout.
//
// If the compositor denies the grab, the popup will be immediately
// dismissed.
//
// This request must be used in response to some sort of user action like a
// button press, key press, or touch down event. The serial number of the
// event should be passed as 'serial'.
//
// The parent of a grabbing popup must either be an xdg_toplevel surface or
// another xdg_popup with an explicit grab. If the parent is another
// xdg_popup it means that the popups are nested, with this popup now being
// the topmost popup.
//
// Nested popups must be destroyed in the reverse order they were created
// in, e.g. the only popup you are allowed to destroy at all times is the
// topmost one.
//
// When compositors choose to dismiss a popup, they may dismiss every
// nested grabbing popup as well. When a compositor dismisses popups, it
// will follow the same dismissing order as required from the client.
//
// The parent of a grabbing popup must either be another xdg_popup with an
// active explicit grab, or an xdg_popup or xdg_toplevel, if there are no
// explicit grabs already taken.
//
// If the topmost grabbing popup is destroyed, the grab will be returned to
// the parent of the popup, if that parent previously had an explicit grab.
//
// If the parent is a grabbing popup which has already been dismissed, this
// popup will be immediately dismissed. If the parent is a popup that did
// not take an explicit grab, an error will be raised.
//
// During a popup grab, the client owning the grab will receive pointer
// and touch events for all their surfaces as normal (similar to an
// "owner-events" grab in X11 parlance), while the top most grabbing popup
// will always have keyboard focus.
//
func (p *Popup) Grab(seat *wl.Seat, serial uint32) error {
	return p.Context().SendRequest(p, 1, seat, serial)
}

const (
	PopupErrorInvalidGrab = 0
)
