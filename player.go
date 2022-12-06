package mahjongo

type Player struct {
	spareTiles    []Tile
	completedSets [][]Tile
	flowers       []Tile
}

func (p *Player) winsWithTile(tile Tile) bool {
	return false
}

func (p *Player) hasWon() bool {
	return false
}

func (p *Player) winningMultiplier() int {
	return -1
}
