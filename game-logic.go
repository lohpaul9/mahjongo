package mahjongo

type GameClient struct {
	playerTurn             int
	tableTiles             []Tile
	tileDrawQueue          []Tile
	flowerReplacementQueue []Tile
	wind                   int
	players                [4]*Player
	isWon                  bool
	lastThrown             Tile
	denomination           int
}

func (g *GameClient) Init() (players [4]*Player) {
	return
}

func (g *GameClient) ViewTable() (tableTiles []Tile, lastThrown Tile) {
	return
}

func (g *GameClient) PlayerView(player int) (ownTiles []Tile, isTurn bool) {
	return
}

func (g *GameClient) PengAndDiscard(player int, toPeng Tile, tilesToJoin []Tile, toDiscard Tile) (e error) {
	return
}

func (g *GameClient) ChiAndDiscard(player int, toChi Tile, tilesToJoin []Tile, toDiscard Tile) (e error) {
	return
}

func (g *GameClient) Draw(player int) (ownTiles []Tile, newTile Tile) {
	return
}

func (g *GameClient) Discard(player int, toDiscard Tile) (ownTiles []Tile, newTile Tile) {
	return
}

func (g *GameClient) GetWinner() (player int) {
	return
}

func (g *GameClient) GetWinningEarnings() (earnings int) {
	return
}
