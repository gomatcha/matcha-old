package imageview

import (
	"bytes"
	"context"
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"net/http"

	"github.com/overcyn/mochi/paint"
	"github.com/overcyn/mochi/view"
	"golang.org/x/image/bmp"
)

type URLImageView struct {
	*view.Embed
	Painter    paint.Painter
	ResizeMode ResizeMode
	URL        string
	stage      view.Stage
	// Image request
	url        string
	cancelFunc context.CancelFunc
	image      image.Image
	err        error
}

func NewURLImageView(ctx *view.Context, key interface{}) *URLImageView {
	v, ok := ctx.Prev(key).(*URLImageView)
	if !ok {
		v = &URLImageView{
			Embed: view.NewEmbed(ctx.NewId(key)),
		}
	}
	return v
}

func (v *URLImageView) Build(ctx *view.Context) *view.Model {
	v.reload()

	n := &view.Model{}
	n.Painter = v.Painter

	chl := NewImageView(ctx, 0)
	chl.ResizeMode = v.ResizeMode
	chl.Image = v.image
	n.Add(chl)

	return n
}

func (v *URLImageView) Lifecycle(from, to view.Stage) {
	v.stage = to
	v.reload()
}

func (v *URLImageView) reload() {
	if v.stage < view.StageMounted {
		v.cancel()
		return
	}

	if v.URL != v.url || v.cancelFunc == nil {
		v.cancel()

		c, cancelFunc := context.WithCancel(context.Background())
		v.url = v.URL
		v.cancelFunc = cancelFunc
		v.image = nil
		v.err = nil
		go func(url string) {
			image, err := loadImageURL(url)

			v.Lock()
			defer v.Unlock()

			select {
			case <-c.Done():
			default:
				v.cancelFunc()
				v.image = image
				v.err = err
				v.Update(nil)
			}
		}(v.url)
	}
}

func (v *URLImageView) cancel() {
	if v.cancelFunc != nil {
		v.cancelFunc()
		v.cancelFunc = nil
	}
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
	*view.Embed
	Painter    paint.Painter
	Image      image.Image
	ResizeMode ResizeMode
	image      image.Image
	bytes      []byte
}

func NewImageView(ctx *view.Context, key interface{}) *ImageView {
	v, ok := ctx.Prev(key).(*ImageView)
	if !ok {
		v = &ImageView{
			Embed: view.NewEmbed(ctx.NewId(key)),
		}
	}
	return v
}

func (v *ImageView) Build(ctx *view.Context) *view.Model {
	if v.Image != v.image {
		v.image = v.Image

		buf := &bytes.Buffer{}
		err := bmp.Encode(buf, v.image)
		if err != nil {
			fmt.Println("ImageView encoding error:", err)
		}
		v.bytes = buf.Bytes()
	}

	n := &view.Model{
		Painter:    v.Painter,
		BridgeName: "github.com/overcyn/mochi/view/imageview",
		BridgeState: struct {
			Bytes      []byte
			ResizeMode ResizeMode
		}{
			Bytes:      v.bytes,
			ResizeMode: v.ResizeMode,
		},
	}
	return n
}

func (v *ImageView) String() string {
	return fmt.Sprintf("&ImageView{%p}", v)
}
