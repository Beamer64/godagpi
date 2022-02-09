package dagpi

import (
	"encoding/json"
	"errors"
	"fmt"
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

// Attempting to get an image's buffer
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

// As new routes are created in the API, their method calls will be added to the bottom of their respective region

//region Data API calls

// WTP returns an interface with all the Pokemon data
// Docs: https://dagpi.docs.apiary.io/#reference/data-api/whos-that-pokemon/who's-that-pokemon?
func (c *Client) WTP() (interface{}, error) {
	data, err := httpGet("https://api.dagpi.xyz/data/wtp", c)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// Roast returns an interface containing a roast
// Docs: https://dagpi.docs.apiary.io/#reference/data-api/roast/roast
func (c *Client) Roast() (interface{}, error) {
	data, err := httpGet("https://api.dagpi.xyz/data/roast", c)
	roast := data["roast"]
	if err != nil {
		return nil, err
	}

	return roast, nil
}

// Joke returns an interface containing a joke & id
// Docs: https://dagpi.docs.apiary.io/#reference/data-api/joke/joke
func (c *Client) Joke() (interface{}, error) {
	data, err := httpGet("https://api.dagpi.xyz/data/joke", c)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// Fact returns an interface containing a fact
// Docs: https://dagpi.docs.apiary.io/#reference/data-api/fact/fact
func (c *Client) Fact() (interface{}, error) {
	data, err := httpGet("https://api.dagpi.xyz/data/fact", c)
	fact := data["fact"]
	if err != nil {
		return nil, err
	}

	return fact, nil
}

// Eightball returns an interface containing a response to 8ball question
// Docs: https://dagpi.docs.apiary.io/#reference/data-api/8ball/8ball
func (c *Client) Eightball() (interface{}, error) {
	data, err := httpGet("https://api.dagpi.xyz/data/8ball", c)
	response := data["response"]
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Yomama returns an interface containing a description of yomama
// Docs: https://dagpi.docs.apiary.io/#reference/data-api/yomama/yomama
func (c *Client) Yomama() (interface{}, error) {
	data, err := httpGet("https://api.dagpi.xyz/data/yomama", c)
	description := data["description"]
	if err != nil {
		return nil, err
	}

	return description, nil
}

// RandomWaifu returns an interface containing data of a random waifu
// Docs: https://dagpi.docs.apiary.io/#reference/data-api/random-waifu/random-waifu
func (c *Client) RandomWaifu() (interface{}, error) {
	data, err := httpGet("https://api.dagpi.xyz/data/waifu", c)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// Waifu returns an interface containing data of a given waifu
// Docs: https://dagpi.docs.apiary.io/#reference/data-api/waifu-saerch/waifu-saerch
func (c *Client) Waifu(waifuName string) (interface{}, error) {
	data, err := httpGet("https://api.dagpi.xyz/data/"+waifuName, c)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// PickupLine returns an interface containing category & joke
// Docs: https://dagpi.docs.apiary.io/#reference/data-api/pickup-line/pickup-line
func (c *Client) PickupLine() (interface{}, error) {
	data, err := httpGet("https://api.dagpi.xyz/data/pickupline", c)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// HeadLine returns an interface containing text and a bool, 'fake'
// Docs: https://dagpi.docs.apiary.io/#reference/data-api/headline/headline
func (c *Client) HeadLine() (interface{}, error) {
	data, err := httpGet("https://api.dagpi.xyz/data/headline", c)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// GTL returns an interface containing data of a random logo (Guess the Logo)
// Docs: https://dagpi.docs.apiary.io/#reference/data-api/guess-the-logo/guess-the-logo
func (c *Client) GTL() (interface{}, error) {
	data, err := httpGet("https://api.dagpi.xyz/data/logo", c)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// Flag returns an interface containing data of a random flag
// Docs: https://dagpi.docs.apiary.io/#reference/data-api/flag/flag
func (c *Client) Flag() (interface{}, error) {
	data, err := httpGet("https://api.dagpi.xyz/data/flag", c)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// Captcha get a random captcha and answer
// Docs: https://dagpi.docs.apiary.io/#reference/data-api/captcha/captcha
func (c *Client) Captcha() (interface{}, error) {
	data, err := httpGet("https://api.dagpi.xyz/data/captcha", c)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// Typeracer get a sentence on an image, with a sentence to create typeracer games
// Docs: https://dagpi.docs.apiary.io/#reference/data-api/typeracer/typeracer
func (c *Client) Typeracer() (interface{}, error) {
	data, err := httpGet("https://api.dagpi.xyz/data/typeracer", c)
	if err != nil {
		return nil, err
	}

	return data, nil
}

//endregion

//region Image Manipulation API calls

// Pixelate Allows you to pixelate an image.
// Docs: https://dagpi.docs.apiary.io/#reference/images-api/pixel/pixel
func (c *Client) Pixelate(url string) ([]byte, error) {
	buffer, err := getImageBuffer("https://api.dagpi.xyz/image/pixel/?url="+url, c)
	if err != nil {
		return nil, err
	}

	return buffer, nil
}

// Mirror an image along the y-axis
// Docs: https://dagpi.docs.apiary.io/#reference/images-api/mirror/mirror
func (c *Client) Mirror(url string) ([]byte, error) {
	buffer, err := getImageBuffer("https://api.dagpi.xyz/image/mirror/?url="+url, c)
	if err != nil {
		return nil, err
	}

	return buffer, nil
}

// FlipImage flip an image
// Docs: https://dagpi.docs.apiary.io/#reference/images-api/flip/flip
func (c *Client) FlipImage(url string) ([]byte, error) {
	buffer, err := getImageBuffer("https://api.dagpi.xyz/image/flip/?url="+url, c)
	if err != nil {
		return nil, err
	}

	return buffer, nil
}

// Colors Allows you to get an Image with the colors present in the image.
// Docs: https://dagpi.docs.apiary.io/#reference/images-api/colors/colors
func (c *Client) Colors(url string) ([]byte, error) {
	buffer, err := getImageBuffer("https://api.dagpi.xyz/image/colors/?url="+url, c)
	if err != nil {
		return nil, err
	}

	return buffer, nil
}

// America Let the star-spangled banner of the free and the brave soar.
// Docs:  https://dagpi.docs.apiary.io/#reference/images-api/america/america
func (c *Client) America(url string) ([]byte, error) {
	buffer, err := getImageBuffer("https://api.dagpi.xyz/image/america/?url="+url, c)
	if err != nil {
		return nil, err
	}

	return buffer, nil
}

// Communism Support the soviet union comrade. Let the red flag fly!
// Docs: https://dagpi.docs.apiary.io/#reference/images-api/communism/communism
func (c *Client) Communism(url string) ([]byte, error) {
	buffer, err := getImageBuffer("https://api.dagpi.xyz/image/communism/?url="+url, c)
	if err != nil {
		return nil, err
	}

	return buffer, nil
}

// Triggered Allows you to get a triggered gif.
// Docs: https://dagpi.docs.apiary.io/#reference/images-api/triggered/triggered
func (c *Client) Triggered(url string) ([]byte, error) {
	buffer, err := getImageBuffer("https://api.dagpi.xyz/image/triggered/?url="+url, c)
	if err != nil {
		return nil, err
	}

	return buffer, nil
}

// ExpandImage animation that streches an image.
// Docs: https://dagpi.docs.apiary.io/#reference/images-api/expand/expand
func (c *Client) ExpandImage(url string) ([]byte, error) {
	buffer, err := getImageBuffer("https://api.dagpi.xyz/image/expand/?url="+url, c)
	if err != nil {
		return nil, err
	}

	return buffer, nil
}

// Wasted Allows you to get an image with GTA V Wasted screen.
// Docs: https://dagpi.docs.apiary.io/#reference/images-api/wasted/wasted
func (c *Client) Wasted(url string) ([]byte, error) {
	buffer, err := getImageBuffer("https://api.dagpi.xyz/image/wasted/?url="+url, c)
	if err != nil {
		return nil, err
	}

	return buffer, nil
}

// Sketch Cool efffect that shows how an image would have been created by an artist.
// Docs: https://dagpi.docs.apiary.io/#reference/images-api/sketch/sketch
func (c *Client) Sketch(url string) ([]byte, error) {
	buffer, err := getImageBuffer("https://api.dagpi.xyz/image/sketch/?url="+url, c)
	if err != nil {
		return nil, err
	}

	return buffer, nil
}

// SpinImage You spin me right round baby.
// Docs: https://dagpi.docs.apiary.io/#reference/images-api/spin/spin
func (c *Client) SpinImage(url string) ([]byte, error) {
	buffer, err := getImageBuffer("https://api.dagpi.xyz/image/spin/?url="+url, c)
	if err != nil {
		return nil, err
	}

	return buffer, nil
}

// PetPet Pet pet gif
// Docs: https://dagpi.docs.apiary.io/#reference/images-api/petpet/petpet
func (c *Client) PetPet(url string) ([]byte, error) {
	buffer, err := getImageBuffer("https://api.dagpi.xyz/image/petpet/?url="+url, c)
	if err != nil {
		return nil, err
	}

	return buffer, nil
}

// Bonk Get bonked on my cheems
// Docs: https://dagpi.docs.apiary.io/#reference/images-api/bonk/bonk
func (c *Client) Bonk(url string) ([]byte, error) {
	buffer, err := getImageBuffer("https://api.dagpi.xyz/image/bonk/?url="+url, c)
	if err != nil {
		return nil, err
	}

	return buffer, nil
}

// Bomb Explosion
// Docs: https://dagpi.docs.apiary.io/#reference/images-api/bomb/bomb
func (c *Client) Bomb(url string) ([]byte, error) {
	buffer, err := getImageBuffer("https://api.dagpi.xyz/image/bomb/?url="+url, c)
	if err != nil {
		return nil, err
	}

	return buffer, nil
}

// Shake a gif by having it wiggle.
// Docs: https://dagpi.docs.apiary.io/#reference/images-api/shake/shake
func (c *Client) Shake(url string) ([]byte, error) {
	buffer, err := getImageBuffer("https://api.dagpi.xyz/image/shake/?url="+url, c)
	if err != nil {
		return nil, err
	}

	return buffer, nil
}

// Invert Allows you to get an image with an inverted color effect.
// Docs: https://dagpi.docs.apiary.io/#reference/images-api/invert/invert
func (c *Client) Invert(url string) ([]byte, error) {
	buffer, err := getImageBuffer("https://api.dagpi.xyz/image/invert/?url="+url, c)
	if err != nil {
		return nil, err
	}

	return buffer, nil
}

// Sobel Allows you to get an image with the sobel effect.
// Docs: https://dagpi.docs.apiary.io/#reference/images-api/sobel/sobel
func (c *Client) Sobel(url string) ([]byte, error) {
	buffer, err := getImageBuffer("https://api.dagpi.xyz/image/sobel/?url="+url, c)
	if err != nil {
		return nil, err
	}

	return buffer, nil
}

// Hog Histogram of Oriented Gradients for an image.
// Docs: https://dagpi.docs.apiary.io/#reference/images-api/hog/hog
func (c *Client) Hog(url string) ([]byte, error) {
	buffer, err := getImageBuffer("https://api.dagpi.xyz/image/hog/?url="+url, c)
	if err != nil {
		return nil, err
	}

	return buffer, nil
}

// Triangle Cool triangle effect for an image.
// Docs: https://dagpi.docs.apiary.io/#reference/images-api/triangle/triangle
func (c *Client) Triangle(url string) ([]byte, error) {
	buffer, err := getImageBuffer("https://api.dagpi.xyz/image/triangle/?url="+url, c)
	if err != nil {
		return nil, err
	}

	return buffer, nil
}

// Blur Blurs a given image.
// Docs: https://dagpi.docs.apiary.io/#reference/images-api/blur/blur
func (c *Client) Blur(url string) ([]byte, error) {
	buffer, err := getImageBuffer("https://api.dagpi.xyz/image/blur/?url="+url, c)
	if err != nil {
		return nil, err
	}

	return buffer, nil
}

// RGB Get an RGB graph of an image's colors.
// Docs: https://dagpi.docs.apiary.io/#reference/images-api/rgb/rgb
func (c *Client) RGB(url string) ([]byte, error) {
	buffer, err := getImageBuffer("https://api.dagpi.xyz/image/rgb/?url="+url, c)
	if err != nil {
		return nil, err
	}

	return buffer, nil
}

// Angel Image on the Angels face.
// Docs: https://dagpi.docs.apiary.io/#reference/images-api/angel/angel
func (c *Client) Angel(url string) ([]byte, error) {
	buffer, err := getImageBuffer("https://api.dagpi.xyz/image/angel/?url="+url, c)
	if err != nil {
		return nil, err
	}

	return buffer, nil
}

// Satan Put an image on the devil.
// Docs: https://dagpi.docs.apiary.io/#reference/images-api/satan/satan
func (c *Client) Satan(url string) ([]byte, error) {
	buffer, err := getImageBuffer("https://api.dagpi.xyz/image/satan/?url="+url, c)
	if err != nil {
		return nil, err
	}

	return buffer, nil
}

// Delete Generates a Windows error meme based on a given image.
// Docs: https://dagpi.docs.apiary.io/#reference/images-api/delete/delete
func (c *Client) Delete(url string) ([]byte, error) {
	buffer, err := getImageBuffer("https://api.dagpi.xyz/image/delete/?url="+url, c)
	if err != nil {
		return nil, err
	}

	return buffer, nil
}

// Fedora Tips fedora in appreciation. Perry the Platypus.
// Docs: https://dagpi.docs.apiary.io/#reference/images-api/fedora/fedora
func (c *Client) Fedora(url string) ([]byte, error) {
	buffer, err := getImageBuffer("https://api.dagpi.xyz/image/fedora/?url="+url, c)
	if err != nil {
		return nil, err
	}

	return buffer, nil
}

// Hitler ?????
// Docs: https://dagpi.docs.apiary.io/#reference/images-api/hitler/hitler
func (c *Client) Hitler(url string) ([]byte, error) {
	buffer, err := getImageBuffer("https://api.dagpi.xyz/image/hitler/?url="+url, c)
	if err != nil {
		return nil, err
	}

	return buffer, nil
}

// Lego Every group of pixels is a lego brick
// Docs: https://dagpi.docs.apiary.io/#reference/images-api/lego/lego
func (c *Client) Lego(url string) ([]byte, error) {
	buffer, err := getImageBuffer("https://api.dagpi.xyz/image/lego/?url="+url, c)
	if err != nil {
		return nil, err
	}

	return buffer, nil
}

// Wanted poster of an image.
// Docs: https://dagpi.docs.apiary.io/#reference/images-api/wanted/wanted
func (c *Client) Wanted(url string) ([]byte, error) {
	buffer, err := getImageBuffer("https://api.dagpi.xyz/image/wanted/?url="+url, c)
	if err != nil {
		return nil, err
	}

	return buffer, nil
}

// Stringify Turn your image into a ball of yarn.
// Docs: https://dagpi.docs.apiary.io/#reference/images-api/stringify/stringify
func (c *Client) Stringify(url string) ([]byte, error) {
	buffer, err := getImageBuffer("https://api.dagpi.xyz/image/stringify/?url="+url, c)
	if err != nil {
		return nil, err
	}

	return buffer, nil
}

// Burn Light your image on fire
// Docs: https://dagpi.docs.apiary.io/#reference/images-api/burn/burn
func (c *Client) Burn(url string) ([]byte, error) {
	buffer, err := getImageBuffer("https://api.dagpi.xyz/image/burn/?url="+url, c)
	if err != nil {
		return nil, err
	}

	return buffer, nil
}

// Earth The green and blue of the earth
// Docs: https://dagpi.docs.apiary.io/#reference/images-api/earth/earth
func (c *Client) Earth(url string) ([]byte, error) {
	buffer, err := getImageBuffer("https://api.dagpi.xyz/image/earth/?url="+url, c)
	if err != nil {
		return nil, err
	}

	return buffer, nil
}

// Freeze Blue ice like tint.
// Docs: https://dagpi.docs.apiary.io/#reference/images-api/freeze/freeze
func (c *Client) Freeze(url string) ([]byte, error) {
	buffer, err := getImageBuffer("https://api.dagpi.xyz/image/freeze/?url="+url, c)
	if err != nil {
		return nil, err
	}

	return buffer, nil
}

// Ground The poower of the earth
// Docs: https://dagpi.docs.apiary.io/#reference/images-api/earth/earth
func (c *Client) Ground(url string) ([]byte, error) {
	buffer, err := getImageBuffer("https://api.dagpi.xyz/image/ground/?url="+url, c)
	if err != nil {
		return nil, err
	}

	return buffer, nil
}

// Mosiac Turn an image into a roman mosiac.
// Docs: https://dagpi.docs.apiary.io/#reference/images-api/mosiac/mosiac
func (c *Client) Mosiac(url string) ([]byte, error) {
	buffer, err := getImageBuffer("https://api.dagpi.xyz/image/mosiac/?url="+url, c)
	if err != nil {
		return nil, err
	}

	return buffer, nil
}

// Sithlord Put an image on the Laughs in Sithlord meme.
// Docs: https://dagpi.docs.apiary.io/#reference/images-api/sithlord/sithlord
func (c *Client) Sithlord(url string) ([]byte, error) {
	buffer, err := getImageBuffer("https://api.dagpi.xyz/image/sith/?url="+url, c)
	if err != nil {
		return nil, err
	}

	return buffer, nil
}

// Jail Put an image behind bars.
// Docs: https://dagpi.docs.apiary.io/#reference/images-api/jail/jail
func (c *Client) Jail(url string) ([]byte, error) {
	buffer, err := getImageBuffer("https://api.dagpi.xyz/image/jail/?url="+url, c)
	if err != nil {
		return nil, err
	}

	return buffer, nil
}

// Shatter Put an image behind bars.
// Docs: https://dagpi.docs.apiary.io/#reference/images-api/shatter/shatter
func (c *Client) Shatter(url string) ([]byte, error) {
	buffer, err := getImageBuffer("https://api.dagpi.xyz/image/shatter/?url="+url, c)
	if err != nil {
		return nil, err
	}

	return buffer, nil
}

// Pride Flag of your choice over an Image!
// Available Choices: Asexual, Bisexual, Gay, Genderfluid, Genderqueer, Intersex, Lesbian, Nonbinary, Progress, Pan, Trans
// Docs: https://dagpi.docs.apiary.io/#reference/images-api/pride/pride
func (c *Client) Pride(url string, flag string) ([]byte, error) {
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
		}
	}
	err := errors.New(fmt.Sprintf("the '%s' pride flag is not accepted\nfunc Pride(url string, flag string) ([]byte, error)", flag))

	return nil, err
}

// Trash Image is trash.
// Docs: https://dagpi.docs.apiary.io/#reference/images-api/trash/trash
func (c *Client) Trash(url string) ([]byte, error) {
	buffer, err := getImageBuffer("https://api.dagpi.xyz/image/trash/?url="+url, c)
	if err != nil {
		return nil, err
	}

	return buffer, nil
}

// Deepfry an image.
// Docs: https://dagpi.docs.apiary.io/#reference/images-api/deepfry/deepfry
func (c *Client) Deepfry(url string) ([]byte, error) {
	buffer, err := getImageBuffer("https://api.dagpi.xyz/image/deepfry/?url="+url, c)
	if err != nil {
		return nil, err
	}

	return buffer, nil
}

// Ascii Cool hackerman effect for an image.
// Docs: https://dagpi.docs.apiary.io/#reference/images-api/ascii/ascii
func (c *Client) Ascii(url string) ([]byte, error) {
	buffer, err := getImageBuffer("https://api.dagpi.xyz/image/ascii/?url="+url, c)
	if err != nil {
		return nil, err
	}

	return buffer, nil
}

// Charcoal Image into a charcoal drawing.
// Docs: https://dagpi.docs.apiary.io/#reference/images-api/charcoal/charcoal
func (c *Client) Charcoal(url string) ([]byte, error) {
	buffer, err := getImageBuffer("https://api.dagpi.xyz/image/charcoal/?url="+url, c)
	if err != nil {
		return nil, err
	}

	return buffer, nil
}

// Posterize Posterizes an image.
// Docs: https://dagpi.docs.apiary.io/#reference/images-api/posterize/posterize
func (c *Client) Posterize(url string) ([]byte, error) {
	buffer, err := getImageBuffer("https://api.dagpi.xyz/image/poster/?url="+url, c)
	if err != nil {
		return nil, err
	}

	return buffer, nil
}

// Sepia Tone an image.
// Docs: https://dagpi.docs.apiary.io/#reference/images-api/sepia/sepia
func (c *Client) Sepia(url string) ([]byte, error) {
	buffer, err := getImageBuffer("https://api.dagpi.xyz/image/sepia/?url="+url, c)
	if err != nil {
		return nil, err
	}

	return buffer, nil
}

// Swirl an image.
// Docs: https://dagpi.docs.apiary.io/#reference/images-api/swirl/swirl
func (c *Client) Swirl(url string) ([]byte, error) {
	buffer, err := getImageBuffer("https://api.dagpi.xyz/image/swirl/?url="+url, c)
	if err != nil {
		return nil, err
	}

	return buffer, nil
}

// Paint Turn an image into art.
// Docs: https://dagpi.docs.apiary.io/#reference/images-api/paint/paint
func (c *Client) Paint(url string) ([]byte, error) {
	buffer, err := getImageBuffer("https://api.dagpi.xyz/image/paint/?url="+url, c)
	if err != nil {
		return nil, err
	}

	return buffer, nil
}

// Night Turn a day into night.
// Docs: https://dagpi.docs.apiary.io/#reference/images-api/night/night
func (c *Client) Night(url string) ([]byte, error) {
	buffer, err := getImageBuffer("https://api.dagpi.xyz/image/night/?url="+url, c)
	if err != nil {
		return nil, err
	}

	return buffer, nil
}

// Rainbow Some trippy light effects.
// Docs: https://dagpi.docs.apiary.io/#reference/images-api/rainbow/rainbow
func (c *Client) Rainbow(url string) ([]byte, error) {
	buffer, err := getImageBuffer("https://api.dagpi.xyz/image/rainbow/?url="+url, c)
	if err != nil {
		return nil, err
	}

	return buffer, nil
}

// Magik The much loved magik endpoint.
// Docs: https://dagpi.docs.apiary.io/#reference/images-api/magik/magik
func (c *Client) Magik(url string) ([]byte, error) {
	buffer, err := getImageBuffer("https://api.dagpi.xyz/image/magik/?url="+url, c)
	if err != nil {
		return nil, err
	}

	return buffer, nil
}

// FivegOneg The meme.
// Docs: https://dagpi.docs.apiary.io/#reference/images-api/five-guys-one-girl/five-guys-one-girl
func (c *Client) FivegOneg(url1 string, url2 string) ([]byte, error) {
	buffer, err := getImageBuffer("https://api.dagpi.xyz/image/5g1g/?url="+url1+"&url2="+url2, c)
	if err != nil {
		return nil, err
	}

	return buffer, nil
}

// WhyAreYouGay The meme.
// Docs: https://dagpi.docs.apiary.io/#reference/images-api/why-are-you-gay/why-are-you-gay
func (c *Client) WhyAreYouGay(url1 string, url2 string) ([]byte, error) {
	buffer, err := getImageBuffer("https://api.dagpi.xyz/image/whyareyougay/?url="+url1+"&url2="+url2, c)
	if err != nil {
		return nil, err
	}

	return buffer, nil
}

// Slap Have one image slap another.
// Docs: https://dagpi.docs.apiary.io/#reference/images-api/slap/slap
func (c *Client) Slap(url1 string, url2 string) ([]byte, error) {
	buffer, err := getImageBuffer("https://api.dagpi.xyz/image/slap/?url="+url1+"&url2="+url2, c)
	if err != nil {
		return nil, err
	}

	return buffer, nil
}

// Obama The meme.
// Docs: https://dagpi.docs.apiary.io/#reference/images-api/obama/obama
func (c *Client) Obama(url1 string, url2 string) ([]byte, error) {
	buffer, err := getImageBuffer("https://api.dagpi.xyz/image/obama/?url="+url1+"&url2="+url2, c)
	if err != nil {
		return nil, err
	}

	return buffer, nil
}

// Tweet The meme.
// Docs: https://dagpi.docs.apiary.io/#reference/images-api/tweet/tweet
func (c *Client) Tweet(url string, username string, text string) ([]byte, error) {
	buffer, err := getImageBuffer("https://api.dagpi.xyz/image/tweet/?url="+url+"&username="+username+"&text="+text, c)
	if err != nil {
		return nil, err
	}

	return buffer, nil
}

// YouTubeComment Generate realistic YouTube messages
// Docs: https://dagpi.docs.apiary.io/#reference/images-api/youtube-comment/youtube-comment
func (c *Client) YouTubeComment(url string, username string, text string, darkMode bool) ([]byte, error) {
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

// Discord Generate realistic discord messages
// Docs: https://dagpi.docs.apiary.io/#reference/images-api/discord/discord
func (c *Client) Discord(url string, username string, text string, darkMode bool) ([]byte, error) {
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

// Retromeme The good old memes. Generated.
// Docs: https://dagpi.docs.apiary.io/#reference/images-api/retromeme/retromeme
func (c *Client) Retromeme(url string, topText string, bottomText string) ([]byte, error) {
	buffer, err := getImageBuffer("https://api.dagpi.xyz/image/retromeme/?url="+url+"&top_text="+topText+"&bottom_text="+bottomText, c)
	if err != nil {
		return nil, err
	}

	return buffer, nil
}

// Motivational The black background with top and bottom motivational text.
// Docs: https://dagpi.docs.apiary.io/#reference/images-api/motivational/motivational
func (c *Client) Motivational(url string, topText string, bottomText string) ([]byte, error) {
	buffer, err := getImageBuffer("https://api.dagpi.xyz/image/motiv/?url="+url+"&top_text="+topText+"&bottom_text="+bottomText, c)
	if err != nil {
		return nil, err
	}

	return buffer, nil
}

// Modernmeme A modern meme generation system that allows reddit ready memes with just one endpoint.
// Docs: https://dagpi.docs.apiary.io/#reference/images-api/modernmeme/modernmeme
func (c *Client) Modernmeme(url string, text string) ([]byte, error) {
	buffer, err := getImageBuffer("https://api.dagpi.xyz/image/modernmeme/?url="+url+"&text="+text, c)
	if err != nil {
		return nil, err
	}

	return buffer, nil
}

// Elmo Burning Elmo Meme
// Docs: todo add docs when available
func (c *Client) Elmo(url string) ([]byte, error) {
	buffer, err := getImageBuffer("https://api.dagpi.xyz/image/elmo/?url="+url, c)
	if err != nil {
		return nil, err
	}

	return buffer, nil
}

// TvStatic Its TV static
// Docs: todo add docs when available
func (c *Client) TvStatic(url string) ([]byte, error) {
	buffer, err := getImageBuffer("https://api.dagpi.xyz/image/tv/?url="+url, c)
	if err != nil {
		return nil, err
	}

	return buffer, nil
}

// Rain Its TV static
// Docs: todo add docs when available
func (c *Client) Rain(url string) ([]byte, error) {
	buffer, err := getImageBuffer("https://api.dagpi.xyz/image/rain/?url="+url, c)
	if err != nil {
		return nil, err
	}

	return buffer, nil
}

// Glitch todo add description when available
// Docs: todo add docs when available
func (c *Client) Glitch(url string) ([]byte, error) {
	buffer, err := getImageBuffer("https://api.dagpi.xyz/image/glitch/?url="+url, c)
	if err != nil {
		return nil, err
	}

	return buffer, nil
}

// GlitchStatic todo add description when available
// Docs: todo add docs when available
func (c *Client) GlitchStatic(url string) ([]byte, error) {
	buffer, err := getImageBuffer("https://api.dagpi.xyz/image/glitchstatic/?url="+url, c)
	if err != nil {
		return nil, err
	}

	return buffer, nil
}

// Album Make an Album cover!
// Docs: todo add docs when available
func (c *Client) Album(url string) ([]byte, error) {
	buffer, err := getImageBuffer("https://api.dagpi.xyz/image/album/?url="+url, c)
	if err != nil {
		return nil, err
	}

	return buffer, nil
}

//endregion
