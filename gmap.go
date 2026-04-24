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

func (gm GMap) show() {
	for y := range int32(gm.Height) {
		for x := range int32(gm.Width) {
			for _, layer := range gm.Layer {
				tile := layer.IData[y * int32(gm.Width) + x]
				if tile > 0 {
					screen.blitt(screen.tileset, tile-1, float32(x*screen.tsize), float32(y*screen.tsize))
				}
			}
		}
	}
}
