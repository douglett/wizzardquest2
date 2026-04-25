package main
import "fmt"
import ray "github.com/gen2brain/raylib-go/raylib"

var screen Screen = Screen{
	width: 160, height: 160, zoom: 4, tsize: 16,
	winname: "WizzardQuest2",
}
var gmap GMap = GMap{}

type Mob struct {
	x, y, tile  int
}

// set screen center to player 
func (mob Mob) centeron() {
	screen.offsetx = -(mob.x - (screen.width - screen.tsize) / 2)
	screen.offsety = -(mob.y - (screen.height - screen.tsize) / 2)
}

var player Mob = Mob{ x: 4*screen.tsize, y: 4*screen.tsize, tile: 14 }

func main() {
	fmt.Println("hello world")
	screen.create()
	defer screen.destroy()
	gmap.load("assets/world.tmx")

	var posx, posy int = player.x, player.y
	// move := -1

	for !ray.WindowShouldClose() {
		if ray.IsKeyPressed(ray.KeySpace) {
			ray.PlaySound(screen.sound)
		}

		if ray.IsKeyDown(ray.KeyUp)    { posy-- }
		if ray.IsKeyDown(ray.KeyDown)  { posy++ }
		if ray.IsKeyDown(ray.KeyLeft)  { posx-- }
		if ray.IsKeyDown(ray.KeyRight) { posx++ }

		// switch move {
		// case 0:
		// 	player.oy--
		// 	if player.oy <= -16 { player.oy = 0; player.y--; move = -1 }
		// case 1:
		// 	player.ox++
		// 	if player.ox >= 16 { player.ox = 0; player.x++; move = -1 }
		// case 2:
		// 	player.oy++
		// 	if player.oy >= 16 { player.oy = 0; player.y++; move = -1 }
		// case 3:
		// 	player.ox--
		// 	if player.ox <= -16 { player.ox = 0; player.x--; move = -1 }
		// case -1:
		// 	if ray.IsKeyDown(ray.KeyUp)    { move = 0 }
		// 	if ray.IsKeyDown(ray.KeyRight) { move = 1 }
		// 	if ray.IsKeyDown(ray.KeyDown)  { move = 2 }
		// 	if ray.IsKeyDown(ray.KeyLeft)  { move = 3 }
		// }

		screen.begin()
			// screen.offsetx = -posx
			// screen.offsety = -posy

			// set player position
			player.x = posx
			player.y = posy
			player.centeron()

			gmap.show(0, 0)
			screen.blitt(screen.tileset, player.tile, player.x, player.y)
		screen.flip()
	}
}
