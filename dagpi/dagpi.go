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

var client = Client{}

// request to get data
func httpGet(url string) (map[string]interface{}, error) {
	httpClient := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", client.Auth)
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
func getImageBuffer(url string) ([]byte, error) {
	httpClient := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", client.Auth)
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

//region Data API calls

// WTP returns an interface with all the Pokemon data
// Docs: https://dagpi.docs.apiary.io/#reference/data-api/whos-that-pokemon/who's-that-pokemon?
func WTP() (interface{}, error) {
	data, err := httpGet("https://api.dagpi.xyz/data/wtp")
	if err != nil {
		return nil, err
	}

	return data, nil
}

// Roast returns an interface containing a roast
// Docs: https://dagpi.docs.apiary.io/#reference/data-api/roast/roast
func Roast() (interface{}, error) {
	data, err := httpGet("https://api.dagpi.xyz/data/roast")
	roast := data["roast"]
	if err != nil {
		return nil, err
	}

	return roast, nil
}

// Joke returns an interface containing a joke & id
// Docs: https://dagpi.docs.apiary.io/#reference/data-api/joke/joke
func Joke() (interface{}, error) {
	data, err := httpGet("https://api.dagpi.xyz/data/joke")
	if err != nil {
		return nil, err
	}

	return data, nil
}

// Fact returns an interface containing a fact
// Docs: https://dagpi.docs.apiary.io/#reference/data-api/fact/fact
func Fact() (interface{}, error) {
	data, err := httpGet("https://api.dagpi.xyz/data/fact")
	fact := data["fact"]
	if err != nil {
		return nil, err
	}

	return fact, nil
}

// Eightball returns an interface containing a response to 8ball question
// Docs: https://dagpi.docs.apiary.io/#reference/data-api/8ball/8ball
func Eightball() (interface{}, error) {
	data, err := httpGet("https://api.dagpi.xyz/data/8ball")
	response := data["response"]
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Yomama returns an interface containing a description of yomama
// Docs: https://dagpi.docs.apiary.io/#reference/data-api/yomama/yomama
func Yomama() (interface{}, error) {
	data, err := httpGet("https://api.dagpi.xyz/data/yomama")
	description := data["description"]
	if err != nil {
		return nil, err
	}

	return description, nil
}

// RandomWaifu returns an interface containing data of a random waifu
// Docs: https://dagpi.docs.apiary.io/#reference/data-api/random-waifu/random-waifu
func RandomWaifu() (interface{}, error) {
	data, err := httpGet("https://api.dagpi.xyz/data/waifu")
	if err != nil {
		return nil, err
	}

	return data, nil
}

// Waifu returns an interface containing data of a given waifu
// Docs: https://dagpi.docs.apiary.io/#reference/data-api/waifu-saerch/waifu-saerch
func Waifu(waifuName string) (interface{}, error) {
	data, err := httpGet("https://api.dagpi.xyz/data/" + waifuName)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// PickupLine returns an interface containing category & joke
// Docs: https://dagpi.docs.apiary.io/#reference/data-api/pickup-line/pickup-line
func PickupLine() (interface{}, error) {
	data, err := httpGet("https://api.dagpi.xyz/data/pickupline")
	if err != nil {
		return nil, err
	}

	return data, nil
}

// HeadLine returns an interface containing text and a bool, 'fake'
// Docs: https://dagpi.docs.apiary.io/#reference/data-api/headline/headline
func HeadLine() (interface{}, error) {
	data, err := httpGet("https://api.dagpi.xyz/data/headline")
	if err != nil {
		return nil, err
	}

	return data, nil
}

// GTL returns an interface containing data of a random logo (Guess the Logo)
// Docs: https://dagpi.docs.apiary.io/#reference/data-api/guess-the-logo/guess-the-logo
func GTL() (interface{}, error) {
	data, err := httpGet("https://api.dagpi.xyz/data/logo")
	if err != nil {
		return nil, err
	}

	return data, nil
}

// Flag returns an interface containing data of a random flag
// Docs: https://dagpi.docs.apiary.io/#reference/data-api/flag/flag
func Flag() (interface{}, error) {
	data, err := httpGet("https://api.dagpi.xyz/data/flag")
	if err != nil {
		return nil, err
	}

	return data, nil
}

// Captcha get a random captcha and answer
// Docs: https://dagpi.docs.apiary.io/#reference/data-api/captcha/captcha
func Captcha() (interface{}, error) {
	data, err := httpGet("https://api.dagpi.xyz/data/captcha")
	if err != nil {
		return nil, err
	}

	return data, nil
}

// Typeracer get a sentence on an image, with a sentence to create typeracer games
// Docs: https://dagpi.docs.apiary.io/#reference/data-api/typeracer/typeracer
func Typeracer() (interface{}, error) {
	data, err := httpGet("https://api.dagpi.xyz/data/typeracer")
	if err != nil {
		return nil, err
	}

	return data, nil
}

//endregion

//region Image Manipulation API calls

// Pixelate Allows you to pixelate an image.
// Docs: https://dagpi.docs.apiary.io/#reference/images-api/pixel/pixel
func Pixelate(url string) ([]byte, error) {
	buffer, err := getImageBuffer("https://api.dagpi.xyz/image/pixel/?url=" + url)
	if err != nil {
		return nil, err
	}

	return buffer, nil
}

// Mirror an image along the y-axis
// Docs: https://dagpi.docs.apiary.io/#reference/images-api/mirror/mirror
func Mirror(url string) ([]byte, error) {
	buffer, err := getImageBuffer("https://api.dagpi.xyz/image/mirror/?url=" + url)
	if err != nil {
		return nil, err
	}

	return buffer, nil
}

// FlipImage flip an image
// Docs: https://dagpi.docs.apiary.io/#reference/images-api/flip/flip
func FlipImage(url string) ([]byte, error) {
	buffer, err := getImageBuffer("https://api.dagpi.xyz/image/flip/?url=" + url)
	if err != nil {
		return nil, err
	}

	return buffer, nil
}

// Colors Allows you to get an Image with the colors present in the image.
// Docs: https://dagpi.docs.apiary.io/#reference/images-api/colors/colors
func Colors(url string) ([]byte, error) {
	buffer, err := getImageBuffer("https://api.dagpi.xyz/image/colors/?url=" + url)
	if err != nil {
		return nil, err
	}

	return buffer, nil
}

// America Let the star spangled banner of the free and the brave soar.
// Docs:  https://dagpi.docs.apiary.io/#reference/images-api/america/america
func America(url string) ([]byte, error) {
	buffer, err := getImageBuffer("https://api.dagpi.xyz/image/america/?url=" + url)
	if err != nil {
		return nil, err
	}

	return buffer, nil
}

// Communism Support the soviet union comrade. Let the red flag fly!
// Docs: https://dagpi.docs.apiary.io/#reference/images-api/communism/communism
func Communism(url string) ([]byte, error) {
	buffer, err := getImageBuffer("https://api.dagpi.xyz/image/communism/?url=" + url)
	if err != nil {
		return nil, err
	}

	return buffer, nil
}

// Triggered Allows you to get a triggered gif.
// Docs: https://dagpi.docs.apiary.io/#reference/images-api/triggered/triggered
func Triggered(url string) ([]byte, error) {
	buffer, err := getImageBuffer("https://api.dagpi.xyz/image/triggered/?url=" + url)
	if err != nil {
		return nil, err
	}

	return buffer, nil
}

// ExpandImage animation that streches an image.
// Docs: https://dagpi.docs.apiary.io/#reference/images-api/expand/expand
func ExpandImage(url string) ([]byte, error) {
	buffer, err := getImageBuffer("https://api.dagpi.xyz/image/expand/?url=" + url)
	if err != nil {
		return nil, err
	}

	return buffer, nil
}

// Wasted Allows you to get an image with GTA V Wasted screen.
// Docs: https://dagpi.docs.apiary.io/#reference/images-api/wasted/wasted
func Wasted(url string) ([]byte, error) {
	buffer, err := getImageBuffer("https://api.dagpi.xyz/image/wasted/?url=" + url)
	if err != nil {
		return nil, err
	}

	return buffer, nil
}

// Sketch Cool efffect that shows how an image would have been created by an artist.
// Docs: https://dagpi.docs.apiary.io/#reference/images-api/sketch/sketch
func Sketch(url string) ([]byte, error) {
	buffer, err := getImageBuffer("https://api.dagpi.xyz/image/sketch/?url=" + url)
	if err != nil {
		return nil, err
	}

	return buffer, nil
}

// SpinImage You spin me right round baby.
// Docs: https://dagpi.docs.apiary.io/#reference/images-api/spin/spin
func SpinImage(url string) ([]byte, error) {
	buffer, err := getImageBuffer("https://api.dagpi.xyz/image/spin/?url=" + url)
	if err != nil {
		return nil, err
	}

	return buffer, nil
}

// PetPet Pet pet gif
// Docs: https://dagpi.docs.apiary.io/#reference/images-api/petpet/petpet
func PetPet(url string) ([]byte, error) {
	buffer, err := getImageBuffer("https://api.dagpi.xyz/image/petpet/?url=" + url)
	if err != nil {
		return nil, err
	}

	return buffer, nil
}

// Bonk Get bonked on my cheems
// Docs: https://dagpi.docs.apiary.io/#reference/images-api/bonk/bonk
func Bonk(url string) ([]byte, error) {
	buffer, err := getImageBuffer("https://api.dagpi.xyz/image/bonk/?url=" + url)
	if err != nil {
		return nil, err
	}

	return buffer, nil
}

// Bomb Explosion
// Docs: https://dagpi.docs.apiary.io/#reference/images-api/bomb/bomb
func Bomb(url string) ([]byte, error) {
	buffer, err := getImageBuffer("https://api.dagpi.xyz/image/bomb/?url=" + url)
	if err != nil {
		return nil, err
	}

	return buffer, nil
}

// Shake a gif by having it wiggle.
// Docs: https://dagpi.docs.apiary.io/#reference/images-api/shake/shake
func Shake(url string) ([]byte, error) {
	buffer, err := getImageBuffer("https://api.dagpi.xyz/image/shake/?url=" + url)
	if err != nil {
		return nil, err
	}

	return buffer, nil
}

// Invert Allows you to get an image with an inverted color effect.
// Docs: https://dagpi.docs.apiary.io/#reference/images-api/invert/invert
func Invert(url string) ([]byte, error) {
	buffer, err := getImageBuffer("https://api.dagpi.xyz/image/invert/?url=" + url)
	if err != nil {
		return nil, err
	}

	return buffer, nil
}

// Sobel Allows you to get an image with the sobel effect.
// Docs: https://dagpi.docs.apiary.io/#reference/images-api/sobel/sobel
func Sobel(url string) ([]byte, error) {
	buffer, err := getImageBuffer("https://api.dagpi.xyz/image/sobel/?url=" + url)
	if err != nil {
		return nil, err
	}

	return buffer, nil
}

// Hog Histogram of Oriented Gradients for an image.
// Docs: https://dagpi.docs.apiary.io/#reference/images-api/hog/hog
func Hog(url string) ([]byte, error) {
	buffer, err := getImageBuffer("https://api.dagpi.xyz/image/hog/?url=" + url)
	if err != nil {
		return nil, err
	}

	return buffer, nil
}

// Triangle Cool triangle effect for an image.
// Docs: https://dagpi.docs.apiary.io/#reference/images-api/triangle/triangle
func Triangle(url string) ([]byte, error) {
	buffer, err := getImageBuffer("https://api.dagpi.xyz/image/triangle/?url=" + url)
	if err != nil {
		return nil, err
	}

	return buffer, nil
}

// Blur Blurs a given image.
// Docs: https://dagpi.docs.apiary.io/#reference/images-api/blur/blur
func Blur(url string) ([]byte, error) {
	buffer, err := getImageBuffer("https://api.dagpi.xyz/image/blur/?url=" + url)
	if err != nil {
		return nil, err
	}

	return buffer, nil
}

// RGB Get an RGB graph of an image's colors.
// Docs: https://dagpi.docs.apiary.io/#reference/images-api/rgb/rgb
func RGB(url string) ([]byte, error) {
	buffer, err := getImageBuffer("https://api.dagpi.xyz/image/rgb/?url=" + url)
	if err != nil {
		return nil, err
	}

	return buffer, nil
}

// Angel Image on the Angels face.
// Docs: https://dagpi.docs.apiary.io/#reference/images-api/angel/angel
func Angel(url string) ([]byte, error) {
	buffer, err := getImageBuffer("https://api.dagpi.xyz/image/angel/?url=" + url)
	if err != nil {
		return nil, err
	}

	return buffer, nil
}

// Satan Put an image on the devil.
// Docs: https://dagpi.docs.apiary.io/#reference/images-api/satan/satan
func Satan(url string) ([]byte, error) {
	buffer, err := getImageBuffer("https://api.dagpi.xyz/image/satan/?url=" + url)
	if err != nil {
		return nil, err
	}

	return buffer, nil
}

// Delete Generates a windows error meme based on a given image.
// Docs: https://dagpi.docs.apiary.io/#reference/images-api/delete/delete
func Delete(url string) ([]byte, error) {
	buffer, err := getImageBuffer("https://api.dagpi.xyz/image/delete/?url=" + url)
	if err != nil {
		return nil, err
	}

	return buffer, nil
}

// Fedora Tips fedora in appreciation. Perry the Platypus.
// Docs: https://dagpi.docs.apiary.io/#reference/images-api/fedora/fedora
func Fedora(url string) ([]byte, error) {
	buffer, err := getImageBuffer("https://api.dagpi.xyz/image/fedora/?url=" + url)
	if err != nil {
		return nil, err
	}

	return buffer, nil
}

// Hitler ?????
// Docs: https://dagpi.docs.apiary.io/#reference/images-api/hitler/hitler
func Hitler(url string) ([]byte, error) {
	buffer, err := getImageBuffer("https://api.dagpi.xyz/image/hitler/?url=" + url)
	if err != nil {
		return nil, err
	}

	return buffer, nil
}

// Lego Every group of pixels is a lego brick
// Docs: https://dagpi.docs.apiary.io/#reference/images-api/lego/lego
func Lego(url string) ([]byte, error) {
	buffer, err := getImageBuffer("https://api.dagpi.xyz/image/lego/?url=" + url)
	if err != nil {
		return nil, err
	}

	return buffer, nil
}

// Wanted poster of an image.
// Docs: https://dagpi.docs.apiary.io/#reference/images-api/wanted/wanted
func Wanted(url string) ([]byte, error) {
	buffer, err := getImageBuffer("https://api.dagpi.xyz/image/wanted/?url=" + url)
	if err != nil {
		return nil, err
	}

	return buffer, nil
}

// Stringify Turn your image into a ball of yarn.
// Docs: https://dagpi.docs.apiary.io/#reference/images-api/stringify/stringify
func Stringify(url string) ([]byte, error) {
	buffer, err := getImageBuffer("https://api.dagpi.xyz/image/stringify/?url=" + url)
	if err != nil {
		return nil, err
	}

	return buffer, nil
}

// Burn Light your image on fire
// Docs: https://dagpi.docs.apiary.io/#reference/images-api/burn/burn
func Burn(url string) ([]byte, error) {
	buffer, err := getImageBuffer("https://api.dagpi.xyz/image/burn/?url=" + url)
	if err != nil {
		return nil, err
	}

	return buffer, nil
}

// Earth The green and blue of the earth
// Docs: https://dagpi.docs.apiary.io/#reference/images-api/earth/earth
func Earth(url string) ([]byte, error) {
	buffer, err := getImageBuffer("https://api.dagpi.xyz/image/earth/?url=" + url)
	if err != nil {
		return nil, err
	}

	return buffer, nil
}

// Freeze Blue ice like tint.
// Docs: https://dagpi.docs.apiary.io/#reference/images-api/freeze/freeze
func Freeze(url string) ([]byte, error) {
	buffer, err := getImageBuffer("https://api.dagpi.xyz/image/freeze/?url=" + url)
	if err != nil {
		return nil, err
	}

	return buffer, nil
}

// Ground The poower of the earth
// Docs: https://dagpi.docs.apiary.io/#reference/images-api/earth/earth
func Ground(url string) ([]byte, error) {
	buffer, err := getImageBuffer("https://api.dagpi.xyz/image/ground/?url=" + url)
	if err != nil {
		return nil, err
	}

	return buffer, nil
}

// Mosiac Turn an image into a roman mosiac.
// Docs: https://dagpi.docs.apiary.io/#reference/images-api/mosiac/mosiac
func Mosiac(url string) ([]byte, error) {
	buffer, err := getImageBuffer("https://api.dagpi.xyz/image/mosiac/?url=" + url)
	if err != nil {
		return nil, err
	}

	return buffer, nil
}

// Sithlord Put an image on the Laughs in Sithlord meme.
// Docs: https://dagpi.docs.apiary.io/#reference/images-api/sithlord/sithlord
func Sithlord(url string) ([]byte, error) {
	buffer, err := getImageBuffer("https://api.dagpi.xyz/image/sith/?url=" + url)
	if err != nil {
		return nil, err
	}

	return buffer, nil
}

// Jail Put an image behind bars.
// Docs: https://dagpi.docs.apiary.io/#reference/images-api/jail/jail
func Jail(url string) ([]byte, error) {
	buffer, err := getImageBuffer("https://api.dagpi.xyz/image/jail/?url=" + url)
	if err != nil {
		return nil, err
	}

	return buffer, nil
}

// Shatter Put an image behind bars.
// Docs: https://dagpi.docs.apiary.io/#reference/images-api/shatter/shatter
func Shatter(url string) ([]byte, error) {
	buffer, err := getImageBuffer("https://api.dagpi.xyz/image/shatter/?url=" + url)
	if err != nil {
		return nil, err
	}

	return buffer, nil
}

// Pride Flag of your choice over an Image!
// Available Choices: Asexual, Bisexual, Gay, Genderfluid, Genderqueer, Intersex, Lesbian, Nonbinary, Progress, Pan, Trans
// Docs: https://dagpi.docs.apiary.io/#reference/images-api/pride/pride
func Pride(url string, flag string) ([]byte, error) {
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
			imgBuffer, err := getImageBuffer("https://api.dagpi.xyz/image/pride/?url=" + url + "&flag=" + flag)
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

// Trash Image is trash.
// Docs: https://dagpi.docs.apiary.io/#reference/images-api/trash/trash
func Trash(url string) ([]byte, error) {
	buffer, err := getImageBuffer("https://api.dagpi.xyz/image/trash/?url=" + url)
	if err != nil {
		return nil, err
	}

	return buffer, nil
}

// Deepfry an image.
// Docs: https://dagpi.docs.apiary.io/#reference/images-api/deepfry/deepfry
func Deepfry(url string) ([]byte, error) {
	buffer, err := getImageBuffer("https://api.dagpi.xyz/image/deepfry/?url=" + url)
	if err != nil {
		return nil, err
	}

	return buffer, nil
}

// Ascii Cool hackerman effect for an image.
// Docs: https://dagpi.docs.apiary.io/#reference/images-api/ascii/ascii
func Ascii(url string) ([]byte, error) {
	buffer, err := getImageBuffer("https://api.dagpi.xyz/image/ascii/?url=" + url)
	if err != nil {
		return nil, err
	}

	return buffer, nil
}

// Charcoal Image into a charcoal drawing.
// Docs: https://dagpi.docs.apiary.io/#reference/images-api/charcoal/charcoal
func Charcoal(url string) ([]byte, error) {
	buffer, err := getImageBuffer("https://api.dagpi.xyz/image/charcoal/?url=" + url)
	if err != nil {
		return nil, err
	}

	return buffer, nil
}

// Posterize Posterizes an image.
// Docs: https://dagpi.docs.apiary.io/#reference/images-api/posterize/posterize
func Posterize(url string) ([]byte, error) {
	buffer, err := getImageBuffer("https://api.dagpi.xyz/image/poster/?url=" + url)
	if err != nil {
		return nil, err
	}

	return buffer, nil
}

// Sepia Tone an image.
// Docs: https://dagpi.docs.apiary.io/#reference/images-api/sepia/sepia
func Sepia(url string) ([]byte, error) {
	buffer, err := getImageBuffer("https://api.dagpi.xyz/image/sepia/?url=" + url)
	if err != nil {
		return nil, err
	}

	return buffer, nil
}

// Swirl an image.
// Docs: https://dagpi.docs.apiary.io/#reference/images-api/swirl/swirl
func Swirl(url string) ([]byte, error) {
	buffer, err := getImageBuffer("https://api.dagpi.xyz/image/swirl/?url=" + url)
	if err != nil {
		return nil, err
	}

	return buffer, nil
}

// Paint Turn an image into art.
// Docs: https://dagpi.docs.apiary.io/#reference/images-api/paint/paint
func Paint(url string) ([]byte, error) {
	buffer, err := getImageBuffer("https://api.dagpi.xyz/image/paint/?url=" + url)
	if err != nil {
		return nil, err
	}

	return buffer, nil
}

// Night Turn an day into night.
// Docs: https://dagpi.docs.apiary.io/#reference/images-api/night/night
func Night(url string) ([]byte, error) {
	buffer, err := getImageBuffer("https://api.dagpi.xyz/image/night/?url=" + url)
	if err != nil {
		return nil, err
	}

	return buffer, nil
}

// Rainbow Some trippy light effects.
// Docs: https://dagpi.docs.apiary.io/#reference/images-api/rainbow/rainbow
func Rainbow(url string) ([]byte, error) {
	buffer, err := getImageBuffer("https://api.dagpi.xyz/image/rainbow/?url=" + url)
	if err != nil {
		return nil, err
	}

	return buffer, nil
}

// Magik The much loved magik endpoint.
// Docs: https://dagpi.docs.apiary.io/#reference/images-api/magik/magik
func Magik(url string) ([]byte, error) {
	buffer, err := getImageBuffer("https://api.dagpi.xyz/image/magik/?url=" + url)
	if err != nil {
		return nil, err
	}

	return buffer, nil
}

// FivegOneg The meme.
// Docs: https://dagpi.docs.apiary.io/#reference/images-api/five-guys-one-girl/five-guys-one-girl
func FivegOneg(url1 string, url2 string) ([]byte, error) {
	buffer, err := getImageBuffer("https://api.dagpi.xyz/image/5g1g/?url=" + url1 + "&url2=" + url2)
	if err != nil {
		return nil, err
	}

	return buffer, nil
}

// WhyAreYouGay The meme.
// Docs: https://dagpi.docs.apiary.io/#reference/images-api/why-are-you-gay/why-are-you-gay
func WhyAreYouGay(url1 string, url2 string) ([]byte, error) {
	buffer, err := getImageBuffer("https://api.dagpi.xyz/image/whyareyougay/?url=" + url1 + "&url2=" + url2)
	if err != nil {
		return nil, err
	}

	return buffer, nil
}

// Slap Have one image slap another.
// Docs: https://dagpi.docs.apiary.io/#reference/images-api/slap/slap
func Slap(url1 string, url2 string) ([]byte, error) {
	buffer, err := getImageBuffer("https://api.dagpi.xyz/image/slap/?url=" + url1 + "&url2=" + url2)
	if err != nil {
		return nil, err
	}

	return buffer, nil
}

// Obama The meme.
// Docs: https://dagpi.docs.apiary.io/#reference/images-api/obama/obama
func Obama(url1 string, url2 string) ([]byte, error) {
	buffer, err := getImageBuffer("https://api.dagpi.xyz/image/obama/?url=" + url1 + "&url2=" + url2)
	if err != nil {
		return nil, err
	}

	return buffer, nil
}

// Tweet The meme.
// Docs: https://dagpi.docs.apiary.io/#reference/images-api/tweet/tweet
func Tweet(url string, username string, text string) ([]byte, error) {
	buffer, err := getImageBuffer("https://api.dagpi.xyz/image/tweet/?url=" + url + "&username=" + username + "&text=" + text)
	if err != nil {
		return nil, err
	}

	return buffer, nil
}

// YouTubeComment Generate realistic Youtube messages
// Docs: https://dagpi.docs.apiary.io/#reference/images-api/youtube-comment/youtube-comment
func YouTubeComment(url string, username string, text string, darkMode bool) ([]byte, error) {
	if darkMode == true {
		buffer, err := getImageBuffer("https://api.dagpi.xyz/image/yt/?url=" + url + "&username=" + username + "&text=" + text + "&dark=" + "true")
		if err != nil {
			return nil, err
		}

		return buffer, nil
	} else {
		buffer, err := getImageBuffer("https://api.dagpi.xyz/image/yt/?url=" + url + "&username=" + username + "&text=" + text + "&dark=" + "false")
		if err != nil {
			return nil, err
		}

		return buffer, nil
	}
}

// Discord Generate realistic discord messages
// Docs: https://dagpi.docs.apiary.io/#reference/images-api/discord/discord
func Discord(url string, username string, text string, darkMode bool) ([]byte, error) {
	if darkMode == true {
		buffer, err := getImageBuffer("https://api.dagpi.xyz/image/discord/?url=" + url + "&username=" + username + "&text=" + text + "&dark=" + "true")
		if err != nil {
			return nil, err
		}

		return buffer, nil
	} else {
		buffer, err := getImageBuffer("https://api.dagpi.xyz/image/discord/?url=" + url + "&username=" + username + "&text=" + text + "&dark=" + "false")
		if err != nil {
			return nil, err
		}

		return buffer, nil
	}
}

// Retromeme The good old memes. Generated.
// Docs: https://dagpi.docs.apiary.io/#reference/images-api/retromeme/retromeme
func Retromeme(url string, topText string, bottomText string) ([]byte, error) {
	buffer, err := getImageBuffer("https://api.dagpi.xyz/image/retromeme/?url=" + url + "&top_text=" + topText + "&bottom_text=" + bottomText)
	if err != nil {
		return nil, err
	}

	return buffer, nil
}

// Motivational The black background with top and bottom motivational text.
// Docs: https://dagpi.docs.apiary.io/#reference/images-api/motivational/motivational
func Motivational(url string, topText string, bottomText string) ([]byte, error) {
	buffer, err := getImageBuffer("https://api.dagpi.xyz/image/motiv/?url=" + url + "&top_text=" + topText + "&bottom_text=" + bottomText)
	if err != nil {
		return nil, err
	}

	return buffer, nil
}

// Modernmeme A modern meme generation system that allows reddit ready memes with just one endpoint.
// Docs: https://dagpi.docs.apiary.io/#reference/images-api/modernmeme/modernmeme
func Modernmeme(url string, text string) ([]byte, error) {
	buffer, err := getImageBuffer("https://api.dagpi.xyz/image/modernmeme/?url=" + url + "&text=" + text)
	if err != nil {
		return nil, err
	}

	return buffer, nil
}

//endregion
