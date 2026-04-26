package game
import ray "github.com/gen2brain/raylib-go/raylib"

type Battle struct {
	monsters  ray.Texture2D
	frame     int
	hand      int
}

var dialogbox = GMapFrag{
	w: 7,
	h: 2,
	idata: []int{
		4,  8,  8,  8,  8,  8,  5,
		7,  10, 10, 10, 10, 10, 6,
	},
}
var battlebox = GMapFrag{
	w: 5,
	h: 5,
	idata: []int{
		4,  8,  8,  8,  5,
		11, 1,  1,  1,  9,
		11, 1,  1,  1,  9,
		11, 1,  1,  1,  9,
		7,  10, 10, 10, 6,
	},
}

func (bt *Battle) mainloop() {
	bt.frame = 0
	bt.hand = 0
	for !ray.WindowShouldClose() {
		switch {
			case ray.IsKeyPressed(ray.KeyUp):     if bt.hand == 1 || bt.hand == 3 { bt.hand-- }
			case ray.IsKeyPressed(ray.KeyDown):   if bt.hand == 0 || bt.hand == 2 { bt.hand++ }
			case ray.IsKeyPressed(ray.KeyRight):  if bt.hand < 2 { bt.hand += 2 }
			case ray.IsKeyPressed(ray.KeyLeft):   if bt.hand > 1 { bt.hand -= 2 }
		}
		bt.paintall()
		bt.frame++
	}
}

func (bt *Battle) paintall() {
	screen.begin()
	bt.paint()
	screen.flip()
}

func (bt *Battle) paint() {
	overhead.paint()

	// battle screen
	screen.offsetx, screen.offsety = (screen.width-battlebox.width())/2, 20
	battlebox.border(10)
	battlebox.show()
	sprindex := (bt.frame / 30) % 3
	screen.blitt(bt.monsters, sprindex, (battlebox.width()-screen.tsize)/2, (battlebox.height()-screen.tsize)/2)

	// dialog box
	screen.offsetx, screen.offsety = (screen.width-dialogbox.width())/2, 120
	dialogbox.border(2)
	dialogbox.show()

	ray.DrawText("fireball", 38, 124, 2, ColorWhite)
	ray.DrawText("fireball", 38, 136, 2, ColorWhite)
	ray.DrawText("fireball", 88, 124, 2, ColorWhite)
	ray.DrawText("fireball", 88, 136, 2, ColorWhite)
	var tx, ty int32
	switch bt.hand {
		case 0:  tx, ty = 30, 124
		case 1:  tx, ty = 30, 136
		case 2:  tx, ty = 80, 124
		case 3:  tx, ty = 80, 136
	}
	ray.DrawText("@", tx, ty, 10, ray.White)
}
