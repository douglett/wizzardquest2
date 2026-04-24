package main
import ray "github.com/gen2brain/raylib-go/raylib"
import "fmt"

type Screen struct {
	width, height, zoom int32
	winname  string
	camera   ray.Camera2D
	tileset  ray.Texture2D
	sound    ray.Sound
}

func (s *Screen) create() error {
	// defaults
	if s.width <= 0  { s.width = 640 }
	if s.height <= 0 { s.height = 480 }
	if s.zoom <= 0   { s.zoom = 1 }
	s.camera.Zoom = float32(s.zoom)
	// init raylib
	// ray.SetTraceLogLevel(ray.LogInfo)
	ray.SetTraceLogLevel(ray.LogWarning)
	ray.InitWindow(s.width * s.zoom, s.height * s.zoom, s.winname)
	ray.InitAudioDevice()
	ray.SetTargetFPS(60)
	// load assets
	s.tileset = ray.LoadTexture("assets/monotiles.png")
	s.sound = ray.LoadSound("assets/target.ogg")
	// ok
	fmt.Println("Screen initialized:", s.width, s.height)
	return nil
}

func (s Screen) destroy() {
	ray.CloseAudioDevice()
	ray.CloseWindow()
	fmt.Println("Screen destroyed")
}

func (s Screen) begin() {
	ray.BeginDrawing()
	ray.BeginMode2D(s.camera)
	ray.ClearBackground(ray.RayWhite)
}

func (s Screen) flip() {
	// show framerate
	fps := fmt.Sprintf("%d", ray.GetFPS())
	fontw := int32(10)
	txtw := ray.MeasureText(fps, fontw)
	ray.DrawText(fps, s.width - (txtw + 2), 1, fontw, ray.Green)
	// flip
	ray.EndMode2D()
	ray.EndDrawing()
}
