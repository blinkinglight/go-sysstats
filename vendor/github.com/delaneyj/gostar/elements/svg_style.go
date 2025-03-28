// DO NOT EDIT THIS FILE. IT IS AUTOMATICALLY GENERATED.
// package svg style is generated from configuration file.
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

// The <style> SVG element allows style sheets to be embedded directly within SVG
// content
// SVG's style element has the same attributes as the corresponding element in
// HTML (see HTML's <style> element).
type SVGSTYLEElement struct {
	*Element
}

// Create a new SVGSTYLEElement element.
// This will create a new element with the tag
// "style" during rendering.
func SVG_STYLE(children ...ElementRenderer) *SVGSTYLEElement {
	e := NewElement("style", children...)
	e.IsSelfClosing = false
	e.Descendants = children

	return &SVGSTYLEElement{Element: e}
}

func (e *SVGSTYLEElement) Children(children ...ElementRenderer) *SVGSTYLEElement {
	e.Descendants = append(e.Descendants, children...)
	return e
}

func (e *SVGSTYLEElement) IfChildren(condition bool, children ...ElementRenderer) *SVGSTYLEElement {
	if condition {
		e.Descendants = append(e.Descendants, children...)
	}
	return e
}

func (e *SVGSTYLEElement) TernChildren(condition bool, trueChildren, falseChildren ElementRenderer) *SVGSTYLEElement {
	if condition {
		e.Descendants = append(e.Descendants, trueChildren)
	} else {
		e.Descendants = append(e.Descendants, falseChildren)
	}
	return e
}

func (e *SVGSTYLEElement) Attr(name string, value string) *SVGSTYLEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set(name, value)
	return e
}

func (e *SVGSTYLEElement) Attrs(attrs ...string) *SVGSTYLEElement {
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

func (e *SVGSTYLEElement) AttrsMap(attrs map[string]string) *SVGSTYLEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	for k, v := range attrs {
		e.StringAttributes.Set(k, v)
	}
	return e
}

func (e *SVGSTYLEElement) Text(text string) *SVGSTYLEElement {
	e.Descendants = append(e.Descendants, Text(text))
	return e
}

func (e *SVGSTYLEElement) TextF(format string, args ...any) *SVGSTYLEElement {
	return e.Text(fmt.Sprintf(format, args...))
}

func (e *SVGSTYLEElement) IfText(condition bool, text string) *SVGSTYLEElement {
	if condition {
		e.Descendants = append(e.Descendants, Text(text))
	}
	return e
}

func (e *SVGSTYLEElement) IfTextF(condition bool, format string, args ...any) *SVGSTYLEElement {
	if condition {
		e.Descendants = append(e.Descendants, Text(fmt.Sprintf(format, args...)))
	}
	return e
}

func (e *SVGSTYLEElement) Escaped(text string) *SVGSTYLEElement {
	e.Descendants = append(e.Descendants, Escaped(text))
	return e
}

func (e *SVGSTYLEElement) IfEscaped(condition bool, text string) *SVGSTYLEElement {
	if condition {
		e.Descendants = append(e.Descendants, Escaped(text))
	}
	return e
}

func (e *SVGSTYLEElement) EscapedF(format string, args ...any) *SVGSTYLEElement {
	return e.Escaped(fmt.Sprintf(format, args...))
}

func (e *SVGSTYLEElement) IfEscapedF(condition bool, format string, args ...any) *SVGSTYLEElement {
	if condition {
		e.Descendants = append(e.Descendants, EscapedF(format, args...))
	}
	return e
}

func (e *SVGSTYLEElement) CustomData(key, value string) *SVGSTYLEElement {
	if e.CustomDataAttributes == nil {
		e.CustomDataAttributes = treemap.New[string, string]()
	}
	e.CustomDataAttributes.Set(key, value)
	return e
}

func (e *SVGSTYLEElement) IfCustomData(condition bool, key, value string) *SVGSTYLEElement {
	if condition {
		e.CustomData(key, value)
	}
	return e
}

func (e *SVGSTYLEElement) CustomDataF(key, format string, args ...any) *SVGSTYLEElement {
	return e.CustomData(key, fmt.Sprintf(format, args...))
}

func (e *SVGSTYLEElement) IfCustomDataF(condition bool, key, format string, args ...any) *SVGSTYLEElement {
	if condition {
		e.CustomData(key, fmt.Sprintf(format, args...))
	}
	return e
}

func (e *SVGSTYLEElement) CustomDataRemove(key string) *SVGSTYLEElement {
	if e.CustomDataAttributes == nil {
		return e
	}
	e.CustomDataAttributes.Del(key)
	return e
}

// The style sheet language.
func (e *SVGSTYLEElement) TYPE(s string) *SVGSTYLEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("type", s)
	return e
}

func (e *SVGSTYLEElement) TYPEF(format string, args ...any) *SVGSTYLEElement {
	return e.TYPE(fmt.Sprintf(format, args...))
}

func (e *SVGSTYLEElement) IfTYPE(condition bool, s string) *SVGSTYLEElement {
	if condition {
		e.TYPE(s)
	}
	return e
}

func (e *SVGSTYLEElement) IfTYPEF(condition bool, format string, args ...any) *SVGSTYLEElement {
	if condition {
		e.TYPE(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute TYPE from the element.
func (e *SVGSTYLEElement) TYPERemove(s string) *SVGSTYLEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("type")
	return e
}

func (e *SVGSTYLEElement) TYPERemoveF(format string, args ...any) *SVGSTYLEElement {
	return e.TYPERemove(fmt.Sprintf(format, args...))
}

// The intended destination medium for style information.
func (e *SVGSTYLEElement) MEDIA(s string) *SVGSTYLEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("media", s)
	return e
}

func (e *SVGSTYLEElement) MEDIAF(format string, args ...any) *SVGSTYLEElement {
	return e.MEDIA(fmt.Sprintf(format, args...))
}

func (e *SVGSTYLEElement) IfMEDIA(condition bool, s string) *SVGSTYLEElement {
	if condition {
		e.MEDIA(s)
	}
	return e
}

func (e *SVGSTYLEElement) IfMEDIAF(condition bool, format string, args ...any) *SVGSTYLEElement {
	if condition {
		e.MEDIA(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute MEDIA from the element.
func (e *SVGSTYLEElement) MEDIARemove(s string) *SVGSTYLEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("media")
	return e
}

func (e *SVGSTYLEElement) MEDIARemoveF(format string, args ...any) *SVGSTYLEElement {
	return e.MEDIARemove(fmt.Sprintf(format, args...))
}

// The advisory title.
func (e *SVGSTYLEElement) TITLE(s string) *SVGSTYLEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("title", s)
	return e
}

func (e *SVGSTYLEElement) TITLEF(format string, args ...any) *SVGSTYLEElement {
	return e.TITLE(fmt.Sprintf(format, args...))
}

func (e *SVGSTYLEElement) IfTITLE(condition bool, s string) *SVGSTYLEElement {
	if condition {
		e.TITLE(s)
	}
	return e
}

func (e *SVGSTYLEElement) IfTITLEF(condition bool, format string, args ...any) *SVGSTYLEElement {
	if condition {
		e.TITLE(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute TITLE from the element.
func (e *SVGSTYLEElement) TITLERemove(s string) *SVGSTYLEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("title")
	return e
}

func (e *SVGSTYLEElement) TITLERemoveF(format string, args ...any) *SVGSTYLEElement {
	return e.TITLERemove(fmt.Sprintf(format, args...))
}

// Specifies a unique id for an element
func (e *SVGSTYLEElement) ID(s string) *SVGSTYLEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("id", s)
	return e
}

func (e *SVGSTYLEElement) IDF(format string, args ...any) *SVGSTYLEElement {
	return e.ID(fmt.Sprintf(format, args...))
}

func (e *SVGSTYLEElement) IfID(condition bool, s string) *SVGSTYLEElement {
	if condition {
		e.ID(s)
	}
	return e
}

func (e *SVGSTYLEElement) IfIDF(condition bool, format string, args ...any) *SVGSTYLEElement {
	if condition {
		e.ID(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute ID from the element.
func (e *SVGSTYLEElement) IDRemove(s string) *SVGSTYLEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("id")
	return e
}

func (e *SVGSTYLEElement) IDRemoveF(format string, args ...any) *SVGSTYLEElement {
	return e.IDRemove(fmt.Sprintf(format, args...))
}

// Specifies one or more classnames for an element (refers to a class in a style
// sheet)
func (e *SVGSTYLEElement) CLASS(s ...string) *SVGSTYLEElement {
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

func (e *SVGSTYLEElement) IfCLASS(condition bool, s ...string) *SVGSTYLEElement {
	if condition {
		e.CLASS(s...)
	}
	return e
}

// Remove the attribute CLASS from the element.
func (e *SVGSTYLEElement) CLASSRemove(s ...string) *SVGSTYLEElement {
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
func (e *SVGSTYLEElement) STYLEF(k string, format string, args ...any) *SVGSTYLEElement {
	return e.STYLE(k, fmt.Sprintf(format, args...))
}

func (e *SVGSTYLEElement) IfSTYLE(condition bool, k string, v string) *SVGSTYLEElement {
	if condition {
		e.STYLE(k, v)
	}
	return e
}

func (e *SVGSTYLEElement) STYLE(k string, v string) *SVGSTYLEElement {
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

func (e *SVGSTYLEElement) IfSTYLEF(condition bool, k string, format string, args ...any) *SVGSTYLEElement {
	if condition {
		e.STYLE(k, fmt.Sprintf(format, args...))
	}
	return e
}

// Add the attributes in the map to the element.
func (e *SVGSTYLEElement) STYLEMap(m map[string]string) *SVGSTYLEElement {
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
func (e *SVGSTYLEElement) STYLEPairs(pairs ...string) *SVGSTYLEElement {
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

func (e *SVGSTYLEElement) IfSTYLEPairs(condition bool, pairs ...string) *SVGSTYLEElement {
	if condition {
		e.STYLEPairs(pairs...)
	}
	return e
}

// Remove the attribute STYLE from the element.
func (e *SVGSTYLEElement) STYLERemove(keys ...string) *SVGSTYLEElement {
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

func (e *SVGSTYLEElement) DATASTAR_STORE(v any) *SVGSTYLEElement {
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

func (e *SVGSTYLEElement) DATASTAR_REF(expression string) *SVGSTYLEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-ref"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGSTYLEElement) IfDATASTAR_REF(condition bool, expression string) *SVGSTYLEElement {
	if condition {
		e.DATASTAR_REF(expression)
	}
	return e
}

// Remove the attribute DATASTAR_REF from the element.
func (e *SVGSTYLEElement) DATASTAR_REFRemove() *SVGSTYLEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-ref")
	return e
}

// Sets the value of the element

func (e *SVGSTYLEElement) DATASTAR_BIND(key string, expression string) *SVGSTYLEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key = fmt.Sprintf("data-bind-%s", key)

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGSTYLEElement) IfDATASTAR_BIND(condition bool, key string, expression string) *SVGSTYLEElement {
	if condition {
		e.DATASTAR_BIND(key, expression)
	}
	return e
}

// Remove the attribute DATASTAR_BIND from the element.
func (e *SVGSTYLEElement) DATASTAR_BINDRemove() *SVGSTYLEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-bind")
	return e
}

// Sets the value of the element

func (e *SVGSTYLEElement) DATASTAR_MODEL(expression string) *SVGSTYLEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-model"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGSTYLEElement) IfDATASTAR_MODEL(condition bool, expression string) *SVGSTYLEElement {
	if condition {
		e.DATASTAR_MODEL(expression)
	}
	return e
}

// Remove the attribute DATASTAR_MODEL from the element.
func (e *SVGSTYLEElement) DATASTAR_MODELRemove() *SVGSTYLEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-model")
	return e
}

// Sets the textContent of the element

func (e *SVGSTYLEElement) DATASTAR_TEXT(expression string) *SVGSTYLEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-text"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGSTYLEElement) IfDATASTAR_TEXT(condition bool, expression string) *SVGSTYLEElement {
	if condition {
		e.DATASTAR_TEXT(expression)
	}
	return e
}

// Remove the attribute DATASTAR_TEXT from the element.
func (e *SVGSTYLEElement) DATASTAR_TEXTRemove() *SVGSTYLEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-text")
	return e
}

// Sets the event handler of the element

type SVGStyleOnMod customDataKeyModifier

// Debounces the event handler
func SVGStyleOnModDebounce(
	d time.Duration,
) SVGStyleOnMod {
	return func() string {
		return fmt.Sprintf("debounce_%dms", d.Milliseconds())
	}
}

// Throttles the event handler
func SVGStyleOnModThrottle(
	d time.Duration,
) SVGStyleOnMod {
	return func() string {
		return fmt.Sprintf("throttle_%dms", d.Milliseconds())
	}
}

func (e *SVGSTYLEElement) DATASTAR_ON(key string, expression string, modifiers ...SVGStyleOnMod) *SVGSTYLEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key = fmt.Sprintf("data-on-%s", key)

	customMods := lo.Map(modifiers, func(m SVGStyleOnMod, i int) customDataKeyModifier {
		return customDataKeyModifier(m)
	})
	key = customDataKey(key, customMods...)
	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGSTYLEElement) IfDATASTAR_ON(condition bool, key string, expression string, modifiers ...SVGStyleOnMod) *SVGSTYLEElement {
	if condition {
		e.DATASTAR_ON(key, expression, modifiers...)
	}
	return e
}

// Remove the attribute DATASTAR_ON from the element.
func (e *SVGSTYLEElement) DATASTAR_ONRemove() *SVGSTYLEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-on")
	return e
}

// Sets the focus of the element

func (e *SVGSTYLEElement) DATASTAR_FOCUSSet(b bool) *SVGSTYLEElement {
	key := "data-focus"
	e.BoolAttributes.Set(key, b)
	return e
}

func (e *SVGSTYLEElement) DATASTAR_FOCUS() *SVGSTYLEElement {
	return e.DATASTAR_FOCUSSet(true)
}

// Sets the header of for fetch requests

func (e *SVGSTYLEElement) DATASTAR_HEADER(key string, expression string) *SVGSTYLEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key = fmt.Sprintf("data-header-%s", key)

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGSTYLEElement) IfDATASTAR_HEADER(condition bool, key string, expression string) *SVGSTYLEElement {
	if condition {
		e.DATASTAR_HEADER(key, expression)
	}
	return e
}

// Remove the attribute DATASTAR_HEADER from the element.
func (e *SVGSTYLEElement) DATASTAR_HEADERRemove() *SVGSTYLEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-header")
	return e
}

// Sets the indicator selector for fetch requests

func (e *SVGSTYLEElement) DATASTAR_FETCH_INDICATOR(expression string) *SVGSTYLEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-fetch-indicator"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGSTYLEElement) IfDATASTAR_FETCH_INDICATOR(condition bool, expression string) *SVGSTYLEElement {
	if condition {
		e.DATASTAR_FETCH_INDICATOR(expression)
	}
	return e
}

// Remove the attribute DATASTAR_FETCH_INDICATOR from the element.
func (e *SVGSTYLEElement) DATASTAR_FETCH_INDICATORRemove() *SVGSTYLEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-fetch-indicator")
	return e
}

// Sets the visibility of the element

func (e *SVGSTYLEElement) DATASTAR_SHOWSet(b bool) *SVGSTYLEElement {
	key := "data-show"
	e.BoolAttributes.Set(key, b)
	return e
}

func (e *SVGSTYLEElement) DATASTAR_SHOW() *SVGSTYLEElement {
	return e.DATASTAR_SHOWSet(true)
}

// Triggers the callback when the element intersects the viewport

func (e *SVGSTYLEElement) DATASTAR_INTERSECTS(expression string) *SVGSTYLEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-intersects"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGSTYLEElement) IfDATASTAR_INTERSECTS(condition bool, expression string) *SVGSTYLEElement {
	if condition {
		e.DATASTAR_INTERSECTS(expression)
	}
	return e
}

// Remove the attribute DATASTAR_INTERSECTS from the element.
func (e *SVGSTYLEElement) DATASTAR_INTERSECTSRemove() *SVGSTYLEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-intersects")
	return e
}

// Teleports the element to the given selector

func (e *SVGSTYLEElement) DATASTAR_TELEPORTSet(b bool) *SVGSTYLEElement {
	key := "data-teleport"
	e.BoolAttributes.Set(key, b)
	return e
}

func (e *SVGSTYLEElement) DATASTAR_TELEPORT() *SVGSTYLEElement {
	return e.DATASTAR_TELEPORTSet(true)
}

// Scrolls the element into view

func (e *SVGSTYLEElement) DATASTAR_SCROLL_INTO_VIEWSet(b bool) *SVGSTYLEElement {
	key := "data-scroll-into-view"
	e.BoolAttributes.Set(key, b)
	return e
}

func (e *SVGSTYLEElement) DATASTAR_SCROLL_INTO_VIEW() *SVGSTYLEElement {
	return e.DATASTAR_SCROLL_INTO_VIEWSet(true)
}

// Setup the ViewTransitionAPI for the element

func (e *SVGSTYLEElement) DATASTAR_VIEW_TRANSITION(expression string) *SVGSTYLEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-view-transition"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGSTYLEElement) IfDATASTAR_VIEW_TRANSITION(condition bool, expression string) *SVGSTYLEElement {
	if condition {
		e.DATASTAR_VIEW_TRANSITION(expression)
	}
	return e
}

// Remove the attribute DATASTAR_VIEW_TRANSITION from the element.
func (e *SVGSTYLEElement) DATASTAR_VIEW_TRANSITIONRemove() *SVGSTYLEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-view-transition")
	return e
}
