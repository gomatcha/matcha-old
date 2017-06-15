package text

import (
	"fmt"

	"github.com/gogo/protobuf/proto"
	"github.com/overcyn/mochi/layout"
	"github.com/overcyn/mochi/pb"
	"github.com/overcyn/mochi/pb/text"
	"github.com/overcyn/mochibridge"
)

type StyledText struct {
	text  *Text
	style *Style
}

func NewStyledText(text *Text) *StyledText {
	return &StyledText{
		text:  text,
		style: &Style{},
	}
}

func (st *StyledText) Text() *Text {
	return st.text
}

func (st *StyledText) At(a int) *Style {
	return nil
}

func (st *StyledText) Set(s *Style, start, end int) {
	st.style = s
}

func (st *StyledText) Update(s *Style, start, end int) {
}

func (st *StyledText) Size(min layout.Point, max layout.Point) layout.Point {
	sizeFunc := &text.SizeFunc{
		Text:    st.MarshalProtobuf(),
		MinSize: min.MarshalProtobuf(),
		MaxSize: max.MarshalProtobuf(),
	}
	data, err := proto.Marshal(sizeFunc)
	if err != nil {
		return layout.Pt(0, 0)
	}

	pointData := mochibridge.Bridge().Call("sizeForAttributedString:", mochibridge.Bytes(data)).ToInterface().([]byte)
	pbpoint := &pb.Point{}
	err = proto.Unmarshal(pointData, pbpoint)
	if err != nil {
		fmt.Println("size decode error", err)
		return layout.Pt(0, 0)
	}
	return layout.Pt(pbpoint.X, pbpoint.Y)
}

func (st *StyledText) MarshalProtobuf() *text.StyledText {
	return &text.StyledText{
		Text:  st.text.MarshalProtobuf(),
		Style: st.style.MarshalProtobuf(),
	}
}
