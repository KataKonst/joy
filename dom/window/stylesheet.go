package window

import "github.com/matthewmueller/golly/dom/medialist"

// StyleSheet interface
// js:"StyleSheet"
type StyleSheet interface {

	// Disabled prop
	// js:"disabled"
	Disabled() (disabled bool)

	// SetDisabled prop
	// js:"disabled"
	SetDisabled(disabled bool)

	// Href prop
	// js:"href"
	Href() (href string)

	// Media prop
	// js:"media"
	Media() (media *medialist.MediaList)

	// OwnerNode prop
	// js:"ownerNode"
	OwnerNode() (ownerNode Node)

	// ParentStyleSheet prop
	// js:"parentStyleSheet"
	ParentStyleSheet() (parentStyleSheet StyleSheet)

	// Title prop
	// js:"title"
	Title() (title string)

	// Type prop
	// js:"type"
	Type() (kind string)
}
