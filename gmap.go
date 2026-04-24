package main
import "fmt"
import "encoding/xml"
import "os"
import "strings"
import "strconv"

type GMap struct {
	XMLName xml.Name     `xml:"map"`
	Width   int32        `xml:"width,attr"`
	Height  int32        `xml:"height,attr"`
	Layer   []GMapLayer  `xml:"layer"`
}

type GMapLayer struct {
	XMLName xml.Name `xml:"layer"`
	Name    string   `xml:"name,attr"`
	Data    string   `xml:"data"`
	IData   []int32
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
		idata := []int32{}
		for _, v := range data {
			i, _ := strconv.Atoi(strings.TrimSpace(v))
			idata = append(idata, int32(i))
		}
		layer.IData = idata
		layer.Data = ""
	}
	
	return nil
}

func (gm GMap) show(posx, posy float32) {
	for y := range gm.Height {
		for x := range gm.Width {
			for _, layer := range gm.Layer {
				tile := layer.IData[y * gm.Width + x]
				if tile > 0 {
					screen.blitt(screen.tileset, tile-1, float32(x*screen.tsize) + posx, float32(y*screen.tsize) + posy)
				}
			}
		}
	}
}
