// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package json_benchmark

import (
	json "encoding/json"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjson6a975c40DecodeJsonBenchmark(in *jlexer.Lexer, out *UnloadCheckResponse) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "id":
			out.Id = string(in.String())
		case "orders":
			if in.IsNull() {
				in.Skip()
				out.Orders = nil
			} else {
				in.Delim('[')
				if out.Orders == nil {
					if !in.IsDelim(']') {
						out.Orders = make([]OrderCheckResponse, 0, 0)
					} else {
						out.Orders = []OrderCheckResponse{}
					}
				} else {
					out.Orders = (out.Orders)[:0]
				}
				for !in.IsDelim(']') {
					var v1 OrderCheckResponse
					(v1).UnmarshalEasyJSON(in)
					out.Orders = append(out.Orders, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "status":
			out.Status = string(in.String())
		case "transportIdentifier":
			out.TransportIdentifier = string(in.String())
		case "locationCode":
			out.LocationCode = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson6a975c40EncodeJsonBenchmark(out *jwriter.Writer, in UnloadCheckResponse) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.String(string(in.Id))
	}
	{
		const prefix string = ",\"orders\":"
		out.RawString(prefix)
		if in.Orders == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v2, v3 := range in.Orders {
				if v2 > 0 {
					out.RawByte(',')
				}
				(v3).MarshalEasyJSON(out)
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"status\":"
		out.RawString(prefix)
		out.String(string(in.Status))
	}
	{
		const prefix string = ",\"transportIdentifier\":"
		out.RawString(prefix)
		out.String(string(in.TransportIdentifier))
	}
	{
		const prefix string = ",\"locationCode\":"
		out.RawString(prefix)
		out.String(string(in.LocationCode))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v UnloadCheckResponse) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson6a975c40EncodeJsonBenchmark(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v UnloadCheckResponse) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson6a975c40EncodeJsonBenchmark(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *UnloadCheckResponse) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson6a975c40DecodeJsonBenchmark(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *UnloadCheckResponse) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson6a975c40DecodeJsonBenchmark(l, v)
}
func easyjson6a975c40DecodeJsonBenchmark1(in *jlexer.Lexer, out *OrderCheckResponse) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "id":
			out.Id = string(in.String())
		case "code":
			out.Code = string(in.String())
		case "items":
			if in.IsNull() {
				in.Skip()
				out.Items = nil
			} else {
				in.Delim('[')
				if out.Items == nil {
					if !in.IsDelim(']') {
						out.Items = make([]ItemCheckResponse, 0, 0)
					} else {
						out.Items = []ItemCheckResponse{}
					}
				} else {
					out.Items = (out.Items)[:0]
				}
				for !in.IsDelim(']') {
					var v4 ItemCheckResponse
					(v4).UnmarshalEasyJSON(in)
					out.Items = append(out.Items, v4)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "type":
			out.Type = string(in.String())
		case "blitzedItems":
			if in.IsNull() {
				in.Skip()
				out.BlitzedItems = nil
			} else {
				in.Delim('[')
				if out.BlitzedItems == nil {
					if !in.IsDelim(']') {
						out.BlitzedItems = make([]BlitzedItemResponse, 0, 1)
					} else {
						out.BlitzedItems = []BlitzedItemResponse{}
					}
				} else {
					out.BlitzedItems = (out.BlitzedItems)[:0]
				}
				for !in.IsDelim(']') {
					var v5 BlitzedItemResponse
					(v5).UnmarshalEasyJSON(in)
					out.BlitzedItems = append(out.BlitzedItems, v5)
					in.WantComma()
				}
				in.Delim(']')
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson6a975c40EncodeJsonBenchmark1(out *jwriter.Writer, in OrderCheckResponse) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.String(string(in.Id))
	}
	{
		const prefix string = ",\"code\":"
		out.RawString(prefix)
		out.String(string(in.Code))
	}
	{
		const prefix string = ",\"items\":"
		out.RawString(prefix)
		if in.Items == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v6, v7 := range in.Items {
				if v6 > 0 {
					out.RawByte(',')
				}
				(v7).MarshalEasyJSON(out)
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"type\":"
		out.RawString(prefix)
		out.String(string(in.Type))
	}
	{
		const prefix string = ",\"blitzedItems\":"
		out.RawString(prefix)
		if in.BlitzedItems == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v8, v9 := range in.BlitzedItems {
				if v8 > 0 {
					out.RawByte(',')
				}
				(v9).MarshalEasyJSON(out)
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v OrderCheckResponse) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson6a975c40EncodeJsonBenchmark1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v OrderCheckResponse) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson6a975c40EncodeJsonBenchmark1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *OrderCheckResponse) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson6a975c40DecodeJsonBenchmark1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *OrderCheckResponse) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson6a975c40DecodeJsonBenchmark1(l, v)
}
func easyjson6a975c40DecodeJsonBenchmark2(in *jlexer.Lexer, out *ItemCheckResponse) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "id":
			out.CheckId = string(in.String())
		case "code":
			out.Code = string(in.String())
		case "description":
			out.Description = string(in.String())
		case "batch":
			out.Batch = string(in.String())
		case "unitsOfMeasurement":
			if in.IsNull() {
				in.Skip()
				out.UnitsOfMeasurement = nil
			} else {
				in.Delim('[')
				if out.UnitsOfMeasurement == nil {
					if !in.IsDelim(']') {
						out.UnitsOfMeasurement = make([]string, 0, 4)
					} else {
						out.UnitsOfMeasurement = []string{}
					}
				} else {
					out.UnitsOfMeasurement = (out.UnitsOfMeasurement)[:0]
				}
				for !in.IsDelim(']') {
					var v10 string
					v10 = string(in.String())
					out.UnitsOfMeasurement = append(out.UnitsOfMeasurement, v10)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "unitOfMeasurement":
			out.UnitOfMeasurement = string(in.String())
		case "manufactureDate":
			out.ManufactureDate = string(in.String())
		case "document":
			(out.Document).UnmarshalEasyJSON(in)
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson6a975c40EncodeJsonBenchmark2(out *jwriter.Writer, in ItemCheckResponse) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.String(string(in.CheckId))
	}
	{
		const prefix string = ",\"code\":"
		out.RawString(prefix)
		out.String(string(in.Code))
	}
	{
		const prefix string = ",\"description\":"
		out.RawString(prefix)
		out.String(string(in.Description))
	}
	if in.Batch != "" {
		const prefix string = ",\"batch\":"
		out.RawString(prefix)
		out.String(string(in.Batch))
	}
	{
		const prefix string = ",\"unitsOfMeasurement\":"
		out.RawString(prefix)
		if in.UnitsOfMeasurement == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v11, v12 := range in.UnitsOfMeasurement {
				if v11 > 0 {
					out.RawByte(',')
				}
				out.String(string(v12))
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"unitOfMeasurement\":"
		out.RawString(prefix)
		out.String(string(in.UnitOfMeasurement))
	}
	{
		const prefix string = ",\"manufactureDate\":"
		out.RawString(prefix)
		out.String(string(in.ManufactureDate))
	}
	{
		const prefix string = ",\"document\":"
		out.RawString(prefix)
		(in.Document).MarshalEasyJSON(out)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v ItemCheckResponse) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson6a975c40EncodeJsonBenchmark2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ItemCheckResponse) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson6a975c40EncodeJsonBenchmark2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ItemCheckResponse) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson6a975c40DecodeJsonBenchmark2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ItemCheckResponse) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson6a975c40DecodeJsonBenchmark2(l, v)
}
func easyjson6a975c40DecodeJsonBenchmark3(in *jlexer.Lexer, out *DocumentResponse) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "code":
			out.Code = string(in.String())
		case "id":
			out.Id = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson6a975c40EncodeJsonBenchmark3(out *jwriter.Writer, in DocumentResponse) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"code\":"
		out.RawString(prefix[1:])
		out.String(string(in.Code))
	}
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix)
		out.String(string(in.Id))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v DocumentResponse) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson6a975c40EncodeJsonBenchmark3(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v DocumentResponse) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson6a975c40EncodeJsonBenchmark3(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *DocumentResponse) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson6a975c40DecodeJsonBenchmark3(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *DocumentResponse) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson6a975c40DecodeJsonBenchmark3(l, v)
}
func easyjson6a975c40DecodeJsonBenchmark4(in *jlexer.Lexer, out *BlitzedItemResponse) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "itemCode":
			out.ItemCode = string(in.String())
		case "amount":
			out.Amount = string(in.String())
		case "unitOfMeasurement":
			out.UnitOfMeasurement = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson6a975c40EncodeJsonBenchmark4(out *jwriter.Writer, in BlitzedItemResponse) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"itemCode\":"
		out.RawString(prefix[1:])
		out.String(string(in.ItemCode))
	}
	{
		const prefix string = ",\"amount\":"
		out.RawString(prefix)
		out.String(string(in.Amount))
	}
	{
		const prefix string = ",\"unitOfMeasurement\":"
		out.RawString(prefix)
		out.String(string(in.UnitOfMeasurement))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v BlitzedItemResponse) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson6a975c40EncodeJsonBenchmark4(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v BlitzedItemResponse) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson6a975c40EncodeJsonBenchmark4(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *BlitzedItemResponse) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson6a975c40DecodeJsonBenchmark4(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *BlitzedItemResponse) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson6a975c40DecodeJsonBenchmark4(l, v)
}
