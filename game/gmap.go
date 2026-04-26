package game
import "fmt"
import "encoding/xml"
import "os"
import "strings"
import "strconv"

type GMap struct {
	XMLName xml.Name     `xml:"map"`
	Width   int          `xml:"width,attr"`
	Height  int          `xml:"height,attr"`
	Layer   []GMapLayer  `xml:"layer"`
	showCollision  bool
}

type GMapLayer struct {
	XMLName xml.Name `xml:"layer"`
	Name    string   `xml:"name,attr"`
	Data    string   `xml:"data"`
	IData   []int
}

func (gm *GMap) load(fname string) error {
	file, err := os.Open(fname)
	if err != nil {
		fmt.Println("error", err)
		return err
	}
	defer file.Close()

	fmt.Println("decoding map:", fname)

	// get data from xml
	decoder := xml.NewDecoder(file)
	decoder.Decode(&gm)

	// decode layer data
	for k := range gm.Layer {
		layer := &gm.Layer[k]
		data := strings.Split(layer.Data, ",")
		idata := []int{}
		for _, v := range data {
			i, _ := strconv.Atoi(strings.TrimSpace(v))
			idata = append(idata, i)
		}
		layer.IData = idata
		layer.Data = ""
	}
	
	return nil
}

func (gm *GMap) paint() {
	gm.paintat(0, 0)
}

func (gm *GMap) paintat(posx, posy int) {
	layer := &gm.Layer[0]
	coll := &gm.Layer[1]
	for y := range gm.Height {
		for x := range gm.Width {
			// show game tile
			tile := layer.IData[y * gm.Width + x]
			if tile > 0 {
				screen.blitt(screen.tileset, tile - 1, (x * screen.tsize) + posx, (y * screen.tsize) + posy)
			}
			// show collision layer (optional)
			c := coll.IData[y * gm.Width + x]
			if gm.showCollision && c > 0 {
				screen.rect((x * screen.tsize) + posx, (y * screen.tsize) + posy, screen.tsize, screen.tsize, ColorCollision)
			}
		}
	}
}

func (gm *GMap) tile(x, y int) (int, bool) {
	if x < 0 || y < 0 || x >= gm.Width || y >= gm.Height { return 1000, true }
	return gm.Layer[0].IData[y * gm.Width + x], gm.Layer[1].IData[y * gm.Width + x] > 0
}



// === Map Fragments ===

type MapFrag struct {
	w, h  int
	idata  []int
}

func (mf *MapFrag) width()  int { return mf.w * screen.tsize }
func (mf *MapFrag) height() int { return mf.h * screen.tsize }

func (mf *MapFrag) border(pad int) {
	screen.rect(-pad, -pad, mf.width()+pad*2, mf.height()+pad*2, ColorBlack)
}

func (mf *MapFrag) show() {
	for y := range mf.h {
		for x := range mf.w {
			screen.blitt(screen.tileset, mf.idata[y * mf.w + x]-1, x*screen.tsize, y*screen.tsize)
		}
	}
}
