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
	fmt.Println(gmap)

	for !ray.WindowShouldClose() {
		if ray.IsKeyPressed(ray.KeySpace) {
			ray.PlaySound(screen.sound)
		}

		screen.begin()
			screen.blit(screen.tileset, 20, 20)
			screen.blitt(screen.tileset, 1, 100, 100)
			screen.blitt(screen.tileset, 14, 100, 120)
		screen.flip()
	}
}
