# godagpi
Go wrapper for Dagpi API. All functions should have a small desciption even its vague (copy pasta from the Docs) and a link to the call in the documentation.

Initialize your client with

```
var client = dagpi.Client{Auth: "api token"}
```

Api Documentation can be found [here](https://dagpi.docs.apiary.io/).

<h2>Example</h2>

```
test test
```
---

## Functions - Data | Returns Interface of Data

* dagpi.WTP / Who's That Pokemon
* dagpi.Roast
* dagpi.Joke
* dagpi.Fact
* dagpi.Eightball
* dagpi.Yomama
* dagpi.RandomWaifu
* dagpi.Waifu
* dagpi.PickupLine
* dagpi.HeadLine
* dagpi.GTL / Guess The Logo
* dagpi.Flag
* dagpi.Captcha
* dogpi.Typeracer

## Functions - Image Manip | Returns Image Buffer/Bytes []byte

***Intintionally slipped calls:***
```
- Gay (Included in Pride Call)
- Captcha (Captcha call in Data)
- Disolve (problems with the call)
```

* dagpi.Pixelate(imageUrl: https://imghost.com/img) Do the same for all image manip funcs.
* dagpi.Mirror
* dagpi.FlipImage
* dagpi.Colors
* dagpi.America
* dagpi.Communism
* dagpi.Triggered
* dagpi.ExpandImage
* dagpi.Wasted
* dagpi.Sketch
* dagpi.Spin
* dagpi.PetPet
* dagpi.Bonk
* dagpi.Bomb
* dagpi.Shake
* dagpi.Invert
* dagpi.Sobel
* dagpi.Hog
* dagpi.Triangle
* dagpi.Blur
* dagpi.RGB
* dagpi.Angel
* dagpi.Satan
* dagpi.Delete
* dagpi.Fedora
* dagpi.Hitler
* dagpi.Lego
* dagpi.Wanted
* dagpi.Stringify
* dagppi.Burn
* dagpi.Earth
* dagpi.Freeze
* dagpi.Ground
* dagpi.Mosiac
* dagpi.Sithlord
* dagpi.Jail
* dagpi.Shatter
* dagpi.Pride(imageUrl, flag) acceptable Flags are listed in function
* dagpi.Trash
* dagpi.Deepfry
* dagpi.Ascii
* dagpi.Charcoal
* dagpi.Posterize
* dagpi.Sepia
* dagpi.Swirl
* dagpi.Paint
* dagpi.Night
* dagpi.Rainbow
* dagpi.Magik
* dagpi.FivegOneg(imageUrl1, imageUrl2) / Five Guys One Girl
* dagpi.WhyAreYouGay(imageUrl1, imageUrl2)
* dagpi.Slap(imageUrl1, imageUrl2)
* dagpi.Oboma(imageUrl1, imageUrl2)
* dagpi.Tweet(imageUrl, username, text)
* dagpi.Youtube(imageUrl, username, text, darkMode: boolean)
* dagpi.Discord(imageUrl, username, text, darkMode: boolean)
* dagpi.Retromeme(imageUrl, topText, bottomText)
* dagpi.Motivational(imageUrl, topText, bottomText)
* dagpi.Modernmeme(imageUrl, text)