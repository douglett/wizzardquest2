package main
import ray "github.com/gen2brain/raylib-go/raylib"
import "fmt"
import "math"

type Screen struct {
	width, height, zoom, tsize int
	winname     string
	camera      ray.Camera2D
	tileset     ray.Texture2D
	sound       ray.Sound
	offsetx, offsety  int
}

func (s *Screen) create() error {
	// defaults
	if s.width <= 0  { s.width = 640 }
	if s.height <= 0 { s.height = 480 }
	if s.zoom <= 0   { s.zoom = 1 }
	if s.tsize <= 0  { s.tsize = 16 }
	s.camera.Zoom = float32(s.zoom)
	// init raylib
	// ray.SetTraceLogLevel(ray.LogInfo)
	ray.SetTraceLogLevel(ray.LogWarning)
	ray.InitWindow(int32(s.width * s.zoom), int32(s.height * s.zoom), s.winname)
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
	ray.DrawText(fps, int32(s.width) - (txtw + 2), 1, fontw, ray.Green)
	// flip
	ray.EndMode2D()
	ray.EndDrawing()
}

// blit sub-pixel positions
func (s Screen) blitf(tex ray.Texture2D, x, y float64) {
	s.blit(tex, int(math.Round(x)), int(math.Round(y)))
}
func (s Screen) blittf(tex ray.Texture2D, tile int, x, y float64) {
	s.blitt(tex, tile, int(math.Round(x)), int(math.Round(y)))
}

// texture blitting
func (s Screen) blit(tex ray.Texture2D, x, y int) {
	ray.DrawTexture(screen.tileset, int32(x + screen.offsetx), int32(y + screen.offsety), ray.White)
}

// blit texture as tileset
func (s Screen) blitt(tex ray.Texture2D, tile, x, y int) {
	tx := tile % (int(tex.Width) / screen.tsize)
	ty := tile / (int(tex.Width) / screen.tsize)
	src := ray.Rectangle{
		float32(tx * screen.tsize), float32(ty * screen.tsize),
		float32(screen.tsize), float32(screen.tsize),
	}
	dst := ray.Vector2{ float32(x + screen.offsetx), float32(y + screen.offsety) }
	ray.DrawTextureRec(screen.tileset, src, dst, ray.White)
}
