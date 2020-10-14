package goo

type Tag struct {
	Name  string
	Value string
}

func (tag Tag) String() string {
	return tag.Name + "->" + tag.Value
}

type Taggable interface {
	GetTags() []Tag
	GetTagByName(name string) (Tag, error)
}
