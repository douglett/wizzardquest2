package main
import "fmt"
import ray "github.com/gen2brain/raylib-go/raylib"

var screen Screen = Screen{
	width: 160, height: 160, zoom: 4, tsize: 16,
	winname: "WizzardQuest2",
}
var gmap GMap = GMap{}

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
			screen.blitt(screen.tileset, 14, float32((screen.width-screen.tsize)/2), float32((screen.height-screen.tsize)/2))

		screen.flip()
	}
}
