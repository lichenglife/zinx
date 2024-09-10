package ziface

/*

   为解决TCP 流的粘包问题， 按照TLV 封包格式
   Head = datalen + id
   Data = data

   进行每一个数据包的分包、拆包


*/

type IDataPack interface {

	// 获取包头长度
	GetHeadLen() uint32
	// 封包 将数据报文封装成 字节
	Pack(msg IMessage) ([]byte, error)
	//  解包  将字节流 拆解成Message

	UpPack([]byte) (IMessage, error)
}
