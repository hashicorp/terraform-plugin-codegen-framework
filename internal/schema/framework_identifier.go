// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package schema

import (
	"regexp"
	"strings"
	"unicode"
	"unicode/utf8"
)

// FrameworkIdentifier is a string that implements helpful methods for validating and converting identifier names that are valid in Terraform Plugin Framework
type FrameworkIdentifier string

// frameworkIdentifierRegex is used to validate that FrameworkIdentifier adheres to the same naming conventions as
// [Terraform Plugin Framework identifiers].
//
// [Terraform Plugin Framework identifiers]: https://github.com/hashicorp/terraform-plugin-framework/blob/e036d9fbab4b72f8ec671a9d3e94649040e3eeb5/internal/fwschema/attribute_name_validation.go#L61
var frameworkIdentifierRegex = regexp.MustCompile("^[a-z_][a-z0-9_]*$")

// snakeLetters will match to the first letter and an underscore followed by a letter
var snakeLetters = regexp.MustCompile("(^[a-z])|_[a-z0-9]")

// Valid will return whether the identifier string is a valid identifier in Terraform Plugin Framework
func (identifier FrameworkIdentifier) Valid() bool {
	return frameworkIdentifierRegex.MatchString(string(identifier))
}

// ToCamelCase will return a camel case formatted string of the identifier.
// Example:
//   - example_resource_thing -> exampleResourceThing
func (identifier FrameworkIdentifier) ToCamelCase() string {
	pascal := identifier.ToPascalCase()

	// Grab first rune and lower case it
	firstLetter, size := utf8.DecodeRuneInString(pascal)
	if firstLetter == utf8.RuneError && size <= 1 {
		return pascal
	}

	return string(unicode.ToLower(firstLetter)) + pascal[size:]
}

// ToPrefixCamelCase will return a camel case formatted string of the identifier,
// prefixed with a camel-cased version of the supplied name if the identifier is
// a generated custom value method name.
// Example:
//   - equal(something) -> somethingEqual
//   - type(something) -> somethingType
func (identifier FrameworkIdentifier) ToPrefixCamelCase(prefix string) string {
	pascalCase := identifier.ToPascalCase()

	methodNames := identifier.methodNames()

	for _, v := range methodNames {
		if pascalCase == v {
			return FrameworkIdentifier(prefix + pascalCase).ToCamelCase()
		}
	}

	return FrameworkIdentifier(pascalCase).ToCamelCase()
}

// ToPascalCase will return a pascal case formatted string of the identifier.
// Example:
//   - example_resource_thing -> ExampleResourceThing
func (identifier FrameworkIdentifier) ToPascalCase() string {
	return snakeLetters.ReplaceAllStringFunc(string(identifier), func(s string) string {
		return strings.ToUpper(strings.Replace(s, "_", "", -1))
	})
}

// ToPrefixPascalCase will return a pascal case formatted string of the identifier,
// prefixed with a pascal-cased version of the supplied name if the identifier is
// a generated custom value method name.
// Example:
//   - equal(something) -> SomethingEqual
//   - type(something) -> SomethingType
func (identifier FrameworkIdentifier) ToPrefixPascalCase(prefix string) string {
	pascalCase := identifier.ToPascalCase()

	methodNames := identifier.methodNames()

	for _, v := range methodNames {
		if pascalCase == v {
			return FrameworkIdentifier(prefix).ToPascalCase() + pascalCase
		}
	}

	return pascalCase
}

// ToString returns the FrameworkIdentifier as a string without any formatting.
// Example:
//   - example_resource_thing -> example_resource_thing
func (identifier FrameworkIdentifier) ToString() string {
	return string(identifier)
}

// methodNames returns a slice containing generated method names for custom value
// types.
func (identifier FrameworkIdentifier) methodNames() []string {
	return []string{
		"AttributeTypes",
		"Equal",
		"IsNull",
		"IsUnknown",
		"String",
		"ToObjectValue",
		"ToTerraformValue",
		"Type",
	}
}
