package text

type StyledText struct {
	text *Text
	// style *Style
}

func NewStyledText(text *Text) *StyledText {
	return &StyledText{
		text: text,
	}
}

func (st *StyledText) Text() *Text {
	return st.text
}

func (st *StyledText) At(a int) *Style {
	return nil
}

func (st *StyledText) Set(s *Style, start, end int) {
}

func (st *StyledText) Update(s *Style, start, end int) {
}

// func (st *StyledText) Size(min layout.Point, max layout.Point) layout.Point {
// 	pbFunc := &pb.SizeFunc{
// 		Text:    t.MarshalProtobuf(),
// 		MinSize: min.MarshalProtobuf(),
// 		MaxSize: max.MarshalProtobuf(),
// 	}
// 	data, err := proto.Marshal(pbFunc)
// 	if err != nil {
// 		return layout.Pt(0, 0)
// 	}

// 	pointData := mochibridge.Bridge().Call("sizeForAttributedString:", mochibridge.Bytes(data)).ToInterface().([]byte)
// 	pbpoint := &pb.Point{}
// 	err = proto.Unmarshal(pointData, pbpoint)
// 	if err != nil {
// 		fmt.Println("size decode error", err)
// 		return layout.Pt(0, 0)
// 	}
// 	return layout.Pt(pbpoint.X, pbpoint.Y)
// }

// func (st *StyledText) MarshalProtobuf() *pb.Text {
// 	return &pb.Text{
// 		Text:  t.str,
// 		Style: t.style.MarshalProtobuf(),
// 	}
// }
