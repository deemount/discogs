package constants

const (

	// IDREGEX defines query string
	IDREGEX = `{id:[0-9]+$}`

	// PAGEREGEX defines query string
	PAGEREGEX = `{page:[0-9]+$}`

	// PERPAGEREGEX defines query string
	// Notice:
	// Only non-capturing groups are accepted: e.g.
	// (?:pattern:pattern:pattern) instead of (pattern|pattern|pattern)
	PERPAGEREGEX = `{per_page:(?:50:100:250)$}`

	// SORTREGEX defines query string
	// Notice:
	// Only non-capturing groups are accepted: e.g.
	// (?:pattern:pattern:pattern) instead of (pattern|pattern|pattern)
	SORTREGEX = `{sort:(?:year:title:format)$}`

	// SORTORDERREGEX defines query string
	// Notice:
	// Only non-capturing groups are accepted: e.g.
	// (?:pattern:pattern) instead of (pattern|pattern)
	SORTORDERREGEX = `{sort_order:(?:asc:desc)$}`
)
