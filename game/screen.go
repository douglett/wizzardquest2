package game
import ray "github.com/gen2brain/raylib-go/raylib"
import "fmt"
import "math"

type Screen struct {
	width, height, zoom, tsize int
	winname     string
	bgcolor     ray.Color
	camera      ray.Camera2D
	tileset     ray.Texture2D
	sound       ray.Sound
	offsetx, offsety  int
}

func (scr *Screen) create() error {
	// defaults
	if scr.width <= 0  { scr.width = 640 }
	if scr.height <= 0 { scr.height = 480 }
	if scr.zoom <= 0   { scr.zoom = 1 }
	if scr.tsize <= 0  { scr.tsize = 16 }
	if scr.bgcolor.A == 0  { scr.bgcolor = ray.RayWhite }
	scr.camera.Zoom = float32(scr.zoom)
	// init raylib
	// ray.SetTraceLogLevel(ray.LogInfo)
	ray.SetTraceLogLevel(ray.LogWarning)
	ray.InitWindow(int32(scr.width * scr.zoom), int32(scr.height * scr.zoom), scr.winname)
	ray.InitAudioDevice()
	ray.SetTargetFPS(60)
	// ok
	fmt.Println("Screen initialized:", scr.width, scr.height)
	return nil
}

func (scr *Screen) destroy() {
	ray.CloseAudioDevice()
	ray.CloseWindow()
	fmt.Println("Screen destroyed")
}

func (scr *Screen) begin() {
	ray.BeginDrawing()
	ray.BeginMode2D(scr.camera)
	ray.ClearBackground(scr.bgcolor)
}

func (scr *Screen) flip() {
	// show framerate
	fps := fmt.Sprintf("%d", ray.GetFPS())
	fontw := int32(10)
	txtw := ray.MeasureText(fps, fontw)
	ray.DrawText(fps, int32(scr.width) - (txtw + 2), 1, fontw, ray.Green)
	// flip
	ray.EndMode2D()
	ray.EndDrawing()
}

// blit sub-pixel positions
func (scr *Screen) blitf(tex ray.Texture2D, x, y float64) {
	scr.blit(tex, int(math.Round(x)), int(math.Round(y)))
}
func (scr *Screen) blittf(tex ray.Texture2D, tile int, x, y float64) {
	scr.blitt(tex, tile, int(math.Round(x)), int(math.Round(y)))
}

// texture blitting
func (scr *Screen) blit(tex ray.Texture2D, x, y int) {
	ray.DrawTexture(tex, int32(x + scr.offsetx), int32(y + scr.offsety), ray.White)
}

// blit texture as tileset
func (scr *Screen) blitt(tex ray.Texture2D, tile, x, y int) {
	tx := tile % (int(tex.Width) / scr.tsize)
	ty := tile / (int(tex.Width) / scr.tsize)
	src := ray.Rectangle{
		float32(tx * scr.tsize), float32(ty * scr.tsize),
		float32(scr.tsize), float32(scr.tsize),
	}
	dst := ray.Vector2{ float32(x + scr.offsetx), float32(y + scr.offsety) }
	ray.DrawTextureRec(tex, src, dst, ray.White)
}

func (scr *Screen) rect(x, y, w, h int, color ray.Color) {
	ray.DrawRectangle(int32(x + scr.offsetx), int32(y + scr.offsety), int32(w), int32(h), color)
}
