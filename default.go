package validation

import (
	"regexp"
)

var Default = New()

func Clear() {
	Default.Clear()
}

func HasError() bool {
	return Default.HasError()
}

func HasErrors() bool {
	return Default.HasErrors()
}

// ErrorMap Return the errors mapped by key.
// If there are multiple validation errors associated with a single key, the
// first one "wins".  (Typically the first validation will be the more basic).
func ErrorMap() map[string]*ValidationError {
	return Default.ErrorsMap()
}

func Error() *ValidationError {
	return Default.Error()
}

// Required Test that the argument is non-nil and non-empty (if string or list)
func Required(obj interface{}, key string) *ValidationResult {
	return Default.Required(obj, key)
}

// Min Test that the obj is greater than min if obj's type is int
func Min(obj interface{}, min int, key string) *ValidationResult {
	return Default.Min(obj, min, key)
}

// Max Test that the obj is less than max if obj's type is int
func Max(obj interface{}, max int, key string) *ValidationResult {
	return Default.Max(obj, max, key)
}

// Range Test that the obj is between mni and max if obj's type is int
func Range(obj interface{}, min, max int, key string) *ValidationResult {
	return Default.Range(obj, min, max, key)
}

func MinSize(obj interface{}, min int, key string) *ValidationResult {
	return Default.MinSize(obj, min, key)
}

func MaxSize(obj interface{}, max int, key string) *ValidationResult {
	return Default.MaxSize(obj, max, key)
}

func Length(obj interface{}, n int, key string) *ValidationResult {
	return Default.Length(obj, n, key)
}

func Alpha(obj interface{}, key string) *ValidationResult {
	return Default.Alpha(obj, key)
}

func Numeric(obj interface{}, key string) *ValidationResult {
	return Default.Numeric(obj, key)
}

func AlphaNumeric(obj interface{}, key string) *ValidationResult {
	return Default.AlphaNumeric(obj, key)
}

func Match(obj interface{}, regex *regexp.Regexp, key string) *ValidationResult {
	return Default.Match(obj, key)
}

func NoMatch(obj interface{}, regex *regexp.Regexp, key string) *ValidationResult {
	return Default.NoMatch(obj, key)
}

func AlphaDash(obj interface{}, key string) *ValidationResult {
	return Default.AlphaDash(obj, key)
}

func Email(obj interface{}, key string) *ValidationResult {
	return Default.Email(obj, key)
}

func Ip(obj interface{}, key string) *ValidationResult {
	return Default.Ip(obj, key)
}

func Base64(obj interface{}, key string) *ValidationResult {
	return Default.Base64(obj, key)
}

func Mobile(obj interface{}, key string) *ValidationResult {
	return Default.Mobile(obj, key)
}

func Tel(obj interface{}, key string) *ValidationResult {
	return Default.Tel(obj, key)
}

func Phone(obj interface{}, key string) *ValidationResult {
	return Default.Phone(obj, key)
}

func ZipCode(obj interface{}, key string) *ValidationResult {
	return Default.ZipCode(obj, key)
}

func SetError(fieldName string, errMsg string) *ValidationError {
	return Default.SetError(fieldName, errMsg)
}

// Check Apply a group of validators to a field, in order, and return the
// ValidationResult from the first one that fails, or the last one that
// succeeds.
func Check(obj interface{}, checks ...Validator) *ValidationResult {
	return Default.Check(obj, checks...)
}

// Valid the obj parameter must be a struct or a struct pointer
func Valid(obj interface{}, args ...string) (bool, error) {
	return Default.Valid(obj, args...)
}

// ValidResult the obj parameter must be a struct or a struct pointer
func ValidResult(obj interface{}, args ...string) (bool, map[string]string) {
	return Default.ValidResult(obj, args...)
}

func ErrMap() map[string]string {
	return Default.ErrMap()
}

func ValidSimple(name string, val string, rule string) (bool, error) {
	return Default.ValidSimple(name, val, rule)
}

func OkBy(name string, val string, rule string) bool {
	return Default.OkBy(name, val, rule)
}

func Ok(obj interface{}, args ...string) bool {
	return Default.Ok(obj, args...)
}
