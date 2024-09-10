package znet

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"lichenglife/zinx/utils"
	"lichenglife/zinx/ziface"
)

/*


实现数据包的封包、拆包方法
*/

// 定义空结构体
type DataPack struct{}

func NewDataPack() ziface.IDataPack {

	return &DataPack{}
}

// 获取包头长度
func (d *DataPack) GetHeadLen() uint32 {
	// 固定长度  ID uint32(4字节) + dataLen(4字节)

	return 8
}

// 封包 将数据报文封装成 字节
func (d *DataPack) Pack(msg ziface.IMessage) ([]byte, error) {

	//  定义字节缓冲区
	buff := bytes.NewBuffer([]byte{})

	// dataLen
	if err := binary.Write(buff, binary.LittleEndian, msg.GetDataLen()); err != nil {
		fmt.Println("pack Message error", err)
		return nil, err
	}
	if err := binary.Write(buff, binary.LittleEndian, msg.GetMsgId()); err != nil {
		fmt.Println("pack Message error", err)
		return nil, err
	}
	if err := binary.Write(buff, binary.LittleEndian, msg.GetData()); err != nil {
		fmt.Println("pack Message error", err)
		return nil, err
	}

	return buff.Bytes(), nil

}

//  解包  将字节流 拆解成Message

func (d *DataPack) UpPack(data []byte) (ziface.IMessage, error) {
	//  从字节流中读取数据
	databuff := bytes.NewReader(data)

	//  定义空结构体,然后通过指针进行赋值
	m := &Message{}
	if err := binary.Read(databuff, binary.LittleEndian, &m.DataLen); err != nil {
		fmt.Println("unpack message error", err)

	}

	if err := binary.Read(databuff, binary.LittleEndian, &m.Id); err != nil {
		fmt.Println("unpack message error", err)

	}

	if utils.GlobalObject.MaxPacketSize > 0 && utils.GlobalObject.MaxPacketSize < m.GetDataLen() {

		err := errors.New("tcp maxpacksize is too  large ")

		return nil, err
	}
	if err := binary.Read(databuff, binary.LittleEndian, &m.Data); err != nil {
		fmt.Println("unpack message error", err)

	}

	return m, nil
}
