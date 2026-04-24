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

	var posx, posy float32

	for !ray.WindowShouldClose() {
		if ray.IsKeyPressed(ray.KeySpace) {
			ray.PlaySound(screen.sound)
		}

		if ray.IsKeyDown(ray.KeyUp)    { posy-- }
		if ray.IsKeyDown(ray.KeyDown)  { posy++ }
		if ray.IsKeyDown(ray.KeyLeft)  { posx-- }
		if ray.IsKeyDown(ray.KeyRight) { posx++ }

		screen.begin()
			gmap.show(-posx, -posy)
			// screen.blitt(screen.tileset, 14, float32((screen.width-screen.tsize)/2), float32((screen.height-screen.tsize)/2))
			screen.blitt(screen.tileset, player.tile, float32(player.x*screen.tsize)+player.ox, float32(player.y*screen.tsize)+player.oy)

		screen.flip()
	}
}
