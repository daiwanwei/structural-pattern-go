package forests

import "errors"

var treeFactoryInstance *TreeFactory

type Tree struct {
	PositionX   int
	PositionY   int
	TreeVariety TreeVariety
}

func NewTree(name string) (*Tree, error) {
	factory, err := GetTreeFactory()
	if err != nil {
		return nil, err
	}
	variety, err := factory.GetTreeVariety(name)
	if err != nil {
		return nil, err
	}
	return &Tree{
		0, 0, variety,
	}, nil
}

func (tree *Tree) NewLocation(x, y int) {
	tree.PositionX = x
	tree.PositionY = y
}

type TreeVariety interface {
	GetName() string
	GetColor() string
	GetTexture() string
}

type TreeFactory struct {
	variety map[string]TreeVariety
}

func GetTreeFactory() (*TreeFactory, error) {
	if treeFactoryInstance == nil {
		treeFactoryInstance = &TreeFactory{
			variety: make(map[string]TreeVariety),
		}
	}
	return treeFactoryInstance, nil
}

func (factory *TreeFactory) GetTreeVariety(name string) (TreeVariety, error) {
	if factory.variety[name] != nil {
		return factory.variety[name], nil
	}
	switch name {
	case "Ficus":
		ficus := NewFicus()
		factory.variety["Ficus"] = ficus
		return NewFicus(), nil
	default:
		return nil, errors.New("沒這品種")
	}
}
