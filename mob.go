package main

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
