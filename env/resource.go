package env

import (
	"bytes"
	"fmt"
	"image"
	"image/color"

	"golang.org/x/image/colornames"

	"github.com/gogo/protobuf/proto"
	"github.com/overcyn/matcha/layout"
	"github.com/overcyn/matcha/pb/env"
	pb "github.com/overcyn/matcha/pb/layout"
	"github.com/overcyn/matchabridge"
)

type Resource struct {
	path string
}

func Load(path string) (*Resource, error) {
	return &Resource{path: path}, nil
}

func MustLoad(path string) *Resource {
	res, err := Load(path)
	if err != nil {
		panic(err.Error())
	}
	return res
}

func (r *Resource) Size() layout.Point {
	pointData := matchabridge.Bridge().Call("sizeForResource:", matchabridge.String(r.path)).ToInterface().([]byte)
	pbpoint := &pb.Point{}
	err := proto.Unmarshal(pointData, pbpoint)
	if err != nil {
		fmt.Println("size decode error", err)
		return layout.Pt(0, 0)
	}
	return layout.Pt(pbpoint.X, pbpoint.Y)
}

func (r *Resource) MarshalProtobuf() *env.Resource {
	return &env.Resource{
		Path: r.path,
	}
}

type ImageResource struct {
	path  string
	rect  image.Rectangle
	image image.Image
}

func LoadImage(path string) (*ImageResource, error) {
	pointData := matchabridge.Bridge().Call("sizeForResource:", matchabridge.String(path)).ToInterface().([]byte)
	pbpoint := &pb.Point{}
	err := proto.Unmarshal(pointData, pbpoint)
	if err != nil {
		return nil, err
	}

	return &ImageResource{
		path:  path,
		rect:  image.Rect(0, 0, int(pbpoint.X), int(pbpoint.Y)),
		image: nil,
	}, nil
}

func MustLoadImage(path string) *ImageResource {
	res, err := LoadImage(path)
	if err != nil {
		panic(err.Error())
	}
	return res
}

func (res *ImageResource) ColorModel() color.Model {
	if res.image == nil {
		res.load()
	}
	return res.image.ColorModel()
}

func (res *ImageResource) Bounds() image.Rectangle {
	return res.rect
}

func (res *ImageResource) At(x, y int) color.Color {
	if res.image == nil {
		res.load()
	}
	return res.image.At(x, y)
}

func (res *ImageResource) load() {
	data := matchabridge.Bridge().Call("imageForResource:", matchabridge.String(res.path)).ToInterface().([]byte)
	reader := bytes.NewReader(data)
	img, _, err := image.Decode(reader)
	if err != nil {
		res.image = image.NewUniform(colornames.Black)
		return
	}
	res.image = img
}

func (res *ImageResource) MarshalProtobuf() *env.ImageResource {
	return &env.ImageResource{
		Path: res.path,
	}
}
