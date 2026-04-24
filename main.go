package main
import "fmt"
import ray "github.com/gen2brain/raylib-go/raylib"

var screen Screen = Screen{
	width: 160, height: 160, zoom: 4, tsize: 16,
	winname: "WizzardQuest2",
}
var gmap GMap = GMap{}

type Mob struct {
	x, y, tile  int32
	ox, oy      float32
}
var player Mob = Mob{ x: 4, y: 4, tile: 14 }

func main() {
	fmt.Println("hello world")
	screen.create()
	defer screen.destroy()
	gmap.load("assets/world.tmx")

	// var posx, posy float32
	move := -1

	for !ray.WindowShouldClose() {
		if ray.IsKeyPressed(ray.KeySpace) {
			ray.PlaySound(screen.sound)
		}

		// if ray.IsKeyDown(ray.KeyUp)    { posy-- }
		// if ray.IsKeyDown(ray.KeyDown)  { posy++ }
		// if ray.IsKeyDown(ray.KeyLeft)  { posx-- }
		// if ray.IsKeyDown(ray.KeyRight) { posx++ }

		switch move {
		case 0:
			player.oy--
			if player.oy <= -16 { player.oy = 0; player.y--; move = -1 }
		case 1:
			player.ox++
			if player.ox >= 16 { player.ox = 0; player.x++; move = -1 }
		case 2:
			player.oy++
			if player.oy >= 16 { player.oy = 0; player.y++; move = -1 }
		case 3:
			player.ox--
			if player.ox <= -16 { player.ox = 0; player.x--; move = -1 }
		case -1:
			if ray.IsKeyDown(ray.KeyUp)    { move = 0 }
			if ray.IsKeyDown(ray.KeyRight) { move = 1 }
			if ray.IsKeyDown(ray.KeyDown)  { move = 2 }
			if ray.IsKeyDown(ray.KeyLeft)  { move = 3 }
		}

		screen.begin()
			// player.ox = posx
			// player.oy = posy
			screen.offx = float32((screen.width-screen.tsize)/2) - (float32(player.x*screen.tsize) + player.ox)
			screen.offy = float32((screen.height-screen.tsize)/2) - (float32(player.y*screen.tsize) + player.oy)
			gmap.show(0, 0)
			screen.blitt(screen.tileset, player.tile, float32(player.x*screen.tsize)+player.ox, float32(player.y*screen.tsize)+player.oy)
		screen.flip()
	}
}
