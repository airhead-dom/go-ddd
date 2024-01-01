package mapper

type Mapper interface {
	Map(interface{}, interface{}) []error
}
