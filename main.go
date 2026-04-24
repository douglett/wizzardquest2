package main
import "fmt"
import ray "github.com/gen2brain/raylib-go/raylib"

var screen Screen = Screen{
	width: 160, height: 160, zoom: 4,
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
			ray.DrawTexture(screen.tileset, 20, 20, ray.White)
			ray.DrawTextureRec(screen.tileset, ray.Rectangle{16, 0, 16, 16}, ray.Vector2{100, 100}, ray.White)
		screen.flip()
	}
}
