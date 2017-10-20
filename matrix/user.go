package matrix

import (
	"bytes"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"

	"github.com/Nordgedanken/Morpheus/util"
	"github.com/disintegration/letteravatar"
	"github.com/matrix-org/gomatrix"
	"github.com/therecipe/qt/gui"
	"github.com/tidwall/buntdb"

	// image/jpeg needed to load jpeg images
	_ "image/jpeg"
	// golang.org/x/image/webp needed to load webp images
	_ "golang.org/x/image/webp"
	// golang.org/x/image/bmp needed to load bmp images
	_ "golang.org/x/image/bmp"
)

var localLog *log.Logger

func init() {
	var file *os.File
	var err error

	localLog = util.Logger()
	localLog, file, err = util.StartFileLog(localLog)
	if err != nil {
		localLog.Fatalln(err)
	}
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

func generateGenericImages(displayname string, size int) (imgData string, err error) {
	identifier := displayname
	if (identifier[0] == '#' || identifier[0] == '!' || identifier[0] == '@') && len(identifier) > 1 {
		identifier = identifier[1:]
	}

	avatarChar, _ := utf8.DecodeRuneInString(identifier)
	img, LetterAvatarErr := letteravatar.Draw(size, unicode.ToUpper(avatarChar), nil)

	if LetterAvatarErr != nil {
		err = LetterAvatarErr
		return
	}

	localLog.Println("generatedIMGBounds: ", img.Bounds())
	buf := new(bytes.Buffer)
	EncodeErr := png.Encode(buf, img)
	if EncodeErr != nil {
		err = EncodeErr
		return
	}
	imgData = buf.String()
	return
}

// GetOwnUserAvatar returns a *gui.QPixmap of an UserAvatar
func GetOwnUserAvatar(cli *gomatrix.Client) (avatar *gui.QPixmap, err error) {
	avatar, err = GetUserAvatar(cli, cli.UserID, 61)
	return
}

// GetUserAvatar returns a *gui.QPixmap of an UserAvatar
func GetUserAvatar(cli *gomatrix.Client, mxid string, size int) (avatarResp *gui.QPixmap, err error) {
	db, DBOpenErr := OpenCacheDB()
	if DBOpenErr != nil {
		localLog.Fatalln(DBOpenErr)
	}
	defer db.Close()

	// Init local vars
	var avatarData string
	var IMGdata string

	// Get cache
	DBErr := db.View(func(tx *buntdb.Tx) error {
		QueryErr := tx.AscendKeys("user|"+mxid+"|avatarData"+strconv.Itoa(size)+"x"+strconv.Itoa(size),
			func(key, value string) bool {
				avatarData = value
				return true
			})
		return QueryErr
	})
	if DBErr != nil {
		err = DBErr
		return
	}

	//If cache is empty do a ServerQuery
	if avatarData == "" {
		// Get avatarURL
		urlPath := cli.BuildURL("profile", mxid, "avatar_url")
		s := struct {
			AvatarURL string `json:"avatar_url"`
		}{}

		_, ReqErr := cli.MakeRequest("GET", urlPath, nil, &s)
		if ReqErr != nil {
			err = ReqErr
			return
		}
		avatarURL := s.AvatarURL

		// If avatarURL is not empty (aka. has a avatar set) download it at the size of 100x100. Else make the data string empty
		if avatarURL != "" {
			hsURL := cli.HomeserverURL.String()
			avatarURLSplits := strings.Split(strings.Replace(avatarURL, "mxc://", "", -1), "/")

			urlPath := hsURL + "/_matrix/media/r0/thumbnail/" + avatarURLSplits[0] + "/" + avatarURLSplits[1] + "?width=" + strconv.Itoa(size) + "&height=" + strconv.Itoa(size) + "&method=crop"

			data, ReqErr := cli.MakeRequest("GET", urlPath, nil, nil)
			if ReqErr != nil {
				err = ReqErr
				return
			}
			IMGdata = string(data[:])
		} else {
			localLog.Println("Generating Avatar")
			DisplayNameResp, _ := cli.GetDisplayName(mxid)
			DisplayName := DisplayNameResp.DisplayName
			var GenerateImgErr error
			IMGdata, GenerateImgErr = generateGenericImages(DisplayName, size)
			if GenerateImgErr != nil {
				err = GenerateImgErr
				return
			}
		}

		// Update cache
		DBerr := db.Update(func(tx *buntdb.Tx) error {
			_, _, DBSetErr := tx.Set("user|"+mxid+"|avatarData"+strconv.Itoa(size)+"x"+strconv.Itoa(size), IMGdata, nil)
			return DBSetErr
		})
		if DBerr != nil {
			err = DBErr
			return
		}
	} else {
		IMGdata = avatarData
	}

	r := bytes.NewReader([]byte(IMGdata))
	srcIMG, _, DecodeErr := image.Decode(r)
	if DecodeErr != nil {
		err = DecodeErr
	}

	// Convert avatarimage to QPixmap for usage in QT
	canvas := image.NewRGBA(srcIMG.Bounds())
	cx := srcIMG.Bounds().Min.X + srcIMG.Bounds().Dx()/2
	cy := srcIMG.Bounds().Min.Y + srcIMG.Bounds().Dy()/2
	draw.DrawMask(canvas, canvas.Bounds(), srcIMG, image.ZP, &circle{image.Point{cx, cy}, cx}, image.ZP, draw.Over)

	avatar := gui.NewQPixmap()
	buf := new(bytes.Buffer)
	ConvErr := png.Encode(buf, canvas)
	if ConvErr != nil {
		err = ConvErr
	}

	str := buf.String()
	avatar.LoadFromData(str, uint(len(str)), "", 0)
	avatarResp = avatar
	return
}
