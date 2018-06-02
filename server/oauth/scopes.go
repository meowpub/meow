package oauth

// SpecialScopeWeb is a special scope issued only to the web frontend, through the login page.
var SpecialScopeWeb = &Scope{"web", "", nil}

// Access to your inbox.
var ScopeInbox = &Scope{"inbox", "Read your timeline", nil}

// Access to your outbox.
var ScopeOutbox = &Scope{"outbox", "Read your posts", ScopeOutboxWrite}
var ScopeOutboxWrite = &Scope{"outbox:write", "Post on your behalf", nil}

// Whitelist of scopes that API clients can request. Remember to add new scopes here!
var PublicScopes = []*Scope{
	ScopeInbox,
	ScopeOutbox,
	ScopeOutboxWrite,
}
