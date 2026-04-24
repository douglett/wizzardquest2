package main
import ray "github.com/gen2brain/raylib-go/raylib"
import "fmt"

type Screen struct {
	width, height, zoom int32
	camera ray.Camera2D
}

func (t *Screen) create() error {
	// defaults
	t.width = 160
	t.height = 160
	t.zoom = 4
	t.camera.Zoom = float32(t.zoom)
	// init raylib
	// ray.SetTraceLogLevel(ray.LogInfo)
	ray.SetTraceLogLevel(ray.LogWarning)
	ray.InitWindow(t.width * t.zoom, t.height * t.zoom, "raylib [core] example - basic window")
	ray.InitAudioDevice()
	ray.SetTargetFPS(60)
	// ok
	fmt.Println("Screen initialized:", t.width, t.height)
	return nil
}

func (t Screen) destroy() {
	ray.CloseAudioDevice()
	ray.CloseWindow()
	fmt.Println("Screen destroyed")
}

func (t Screen) begin() {
	ray.BeginDrawing()
	ray.BeginMode2D(t.camera)
	ray.ClearBackground(ray.RayWhite)
}

func (t Screen) flip() {
	// show framerate
	fps := fmt.Sprintf("%d", ray.GetFPS())
	fontw := int32(10)
	txtw := ray.MeasureText(fps, fontw)
	ray.DrawText(fps, t.width - (txtw + 2), 1, fontw, ray.Green)
	// flip
	ray.EndMode2D()
	ray.EndDrawing()
}
