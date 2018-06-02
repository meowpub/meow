package secrets

var sessionKey = declare("session", 64)

func SessionKey() []byte { return sessionKey.access() }

var csrfKey = declare("csrf", 64)

func CSRFKey() []byte { return csrfKey.access() }
