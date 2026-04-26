package game
import ray "github.com/gen2brain/raylib-go/raylib"

// useful colors
var ColorCollision = ray.Color{255, 0, 0, 100}
var ColorBlack     = ray.Color{16, 8, 32, 255}
var ColorWhite     = ray.Color{255, 255, 255, 255}

// game state
var screen = Screen{
	width: 160, height: 160, zoom: 4, tsize: 16,
	winname: "WizzardQuest2",
	bgcolor: ColorBlack,
}
var gmap = GMap{}
var overhead = Overhead{}
var battle = Battle{}
var player = Mob{ x: 4*screen.tsize, y: 4*screen.tsize, tile: 14 }


// start game
func Start() {
	screen.create()
	defer screen.destroy()
	gmap.load("assets/world.tmx")
	// gmap.showCollision = true

	overhead.mainloop()
}


// Overhead game scene
type Overhead struct {}

func (oh *Overhead) mainloop() {
	for !ray.WindowShouldClose() {
		if ray.IsKeyPressed(ray.KeySpace) {
			ray.PlaySound(screen.sound)
		}

		// player walk
		switch {
			case ray.IsKeyDown(ray.KeyUp):     oh.walk( 0, -1)
			case ray.IsKeyDown(ray.KeyRight):  oh.walk( 1,  0)
			case ray.IsKeyDown(ray.KeyDown):   oh.walk( 0,  1)
			case ray.IsKeyDown(ray.KeyLeft):   oh.walk(-1,  0)
			case ray.IsKeyDown(ray.KeyEnter):  battle.mainloop()
		}

		// repaint screen
		oh.paint()
	}
}

func (oh *Overhead) paint() {
	screen.begin()
		player.centeron()
		gmap.paint()
		screen.blitt(screen.tileset, player.tile, player.x, player.y)
	screen.flip()
}

func (oh *Overhead) walk(dx, dy int) {
	tx, ty := player.pos()
	if _, c := gmap.tile(tx + dx, ty + dy); c == true { return }
	dist := 0
	for dist < screen.tsize {
		player.x += dx
		player.y += dy
		dist++
		if (dist < screen.tsize) { oh.paint() }
	}
}
