package env

import (
	"fmt"

	"github.com/gogo/protobuf/proto"
	"github.com/overcyn/mochi/layout"
	"github.com/overcyn/mochi/pb"
	"github.com/overcyn/mochi/pb/env"
	"github.com/overcyn/mochibridge"
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
	pointData := mochibridge.Bridge().Call("sizeForResource:", mochibridge.String(r.path)).ToInterface().([]byte)
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
