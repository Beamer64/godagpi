package dagpi

import (
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

// Client Struct
type Client struct {
	Auth string
}

// request to get data
func httpGet(url string, c *Client) (map[string]interface{}, error) {
	httpClient := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", c.Auth)
	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer func(Body io.ReadCloser) {
		err = Body.Close()
	}(resp.Body)
	if err != nil {
		return nil, err
	}

	var data map[string]interface{}
	decoder := json.NewDecoder(resp.Body)

	err = decoder.Decode(&data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// attempting to get an image's buffer
func getImageBuffer(url string, c *Client) ([]byte, error) {
	httpClient := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", c.Auth)
	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer func(Body io.ReadCloser) {
		err = Body.Close()
	}(resp.Body)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

// as new routes are created in the aPI, their method calls will be added to the bottom of their respective region

//region Data aPI calls

// wtp returns an interface with all the Pokemon data
// Docs: https://dagpi.docs.apiary.io/#reference/data-api/whos-that-pokemon/who's-that-pokemon?
func (c *Client) wtp() (interface{}, error) {
	data, err := httpGet("https://api.dagpi.xyz/data/wtp", c)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// roast returns an interface containing a roast
// Docs: https://dagpi.docs.apiary.io/#reference/data-api/roast/roast
func (c *Client) roast() (interface{}, error) {
	data, err := httpGet("https://api.dagpi.xyz/data/roast", c)
	roast := data["roast"]
	if err != nil {
		return nil, err
	}

	return roast, nil
}

// joke returns an interface containing a joke & id
// Docs: https://dagpi.docs.apiary.io/#reference/data-api/joke/joke
func (c *Client) joke() (interface{}, error) {
	data, err := httpGet("https://api.dagpi.xyz/data/joke", c)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// fact returns an interface containing a fact
// Docs: https://dagpi.docs.apiary.io/#reference/data-api/fact/fact
func (c *Client) fact() (interface{}, error) {
	data, err := httpGet("https://api.dagpi.xyz/data/fact", c)
	fact := data["fact"]
	if err != nil {
		return nil, err
	}

	return fact, nil
}

// eightball returns an interface containing a response to 8ball question
// Docs: https://dagpi.docs.apiary.io/#reference/data-api/8ball/8ball
func (c *Client) eightball() (interface{}, error) {
	data, err := httpGet("https://api.dagpi.xyz/data/8ball", c)
	response := data["response"]
	if err != nil {
		return nil, err
	}

	return response, nil
}

// yomama returns an interface containing a description of yomama
// Docs: https://dagpi.docs.apiary.io/#reference/data-api/yomama/yomama
func (c *Client) yomama() (interface{}, error) {
	data, err := httpGet("https://api.dagpi.xyz/data/yomama", c)
	description := data["description"]
	if err != nil {
		return nil, err
	}

	return description, nil
}

// randomWaifu returns an interface containing data of a random waifu
// Docs: https://dagpi.docs.apiary.io/#reference/data-api/random-waifu/random-waifu
func (c *Client) randomWaifu() (interface{}, error) {
	data, err := httpGet("https://api.dagpi.xyz/data/waifu", c)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// waifu returns an interface containing data of a given waifu
// Docs: https://dagpi.docs.apiary.io/#reference/data-api/waifu-saerch/waifu-saerch
func (c *Client) waifu(waifuName string) (interface{}, error) {
	data, err := httpGet("https://api.dagpi.xyz/data/"+waifuName, c)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// pickupLine returns an interface containing category & joke
// Docs: https://dagpi.docs.apiary.io/#reference/data-api/pickup-line/pickup-line
func (c *Client) pickupLine() (interface{}, error) {
	data, err := httpGet("https://api.dagpi.xyz/data/pickupline", c)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// headLine returns an interface containing text and a bool, 'fake'
// Docs: https://dagpi.docs.apiary.io/#reference/data-api/headline/headline
func (c *Client) headLine() (interface{}, error) {
	data, err := httpGet("https://api.dagpi.xyz/data/headline", c)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// gtl returns an interface containing data of a random logo (Guess the Logo)
// Docs: https://dagpi.docs.apiary.io/#reference/data-api/guess-the-logo/guess-the-logo
func (c *Client) gtl() (interface{}, error) {
	data, err := httpGet("https://api.dagpi.xyz/data/logo", c)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// flag returns an interface containing data of a random flag
// Docs: https://dagpi.docs.apiary.io/#reference/data-api/flag/flag
func (c *Client) flag() (interface{}, error) {
	data, err := httpGet("https://api.dagpi.xyz/data/flag", c)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// captcha get a random captcha and answer
// Docs: https://dagpi.docs.apiary.io/#reference/data-api/captcha/captcha
func (c *Client) captcha() (interface{}, error) {
	data, err := httpGet("https://api.dagpi.xyz/data/captcha", c)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// typeracer get a sentence on an image, with a sentence to create typeracer games
// Docs: https://dagpi.docs.apiary.io/#reference/data-api/typeracer/typeracer
func (c *Client) typeracer() (interface{}, error) {
	data, err := httpGet("https://api.dagpi.xyz/data/typeracer", c)
	if err != nil {
		return nil, err
	}

	return data, nil
}

//endregion

//region Image Manipulation aPI calls

// pixelate allows you to pixelate an image.
// Docs: https://dagpi.docs.apiary.io/#reference/images-api/pixel/pixel
func (c *Client) pixelate(url string) ([]byte, error) {
	buffer, err := getImageBuffer("https://api.dagpi.xyz/image/pixel/?url="+url, c)
	if err != nil {
		return nil, err
	}

	return buffer, nil
}

// mirror an image along the y-axis
// Docs: https://dagpi.docs.apiary.io/#reference/images-api/mirror/mirror
func (c *Client) mirror(url string) ([]byte, error) {
	buffer, err := getImageBuffer("https://api.dagpi.xyz/image/mirror/?url="+url, c)
	if err != nil {
		return nil, err
	}

	return buffer, nil
}

// flipImage flip an image
// Docs: https://dagpi.docs.apiary.io/#reference/images-api/flip/flip
func (c *Client) flipImage(url string) ([]byte, error) {
	buffer, err := getImageBuffer("https://api.dagpi.xyz/image/flip/?url="+url, c)
	if err != nil {
		return nil, err
	}

	return buffer, nil
}

// colors allows you to get an Image with the colors present in the image.
// Docs: https://dagpi.docs.apiary.io/#reference/images-api/colors/colors
func (c *Client) colors(url string) ([]byte, error) {
	buffer, err := getImageBuffer("https://api.dagpi.xyz/image/colors/?url="+url, c)
	if err != nil {
		return nil, err
	}

	return buffer, nil
}

// america Let the star-spangled banner of the free and the brave soar.
// Docs:  https://dagpi.docs.apiary.io/#reference/images-api/america/america
func (c *Client) america(url string) ([]byte, error) {
	buffer, err := getImageBuffer("https://api.dagpi.xyz/image/america/?url="+url, c)
	if err != nil {
		return nil, err
	}

	return buffer, nil
}

// communism Support the soviet union comrade. Let the red flag fly!
// Docs: https://dagpi.docs.apiary.io/#reference/images-api/communism/communism
func (c *Client) communism(url string) ([]byte, error) {
	buffer, err := getImageBuffer("https://api.dagpi.xyz/image/communism/?url="+url, c)
	if err != nil {
		return nil, err
	}

	return buffer, nil
}

// triggered allows you to get a triggered gif.
// Docs: https://dagpi.docs.apiary.io/#reference/images-api/triggered/triggered
func (c *Client) triggered(url string) ([]byte, error) {
	buffer, err := getImageBuffer("https://api.dagpi.xyz/image/triggered/?url="+url, c)
	if err != nil {
		return nil, err
	}

	return buffer, nil
}

// expandImage animation that streches an image.
// Docs: https://dagpi.docs.apiary.io/#reference/images-api/expand/expand
func (c *Client) expandImage(url string) ([]byte, error) {
	buffer, err := getImageBuffer("https://api.dagpi.xyz/image/expand/?url="+url, c)
	if err != nil {
		return nil, err
	}

	return buffer, nil
}

// wasted allows you to get an image with GTA V Wasted screen.
// Docs: https://dagpi.docs.apiary.io/#reference/images-api/wasted/wasted
func (c *Client) wasted(url string) ([]byte, error) {
	buffer, err := getImageBuffer("https://api.dagpi.xyz/image/wasted/?url="+url, c)
	if err != nil {
		return nil, err
	}

	return buffer, nil
}

// sketch Cool effect that shows how an image would have been created by an artist.
// Docs: https://dagpi.docs.apiary.io/#reference/images-api/sketch/sketch
func (c *Client) sketch(url string) ([]byte, error) {
	buffer, err := getImageBuffer("https://api.dagpi.xyz/image/sketch/?url="+url, c)
	if err != nil {
		return nil, err
	}

	return buffer, nil
}

// spinImage You spin me right round baby.
// Docs: https://dagpi.docs.apiary.io/#reference/images-api/spin/spin
func (c *Client) spinImage(url string) ([]byte, error) {
	buffer, err := getImageBuffer("https://api.dagpi.xyz/image/spin/?url="+url, c)
	if err != nil {
		return nil, err
	}

	return buffer, nil
}

// petPet Pet pet gif
// Docs: https://dagpi.docs.apiary.io/#reference/images-api/petpet/petpet
func (c *Client) petPet(url string) ([]byte, error) {
	buffer, err := getImageBuffer("https://api.dagpi.xyz/image/petpet/?url="+url, c)
	if err != nil {
		return nil, err
	}

	return buffer, nil
}

// bonk Get bonked on my cheems
// Docs: https://dagpi.docs.apiary.io/#reference/images-api/bonk/bonk
func (c *Client) bonk(url string) ([]byte, error) {
	buffer, err := getImageBuffer("https://api.dagpi.xyz/image/bonk/?url="+url, c)
	if err != nil {
		return nil, err
	}

	return buffer, nil
}

// bomb explosion
// Docs: https://dagpi.docs.apiary.io/#reference/images-api/bomb/bomb
func (c *Client) bomb(url string) ([]byte, error) {
	buffer, err := getImageBuffer("https://api.dagpi.xyz/image/bomb/?url="+url, c)
	if err != nil {
		return nil, err
	}

	return buffer, nil
}

// shake a gif by having it wiggle.
// Docs: https://dagpi.docs.apiary.io/#reference/images-api/shake/shake
func (c *Client) shake(url string) ([]byte, error) {
	buffer, err := getImageBuffer("https://api.dagpi.xyz/image/shake/?url="+url, c)
	if err != nil {
		return nil, err
	}

	return buffer, nil
}

// invert allows you to get an image with an inverted color effect.
// Docs: https://dagpi.docs.apiary.io/#reference/images-api/invert/invert
func (c *Client) invert(url string) ([]byte, error) {
	buffer, err := getImageBuffer("https://api.dagpi.xyz/image/invert/?url="+url, c)
	if err != nil {
		return nil, err
	}

	return buffer, nil
}

// sobel allows you to get an image with the sobel effect.
// Docs: https://dagpi.docs.apiary.io/#reference/images-api/sobel/sobel
func (c *Client) sobel(url string) ([]byte, error) {
	buffer, err := getImageBuffer("https://api.dagpi.xyz/image/sobel/?url="+url, c)
	if err != nil {
		return nil, err
	}

	return buffer, nil
}

// hog Histogram of Oriented Gradients for an image.
// Docs: https://dagpi.docs.apiary.io/#reference/images-api/hog/hog
func (c *Client) hog(url string) ([]byte, error) {
	buffer, err := getImageBuffer("https://api.dagpi.xyz/image/hog/?url="+url, c)
	if err != nil {
		return nil, err
	}

	return buffer, nil
}

// triangle Cool triangle effect for an image.
// Docs: https://dagpi.docs.apiary.io/#reference/images-api/triangle/triangle
func (c *Client) triangle(url string) ([]byte, error) {
	buffer, err := getImageBuffer("https://api.dagpi.xyz/image/triangle/?url="+url, c)
	if err != nil {
		return nil, err
	}

	return buffer, nil
}

// blur blurs a given image.
// Docs: https://dagpi.docs.apiary.io/#reference/images-api/blur/blur
func (c *Client) blur(url string) ([]byte, error) {
	buffer, err := getImageBuffer("https://api.dagpi.xyz/image/blur/?url="+url, c)
	if err != nil {
		return nil, err
	}

	return buffer, nil
}

// rgb Get an RGB graph of an image's colors.
// Docs: https://dagpi.docs.apiary.io/#reference/images-api/rgb/rgb
func (c *Client) rgb(url string) ([]byte, error) {
	buffer, err := getImageBuffer("https://api.dagpi.xyz/image/rgb/?url="+url, c)
	if err != nil {
		return nil, err
	}

	return buffer, nil
}

// angel Image on the angels face.
// Docs: https://dagpi.docs.apiary.io/#reference/images-api/angel/angel
func (c *Client) angel(url string) ([]byte, error) {
	buffer, err := getImageBuffer("https://api.dagpi.xyz/image/angel/?url="+url, c)
	if err != nil {
		return nil, err
	}

	return buffer, nil
}

// satan Put an image on the devil.
// Docs: https://dagpi.docs.apiary.io/#reference/images-api/satan/satan
func (c *Client) satan(url string) ([]byte, error) {
	buffer, err := getImageBuffer("https://api.dagpi.xyz/image/satan/?url="+url, c)
	if err != nil {
		return nil, err
	}

	return buffer, nil
}

// delete Generates a Windows error meme based on a given image.
// Docs: https://dagpi.docs.apiary.io/#reference/images-api/delete/delete
func (c *Client) delete(url string) ([]byte, error) {
	buffer, err := getImageBuffer("https://api.dagpi.xyz/image/delete/?url="+url, c)
	if err != nil {
		return nil, err
	}

	return buffer, nil
}

// fedora Tips fedora in appreciation. Perry the Platypus.
// Docs: https://dagpi.docs.apiary.io/#reference/images-api/fedora/fedora
func (c *Client) fedora(url string) ([]byte, error) {
	buffer, err := getImageBuffer("https://api.dagpi.xyz/image/fedora/?url="+url, c)
	if err != nil {
		return nil, err
	}

	return buffer, nil
}

// hitler ?????
// Docs: https://dagpi.docs.apiary.io/#reference/images-api/hitler/hitler
func (c *Client) hitler(url string) ([]byte, error) {
	buffer, err := getImageBuffer("https://api.dagpi.xyz/image/hitler/?url="+url, c)
	if err != nil {
		return nil, err
	}

	return buffer, nil
}

// lego every group of pixels is a lego brick
// Docs: https://dagpi.docs.apiary.io/#reference/images-api/lego/lego
func (c *Client) lego(url string) ([]byte, error) {
	buffer, err := getImageBuffer("https://api.dagpi.xyz/image/lego/?url="+url, c)
	if err != nil {
		return nil, err
	}

	return buffer, nil
}

// wanted poster of an image.
// Docs: https://dagpi.docs.apiary.io/#reference/images-api/wanted/wanted
func (c *Client) wanted(url string) ([]byte, error) {
	buffer, err := getImageBuffer("https://api.dagpi.xyz/image/wanted/?url="+url, c)
	if err != nil {
		return nil, err
	}

	return buffer, nil
}

// stringify Turn your image into a ball of yarn.
// Docs: https://dagpi.docs.apiary.io/#reference/images-api/stringify/stringify
func (c *Client) stringify(url string) ([]byte, error) {
	buffer, err := getImageBuffer("https://api.dagpi.xyz/image/stringify/?url="+url, c)
	if err != nil {
		return nil, err
	}

	return buffer, nil
}

// burn Light your image on fire
// Docs: https://dagpi.docs.apiary.io/#reference/images-api/burn/burn
func (c *Client) burn(url string) ([]byte, error) {
	buffer, err := getImageBuffer("https://api.dagpi.xyz/image/burn/?url="+url, c)
	if err != nil {
		return nil, err
	}

	return buffer, nil
}

// earth The green and blue of the earth
// Docs: https://dagpi.docs.apiary.io/#reference/images-api/earth/earth
func (c *Client) earth(url string) ([]byte, error) {
	buffer, err := getImageBuffer("https://api.dagpi.xyz/image/earth/?url="+url, c)
	if err != nil {
		return nil, err
	}

	return buffer, nil
}

// freeze Blue ice like tint.
// Docs: https://dagpi.docs.apiary.io/#reference/images-api/freeze/freeze
func (c *Client) freeze(url string) ([]byte, error) {
	buffer, err := getImageBuffer("https://api.dagpi.xyz/image/freeze/?url="+url, c)
	if err != nil {
		return nil, err
	}

	return buffer, nil
}

// ground The power of the earth
// Docs: https://dagpi.docs.apiary.io/#reference/images-api/earth/earth
func (c *Client) ground(url string) ([]byte, error) {
	buffer, err := getImageBuffer("https://api.dagpi.xyz/image/ground/?url="+url, c)
	if err != nil {
		return nil, err
	}

	return buffer, nil
}

// mosiac Turn an image into a roman mosiac.
// Docs: https://dagpi.docs.apiary.io/#reference/images-api/mosiac/mosiac
func (c *Client) mosiac(url string) ([]byte, error) {
	buffer, err := getImageBuffer("https://api.dagpi.xyz/image/mosiac/?url="+url, c)
	if err != nil {
		return nil, err
	}

	return buffer, nil
}

// sithlord Put an image on the Laughs in Sithlord meme.
// Docs: https://dagpi.docs.apiary.io/#reference/images-api/sithlord/sithlord
func (c *Client) sithlord(url string) ([]byte, error) {
	buffer, err := getImageBuffer("https://api.dagpi.xyz/image/sith/?url="+url, c)
	if err != nil {
		return nil, err
	}

	return buffer, nil
}

// jail Put an image behind bars.
// Docs: https://dagpi.docs.apiary.io/#reference/images-api/jail/jail
func (c *Client) jail(url string) ([]byte, error) {
	buffer, err := getImageBuffer("https://api.dagpi.xyz/image/jail/?url="+url, c)
	if err != nil {
		return nil, err
	}

	return buffer, nil
}

// shatter Put an image behind bars.
// Docs: https://dagpi.docs.apiary.io/#reference/images-api/shatter/shatter
func (c *Client) shatter(url string) ([]byte, error) {
	buffer, err := getImageBuffer("https://api.dagpi.xyz/image/shatter/?url="+url, c)
	if err != nil {
		return nil, err
	}

	return buffer, nil
}

// pride Flag of your choice over an Image!
// available Choices: Asexual, Bisexual, Gay, Genderfluid, Genderqueer, Intersex, Lesbian, Nonbinary, Progress, Pan, Trans
// Docs: https://dagpi.docs.apiary.io/#reference/images-api/pride/pride
func (c *Client) pride(url string, flag string) ([]byte, error) {
	acceptableFlags := []string{
		"asexual",
		"bisexual",
		"gay",
		"genderfluid",
		"genderqueer",
		"intersex",
		"lesbian",
		"nonbinary",
		"progress",
		"pan",
		"trans",
	}
	for _, acceptableFlag := range acceptableFlags {
		if acceptableFlag == strings.ToLower(flag) {
			imgBuffer, err := getImageBuffer("https://api.dagpi.xyz/image/pride/?url="+url+"&flag="+flag, c)
			if err != nil {
				return nil, err
			}

			return imgBuffer, nil
		} else {
			err := errors.New("that pride flag is not accepted\nfunc Pride(url string, flag string) ([]byte, error)")

			return nil, err
		}
	}

	return nil, nil
}

// trash Image is trash.
// Docs: https://dagpi.docs.apiary.io/#reference/images-api/trash/trash
func (c *Client) trash(url string) ([]byte, error) {
	buffer, err := getImageBuffer("https://api.dagpi.xyz/image/trash/?url="+url, c)
	if err != nil {
		return nil, err
	}

	return buffer, nil
}

// deepfry an image.
// Docs: https://dagpi.docs.apiary.io/#reference/images-api/deepfry/deepfry
func (c *Client) deepfry(url string) ([]byte, error) {
	buffer, err := getImageBuffer("https://api.dagpi.xyz/image/deepfry/?url="+url, c)
	if err != nil {
		return nil, err
	}

	return buffer, nil
}

// ascii Cool hackerman effect for an image.
// Docs: https://dagpi.docs.apiary.io/#reference/images-api/ascii/ascii
func (c *Client) ascii(url string) ([]byte, error) {
	buffer, err := getImageBuffer("https://api.dagpi.xyz/image/ascii/?url="+url, c)
	if err != nil {
		return nil, err
	}

	return buffer, nil
}

// charcoal Image into a charcoal drawing.
// Docs: https://dagpi.docs.apiary.io/#reference/images-api/charcoal/charcoal
func (c *Client) charcoal(url string) ([]byte, error) {
	buffer, err := getImageBuffer("https://api.dagpi.xyz/image/charcoal/?url="+url, c)
	if err != nil {
		return nil, err
	}

	return buffer, nil
}

// posterize Posterizes an image.
// Docs: https://dagpi.docs.apiary.io/#reference/images-api/posterize/posterize
func (c *Client) posterize(url string) ([]byte, error) {
	buffer, err := getImageBuffer("https://api.dagpi.xyz/image/poster/?url="+url, c)
	if err != nil {
		return nil, err
	}

	return buffer, nil
}

// sepia Tone an image.
// Docs: https://dagpi.docs.apiary.io/#reference/images-api/sepia/sepia
func (c *Client) sepia(url string) ([]byte, error) {
	buffer, err := getImageBuffer("https://api.dagpi.xyz/image/sepia/?url="+url, c)
	if err != nil {
		return nil, err
	}

	return buffer, nil
}

// swirl an image.
// Docs: https://dagpi.docs.apiary.io/#reference/images-api/swirl/swirl
func (c *Client) swirl(url string) ([]byte, error) {
	buffer, err := getImageBuffer("https://api.dagpi.xyz/image/swirl/?url="+url, c)
	if err != nil {
		return nil, err
	}

	return buffer, nil
}

// paint Turn an image into art.
// Docs: https://dagpi.docs.apiary.io/#reference/images-api/paint/paint
func (c *Client) paint(url string) ([]byte, error) {
	buffer, err := getImageBuffer("https://api.dagpi.xyz/image/paint/?url="+url, c)
	if err != nil {
		return nil, err
	}

	return buffer, nil
}

// night Turn a day into night.
// Docs: https://dagpi.docs.apiary.io/#reference/images-api/night/night
func (c *Client) night(url string) ([]byte, error) {
	buffer, err := getImageBuffer("https://api.dagpi.xyz/image/night/?url="+url, c)
	if err != nil {
		return nil, err
	}

	return buffer, nil
}

// rainbow Some trippy light effects.
// Docs: https://dagpi.docs.apiary.io/#reference/images-api/rainbow/rainbow
func (c *Client) rainbow(url string) ([]byte, error) {
	buffer, err := getImageBuffer("https://api.dagpi.xyz/image/rainbow/?url="+url, c)
	if err != nil {
		return nil, err
	}

	return buffer, nil
}

// magik The much loved magik endpoint.
// Docs: https://dagpi.docs.apiary.io/#reference/images-api/magik/magik
func (c *Client) magik(url string) ([]byte, error) {
	buffer, err := getImageBuffer("https://api.dagpi.xyz/image/magik/?url="+url, c)
	if err != nil {
		return nil, err
	}

	return buffer, nil
}

// fivegOneg The meme.
// Docs: https://dagpi.docs.apiary.io/#reference/images-api/five-guys-one-girl/five-guys-one-girl
func (c *Client) fivegOneg(url1 string, url2 string) ([]byte, error) {
	buffer, err := getImageBuffer("https://api.dagpi.xyz/image/5g1g/?url="+url1+"&url2="+url2, c)
	if err != nil {
		return nil, err
	}

	return buffer, nil
}

// whyAreYouGay The meme.
// Docs: https://dagpi.docs.apiary.io/#reference/images-api/why-are-you-gay/why-are-you-gay
func (c *Client) whyAreYouGay(url1 string, url2 string) ([]byte, error) {
	buffer, err := getImageBuffer("https://api.dagpi.xyz/image/whyareyougay/?url="+url1+"&url2="+url2, c)
	if err != nil {
		return nil, err
	}

	return buffer, nil
}

// slap Have one image slap another.
// Docs: https://dagpi.docs.apiary.io/#reference/images-api/slap/slap
func (c *Client) slap(url1 string, url2 string) ([]byte, error) {
	buffer, err := getImageBuffer("https://api.dagpi.xyz/image/slap/?url="+url1+"&url2="+url2, c)
	if err != nil {
		return nil, err
	}

	return buffer, nil
}

// obama The meme.
// Docs: https://dagpi.docs.apiary.io/#reference/images-api/obama/obama
func (c *Client) obama(url1 string, url2 string) ([]byte, error) {
	buffer, err := getImageBuffer("https://api.dagpi.xyz/image/obama/?url="+url1+"&url2="+url2, c)
	if err != nil {
		return nil, err
	}

	return buffer, nil
}

// tweet The meme.
// Docs: https://dagpi.docs.apiary.io/#reference/images-api/tweet/tweet
func (c *Client) tweet(url string, username string, text string) ([]byte, error) {
	buffer, err := getImageBuffer("https://api.dagpi.xyz/image/tweet/?url="+url+"&username="+username+"&text="+text, c)
	if err != nil {
		return nil, err
	}

	return buffer, nil
}

// youTubeComment Generate realistic YouTube messages
// Docs: https://dagpi.docs.apiary.io/#reference/images-api/youtube-comment/youtube-comment
func (c *Client) youTubeComment(url string, username string, text string, darkMode bool) ([]byte, error) {
	if darkMode == true {
		buffer, err := getImageBuffer("https://api.dagpi.xyz/image/yt/?url="+url+"&username="+username+"&text="+text+"&dark="+"true", c)
		if err != nil {
			return nil, err
		}

		return buffer, nil
	} else {
		buffer, err := getImageBuffer("https://api.dagpi.xyz/image/yt/?url="+url+"&username="+username+"&text="+text+"&dark="+"false", c)
		if err != nil {
			return nil, err
		}

		return buffer, nil
	}
}

// discord Generate realistic discord messages
// Docs: https://dagpi.docs.apiary.io/#reference/images-api/discord/discord
func (c *Client) discord(url string, username string, text string, darkMode bool) ([]byte, error) {
	if darkMode == true {
		buffer, err := getImageBuffer("https://api.dagpi.xyz/image/discord/?url="+url+"&username="+username+"&text="+text+"&dark="+"true", c)
		if err != nil {
			return nil, err
		}

		return buffer, nil
	} else {
		buffer, err := getImageBuffer("https://api.dagpi.xyz/image/discord/?url="+url+"&username="+username+"&text="+text+"&dark="+"false", c)
		if err != nil {
			return nil, err
		}

		return buffer, nil
	}
}

// retromeme The good old memes. Generated.
// Docs: https://dagpi.docs.apiary.io/#reference/images-api/retromeme/retromeme
func (c *Client) retromeme(url string, topText string, bottomText string) ([]byte, error) {
	buffer, err := getImageBuffer("https://api.dagpi.xyz/image/retromeme/?url="+url+"&top_text="+topText+"&bottom_text="+bottomText, c)
	if err != nil {
		return nil, err
	}

	return buffer, nil
}

// motivational The black background with top and bottom motivational text.
// Docs: https://dagpi.docs.apiary.io/#reference/images-api/motivational/motivational
func (c *Client) motivational(url string, topText string, bottomText string) ([]byte, error) {
	buffer, err := getImageBuffer("https://api.dagpi.xyz/image/motiv/?url="+url+"&top_text="+topText+"&bottom_text="+bottomText, c)
	if err != nil {
		return nil, err
	}

	return buffer, nil
}

// modernmeme a modern meme generation system that allows reddit ready memes with just one endpoint.
// Docs: https://dagpi.docs.apiary.io/#reference/images-api/modernmeme/modernmeme
func (c *Client) modernmeme(url string, text string) ([]byte, error) {
	buffer, err := getImageBuffer("https://api.dagpi.xyz/image/modernmeme/?url="+url+"&text="+text, c)
	if err != nil {
		return nil, err
	}

	return buffer, nil
}

// elmo Burning elmo Meme
// Docs: todo add docs when available
func (c *Client) elmo(url string) ([]byte, error) {
	buffer, err := getImageBuffer("https://api.dagpi.xyz/image/elmo/?url="+url, c)
	if err != nil {
		return nil, err
	}

	return buffer, nil
}

// tvStatic Its TV static
// Docs: todo add docs when available
func (c *Client) tvStatic(url string) ([]byte, error) {
	buffer, err := getImageBuffer("https://api.dagpi.xyz/image/tv/?url="+url, c)
	if err != nil {
		return nil, err
	}

	return buffer, nil
}

// glitch todo add description when available
// Docs: todo add docs when available
func (c *Client) glitch(url string) ([]byte, error) {
	buffer, err := getImageBuffer("https://api.dagpi.xyz/image/glitch/?url="+url, c)
	if err != nil {
		return nil, err
	}

	return buffer, nil
}

// glitchStatic todo add description when available
// Docs: todo add docs when available
func (c *Client) glitchStatic(url string) ([]byte, error) {
	buffer, err := getImageBuffer("https://api.dagpi.xyz/image/glitchstatic/?url="+url, c)
	if err != nil {
		return nil, err
	}

	return buffer, nil
}

// album Make an album cover!
// Docs: todo add docs when available
func (c *Client) album(url string) ([]byte, error) {
	buffer, err := getImageBuffer("https://api.dagpi.xyz/image/album/?url="+url, c)
	if err != nil {
		return nil, err
	}

	return buffer, nil
}

//endregion
