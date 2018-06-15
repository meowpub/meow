package ns

func AS(postfix string) string {
	return "https://www.w3.org/ns/activitystreams#" + postfix
}

func LDP(postfix string) string {
	return "http://www.w3.org/ns/ldp#" + postfix
}
