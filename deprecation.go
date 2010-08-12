package main

import (
	"fmt"
	"regexp"
)

// DeprecationLint is a stateless linter that looks for and reports deprecated
// constructs. New DebrecationLints should be created with NewDeprecationLint
// or a related method.
type DeprecationLint struct {
	Form string // a regular expression representing the form
	Reason string // the reason for deprecation, or an alternative form
}

func NewDeprecationLint(re string, reason string) (lint DeprecationLint) {
	lint.Form, lint.Reason = re, reason
	return
}

func FuncDeprecationLint(fn string, reason string) (lint DeprecationLint) {
	return NewDeprecationLint("[^a-zA-Z0-9_]"+fn+" *\\(", reason)
}


func (l DeprecationLint) Lint(line string) (msg string, err bool) {
	if m, _ := regexp.MatchString(l.Form, line); m {
		err = true
		// XXX once the regexp patch is out, we can give more info
		msg = fmt.Sprintf("deprecated: %s", l.Reason)
	}
	return
}
