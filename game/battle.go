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
	moffy      int
	// battle stats
	player     BattleMob       
	enemy      BattleMob       
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
	// reset
	bt.frame = 0
	bt.hand = 0
	bt.actions = [4]string{"fireball", "run", "test", "test"}
	// add mobs
	bt.player = BMPlayer
	bt.player.hp = bt.player.maxhp()
	bt.enemy = BMSlime
	bt.enemy.hp = bt.enemy.maxhp()
	// mainloop
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
		// bt.frame++
	}
}

func (bt *Battle) cast() {
	const DamageFireball = 8

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
	
	// calculate damage
	damage := maxi(DamageFireball - bt.enemy.def, 1)
	bt.enemy.hp = maxi(bt.enemy.hp - damage, 0)

	fmt.Println("showing damage")
	for i := 0; i < 50; i++ {
		screen.begin()
		bt.paint()
		// show damage
		screen.offsetx, screen.offsety = offx, offy
		j := int(math.Min(float64(i), 25))
		screen.text(fmt.Sprintf("%d", damage), midx+5, midy-j)
		// flip
		screen.flip()
	}
}

func (bt *Battle) enemyattack() {
	// jump monster
	dist := 10
	for i := 0; i < dist; i++ {
		bt.moffy = -i
		bt.paintall()
	}
	for i := dist; i >= 0; i-- {
		bt.moffy = -i
		bt.paintall()
	}

	// shake screen
	dist = 2
	for i := 0; i < 50; i++ {
		if i % 3 == 0 {
			bt.offx = (rand.Intn(2) - 1) * dist
			bt.offy = (rand.Intn(2) - 1) * dist
		}
		bt.paintall()
	}
	bt.offx, bt.offy = 0, 0

	// calculate damage
	damage := maxi(bt.enemy.atk - bt.player.def, 1)
	bt.player.hp = maxi(bt.player.hp - damage, 0)

	// show damage
	for i := 0; i < 100; i++ {
		screen.begin()
		bt.paint()
		screen.text(fmt.Sprintf("You take %d\ndamage!", damage), 7, 3)
		screen.flip()
	}
}

func (bt *Battle) paintall() {
	screen.begin()
	bt.paint()
	screen.flip()
}

func (bt *Battle) paint() {
	bt.frame++
	overhead.paint()

	// battle screen
	screen.offsetx, screen.offsety = (screen.width-battlebox.width())/2, 20
	battlebox.border(10)
	screen.offsetx += bt.offx
	screen.offsety += bt.offy
	battlebox.show()
	monindex := (bt.frame / 30) % 3
	monx := (battlebox.width() - screen.tsize) / 2
	mony := (battlebox.height() - screen.tsize) / 2 + bt.moffy
	screen.blitt(bt.monsters, monindex, monx, mony)

	// dialog box
	screen.offsetx, screen.offsety = (screen.width-dialogbox.width())/2, 120
	dialogbox.border(2)
	dialogbox.show()

	// player health
	px, py := 2, -10
	screen.rect(px-2, py-1, 60, 11, ColorBlack)
	screen.text(fmt.Sprintf("%d/%d", bt.player.hp, bt.player.maxhp()), px, py)
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
