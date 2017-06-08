package imageview

import (
	"fmt"
	"image"
	_ "image/png"
	"os"

	"github.com/overcyn/mochi/env"
	"github.com/overcyn/mochi/layout/constraint"
	"github.com/overcyn/mochi/paint"
	"github.com/overcyn/mochi/view"
	"github.com/overcyn/mochi/view/imageview"
	"github.com/overcyn/mochibridge"
	"golang.org/x/image/colornames"
)

func init() {
	mochibridge.RegisterFunc("github.com/overcyn/mochi/examples/imageview New", func() *view.Root {
		return view.NewRoot(New(nil, nil))
	})
}

type ImageView struct {
	*view.Embed
	image image.Image
}

func New(ctx *view.Context, key interface{}) *ImageView {
	if v, ok := ctx.Prev(key).(*ImageView); ok {
		return v
	}

	dir, err := env.AssetsDir()
	fmt.Println("ENV", dir, err)

	file, err := os.Open(dir + "/TableArrow.png")
	img, str, err := image.Decode(file)

	return &ImageView{
		Embed: view.NewEmbed(ctx.NewId(key)),
		image: img,
	}
}

func (v *ImageView) Build(ctx *view.Context) *view.Model {
	l := constraint.New()

	chl := imageview.NewURLImageView(ctx, 0)
	chl.Painter = &paint.Style{BackgroundColor: colornames.Cyan}
	chl.URL = "https://www.google.com/images/branding/googlelogo/2x/googlelogo_color_272x92dp.png"
	chl.ResizeMode = imageview.ResizeModeFit
	l.Add(chl, func(s *constraint.Solver) {
		s.TopEqual(constraint.Const(100))
		s.LeftEqual(constraint.Const(100))
		s.WidthEqual(constraint.Const(200))
		s.HeightEqual(constraint.Const(200))
	})

	chl2 := imageview.NewImageView(ctx, 1)
	chl2.Image = v.image
	chl2.Painter = &paint.Style{BackgroundColor: colornames.Blue}
	// chl2.URL = "https://www.google.com/images/branding/googlelogo/2x/googlelogo_color_272x92dp.png"
	chl2.ResizeMode = imageview.ResizeModeFit
	l.Add(chl2, func(s *constraint.Solver) {
		s.TopEqual(constraint.Const(300))
		s.LeftEqual(constraint.Const(100))
		s.WidthEqual(constraint.Const(200))
		s.HeightEqual(constraint.Const(200))
	})

	return &view.Model{
		Children: []view.View{chl, chl2},
		Layouter: l,
		Painter:  &paint.Style{BackgroundColor: colornames.Green},
	}
}
