package urlimageview

import (
	"context"
	"errors"
	"image"
	"image/color"
	"net/http"
	"os"

	"github.com/overcyn/mochi"
	"github.com/overcyn/mochi/comm"
	"github.com/overcyn/mochi/layout"
	"github.com/overcyn/mochi/view"
	"github.com/overcyn/mochi/view/imageview"
)

// Layouter that returns the child's layout
type layouter struct{}

func (l layouter) Layout(ctx *layout.Context) (layout.Guide, map[mochi.Id]layout.Guide) {
	g := layout.Guide{Frame: layout.Rect{Max: ctx.MaxSize}}
	gs := map[mochi.Id]layout.Guide{}
	for _, id := range ctx.ChildIds {
		f := ctx.LayoutChild(id, ctx.MinSize, ctx.MaxSize)
		gs[id] = f
		g.Frame = f.Frame
	}
	return g, gs
}

func (l layouter) Notify(f func()) comm.Id {
	return 0 // no-op
}

func (l layouter) Unnotify(id comm.Id) {
	// no-op
}

type View struct {
	*view.Embed
	ResizeMode imageview.ResizeMode
	URL        string
	Path       string
	Tint       color.Color
	stage      view.Stage
	// Image request
	url        string
	path       string
	cancelFunc context.CancelFunc
	image      image.Image
	err        error
}

func New(ctx *view.Context, key interface{}) *View {
	if v, ok := ctx.Prev(key).(*View); ok {
		return v
	}
	return &View{
		Embed: view.NewEmbed(ctx.NewId(key)),
	}
}

func (v *View) Build(ctx *view.Context) *view.Model {
	v.reload()

	chl := imageview.New(ctx, 0)
	chl.ResizeMode = v.ResizeMode
	chl.Image = v.image
	chl.Tint = v.Tint

	return &view.Model{
		Layouter: layouter{},
		Children: []view.View{chl},
	}
}

func (v *View) Lifecycle(from, to view.Stage) {
	v.stage = to
	v.reload()
}

func (v *View) reload() {
	if v.stage < view.StageMounted {
		v.cancel()
		return
	}

	if v.URL != v.url || v.Path != v.path || v.cancelFunc == nil {
		v.cancel()

		c, cancelFunc := context.WithCancel(context.Background())
		v.url = v.URL
		v.path = v.Path
		v.cancelFunc = cancelFunc
		v.image = nil
		v.err = nil
		go func(url, path string) {
			image, err := loadImageURL(url, path)

			view.MainMu.Lock()
			defer view.MainMu.Unlock()

			select {
			case <-c.Done():
			default:
				v.cancelFunc()
				v.image = image
				v.err = err
				v.Update()
			}
		}(v.url, v.path)
	}
}

func (v *View) cancel() {
	if v.cancelFunc != nil {
		v.cancelFunc()
		v.cancelFunc = nil
	}
}

func loadImageURL(url, path string) (image.Image, error) {
	if len(url) > 0 {
		resp, err := http.Get(url)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()
		img, _, err := image.Decode(resp.Body)
		return img, err
	} else if len(path) > 0 {
		file, err := os.Open(path)
		if err != nil {
			return nil, err
		}
		defer file.Close()
		img, _, err := image.Decode(file)
		return img, err
	}
	return nil, errors.New("URLImageView.loadImageURL: No url or path")
}
