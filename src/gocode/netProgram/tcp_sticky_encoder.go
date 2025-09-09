package netProgram

import (
	"bytes"
	"encoding/binary"
	"errors"
	"io"
)

// 定义编码器（发送端）
type Encoder struct {
	// 编码结束后，写入的目标
	w io.Writer
}

func NewEncoder(w io.Writer) *Encoder {
	return &Encoder{w: w}
}

// 编码，将编码的结果，写入到w io.writer
func (enc *Encoder) Encode(message string) error {
	// 1 获取message的长度
	l := int32(len(message))

	buf := new(bytes.Buffer)
	// 2 在数据包中写入长度
	// 需要二进制的写入操作，需要将数据以bit的形式写入
	if err := binary.Write(buf, binary.LittleEndian, l); err != nil {
		return err
	}

	// 3 将数据body写入
	if err := binary.Write(buf, binary.LittleEndian, []byte(message)); err != nil {
		return err
	}
	// 4 利用io.Writer发送数据
	if _, err := enc.w.Write(buf.Bytes()); err != nil {
		return err
	}
	return nil
}

// 定义解码器（接收端）

type Decoder struct {
	r io.Reader
}

func NewDecoder(r io.Reader) *Decoder {
	return &Decoder{r: r}
}

func (dec *Decoder) Decode(message *string) error {
	// 读取前四个字节，读取header
	header := make([]byte, 4)
	hn, err := dec.r.Read(header)
	if err != nil {
		return err
	}
	if hn != 4 {
		return errors.New("header is not enough")
	}

	// 将前四个字节转换为int32类型，确定了body的长度
	var l int32
	buf := bytes.NewBuffer(header)
	if err := binary.Read(buf, binary.LittleEndian, &l); err != nil {
		return err
	}

	//  读取body
	body := make([]byte, l)
	bn, err := dec.r.Read(body)
	if err != nil {
		return err
	}
	if bn != int(l) {
		return errors.New("body is not enough")
	}
	*message = string(body)
	return nil
}
