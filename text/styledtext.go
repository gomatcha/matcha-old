package text

import (
	"fmt"

	"github.com/gogo/protobuf/proto"
	"gomatcha.io/bridge"
	"gomatcha.io/matcha/layout"
	pb "gomatcha.io/matcha/pb/layout"
	"gomatcha.io/matcha/pb/text"
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

	pointData := bridge.Bridge().Call("sizeForAttributedString:", bridge.Bytes(data)).ToInterface().([]byte)
	pbpoint := &pb.Point{}
	err = proto.Unmarshal(pointData, pbpoint)
	if err != nil {
		fmt.Println("StyledText.Size(): Decode error", err)
		return layout.Pt(0, 0)
	}
	return layout.Pt(pbpoint.X, pbpoint.Y)
}

func (st *StyledText) MarshalProtobuf() *text.StyledText {
	if st == nil {
		return nil
	}
	return &text.StyledText{
		Text:  st.text.MarshalProtobuf(),
		Style: st.style.MarshalProtobuf(),
	}
}
