package typedef

func Any() Type[any] {
	return &anyType{}
}

type anyType struct {
}

func (t *anyType) Kind() string {
	return "any"
}
