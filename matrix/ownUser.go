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
	_ "image/gif"
	_ "image/jpeg"
	"image/png"
	"math/rand"
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

func generateGenericImages(displayname string) string {
	// Create an 100 x 100 image
	img := image.NewRGBA(image.Rect(0, 0, 100, 100))

	// Draw a red dot at (2, 3)
	img.Set(50, 50, color.RGBA{uint8(rand.Intn(200)), uint8(rand.Intn(255)), uint8(rand.Intn(255)), 255})

	buf := new(bytes.Buffer)
	png.Encode(buf, img)
	return string(buf.Bytes())
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
			avatarURLSplits := strings.Split(strings.Replace(avatarURL, "mxc://", "", -1), "/")

			urlPath := hsURL + "/_matrix/media/r0/thumbnail/" + avatarURLSplits[0] + "/" + avatarURLSplits[1] + "?width=100&height=100"

			data, err := cli.MakeRequest("GET", urlPath, nil, nil)
			if err != nil {
				localLog.Println(err)
			}
			IMGdata = string(data[:])
		} else {
			//TODO Generate default image (Step: AfterUI)
			DisplayNameResp, _ := cli.GetOwnDisplayName()
			DisplayName := DisplayNameResp.DisplayName
			IMGdata = generateGenericImages(DisplayName)
		}

		// Update cache
		DBerr := db.Update(func(tx *buntdb.Tx) error {
			tx.Set("user:"+cli.UserID+":avatarData100x100", IMGdata, nil)
			return nil
		})
		if DBerr != nil {
			localLog.Fatalln("DB ERROR: ", DBerr)
		}
	} else {
		IMGdata = avatarData
	}

	r := bytes.NewReader([]byte(IMGdata))
	srcIMG, _, err := image.Decode(r)
	if err != nil {
		localLog.Println("Decoder error: ", err)
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
		localLog.Println("Converting error: ", Converr)
	}

	str := fmt.Sprintf("%s", buf)
	avatar.LoadFromData(str, uint(len(str)), "", 0)
	return avatar
}

// getOwnUserAvatar returns a *gui.QPixmap of an UserAvatar
func GetUserAvatar(cli *gomatrix.Client, mxid string) *gui.QPixmap {
	// Init local vars
	var avatarData string
	var IMGdata string

	// Get cache
	db.View(func(tx *buntdb.Tx) error {
		QueryErr := tx.AscendKeys("user:"+mxid+":avatarData100x100",
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
		urlPath := cli.BuildURL("profile", mxid, "avatar_url")
		s := struct {
			AvatarURL string `json:"avatar_url"`
		}{}

		cli.MakeRequest("GET", urlPath, nil, &s)
		avatarURL := s.AvatarURL

		// If avatarURL is not empty (aka. has a avatar set) download it at the size of 100x100. Else make the data string empty
		if avatarURL != "" {
			hsURL := cli.HomeserverURL.String()
			avatarURLSplits := strings.Split(strings.Replace(avatarURL, "mxc://", "", -1), "/")

			urlPath := hsURL + "/_matrix/media/r0/thumbnail/" + avatarURLSplits[0] + "/" + avatarURLSplits[1] + "?width=100&height=100"

			data, err := cli.MakeRequest("GET", urlPath, nil, nil)
			if err != nil {
				localLog.Println(err)
			}
			IMGdata = string(data[:])
		} else {
			//TODO Generate default image (Step: AfterUI)
			DisplayNameResp, _ := cli.GetOwnDisplayName()
			DisplayName := DisplayNameResp.DisplayName
			IMGdata = generateGenericImages(DisplayName)
		}

		// Update cache
		DBerr := db.Update(func(tx *buntdb.Tx) error {
			tx.Set("user:"+mxid+":avatarData100x100", IMGdata, nil)
			return nil
		})
		if DBerr != nil {
			localLog.Fatalln("DB ERROR: ", DBerr)
		}
	} else {
		IMGdata = avatarData
	}

	r := bytes.NewReader([]byte(IMGdata))
	srcIMG, _, err := image.Decode(r)
	if err != nil {
		localLog.Println("Decoder error: ", err)
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
		localLog.Println("Converting error: ", Converr)
	}

	str := fmt.Sprintf("%s", buf)
	avatar.LoadFromData(str, uint(len(str)), "", 0)
	return avatar
}
