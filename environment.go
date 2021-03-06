/*
 * Copyright 2018-2020 the original author or authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package libcnb

import (
	"fmt"
)

// Environment represents the file-based environment variable specification.
type Environment map[string]string

// Append formats using the default formats for its operands and appends the value of this environment variable to any
// previous declarations of the value without any delimitation. Spaces are added between operands when neither is a
// string. If delimitation is important during concatenation, callers are required to add it.
func (e Environment) Append(name string, a ...interface{}) {
	e[fmt.Sprintf("%s.append", name)] = fmt.Sprint(a...)
}

// Appendf formats according to a format specifier and appends the value of this environment variable to any previous
// declarations of the value without any delimitation.  If delimitation is important during concatenation, callers are
// required to add it.
func (e Environment) Appendf(name string, format string, a ...interface{}) {
	e[fmt.Sprintf("%s.append", name)] = fmt.Sprintf(format, a...)
}

// Default formats using the default formats for its operands and sets a default for an environment variable with this
// value. Spaces are added between operands when neither is a string.
func (e Environment) Default(name string, a ...interface{}) {
	e[fmt.Sprintf("%s.default", name)] = fmt.Sprint(a...)
}

// Defaultf formats according to a format specifier and sets a default for an environment variable with this value.
func (e Environment) Defaultf(name string, format string, a ...interface{}) {
	e[fmt.Sprintf("%s.default", name)] = fmt.Sprintf(format, a...)
}

// Delimiter sets a delimiter for an environment variable with this value.
func (e Environment) Delimiter(name string, delimiter string) {
	e[fmt.Sprintf("%s.delim", name)] = delimiter
}

// Override formats using the default formats for its operands and overrides any existing value for an environment
// variable with this value. Spaces are added between operands when neither is a string.
func (e Environment) Override(name string, a ...interface{}) {
	e[fmt.Sprintf("%s.override", name)] = fmt.Sprint(a...)
}

// Overridef formats according to a format specifier and overrides any existing value for an environment variable with
// this value.
func (e Environment) Overridef(name string, format string, a ...interface{}) {
	e[fmt.Sprintf("%s.override", name)] = fmt.Sprintf(format, a...)
}

// Prepend formats using the default formats for its operands and prepends the value of this environment variable to any
// previous declarations of the value without any delimitation.  Spaces are added between operands when neither is a
// string. If delimitation is important during concatenation, callers are required to add it.
func (e Environment) Prepend(name string, a ...interface{}) {
	e[fmt.Sprintf("%s.prepend", name)] = fmt.Sprint(a...)
}

// Prependf formats using the default formats for its operands and prepends the value of this environment variable to
// any previous declarations of the value without any delimitation.  If delimitation is important during concatenation,
// callers are required to add it.
func (e Environment) Prependf(name string, format string, a ...interface{}) {
	e[fmt.Sprintf("%s.prepend", name)] = fmt.Sprintf(format, a...)
}

// PrependPath formats using the default formats for its operands and prepends the value of this environment variable to
// any previous declarations of the value using the OS path delimiter. Spaces are added between operands when neither is
// a string.
func (e Environment) PrependPath(name string, a ...interface{}) {
	e[name] = fmt.Sprint(a...)
}

// PrependPathf formats according to a format specifier and prepends the value of this environment variable to any
// previous declarations of the value using the OS path delimiter.
func (e Environment) PrependPathf(name string, format string, a ...interface{}) {
	e[name] = fmt.Sprintf(format, a...)
}
