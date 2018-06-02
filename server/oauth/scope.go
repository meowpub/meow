package oauth

// Scope describes an authorization scope.
type Scope struct {
	ID   string // ID, as it can be requested by a client.
	Name string // Human readable name.

	// Scopes can be implied by other scopes, eg. a write scope would imply the corresponding read
	// scope. You can use CheckScope() to recursively resolve scopes and implied scopes.
	ImpliedBy *Scope
}

// CheckScope checks whether a set of scope names includes the Scope, or a scope which implies it.
func CheckScope(has []string, want *Scope) bool {
	var walk func(scope *Scope) bool
	walk = func(scope *Scope) bool {
		if scope == nil {
			return false
		}
		for _, has := range has {
			if has == scope.ID {
				return true
			}
		}
		return walk(scope.ImpliedBy)
	}
	return walk(want)
}
