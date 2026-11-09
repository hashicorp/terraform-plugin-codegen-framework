// Copyright IBM Corp. 2023, 2026
// SPDX-License-Identifier: MPL-2.0

package schema

import (
	"fmt"
	"strings"
)

// NestedTypeName holds the identifier used to name the custom Type and Value
// types generated for a nested attribute or block, when it differs from the
// attribute or block name.
type NestedTypeName struct {
	TypeName FrameworkIdentifier
}

func (n NestedTypeName) GoTypeName(name string) FrameworkIdentifier {
	if n.TypeName != "" {
		return n.TypeName
	}

	return FrameworkIdentifier(name)
}

type GoTypeNamer interface {
	GoTypeName(name string) FrameworkIdentifier
}

type attributeTypeNameSetter interface {
	WithTypeName(FrameworkIdentifier) GeneratorAttribute
}

type blockTypeNameSetter interface {
	WithTypeName(FrameworkIdentifier) GeneratorBlock
}

func goTypeName(v any, name string) FrameworkIdentifier {
	if n, ok := v.(GoTypeNamer); ok {
		return n.GoTypeName(name)
	}

	return FrameworkIdentifier(name)
}

// TypeNames returns the type names of the supplied attributes and blocks, keyed
// by attribute or block name.
func TypeNames(attributes GeneratorAttributes, blocks GeneratorBlocks) map[FrameworkIdentifier]FrameworkIdentifier {
	typeNames := make(map[FrameworkIdentifier]FrameworkIdentifier)

	for k, v := range attributes {
		typeNames[FrameworkIdentifier(k)] = goTypeName(v, k)
	}

	for k, v := range blocks {
		typeNames[FrameworkIdentifier(k)] = goTypeName(v, k)
	}

	return typeNames
}

type typeNameNode struct {
	path     string
	key      string
	parent   *typeNameNode
	prefixed bool
	set      func(FrameworkIdentifier)
}

func (n *typeNameNode) typeName() FrameworkIdentifier {
	if n.prefixed {
		return n.parent.typeName() + "_" + FrameworkIdentifier(n.key)
	}

	return FrameworkIdentifier(n.key)
}

// ResolveTypeNames prefixes the type names of nested attributes and blocks whose
// generated custom Type and Value types would otherwise collide with each other,
// using the type name of their parent. Non-colliding names are left unchanged.
func (g GeneratorSchema) ResolveTypeNames() error {
	var nodes []*typeNameNode

	collectTypeNameNodes(g.Attributes, g.Blocks, nil, &nodes)

	for changed := true; changed; {
		changed = false

		for _, group := range typeNameCollisions(nodes) {
			for _, n := range group {
				for a := n; a != nil; a = a.parent {
					if !a.prefixed && a.parent != nil {
						a.prefixed = true
						changed = true

						break
					}
				}
			}
		}
	}

	if groups := typeNameCollisions(nodes); len(groups) > 0 {
		var msgs []string

		for _, group := range groups {
			paths := make([]string, 0, len(group))

			for _, n := range group {
				paths = append(paths, n.path)
			}

			msgs = append(msgs, fmt.Sprintf("%sType/%sValue for %s", group[0].typeName().ToPascalCase(), group[0].typeName().ToPascalCase(), strings.Join(paths, ", ")))
		}

		return fmt.Errorf("generated type names collide: %s", strings.Join(msgs, "; "))
	}

	for _, n := range nodes {
		if n.prefixed {
			n.set(n.typeName())
		}
	}

	return nil
}

// ResolveTypeNames calls ResolveTypeNames on each schema.
func ResolveTypeNames(schemas map[string]GeneratorSchema) error {
	for name, s := range schemas {
		if err := s.ResolveTypeNames(); err != nil {
			return fmt.Errorf("%s: %w", name, err)
		}
	}

	return nil
}

func collectTypeNameNodes(attributes GeneratorAttributes, blocks GeneratorBlocks, parent *typeNameNode, nodes *[]*typeNameNode) {
	for _, k := range attributes.SortedKeys() {
		s, ok := attributes[k].(attributeTypeNameSetter)

		if !ok {
			continue
		}

		n := newTypeNameNode(k, parent, func(t FrameworkIdentifier) {
			attributes[k] = s.WithTypeName(t)
		})

		*nodes = append(*nodes, n)

		collectChildTypeNameNodes(attributes[k], n, nodes)
	}

	for _, k := range blocks.SortedKeys() {
		s, ok := blocks[k].(blockTypeNameSetter)

		if !ok {
			continue
		}

		n := newTypeNameNode(k, parent, func(t FrameworkIdentifier) {
			blocks[k] = s.WithTypeName(t)
		})

		*nodes = append(*nodes, n)

		collectChildTypeNameNodes(blocks[k], n, nodes)
	}
}

func newTypeNameNode(key string, parent *typeNameNode, set func(FrameworkIdentifier)) *typeNameNode {
	path := key

	if parent != nil {
		path = parent.path + "." + key
	}

	return &typeNameNode{path: path, key: key, parent: parent, set: set}
}

func collectChildTypeNameNodes(v any, parent *typeNameNode, nodes *[]*typeNameNode) {
	var attributes GeneratorAttributes
	var blocks GeneratorBlocks

	if a, ok := v.(Attributes); ok {
		attributes = a.GetAttributes()
	}

	if b, ok := v.(Blocks); ok {
		blocks = b.GetBlocks()
	}

	collectTypeNameNodes(attributes, blocks, parent, nodes)
}

func typeNameCollisions(nodes []*typeNameNode) [][]*typeNameNode {
	var order []string
	groups := make(map[string][]*typeNameNode)

	for _, n := range nodes {
		name := n.typeName().ToPascalCase()

		if _, ok := groups[name]; !ok {
			order = append(order, name)
		}

		groups[name] = append(groups[name], n)
	}

	var collisions [][]*typeNameNode

	for _, name := range order {
		if len(groups[name]) > 1 {
			collisions = append(collisions, groups[name])
		}
	}

	return collisions
}
