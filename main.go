package main
import "fmt"
import ray "github.com/gen2brain/raylib-go/raylib"

func main() {
	fmt.Println("hello world")
	ray.InitWindow(800, 450, "raylib [core] example - basic window")
	defer ray.CloseWindow()

	ray.SetTargetFPS(60)
	camera := ray.Camera2D{ Zoom: 2 }
	// camera.Target.X = 100

	for !ray.WindowShouldClose() {
		ray.BeginDrawing()
		ray.BeginMode2D(camera)

		ray.ClearBackground(ray.RayWhite)
		ray.DrawText("Congrats! You created your first window!", 0, 200, 20, ray.LightGray)
		fps := fmt.Sprintf("%d", ray.GetFPS())
		right := int32(float32(ray.GetScreenWidth()) / camera.Zoom)
		ray.DrawText(fps, right-30, 2, 20, ray.Green)

		ray.EndMode2D()
		ray.EndDrawing()
	}
}
