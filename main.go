package main
import "fmt"
import ray "github.com/gen2brain/raylib-go/raylib"

func main() {
	fmt.Println("hello world")
	ray.InitWindow(800, 450, "raylib [core] example - basic window")
	defer ray.CloseWindow()

	ray.SetTargetFPS(60)
	fps := ""

	for !ray.WindowShouldClose() {
		ray.BeginDrawing()

		ray.ClearBackground(ray.RayWhite)
		ray.DrawText("Congrats! You created your first window!", 190, 200, 20, ray.LightGray)
		fps = fmt.Sprintf("%d", ray.GetFPS())
		ray.DrawText(fps, int32(ray.GetScreenWidth()-30), 2, 20, ray.Green)

		ray.EndDrawing()
	}
}
