package game
import "fmt"
import "math"
import "math/rand"
import ray "github.com/gen2brain/raylib-go/raylib"

type Battle struct {
	monsters   ray.Texture2D
	spells     ray.Texture2D
	frame      int
	hand       int
	actions    [4]string
	offx, offy int
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
	bt.actions = [4]string{"fireball", "run", "test", "test"}
	bt.paintall()
	for !ray.WindowShouldClose() {
		switch {
			case ray.IsKeyPressed(ray.KeyUp):     if bt.hand == 1 || bt.hand == 3 { bt.hand-- }
			case ray.IsKeyPressed(ray.KeyDown):   if bt.hand == 0 || bt.hand == 2 { bt.hand++ }
			case ray.IsKeyPressed(ray.KeyRight):  if bt.hand < 2 { bt.hand += 2 }
			case ray.IsKeyPressed(ray.KeyLeft):   if bt.hand > 1 { bt.hand -= 2 }
			case ray.IsKeyPressed(ray.KeyEnter):
				switch bt.actions[bt.hand] {
					case "":
					case "fireball":
						bt.cast()
						bt.enemyattack()
					// case "run":
					default:  fmt.Printf("unknown selection '%s'\n", bt.actions[bt.hand])
				}
		}
		// paint screen
		screen.begin()
		bt.paint()
		bt.painthand()
		screen.flip()
		bt.frame++
	}
}

func (bt *Battle) cast() {
	// positioning variables
	offx, offy := (screen.width-battlebox.width())/2, 20
	midx := (battlebox.width() - screen.tsize) / 2
	midy := (battlebox.height() - screen.tsize) / 2

	fmt.Println("casting fireball")
	for i := 0; i < 50; i++ {
		screen.begin()
		bt.paint()
		// show fireball
		screen.offsetx, screen.offsety = offx, offy
		screen.blitt(bt.spells, (i/8)%3, midx, midy+8)
		// flip
		screen.flip()
	}
	fmt.Println("showing damage")
	for i := 0; i < 50; i++ {
		screen.begin()
		bt.paint()
		// show damage
		screen.offsetx, screen.offsety = offx, offy
		j := int(math.Min(float64(i), 25))
		screen.text("10", midx, midy-j)
		// flip
		screen.flip()
	}
}

func (bt *Battle) enemyattack() {
	// shake screen
	dist := 2
	for i := 0; i < 50; i++ {
		if i % 3 == 0 {
			bt.offx = (rand.Intn(2) - 1) * dist
			bt.offy = (rand.Intn(2) - 1) * dist
		}
		bt.paintall()
	}
	bt.offx, bt.offy = 0, 0

	// show damage
	for i := 0; i < 100; i++ {
		screen.begin()
		bt.paint()
		screen.text("You take 10\ndamage!", 7, 3)
		screen.flip()
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
	screen.offsetx += bt.offx
	screen.offsety += bt.offy
	battlebox.show()
	sprindex := (bt.frame / 30) % 3
	screen.blitt(bt.monsters, sprindex, (battlebox.width()-screen.tsize)/2, (battlebox.height()-screen.tsize)/2)

	// dialog box
	screen.offsetx, screen.offsety = (screen.width-dialogbox.width())/2, 120
	dialogbox.border(2)
	dialogbox.show()
}

func (bt *Battle) painthand() {
	// paint hand dialog
	screen.text(bt.actions[0], 14, 4)
	screen.text(bt.actions[1], 14, 16)
	screen.text(bt.actions[2], 64, 4)
	screen.text(bt.actions[3], 64, 16)
	var tx, ty int
	switch bt.hand {
		case 0:  tx, ty = 4, 4
		case 1:  tx, ty = 4, 16
		case 2:  tx, ty = 54, 4
		case 3:  tx, ty = 54, 16
	}
	screen.text("@", tx, ty)
}
