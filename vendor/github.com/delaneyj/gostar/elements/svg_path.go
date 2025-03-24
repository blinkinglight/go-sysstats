// DO NOT EDIT THIS FILE. IT IS AUTOMATICALLY GENERATED.
// package svg path is generated from configuration file.
// Description:
package elements

import (
	"fmt"
	"html"
	"time"

	"github.com/goccy/go-json"
	"github.com/igrmk/treemap/v2"
	"github.com/samber/lo"
)

// The <path> SVG element is the generic element to define a shape
// All the basic shapes can be created with a path element.
type SVGPATHElement struct {
	*Element
}

// Create a new SVGPATHElement element.
// This will create a new element with the tag
// "path" during rendering.
func SVG_PATH(children ...ElementRenderer) *SVGPATHElement {
	e := NewElement("path", children...)
	e.IsSelfClosing = false
	e.Descendants = children

	return &SVGPATHElement{Element: e}
}

func (e *SVGPATHElement) Children(children ...ElementRenderer) *SVGPATHElement {
	e.Descendants = append(e.Descendants, children...)
	return e
}

func (e *SVGPATHElement) IfChildren(condition bool, children ...ElementRenderer) *SVGPATHElement {
	if condition {
		e.Descendants = append(e.Descendants, children...)
	}
	return e
}

func (e *SVGPATHElement) TernChildren(condition bool, trueChildren, falseChildren ElementRenderer) *SVGPATHElement {
	if condition {
		e.Descendants = append(e.Descendants, trueChildren)
	} else {
		e.Descendants = append(e.Descendants, falseChildren)
	}
	return e
}

func (e *SVGPATHElement) Attr(name string, value string) *SVGPATHElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set(name, value)
	return e
}

func (e *SVGPATHElement) Attrs(attrs ...string) *SVGPATHElement {
	if len(attrs)%2 != 0 {
		panic("attrs must be a multiple of 2")
	}
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	for i := 0; i < len(attrs); i += 2 {
		k := attrs[i]
		v := attrs[i+1]
		e.StringAttributes.Set(k, v)
	}
	return e
}

func (e *SVGPATHElement) AttrsMap(attrs map[string]string) *SVGPATHElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	for k, v := range attrs {
		e.StringAttributes.Set(k, v)
	}
	return e
}

func (e *SVGPATHElement) Text(text string) *SVGPATHElement {
	e.Descendants = append(e.Descendants, Text(text))
	return e
}

func (e *SVGPATHElement) TextF(format string, args ...any) *SVGPATHElement {
	return e.Text(fmt.Sprintf(format, args...))
}

func (e *SVGPATHElement) IfText(condition bool, text string) *SVGPATHElement {
	if condition {
		e.Descendants = append(e.Descendants, Text(text))
	}
	return e
}

func (e *SVGPATHElement) IfTextF(condition bool, format string, args ...any) *SVGPATHElement {
	if condition {
		e.Descendants = append(e.Descendants, Text(fmt.Sprintf(format, args...)))
	}
	return e
}

func (e *SVGPATHElement) Escaped(text string) *SVGPATHElement {
	e.Descendants = append(e.Descendants, Escaped(text))
	return e
}

func (e *SVGPATHElement) IfEscaped(condition bool, text string) *SVGPATHElement {
	if condition {
		e.Descendants = append(e.Descendants, Escaped(text))
	}
	return e
}

func (e *SVGPATHElement) EscapedF(format string, args ...any) *SVGPATHElement {
	return e.Escaped(fmt.Sprintf(format, args...))
}

func (e *SVGPATHElement) IfEscapedF(condition bool, format string, args ...any) *SVGPATHElement {
	if condition {
		e.Descendants = append(e.Descendants, EscapedF(format, args...))
	}
	return e
}

func (e *SVGPATHElement) CustomData(key, value string) *SVGPATHElement {
	if e.CustomDataAttributes == nil {
		e.CustomDataAttributes = treemap.New[string, string]()
	}
	e.CustomDataAttributes.Set(key, value)
	return e
}

func (e *SVGPATHElement) IfCustomData(condition bool, key, value string) *SVGPATHElement {
	if condition {
		e.CustomData(key, value)
	}
	return e
}

func (e *SVGPATHElement) CustomDataF(key, format string, args ...any) *SVGPATHElement {
	return e.CustomData(key, fmt.Sprintf(format, args...))
}

func (e *SVGPATHElement) IfCustomDataF(condition bool, key, format string, args ...any) *SVGPATHElement {
	if condition {
		e.CustomData(key, fmt.Sprintf(format, args...))
	}
	return e
}

func (e *SVGPATHElement) CustomDataRemove(key string) *SVGPATHElement {
	if e.CustomDataAttributes == nil {
		return e
	}
	e.CustomDataAttributes.Del(key)
	return e
}

// The definition of the outline of a shape.
func (e *SVGPATHElement) D(s string) *SVGPATHElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("d", s)
	return e
}

func (e *SVGPATHElement) DF(format string, args ...any) *SVGPATHElement {
	return e.D(fmt.Sprintf(format, args...))
}

func (e *SVGPATHElement) IfD(condition bool, s string) *SVGPATHElement {
	if condition {
		e.D(s)
	}
	return e
}

func (e *SVGPATHElement) IfDF(condition bool, format string, args ...any) *SVGPATHElement {
	if condition {
		e.D(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute D from the element.
func (e *SVGPATHElement) DRemove(s string) *SVGPATHElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("d")
	return e
}

func (e *SVGPATHElement) DRemoveF(format string, args ...any) *SVGPATHElement {
	return e.DRemove(fmt.Sprintf(format, args...))
}

// The <path> SVG element is the generic element to define a shape
// All the basic shapes can be created with a path element.
func (e *SVGPATHElement) FILL(s string) *SVGPATHElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("fill", s)
	return e
}

func (e *SVGPATHElement) FILLF(format string, args ...any) *SVGPATHElement {
	return e.FILL(fmt.Sprintf(format, args...))
}

func (e *SVGPATHElement) IfFILL(condition bool, s string) *SVGPATHElement {
	if condition {
		e.FILL(s)
	}
	return e
}

func (e *SVGPATHElement) IfFILLF(condition bool, format string, args ...any) *SVGPATHElement {
	if condition {
		e.FILL(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute FILL from the element.
func (e *SVGPATHElement) FILLRemove(s string) *SVGPATHElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("fill")
	return e
}

func (e *SVGPATHElement) FILLRemoveF(format string, args ...any) *SVGPATHElement {
	return e.FILLRemove(fmt.Sprintf(format, args...))
}

// The <path> SVG element is the generic element to define a shape
// All the basic shapes can be created with a path element.
func (e *SVGPATHElement) FILL_OPACITY(f float64) *SVGPATHElement {
	if e.FloatAttributes == nil {
		e.FloatAttributes = treemap.New[string, float64]()
	}
	e.FloatAttributes.Set("fill-opacity", f)
	return e
}

func (e *SVGPATHElement) IfFILL_OPACITY(condition bool, f float64) *SVGPATHElement {
	if condition {
		e.FILL_OPACITY(f)
	}
	return e
}

// The total length for the path, in user units.
func (e *SVGPATHElement) PATH_LENGTH(f float64) *SVGPATHElement {
	if e.FloatAttributes == nil {
		e.FloatAttributes = treemap.New[string, float64]()
	}
	e.FloatAttributes.Set("pathLength", f)
	return e
}

func (e *SVGPATHElement) IfPATH_LENGTH(condition bool, f float64) *SVGPATHElement {
	if condition {
		e.PATH_LENGTH(f)
	}
	return e
}

// Specifies a unique id for an element
func (e *SVGPATHElement) ID(s string) *SVGPATHElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("id", s)
	return e
}

func (e *SVGPATHElement) IDF(format string, args ...any) *SVGPATHElement {
	return e.ID(fmt.Sprintf(format, args...))
}

func (e *SVGPATHElement) IfID(condition bool, s string) *SVGPATHElement {
	if condition {
		e.ID(s)
	}
	return e
}

func (e *SVGPATHElement) IfIDF(condition bool, format string, args ...any) *SVGPATHElement {
	if condition {
		e.ID(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute ID from the element.
func (e *SVGPATHElement) IDRemove(s string) *SVGPATHElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("id")
	return e
}

func (e *SVGPATHElement) IDRemoveF(format string, args ...any) *SVGPATHElement {
	return e.IDRemove(fmt.Sprintf(format, args...))
}

// Specifies one or more classnames for an element (refers to a class in a style
// sheet)
func (e *SVGPATHElement) CLASS(s ...string) *SVGPATHElement {
	if e.DelimitedStrings == nil {
		e.DelimitedStrings = treemap.New[string, *DelimitedBuilder[string]]()
	}
	ds, ok := e.DelimitedStrings.Get("class")
	if !ok {
		ds = NewDelimitedBuilder[string](" ")
		e.DelimitedStrings.Set("class", ds)
	}
	ds.Add(s...)
	return e
}

func (e *SVGPATHElement) IfCLASS(condition bool, s ...string) *SVGPATHElement {
	if condition {
		e.CLASS(s...)
	}
	return e
}

// Remove the attribute CLASS from the element.
func (e *SVGPATHElement) CLASSRemove(s ...string) *SVGPATHElement {
	if e.DelimitedStrings == nil {
		return e
	}
	ds, ok := e.DelimitedStrings.Get("class")
	if !ok {
		return e
	}
	ds.Remove(s...)
	return e
}

// Specifies an inline CSS style for an element
func (e *SVGPATHElement) STYLEF(k string, format string, args ...any) *SVGPATHElement {
	return e.STYLE(k, fmt.Sprintf(format, args...))
}

func (e *SVGPATHElement) IfSTYLE(condition bool, k string, v string) *SVGPATHElement {
	if condition {
		e.STYLE(k, v)
	}
	return e
}

func (e *SVGPATHElement) STYLE(k string, v string) *SVGPATHElement {
	if e.KVStrings == nil {
		e.KVStrings = treemap.New[string, *KVBuilder]()
	}
	kv, ok := e.KVStrings.Get("style")
	if !ok {
		kv = NewKVBuilder(":", ";")
		e.KVStrings.Set("style", kv)
	}
	kv.Add(k, v)
	return e
}

func (e *SVGPATHElement) IfSTYLEF(condition bool, k string, format string, args ...any) *SVGPATHElement {
	if condition {
		e.STYLE(k, fmt.Sprintf(format, args...))
	}
	return e
}

// Add the attributes in the map to the element.
func (e *SVGPATHElement) STYLEMap(m map[string]string) *SVGPATHElement {
	if e.KVStrings == nil {
		e.KVStrings = treemap.New[string, *KVBuilder]()
	}
	kv, ok := e.KVStrings.Get("style")
	if !ok {
		kv = NewKVBuilder(":", ";")
		e.KVStrings.Set("style", kv)
	}
	for k, v := range m {
		kv.Add(k, v)
	}
	return e
}

// Add pairs of attributes to the element.
func (e *SVGPATHElement) STYLEPairs(pairs ...string) *SVGPATHElement {
	if len(pairs)%2 != 0 {
		panic("Must have an even number of pairs")
	}
	if e.KVStrings == nil {
		e.KVStrings = treemap.New[string, *KVBuilder]()
	}
	kv, ok := e.KVStrings.Get("style")
	if !ok {
		kv = NewKVBuilder(":", ";")
		e.KVStrings.Set("style", kv)
	}

	for i := 0; i < len(pairs); i += 2 {
		kv.Add(pairs[i], pairs[i+1])
	}

	return e
}

func (e *SVGPATHElement) IfSTYLEPairs(condition bool, pairs ...string) *SVGPATHElement {
	if condition {
		e.STYLEPairs(pairs...)
	}
	return e
}

// Remove the attribute STYLE from the element.
func (e *SVGPATHElement) STYLERemove(keys ...string) *SVGPATHElement {
	if e.KVStrings == nil {
		return e
	}
	kv, ok := e.KVStrings.Get("style")
	if !ok {
		return e
	}
	for _, k := range keys {
		kv.Remove(k)
	}
	return e
}

// Merges the singleton store with the given object

func (e *SVGPATHElement) DATASTAR_STORE(v any) *SVGPATHElement {
	if e.CustomDataAttributes == nil {
		e.CustomDataAttributes = treemap.New[string, string]()
	}
	b, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}
	e.CustomDataAttributes.Set("store", html.EscapeString(string(b)))
	return e
}

// Sets the reference of the element

func (e *SVGPATHElement) DATASTAR_REF(expression string) *SVGPATHElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-ref"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGPATHElement) IfDATASTAR_REF(condition bool, expression string) *SVGPATHElement {
	if condition {
		e.DATASTAR_REF(expression)
	}
	return e
}

// Remove the attribute DATASTAR_REF from the element.
func (e *SVGPATHElement) DATASTAR_REFRemove() *SVGPATHElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-ref")
	return e
}

// Sets the value of the element

func (e *SVGPATHElement) DATASTAR_BIND(key string, expression string) *SVGPATHElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key = fmt.Sprintf("data-bind-%s", key)

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGPATHElement) IfDATASTAR_BIND(condition bool, key string, expression string) *SVGPATHElement {
	if condition {
		e.DATASTAR_BIND(key, expression)
	}
	return e
}

// Remove the attribute DATASTAR_BIND from the element.
func (e *SVGPATHElement) DATASTAR_BINDRemove() *SVGPATHElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-bind")
	return e
}

// Sets the value of the element

func (e *SVGPATHElement) DATASTAR_MODEL(expression string) *SVGPATHElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-model"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGPATHElement) IfDATASTAR_MODEL(condition bool, expression string) *SVGPATHElement {
	if condition {
		e.DATASTAR_MODEL(expression)
	}
	return e
}

// Remove the attribute DATASTAR_MODEL from the element.
func (e *SVGPATHElement) DATASTAR_MODELRemove() *SVGPATHElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-model")
	return e
}

// Sets the textContent of the element

func (e *SVGPATHElement) DATASTAR_TEXT(expression string) *SVGPATHElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-text"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGPATHElement) IfDATASTAR_TEXT(condition bool, expression string) *SVGPATHElement {
	if condition {
		e.DATASTAR_TEXT(expression)
	}
	return e
}

// Remove the attribute DATASTAR_TEXT from the element.
func (e *SVGPATHElement) DATASTAR_TEXTRemove() *SVGPATHElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-text")
	return e
}

// Sets the event handler of the element

type SVGPathOnMod customDataKeyModifier

// Debounces the event handler
func SVGPathOnModDebounce(
	d time.Duration,
) SVGPathOnMod {
	return func() string {
		return fmt.Sprintf("debounce_%dms", d.Milliseconds())
	}
}

// Throttles the event handler
func SVGPathOnModThrottle(
	d time.Duration,
) SVGPathOnMod {
	return func() string {
		return fmt.Sprintf("throttle_%dms", d.Milliseconds())
	}
}

func (e *SVGPATHElement) DATASTAR_ON(key string, expression string, modifiers ...SVGPathOnMod) *SVGPATHElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key = fmt.Sprintf("data-on-%s", key)

	customMods := lo.Map(modifiers, func(m SVGPathOnMod, i int) customDataKeyModifier {
		return customDataKeyModifier(m)
	})
	key = customDataKey(key, customMods...)
	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGPATHElement) IfDATASTAR_ON(condition bool, key string, expression string, modifiers ...SVGPathOnMod) *SVGPATHElement {
	if condition {
		e.DATASTAR_ON(key, expression, modifiers...)
	}
	return e
}

// Remove the attribute DATASTAR_ON from the element.
func (e *SVGPATHElement) DATASTAR_ONRemove() *SVGPATHElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-on")
	return e
}

// Sets the focus of the element

func (e *SVGPATHElement) DATASTAR_FOCUSSet(b bool) *SVGPATHElement {
	key := "data-focus"
	e.BoolAttributes.Set(key, b)
	return e
}

func (e *SVGPATHElement) DATASTAR_FOCUS() *SVGPATHElement {
	return e.DATASTAR_FOCUSSet(true)
}

// Sets the header of for fetch requests

func (e *SVGPATHElement) DATASTAR_HEADER(key string, expression string) *SVGPATHElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key = fmt.Sprintf("data-header-%s", key)

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGPATHElement) IfDATASTAR_HEADER(condition bool, key string, expression string) *SVGPATHElement {
	if condition {
		e.DATASTAR_HEADER(key, expression)
	}
	return e
}

// Remove the attribute DATASTAR_HEADER from the element.
func (e *SVGPATHElement) DATASTAR_HEADERRemove() *SVGPATHElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-header")
	return e
}

// Sets the indicator selector for fetch requests

func (e *SVGPATHElement) DATASTAR_FETCH_INDICATOR(expression string) *SVGPATHElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-fetch-indicator"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGPATHElement) IfDATASTAR_FETCH_INDICATOR(condition bool, expression string) *SVGPATHElement {
	if condition {
		e.DATASTAR_FETCH_INDICATOR(expression)
	}
	return e
}

// Remove the attribute DATASTAR_FETCH_INDICATOR from the element.
func (e *SVGPATHElement) DATASTAR_FETCH_INDICATORRemove() *SVGPATHElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-fetch-indicator")
	return e
}

// Sets the visibility of the element

func (e *SVGPATHElement) DATASTAR_SHOWSet(b bool) *SVGPATHElement {
	key := "data-show"
	e.BoolAttributes.Set(key, b)
	return e
}

func (e *SVGPATHElement) DATASTAR_SHOW() *SVGPATHElement {
	return e.DATASTAR_SHOWSet(true)
}

// Triggers the callback when the element intersects the viewport

func (e *SVGPATHElement) DATASTAR_INTERSECTS(expression string) *SVGPATHElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-intersects"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGPATHElement) IfDATASTAR_INTERSECTS(condition bool, expression string) *SVGPATHElement {
	if condition {
		e.DATASTAR_INTERSECTS(expression)
	}
	return e
}

// Remove the attribute DATASTAR_INTERSECTS from the element.
func (e *SVGPATHElement) DATASTAR_INTERSECTSRemove() *SVGPATHElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-intersects")
	return e
}

// Teleports the element to the given selector

func (e *SVGPATHElement) DATASTAR_TELEPORTSet(b bool) *SVGPATHElement {
	key := "data-teleport"
	e.BoolAttributes.Set(key, b)
	return e
}

func (e *SVGPATHElement) DATASTAR_TELEPORT() *SVGPATHElement {
	return e.DATASTAR_TELEPORTSet(true)
}

// Scrolls the element into view

func (e *SVGPATHElement) DATASTAR_SCROLL_INTO_VIEWSet(b bool) *SVGPATHElement {
	key := "data-scroll-into-view"
	e.BoolAttributes.Set(key, b)
	return e
}

func (e *SVGPATHElement) DATASTAR_SCROLL_INTO_VIEW() *SVGPATHElement {
	return e.DATASTAR_SCROLL_INTO_VIEWSet(true)
}

// Setup the ViewTransitionAPI for the element

func (e *SVGPATHElement) DATASTAR_VIEW_TRANSITION(expression string) *SVGPATHElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-view-transition"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGPATHElement) IfDATASTAR_VIEW_TRANSITION(condition bool, expression string) *SVGPATHElement {
	if condition {
		e.DATASTAR_VIEW_TRANSITION(expression)
	}
	return e
}

// Remove the attribute DATASTAR_VIEW_TRANSITION from the element.
func (e *SVGPATHElement) DATASTAR_VIEW_TRANSITIONRemove() *SVGPATHElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-view-transition")
	return e
}
