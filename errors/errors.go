package errors

import "errors"

// MissingRequiredArgumentErr missing required argument error
var MissingRequiredArgumentErr = errors.New("missing required argument")

// ArgumentVdValidateErr argument vid validate error
var ArgumentVdValidateErr = errors.New("argument vid validate error")

// FlagDirValidateErr flag date validate error
var FlagDirValidateErr = errors.New("flag dir validate error")
