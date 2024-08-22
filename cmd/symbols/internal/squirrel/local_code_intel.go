package squirrel

import (
	"context"
	"strings"

	sitter "github.com/smacker/go-tree-sitter"

	"github.com/sourcegraph/sourcegraph/internal/types"
)

// SymbolName is a nominal type for symbol names.
type SymbolName string

// Scope is a mapping from symbol name to symbol.
type Scope = map[SymbolName]*PartialSymbol // pointer for mutability

// PartialSymbol is the same as types.Symbol, but with the refs stored in a map to deduplicate.
type PartialSymbol struct {
	Name  string
	Hover string
	Def   types.Range
	// Store refs as a set to avoid duplicates from some tree-sitter queries.
	Refs map[types.Range]struct{}
}

// LocalCodeIntel computes the local code intel payload, which is a list of symbols.
func (s *SquirrelService) LocalCodeIntel(ctx context.Context, repoCommitPath types.RepoCommitPath) (*types.LocalCodeIntelPayload, error) {
	// Parse the file.
	root, err := s.parse(ctx, repoCommitPath)
	if err != nil {
		return nil, err
	}

	// Collect scopes
	scopes := map[NodeId]Scope{}
	forEachCapture(root.LangSpec.localsQuery, *root, func(nameToNode map[string]Node) {
		if node, ok := nameToNode["scope"]; ok {
			scopes[nodeId(node.Node)] = map[SymbolName]*PartialSymbol{}
			return
		}
	})

	// Collect defs
	forEachCapture(root.LangSpec.localsQuery, *root, func(nameToNode map[string]Node) {
		for captureName, node := range nameToNode {
			// Only collect "definition*" captures.
			if strings.HasPrefix(captureName, "definition") {
				// Find the nearest scope (if it exists).
				for cur := node.Node; cur != nil; cur = cur.Parent() {
					// Found the scope.
					if scope, ok := scopes[nodeId(cur)]; ok {
						// Get the symbol name.
						symbolName := SymbolName(strings.ToValidUTF8(node.Content(node.Contents), "�"))

						// Skip the symbol if it's already defined.
						if _, ok := scope[symbolName]; ok {
							break
						}

						// Put the symbol in the scope.
						scope[symbolName] = &PartialSymbol{
							Name:  string(symbolName),
							Hover: findHover(node),
							Def:   nodeToRange(node.Node),
							Refs:  map[types.Range]struct{}{},
						}

						// Stop walking up the tree.
						break
					}
				}
			}
		}
	})

	// Collect refs by walking the entire tree.
	walk(root.Node, func(node *sitter.Node) {
		// Only collect identifiers.
		if !strings.Contains(node.Type(), "identifier") {
			return
		}

		// Get the symbol name.
		symbolName := SymbolName(node.Content(root.Contents))

		// Find the nearest scope (if it exists).
		for cur := node; cur != nil; cur = cur.Parent() {
			if scope, ok := scopes[nodeId(cur)]; ok {
				// Check if it's in the scope.
				if _, ok := scope[symbolName]; !ok {
					// It's not in this scope, so keep walking up the tree.
					continue
				}

				// Put the ref in the scope.
				scope[symbolName].Refs[nodeToRange(node)] = struct{}{}

				// Done.
				return
			}
		}

		// Did not find the symbol in this file, so ignore it.
	})

	// Collect the symbols.
	symbols := []types.Symbol{}
	for _, scope := range scopes {
		for _, partialSymbol := range scope {
			refs := []types.Range{}
			for ref := range partialSymbol.Refs {
				refs = append(refs, ref)
			}
			symbols = append(symbols, types.Symbol{
				Name:  partialSymbol.Name,
				Hover: partialSymbol.Hover,
				Def:   partialSymbol.Def,
				Refs:  refs,
			})
		}
	}

	return &types.LocalCodeIntelPayload{Symbols: symbols}, nil
}
