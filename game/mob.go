package game

// Mob on screen
type Mob struct {
	x, y, tile  int
}

func (mob *Mob) pos() (int, int) {
	return mob.x / screen.tsize, mob.y / screen.tsize
}

func (mob *Mob) paint() {
	screen.blitt(screen.tileset, mob.tile, mob.x, mob.y)
}

// set screen center to player 
func (mob *Mob) centeron() {
	screen.offsetx = -(mob.x - (screen.width - screen.tsize) / 2)
	screen.offsety = -(mob.y - (screen.height - screen.tsize) / 2)
}


// Battle stats for mobs
type BattleMob struct {
	name string
	hp, mp int
	stm, atk, def int
}

func (bm *BattleMob) maxhp() int {
	return bm.stm * 5
}

// types of battle mob. these can't be const because golang is retarded.
var BMPlayer = BattleMob{
	name: "player",
	stm: 4, atk: 0, def: 2,
}

var BMSlime = BattleMob{
	name: "slime",
	stm: 2, atk: 5, def: 1,
}
