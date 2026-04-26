package main
import "fmt"
import ray "github.com/gen2brain/raylib-go/raylib"

var screen Screen = Screen{
	width: 160, height: 160, zoom: 4, tsize: 16,
	winname: "WizzardQuest2",
}
var gmap GMap = GMap{}
var player Mob = Mob{ x: 4*screen.tsize, y: 4*screen.tsize, tile: 14 }

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
		}

		// repaint screen
		repaint()
	}
}

func repaint() {
	screen.begin()
		gmap.paint()
		player.centeron()
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
