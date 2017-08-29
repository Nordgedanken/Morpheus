package matrix

import (
	"log"
	"os"
	"strings"

	"bytes"
	"fmt"
	"github.com/Nordgedanken/Morpheus/util"
	"github.com/matrix-org/gomatrix"
	"github.com/therecipe/qt/gui"
	"github.com/tidwall/buntdb"
	"image"
	"image/color"
	"image/draw"
	"image/png"
)

var localLog *log.Logger

func init() {
	var file *os.File
	localLog = util.Logger()
	localLog, file = util.StartFileLog(localLog)
	defer file.Close()
}

type circle struct {
	p image.Point
	r int
}

func (c *circle) ColorModel() color.Model {
	return color.AlphaModel
}

func (c *circle) Bounds() image.Rectangle {
	return image.Rect(c.p.X-c.r, c.p.Y-c.r, c.p.X+c.r, c.p.Y+c.r)
}

func (c *circle) At(x, y int) color.Color {
	xx, yy, rr := float64(x-c.p.X)+0.5, float64(y-c.p.Y)+0.5, float64(c.r)
	if xx*xx+yy*yy < rr*rr {
		return color.Alpha{255}
	}
	return color.Alpha{0}
}

// getOwnUserAvatar returns a *gui.QPixmap of an UserAvatar
func GetOwnUserAvatar(cli *gomatrix.Client) *gui.QPixmap {
	// Init local vars
	var avatarData string
	var IMGdata string

	// Get cache
	db.View(func(tx *buntdb.Tx) error {
		QueryErr := tx.AscendKeys("user:"+cli.UserID+":avatarData100x100",
			func(key, value string) bool {
				avatarData = value
				return true
			})
		if QueryErr != nil {
			return QueryErr
		}
		return nil
	})

	//If cache is empty do a ServerQuery
	if avatarData == "" {
		// Get avatarURL
		avatarURL, avatarErr := cli.GetAvatarURL()
		if avatarErr != nil {
			localLog.Println(avatarErr)
		}

		// If avatarURL is not empty (aka. has a avatar set) download it at the size of 100x100. Else make the data string empty
		if avatarURL != "" {
			hsURL := cli.HomeserverURL.String()
			avatarURL_splits := strings.Split(strings.Replace(avatarURL, "mxc://", "", -1), "/")

			urlPath := hsURL + "/_matrix/media/r0/thumbnail/" + avatarURL_splits[0] + "/" + avatarURL_splits[1] + "?width=100&height=100"

			data, err := cli.MakeRequest("GET", urlPath, nil, nil)
			if err != nil {
				localLog.Println(err)
			}
			IMGdata = string(data[:])
		} else {
			//TODO Generate default image (Step: AfterUI)
			IMGdata = "0"
		}

		// Update cache
		DBerr := db.Update(func(tx *buntdb.Tx) error {
			tx.Set("user:"+cli.UserID+":avatarData100x100", IMGdata, nil)
			return nil
		})
		if DBerr != nil {
			localLog.Fatalln(DBerr)
		}
	} else {
		IMGdata = avatarData
	}

	r := bytes.NewReader([]byte(IMGdata))
	srcIMG, _, err := image.Decode(r)
	if err != nil {
		localLog.Println(err)
	}

	// Convert avatarimage to QPixmap for usage in QT
	canvas := image.NewRGBA(srcIMG.Bounds())
	cx := srcIMG.Bounds().Min.X + srcIMG.Bounds().Dx()/2
	cy := srcIMG.Bounds().Min.Y + srcIMG.Bounds().Dy()/2
	draw.DrawMask(canvas, canvas.Bounds(), srcIMG, image.ZP, &circle{image.Point{cx, cy}, 110}, image.ZP, draw.Over)
	avatar := gui.NewQPixmap()
	buf := new(bytes.Buffer)
	Converr := png.Encode(buf, canvas)
	if Converr != nil {
		localLog.Println(Converr)
	}

	str := fmt.Sprintf("%s", buf)
	avatar.LoadFromData(str, uint(len(str)), "", 0)
	return avatar
}
