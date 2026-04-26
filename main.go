package main
import "fmt"
import ray "github.com/gen2brain/raylib-go/raylib"

// game state
var screen Screen = Screen{
	width: 160, height: 160, zoom: 4, tsize: 16,
	winname: "WizzardQuest2",
}
var gmap GMap = GMap{}
var player Mob = Mob{ x: 4*screen.tsize, y: 4*screen.tsize, tile: 14 }

// useful colors
var ColorCollision ray.Color = ray.Color{255, 0, 0, 100}
var ColorBlack ray.Color = ray.Color{16, 8, 32, 255}

// start game
func main() {
	fmt.Println("hello world")
	screen.create()
	defer screen.destroy()
	gmap.load("assets/world.tmx")
	gmap.showCollision = true

	for !ray.WindowShouldClose() {
		if ray.IsKeyPressed(ray.KeySpace) {
			ray.PlaySound(screen.sound)
		}

		// player walk
		switch {
			case ray.IsKeyDown(ray.KeyUp):     walk( 0, -1)
			case ray.IsKeyDown(ray.KeyRight):  walk( 1,  0)
			case ray.IsKeyDown(ray.KeyDown):   walk( 0,  1)
			case ray.IsKeyDown(ray.KeyLeft):   walk(-1,  0)
			case ray.IsKeyDown(ray.KeyEnter):  battle()
		}

		// repaint screen
		repaint()
	}
}

func repaint() {
	screen.begin()
		player.centeron()
		gmap.paint()
		screen.blitt(screen.tileset, player.tile, player.x, player.y)
	screen.flip()
}

func walk(dx, dy int) {
	tx, ty := player.pos()
	if _, c := gmap.tile(tx + dx, ty + dy); c == true { return }
	dist := 0
	for dist < screen.tsize {
		player.x += dx
		player.y += dy
		dist++
		if (dist < screen.tsize) { repaint() }
	}
}

func battle() {
	dialog := MapFrag{
		w: 7,
		h: 2,
		idata: []int{
			4,  8,  8,  8,  8,  8,  5,
			7,  10, 10, 10, 10, 10, 6,
		},
	}
	box := MapFrag{
		w: 5,
		h: 5,
		idata: []int{
			4,  8,  8,  8,  5,
			11, 1,  1,  1,  9,
			11, 1,  1,  1,  9,
			11, 1,  1,  1,  9,
			7,  10, 10, 10, 6,
		},
	}
	for !ray.WindowShouldClose() {
		screen.begin()
			player.centeron()
			gmap.paint()

			screen.offsetx, screen.offsety = (screen.width-box.width())/2, 20
			box.border(10)
			box.show()
			screen.offsetx, screen.offsety = (screen.width-dialog.width())/2, 120
			dialog.border(2)
			dialog.show()
		screen.flip()
	}
}
