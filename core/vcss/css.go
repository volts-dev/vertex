package vcss

import (
	"fmt"
	"regexp"
	"strings"
	"sync"

	"github.com/volts-dev/vertex/js"
)

// CSSResult represents a container for CSS text that can be used to create a CSSStyleSheet.
// CSSResult is the return value of css-tagged template literals and unsafeCSS().
type CSSResult struct {
	// CSSText is the CSS text string
	CSSText string

	// strings holds the template strings array for caching purposes
	strings []string

	// styleSheet is cached CSSStyleSheet (lazily created)
	styleSheet interface{} // In Go, this would be the actual stylesheet if using web components

	// marker to identify CSS results
	isCSSResult bool
}

// CSSResultOrNative can be either a CSSResult or a native CSSStyleSheet
type CSSResultOrNative interface{}

// CSSResultArray is an array of CSSResultOrNative or nested arrays
type CSSResultArray []CSSResultOrNative

// CSSResultGroup is a single CSSResult, CSSStyleSheet, or an array/nested arrays of those
type CSSResultGroup interface{}

var (
	// constructionToken ensures CSSResult can only be created via css() or unsafeCSS()
	constructionToken = struct{}{}

	// cssTagCache caches stylesheets by template strings array
	cssTagCache = &sync.Map{}

	// supportsAdoptingStyleSheets indicates if the browser supports adoptedStyleSheets
	// In Go/server context, this is typically false
	SupportsAdoptingStyleSheets = false
)

// textFromCSSResult converts a CSS result value to text
func textFromCSSResult(value interface{}) (string, error) {
	switch v := value.(type) {
	case *CSSResult:
		if v.isCSSResult {
			return v.CSSText, nil
		}
	case CSSResultGroup:
		if result, ok := v.(*CSSResult); ok && result.isCSSResult {
			return result.CSSText, nil
		}
	case int, int32, int64, float32, float64:
		return fmt.Sprintf("%v", v), nil
	case nil:
		return "", nil
	}

	return "", fmt.Errorf(
		"Value passed to 'css' function must be a 'css' function result: %v. "+
			"Use 'unsafeCSS' to pass non-literal values, but take care to ensure page security.",
		value,
	)
}

// UnsafeCSS wraps a value for interpolation in a css tagged template literal.
// This is unsafe because untrusted CSS text can be used to exfiltrate data.
// Take care to only use this with trusted input.
func UnsafeCSS(value interface{}) *CSSResult {
	var cssText string
	if str, ok := value.(string); ok {
		cssText = str
	} else {
		cssText = fmt.Sprintf("%v", value)
	}

	return &CSSResult{
		CSSText:     cssText,
		strings:     nil,
		isCSSResult: true,
	}
}

// CSS is a template literal tag which can be used to set element styles.
// For security reasons, only literal string values and numbers may be used in
// embedded expressions. To incorporate non-literal values, use unsafeCSS().
//
// Example:
//
//	result := CSS("div { color: red; }")
//	result = CSS([]string{"div { color: ", "; }"}, "red")
func CSS(stringsOrArray interface{}, values ...interface{}) *CSSResult {
	var cssText string
	var strings []string

	// Handle different input formats
	switch v := stringsOrArray.(type) {
	case string:
		// Direct string input
		cssText = v
		strings = []string{v}

	case []string:
		// Array of strings (from template literal)
		strings = v

		if len(strings) == 1 {
			cssText = strings[0]
		} else {
			// Combine strings and values
			result := strings[0]
			for i, val := range values {
				textVal, err := textFromCSSResult(val)
				if err != nil {
					panic(err)
				}
				result += textVal
				if i+1 < len(strings) {
					result += strings[i+1]
				}
			}
			cssText = result
		}

	default:
		panic(fmt.Sprintf("Invalid input to CSS function: %T", v))
	}

	return &CSSResult{
		CSSText:     cssText,
		strings:     strings,
		isCSSResult: true,
	}
}

// CSSTemplate is an alias for CSS to match TypeScript naming convention
var CSSTemplate = CSS

// ToString returns the CSS text as a string
func (cr *CSSResult) ToString() string {
	return cr.CSSText
}

// String implements the Stringer interface
func (cr *CSSResult) String() string {
	return cr.CSSText
}

// StyleSheet returns the cached CSSStyleSheet (for web components context)
func (cr *CSSResult) StyleSheet() interface{} {
	if !SupportsAdoptingStyleSheets {
		return nil
	}
	return cr.styleSheet
}

// SetStyleSheet sets the cached stylesheet
func (cr *CSSResult) SetStyleSheet(sheet interface{}) {
	cr.styleSheet = sheet
}

// GetCSSResultMarker returns the CSS result marker
func (cr *CSSResult) GetCSSResultMarker() bool {
	return cr.isCSSResult
}

// ConcatCSS concatenates multiple CSS results and strings into a single CSSResult
func ConcatCSS(items ...interface{}) *CSSResult {
	var allText strings.Builder

	for _, item := range items {
		switch v := item.(type) {
		case *CSSResult:
			if v != nil && v.isCSSResult {
				allText.WriteString(v.CSSText)
			}
		case string:
			allText.WriteString(v)
		case int, int32, int64, float32, float64:
			fmt.Fprintf(&allText, "%v", v)
		default:
			if text, err := textFromCSSResult(v); err == nil {
				allText.WriteString(text)
			}
		}
	}

	return &CSSResult{
		CSSText:     allText.String(),
		isCSSResult: true,
	}
}

// FlattenCSSResultArray flattens nested CSS result arrays
func FlattenCSSResultArray(arr interface{}) []*CSSResult {
	var results []*CSSResult

	switch v := arr.(type) {
	case []*CSSResult:
		results = append(results, v...)
	case *CSSResult:
		results = append(results, v)
	case []interface{}:
		for _, item := range v {
			results = append(results, FlattenCSSResultArray(item)...)
		}
	}

	return results
}

// NormalizeCSS removes unnecessary whitespace from CSS while preserving functionality
func NormalizeCSS(css string) string {
	// Remove comments
	commentRegex := regexp.MustCompile(`/\*[\s\S]*?\*/`)
	css = commentRegex.ReplaceAllString(css, "")

	// Replace multiple spaces/newlines with single space
	whiteSpaceRegex := regexp.MustCompile(`\s+`)
	css = whiteSpaceRegex.ReplaceAllString(css, " ")

	// Remove space around certain characters
	css = strings.ReplaceAll(css, " {", "{")
	css = strings.ReplaceAll(css, "} ", "}")
	css = strings.ReplaceAll(css, ", ", ",")
	css = strings.ReplaceAll(css, ": ", ":")
	css = strings.ReplaceAll(css, "; ", ";")
	css = strings.ReplaceAll(css, " >", ">")
	css = strings.ReplaceAll(css, "> ", ">")
	css = strings.ReplaceAll(css, " +", "+")
	css = strings.ReplaceAll(css, "+ ", "+")
	css = strings.ReplaceAll(css, " ~", "~")
	css = strings.ReplaceAll(css, "~ ", "~")

	return strings.TrimSpace(css)
}

// ===== 便利函数 =====
// InjectGlobalStyles 在文档中注入全局样式
func InjectGlobalStyles(styles []*CSSResult) {
	if len(styles) == 0 {
		return
	}

	doc := js.Global().Get("document")
	head := doc.Get("head")

	for i, style := range styles {
		styleElement := doc.Call("createElement", "style")
		styleElement.Set("id", fmt.Sprintf("global-style-%d", i))
		styleElement.Set("textContent", style.CSSText)
		head.Call("appendChild", styleElement)
	}
}

// RemoveGlobalStyles 从文档中移除全局样式
func RemoveGlobalStyles(styleIDs []string) {
	doc := js.Global().Get("document")

	for _, id := range styleIDs {
		element := doc.Call("getElementById", id)
		if !element.IsNull() && !element.IsUndefined() {
			element.Call("remove")
		}
	}
}

// GetComponentStylesFromRegistry 从样式注册表中获取组件样式
func GetComponentStylesFromRegistry(componentName string) []*CSSResult {
	var styles []*CSSResult

	allStyles := GetAllGlobalStyles()
	for key, style := range allStyles {
		if fmt.Sprintf("%s-", componentName) == key[:len(componentName)+1] {
			styles = append(styles, style)
		}
	}

	return styles
}
