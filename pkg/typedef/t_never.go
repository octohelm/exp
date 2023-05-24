package typedef

func Never() Type[any] {
	return &neverType{}
}

type neverType struct {
}

func (t *neverType) Kind() string {
	return "never"
}
