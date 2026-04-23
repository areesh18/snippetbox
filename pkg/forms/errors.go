package forms

type errors map[string][]string

// Implement an Add() method to add error messages for a given field to the map
func (e errors) Add(field string, message string) {
	e[field]=append(e[field],message)
}
// Implement a Get() method to retrieve the first error message for a given
func (e errors) Get(field string) string{
	es:=e[field]
	if len(es)==0{
		return ""
	}
	return es[0]
}
