package gmap
import "fmt"
import "encoding/xml"
import "os"
import "strings"
import "strconv"

type GMap struct {
	XMLName xml.Name     `xml:"map"`
	Width   string       `xml:"width,attr"`
	Height  string       `xml:"height,attr"`
	Fart    string       `xml:"fart"`
	Layer   []GMapLayer  `xml:"layer"`
}

type GMapLayer struct {
	XMLName xml.Name `xml:"layer"`
	Name    string   `xml:"name,attr"`
	Data    string   `xml:"data"`
	IData   []int
}

var MainMap GMap

func Load(fname string) error {
	file, err := os.Open(fname)
	if err != nil {
		fmt.Println("error", err)
		return err
	}
	defer file.Close()

	fmt.Println("decoding map:", fname)

	// get data from xml
	gm := GMap{}
	decoder := xml.NewDecoder(file)
	decoder.Decode(&gm)

	// decode layer data
	for k := range gm.Layer {
		layer := &gm.Layer[k]
		data := strings.Split(layer.Data, ",")
		idata := []int{}
		for _, v := range data {
			i, _ := strconv.Atoi(v)
			idata = append(idata, i)
		}
		layer.IData = idata
		layer.Data = ""
	}

	fmt.Println(gm)
	MainMap = gm
	return nil
}

func Show() {
	// todo
}
