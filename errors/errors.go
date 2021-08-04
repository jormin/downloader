package errors

import "errors"

// ErrMissingRequiredArgument missing required argument error
var ErrMissingRequiredArgument = errors.New("missing required argument")

// ErrArgumentVdValidate argument vid validate error
var ErrArgumentVdValidate = errors.New("argument vid validate error")

// ErrFlagDirValidate flag date validate error
var ErrFlagDirValidate = errors.New("flag dir validate error")
