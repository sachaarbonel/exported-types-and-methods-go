package animals

// Dog represents information about dogs.
type Dog struct {
	Name         string
	BarkStrength int
	Age          int
	Wouaf        string
}

func (d *Dog) Bark() string {
	wouaf := d.Wouaf
	return wouaf
}
