package forests

type Ficus struct {
	name    string
	color   string
	texture string
}

func NewFicus() *Ficus {
	return &Ficus{
		name:    "Ficus",
		color:   "棕色",
		texture: "粗糙",
	}
}

func (tree *Ficus) GetName() string {
	return tree.name
}

func (tree *Ficus) GetColor() string {
	return tree.color
}

func (tree *Ficus) GetTexture() string {
	return tree.texture
}
