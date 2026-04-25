package main
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

func (gm GMap) paint() {
	gm.paintat(0, 0)
}

func (gm GMap) paintat(posx, posy int) {
	for y := range gm.Height {
		for x := range gm.Width {
			for _, layer := range gm.Layer {
				tile := layer.IData[y * gm.Width + x]
				if tile > 0 {
					screen.blitt(screen.tileset, tile - 1, (x * screen.tsize) + posx, (y * screen.tsize) + posy)
				}
			}
		}
	}
}

func (gm GMap) tile(x, y int) int {
	// if gm.Width == 0 || gm.Height == 0 || len(gm.Layer) == 0 { return 0 }
	if x < 0 || y < 0 || x >= gm.Width || y >= gm.Height { return 1000 }
	return gm.Layer[0].IData[y * gm.Width + x]
}
