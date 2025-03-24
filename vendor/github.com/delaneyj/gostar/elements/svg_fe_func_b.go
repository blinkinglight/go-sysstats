// DO NOT EDIT THIS FILE. IT IS AUTOMATICALLY GENERATED.
// package svg feFuncB is generated from configuration file.
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

// The <feFuncB> SVG filter primitive defines the transfer function for the blue
// component of the input graphic of its parent <feComponentTransfer> element.
type SVGFEFUNCBElement struct {
	*Element
}

// Create a new SVGFEFUNCBElement element.
// This will create a new element with the tag
// "feFuncB" during rendering.
func SVG_FEFUNCB(children ...ElementRenderer) *SVGFEFUNCBElement {
	e := NewElement("feFuncB", children...)
	e.IsSelfClosing = false
	e.Descendants = children

	return &SVGFEFUNCBElement{Element: e}
}

func (e *SVGFEFUNCBElement) Children(children ...ElementRenderer) *SVGFEFUNCBElement {
	e.Descendants = append(e.Descendants, children...)
	return e
}

func (e *SVGFEFUNCBElement) IfChildren(condition bool, children ...ElementRenderer) *SVGFEFUNCBElement {
	if condition {
		e.Descendants = append(e.Descendants, children...)
	}
	return e
}

func (e *SVGFEFUNCBElement) TernChildren(condition bool, trueChildren, falseChildren ElementRenderer) *SVGFEFUNCBElement {
	if condition {
		e.Descendants = append(e.Descendants, trueChildren)
	} else {
		e.Descendants = append(e.Descendants, falseChildren)
	}
	return e
}

func (e *SVGFEFUNCBElement) Attr(name string, value string) *SVGFEFUNCBElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set(name, value)
	return e
}

func (e *SVGFEFUNCBElement) Attrs(attrs ...string) *SVGFEFUNCBElement {
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

func (e *SVGFEFUNCBElement) AttrsMap(attrs map[string]string) *SVGFEFUNCBElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	for k, v := range attrs {
		e.StringAttributes.Set(k, v)
	}
	return e
}

func (e *SVGFEFUNCBElement) Text(text string) *SVGFEFUNCBElement {
	e.Descendants = append(e.Descendants, Text(text))
	return e
}

func (e *SVGFEFUNCBElement) TextF(format string, args ...any) *SVGFEFUNCBElement {
	return e.Text(fmt.Sprintf(format, args...))
}

func (e *SVGFEFUNCBElement) IfText(condition bool, text string) *SVGFEFUNCBElement {
	if condition {
		e.Descendants = append(e.Descendants, Text(text))
	}
	return e
}

func (e *SVGFEFUNCBElement) IfTextF(condition bool, format string, args ...any) *SVGFEFUNCBElement {
	if condition {
		e.Descendants = append(e.Descendants, Text(fmt.Sprintf(format, args...)))
	}
	return e
}

func (e *SVGFEFUNCBElement) Escaped(text string) *SVGFEFUNCBElement {
	e.Descendants = append(e.Descendants, Escaped(text))
	return e
}

func (e *SVGFEFUNCBElement) IfEscaped(condition bool, text string) *SVGFEFUNCBElement {
	if condition {
		e.Descendants = append(e.Descendants, Escaped(text))
	}
	return e
}

func (e *SVGFEFUNCBElement) EscapedF(format string, args ...any) *SVGFEFUNCBElement {
	return e.Escaped(fmt.Sprintf(format, args...))
}

func (e *SVGFEFUNCBElement) IfEscapedF(condition bool, format string, args ...any) *SVGFEFUNCBElement {
	if condition {
		e.Descendants = append(e.Descendants, EscapedF(format, args...))
	}
	return e
}

func (e *SVGFEFUNCBElement) CustomData(key, value string) *SVGFEFUNCBElement {
	if e.CustomDataAttributes == nil {
		e.CustomDataAttributes = treemap.New[string, string]()
	}
	e.CustomDataAttributes.Set(key, value)
	return e
}

func (e *SVGFEFUNCBElement) IfCustomData(condition bool, key, value string) *SVGFEFUNCBElement {
	if condition {
		e.CustomData(key, value)
	}
	return e
}

func (e *SVGFEFUNCBElement) CustomDataF(key, format string, args ...any) *SVGFEFUNCBElement {
	return e.CustomData(key, fmt.Sprintf(format, args...))
}

func (e *SVGFEFUNCBElement) IfCustomDataF(condition bool, key, format string, args ...any) *SVGFEFUNCBElement {
	if condition {
		e.CustomData(key, fmt.Sprintf(format, args...))
	}
	return e
}

func (e *SVGFEFUNCBElement) CustomDataRemove(key string) *SVGFEFUNCBElement {
	if e.CustomDataAttributes == nil {
		return e
	}
	e.CustomDataAttributes.Del(key)
	return e
}

// The type of transfer function.
func (e *SVGFEFUNCBElement) TYPE(c SVGFeFuncBTypeChoice) *SVGFEFUNCBElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("type", string(c))
	return e
}

type SVGFeFuncBTypeChoice string

const (
	// The type of transfer function.
	SVGFeFuncBType_identity SVGFeFuncBTypeChoice = "identity"
	// The type of transfer function.
	SVGFeFuncBType_table SVGFeFuncBTypeChoice = "table"
	// The type of transfer function.
	SVGFeFuncBType_discrete SVGFeFuncBTypeChoice = "discrete"
	// The type of transfer function.
	SVGFeFuncBType_linear SVGFeFuncBTypeChoice = "linear"
	// The type of transfer function.
	SVGFeFuncBType_gamma SVGFeFuncBTypeChoice = "gamma"
)

// Remove the attribute TYPE from the element.
func (e *SVGFEFUNCBElement) TYPERemove(c SVGFeFuncBTypeChoice) *SVGFEFUNCBElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("type")
	return e
}

// Contains the list of <number>s that define the lookup table
// Values must be in the 0-1 range and be equally spaced
// There must be at least two values.
func (e *SVGFEFUNCBElement) TABLE_VALUES(s string) *SVGFEFUNCBElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("tableValues", s)
	return e
}

func (e *SVGFEFUNCBElement) TABLE_VALUESF(format string, args ...any) *SVGFEFUNCBElement {
	return e.TABLE_VALUES(fmt.Sprintf(format, args...))
}

func (e *SVGFEFUNCBElement) IfTABLE_VALUES(condition bool, s string) *SVGFEFUNCBElement {
	if condition {
		e.TABLE_VALUES(s)
	}
	return e
}

func (e *SVGFEFUNCBElement) IfTABLE_VALUESF(condition bool, format string, args ...any) *SVGFEFUNCBElement {
	if condition {
		e.TABLE_VALUES(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute TABLE_VALUES from the element.
func (e *SVGFEFUNCBElement) TABLE_VALUESRemove(s string) *SVGFEFUNCBElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("tableValues")
	return e
}

func (e *SVGFEFUNCBElement) TABLE_VALUESRemoveF(format string, args ...any) *SVGFEFUNCBElement {
	return e.TABLE_VALUESRemove(fmt.Sprintf(format, args...))
}

// The slope attribute indicates the slope of the linear function.
func (e *SVGFEFUNCBElement) SLOPE(f float64) *SVGFEFUNCBElement {
	if e.FloatAttributes == nil {
		e.FloatAttributes = treemap.New[string, float64]()
	}
	e.FloatAttributes.Set("slope", f)
	return e
}

func (e *SVGFEFUNCBElement) IfSLOPE(condition bool, f float64) *SVGFEFUNCBElement {
	if condition {
		e.SLOPE(f)
	}
	return e
}

// The intercept attribute indicates the intercept of the linear function.
func (e *SVGFEFUNCBElement) INTERCEPT(f float64) *SVGFEFUNCBElement {
	if e.FloatAttributes == nil {
		e.FloatAttributes = treemap.New[string, float64]()
	}
	e.FloatAttributes.Set("intercept", f)
	return e
}

func (e *SVGFEFUNCBElement) IfINTERCEPT(condition bool, f float64) *SVGFEFUNCBElement {
	if condition {
		e.INTERCEPT(f)
	}
	return e
}

// The amplitude attribute indicates the amplitude of the cubic function.
func (e *SVGFEFUNCBElement) AMPLITUDE(f float64) *SVGFEFUNCBElement {
	if e.FloatAttributes == nil {
		e.FloatAttributes = treemap.New[string, float64]()
	}
	e.FloatAttributes.Set("amplitude", f)
	return e
}

func (e *SVGFEFUNCBElement) IfAMPLITUDE(condition bool, f float64) *SVGFEFUNCBElement {
	if condition {
		e.AMPLITUDE(f)
	}
	return e
}

// The exponent attribute indicates the exponent of the exponential function.
func (e *SVGFEFUNCBElement) EXPONENT(f float64) *SVGFEFUNCBElement {
	if e.FloatAttributes == nil {
		e.FloatAttributes = treemap.New[string, float64]()
	}
	e.FloatAttributes.Set("exponent", f)
	return e
}

func (e *SVGFEFUNCBElement) IfEXPONENT(condition bool, f float64) *SVGFEFUNCBElement {
	if condition {
		e.EXPONENT(f)
	}
	return e
}

// The offset attribute indicates the offset of the function.
func (e *SVGFEFUNCBElement) OFFSET(f float64) *SVGFEFUNCBElement {
	if e.FloatAttributes == nil {
		e.FloatAttributes = treemap.New[string, float64]()
	}
	e.FloatAttributes.Set("offset", f)
	return e
}

func (e *SVGFEFUNCBElement) IfOFFSET(condition bool, f float64) *SVGFEFUNCBElement {
	if condition {
		e.OFFSET(f)
	}
	return e
}

// Specifies a unique id for an element
func (e *SVGFEFUNCBElement) ID(s string) *SVGFEFUNCBElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("id", s)
	return e
}

func (e *SVGFEFUNCBElement) IDF(format string, args ...any) *SVGFEFUNCBElement {
	return e.ID(fmt.Sprintf(format, args...))
}

func (e *SVGFEFUNCBElement) IfID(condition bool, s string) *SVGFEFUNCBElement {
	if condition {
		e.ID(s)
	}
	return e
}

func (e *SVGFEFUNCBElement) IfIDF(condition bool, format string, args ...any) *SVGFEFUNCBElement {
	if condition {
		e.ID(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute ID from the element.
func (e *SVGFEFUNCBElement) IDRemove(s string) *SVGFEFUNCBElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("id")
	return e
}

func (e *SVGFEFUNCBElement) IDRemoveF(format string, args ...any) *SVGFEFUNCBElement {
	return e.IDRemove(fmt.Sprintf(format, args...))
}

// Specifies one or more classnames for an element (refers to a class in a style
// sheet)
func (e *SVGFEFUNCBElement) CLASS(s ...string) *SVGFEFUNCBElement {
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

func (e *SVGFEFUNCBElement) IfCLASS(condition bool, s ...string) *SVGFEFUNCBElement {
	if condition {
		e.CLASS(s...)
	}
	return e
}

// Remove the attribute CLASS from the element.
func (e *SVGFEFUNCBElement) CLASSRemove(s ...string) *SVGFEFUNCBElement {
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
func (e *SVGFEFUNCBElement) STYLEF(k string, format string, args ...any) *SVGFEFUNCBElement {
	return e.STYLE(k, fmt.Sprintf(format, args...))
}

func (e *SVGFEFUNCBElement) IfSTYLE(condition bool, k string, v string) *SVGFEFUNCBElement {
	if condition {
		e.STYLE(k, v)
	}
	return e
}

func (e *SVGFEFUNCBElement) STYLE(k string, v string) *SVGFEFUNCBElement {
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

func (e *SVGFEFUNCBElement) IfSTYLEF(condition bool, k string, format string, args ...any) *SVGFEFUNCBElement {
	if condition {
		e.STYLE(k, fmt.Sprintf(format, args...))
	}
	return e
}

// Add the attributes in the map to the element.
func (e *SVGFEFUNCBElement) STYLEMap(m map[string]string) *SVGFEFUNCBElement {
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
func (e *SVGFEFUNCBElement) STYLEPairs(pairs ...string) *SVGFEFUNCBElement {
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

func (e *SVGFEFUNCBElement) IfSTYLEPairs(condition bool, pairs ...string) *SVGFEFUNCBElement {
	if condition {
		e.STYLEPairs(pairs...)
	}
	return e
}

// Remove the attribute STYLE from the element.
func (e *SVGFEFUNCBElement) STYLERemove(keys ...string) *SVGFEFUNCBElement {
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

func (e *SVGFEFUNCBElement) DATASTAR_STORE(v any) *SVGFEFUNCBElement {
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

func (e *SVGFEFUNCBElement) DATASTAR_REF(expression string) *SVGFEFUNCBElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-ref"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEFUNCBElement) IfDATASTAR_REF(condition bool, expression string) *SVGFEFUNCBElement {
	if condition {
		e.DATASTAR_REF(expression)
	}
	return e
}

// Remove the attribute DATASTAR_REF from the element.
func (e *SVGFEFUNCBElement) DATASTAR_REFRemove() *SVGFEFUNCBElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-ref")
	return e
}

// Sets the value of the element

func (e *SVGFEFUNCBElement) DATASTAR_BIND(key string, expression string) *SVGFEFUNCBElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key = fmt.Sprintf("data-bind-%s", key)

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEFUNCBElement) IfDATASTAR_BIND(condition bool, key string, expression string) *SVGFEFUNCBElement {
	if condition {
		e.DATASTAR_BIND(key, expression)
	}
	return e
}

// Remove the attribute DATASTAR_BIND from the element.
func (e *SVGFEFUNCBElement) DATASTAR_BINDRemove() *SVGFEFUNCBElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-bind")
	return e
}

// Sets the value of the element

func (e *SVGFEFUNCBElement) DATASTAR_MODEL(expression string) *SVGFEFUNCBElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-model"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEFUNCBElement) IfDATASTAR_MODEL(condition bool, expression string) *SVGFEFUNCBElement {
	if condition {
		e.DATASTAR_MODEL(expression)
	}
	return e
}

// Remove the attribute DATASTAR_MODEL from the element.
func (e *SVGFEFUNCBElement) DATASTAR_MODELRemove() *SVGFEFUNCBElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-model")
	return e
}

// Sets the textContent of the element

func (e *SVGFEFUNCBElement) DATASTAR_TEXT(expression string) *SVGFEFUNCBElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-text"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEFUNCBElement) IfDATASTAR_TEXT(condition bool, expression string) *SVGFEFUNCBElement {
	if condition {
		e.DATASTAR_TEXT(expression)
	}
	return e
}

// Remove the attribute DATASTAR_TEXT from the element.
func (e *SVGFEFUNCBElement) DATASTAR_TEXTRemove() *SVGFEFUNCBElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-text")
	return e
}

// Sets the event handler of the element

type SVGFeFuncBOnMod customDataKeyModifier

// Debounces the event handler
func SVGFeFuncBOnModDebounce(
	d time.Duration,
) SVGFeFuncBOnMod {
	return func() string {
		return fmt.Sprintf("debounce_%dms", d.Milliseconds())
	}
}

// Throttles the event handler
func SVGFeFuncBOnModThrottle(
	d time.Duration,
) SVGFeFuncBOnMod {
	return func() string {
		return fmt.Sprintf("throttle_%dms", d.Milliseconds())
	}
}

func (e *SVGFEFUNCBElement) DATASTAR_ON(key string, expression string, modifiers ...SVGFeFuncBOnMod) *SVGFEFUNCBElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key = fmt.Sprintf("data-on-%s", key)

	customMods := lo.Map(modifiers, func(m SVGFeFuncBOnMod, i int) customDataKeyModifier {
		return customDataKeyModifier(m)
	})
	key = customDataKey(key, customMods...)
	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEFUNCBElement) IfDATASTAR_ON(condition bool, key string, expression string, modifiers ...SVGFeFuncBOnMod) *SVGFEFUNCBElement {
	if condition {
		e.DATASTAR_ON(key, expression, modifiers...)
	}
	return e
}

// Remove the attribute DATASTAR_ON from the element.
func (e *SVGFEFUNCBElement) DATASTAR_ONRemove() *SVGFEFUNCBElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-on")
	return e
}

// Sets the focus of the element

func (e *SVGFEFUNCBElement) DATASTAR_FOCUSSet(b bool) *SVGFEFUNCBElement {
	key := "data-focus"
	e.BoolAttributes.Set(key, b)
	return e
}

func (e *SVGFEFUNCBElement) DATASTAR_FOCUS() *SVGFEFUNCBElement {
	return e.DATASTAR_FOCUSSet(true)
}

// Sets the header of for fetch requests

func (e *SVGFEFUNCBElement) DATASTAR_HEADER(key string, expression string) *SVGFEFUNCBElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key = fmt.Sprintf("data-header-%s", key)

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEFUNCBElement) IfDATASTAR_HEADER(condition bool, key string, expression string) *SVGFEFUNCBElement {
	if condition {
		e.DATASTAR_HEADER(key, expression)
	}
	return e
}

// Remove the attribute DATASTAR_HEADER from the element.
func (e *SVGFEFUNCBElement) DATASTAR_HEADERRemove() *SVGFEFUNCBElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-header")
	return e
}

// Sets the indicator selector for fetch requests

func (e *SVGFEFUNCBElement) DATASTAR_FETCH_INDICATOR(expression string) *SVGFEFUNCBElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-fetch-indicator"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEFUNCBElement) IfDATASTAR_FETCH_INDICATOR(condition bool, expression string) *SVGFEFUNCBElement {
	if condition {
		e.DATASTAR_FETCH_INDICATOR(expression)
	}
	return e
}

// Remove the attribute DATASTAR_FETCH_INDICATOR from the element.
func (e *SVGFEFUNCBElement) DATASTAR_FETCH_INDICATORRemove() *SVGFEFUNCBElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-fetch-indicator")
	return e
}

// Sets the visibility of the element

func (e *SVGFEFUNCBElement) DATASTAR_SHOWSet(b bool) *SVGFEFUNCBElement {
	key := "data-show"
	e.BoolAttributes.Set(key, b)
	return e
}

func (e *SVGFEFUNCBElement) DATASTAR_SHOW() *SVGFEFUNCBElement {
	return e.DATASTAR_SHOWSet(true)
}

// Triggers the callback when the element intersects the viewport

func (e *SVGFEFUNCBElement) DATASTAR_INTERSECTS(expression string) *SVGFEFUNCBElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-intersects"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEFUNCBElement) IfDATASTAR_INTERSECTS(condition bool, expression string) *SVGFEFUNCBElement {
	if condition {
		e.DATASTAR_INTERSECTS(expression)
	}
	return e
}

// Remove the attribute DATASTAR_INTERSECTS from the element.
func (e *SVGFEFUNCBElement) DATASTAR_INTERSECTSRemove() *SVGFEFUNCBElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-intersects")
	return e
}

// Teleports the element to the given selector

func (e *SVGFEFUNCBElement) DATASTAR_TELEPORTSet(b bool) *SVGFEFUNCBElement {
	key := "data-teleport"
	e.BoolAttributes.Set(key, b)
	return e
}

func (e *SVGFEFUNCBElement) DATASTAR_TELEPORT() *SVGFEFUNCBElement {
	return e.DATASTAR_TELEPORTSet(true)
}

// Scrolls the element into view

func (e *SVGFEFUNCBElement) DATASTAR_SCROLL_INTO_VIEWSet(b bool) *SVGFEFUNCBElement {
	key := "data-scroll-into-view"
	e.BoolAttributes.Set(key, b)
	return e
}

func (e *SVGFEFUNCBElement) DATASTAR_SCROLL_INTO_VIEW() *SVGFEFUNCBElement {
	return e.DATASTAR_SCROLL_INTO_VIEWSet(true)
}

// Setup the ViewTransitionAPI for the element

func (e *SVGFEFUNCBElement) DATASTAR_VIEW_TRANSITION(expression string) *SVGFEFUNCBElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-view-transition"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEFUNCBElement) IfDATASTAR_VIEW_TRANSITION(condition bool, expression string) *SVGFEFUNCBElement {
	if condition {
		e.DATASTAR_VIEW_TRANSITION(expression)
	}
	return e
}

// Remove the attribute DATASTAR_VIEW_TRANSITION from the element.
func (e *SVGFEFUNCBElement) DATASTAR_VIEW_TRANSITIONRemove() *SVGFEFUNCBElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-view-transition")
	return e
}
