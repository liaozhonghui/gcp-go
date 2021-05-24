package typeassesert

import "fmt"

type Container []interface{}

func (c *Container) Get() interface{} {
	elem := (*c)[0]
	*c = (*c)[1:]
	return elem
}

func (c *Container) Put(elem interface{}) {
	*c = append(*c, elem)
}

func main() {
	intContainer := &Container{}
	intContainer.Put(7)
	intContainer.Put(42)
	elem, ok := intContainer.Get().(int)
	if !ok {
		fmt.Println("unable to read an int from intContainer.")
	}
	fmt.Printf("type assert example: %d, (%T)\n", elem, elem)
}
