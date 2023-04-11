package forms

// define new errors type
type errors map[string][]string

// Add implement an add method
func (e errors) Add(field, message string) {
	e[field] = append(e[field], message)
}

// Get implement a get method
func (e errors) Get(field string) string {
	es := e[field]
	if len(es) == 0 {
		return ""
	}
	return es[0]
}
