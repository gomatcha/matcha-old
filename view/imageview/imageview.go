package imageview

import (
	"bytes"
	"context"
	"fmt"
	"github.com/overcyn/mochi"
	"golang.org/x/image/bmp"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"net/http"
)

const (
	urlImageViewId int = iota
)

type URLImageView struct {
	*mochi.Embed
	PaintOptions mochi.PaintOptions
	ResizeMode   ResizeMode
	URL          string
	// Image request
	url    string
	cancel context.CancelFunc
	image  image.Image
	err    error
}

func NewURLImageView(c mochi.Config) *URLImageView {
	v, ok := c.Prev.(*URLImageView)
	if !ok {
		v = &URLImageView{}
		v.Embed = c.Embed
	}
	return v
}

func (v *URLImageView) Build(ctx *mochi.BuildContext) *mochi.Node {
	if v.URL != v.url {
		if v.cancel != nil {
			v.cancel()
		}

		c, cancel := context.WithCancel(context.Background())
		v.url = v.URL
		v.cancel = cancel
		v.image = nil
		v.err = nil
		go func(url string) {
			image, err := loadImageURL(url)

			v.Lock()
			defer v.Unlock()

			select {
			case <-c.Done():
			default:
				v.cancel()
				v.image = image
				v.err = err
				v.Update(nil)
			}
		}(v.url)
	}

	n := &mochi.Node{}
	n.Painter = v.PaintOptions

	chl := NewImageView(ctx.Get(urlImageViewId))
	chl.ResizeMode = v.ResizeMode
	chl.Image = v.image
	chl.PaintOptions.BackgroundColor = mochi.RedColor
	n.Add(chl)

	return n
}

func (v *URLImageView) String() string {
	return fmt.Sprintf("&URLImageView{%p URL:%v}", v, v.URL)
}

func loadImageURL(url string) (image.Image, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	img, _, err := image.Decode(resp.Body)
	return img, err
}

// ImageView

type ResizeMode int

const (
	ResizeModeFit ResizeMode = iota
	ResizeModeFill
	ResizeModeStretch
	ResizeModeCenter
)

type ImageView struct {
	*mochi.Embed
	PaintOptions mochi.PaintOptions
	Image        image.Image
	ResizeMode   ResizeMode
	image        image.Image
	bytes        []byte
}

func NewImageView(c mochi.Config) *ImageView {
	v, ok := c.Prev.(*ImageView)
	if !ok {
		v = &ImageView{}
		v.Embed = c.Embed
	}
	return v
}

func (v *ImageView) Build(ctx *mochi.BuildContext) *mochi.Node {
	if v.Image != v.image {
		v.image = v.Image

		buf := &bytes.Buffer{}
		err := bmp.Encode(buf, v.image)
		if err != nil {
			fmt.Println("ImageView encoding error:", err)
		}
		v.bytes = buf.Bytes()
	}

	n := &mochi.Node{}
	n.Painter = v.PaintOptions
	n.Bridge.Name = "github.com/overcyn/mochi/view/imageview"
	n.Bridge.State = struct {
		Bytes      []byte
		ResizeMode ResizeMode
	}{
		Bytes:      v.bytes,
		ResizeMode: v.ResizeMode,
	}
	return n
}

func (v *ImageView) String() string {
	return fmt.Sprintf("&ImageView{%p}", v)
}
