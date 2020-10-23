package blockchain

//
////桶的名称,该桶用于装区块信息
//var BUCKET_NAME = "blocks"
//
////表示最新区块的key名字
//
///*
//区块链结构体实例定义:用于表示代表一条区块链
//*该区块链包含以下功能:
//*①将新产生的区块与已有的区块连接起来,并保存
//*②可以查询某个区块的信息
//*③可以将所有区块进行遍历,输出区块信息
//*/
//
//type BlockChain struct {
//	LastHash []byte //最新区块的hash
//	BoltDb   bolt.DB
//}
//
///*
//*用于创建一条区块链,并返回该区块链实例
//*解释:由于区块链就是由一个一个的区块组成的,因此,如果要创建一条区块链,那么必须先*/
//
//func NewBlockChain() BlockChain {
//	genesls := CreatGenesisBlock() //创世区块
//	db, err := bolt.Open("chain.db", 0600, nil)
//	if err != nil {
//		panic(err.Error())
//	}
//	bl := BlockChain{
//		LastHash: genesls.Hash,
//		BoltDb:   db,
//	}
//	return bl
//}
//
///*
//*调用BlockChain的该SaveBlock方法,该方法可以将一个生成的新区块保存到chain.db文件中*/
//func (bc BlockChain) SvaeBlock(block Block) {
//	db := bc.BoltDb
//
//	//操作chain.db文件
//	db.Update(func(tx *bolt.Tx) error {
//		var tong *bolt.Bucket
//		tong = tx.Bucket([]byte(BUCKET_NAME))
//		if tong == nil { //新建
//			tong, err := tx.CreateBucket([]byte(BUCKET_NAME))
//			if err != nil {
//				return err
//			}
//			//先查看获取看桶中是否已包含要保存的区块
//			lastBlock := tong.Get([]byte("lasthash"))
//			blockHash, err := block.Serialize()
//			if err != nil {
//				return err
//			}
//			if lastBlock == nil { //未获取到最新hash值,桶中没有数据,是空的
//				tong.Put(block.Hash, blockHash) //存区块的信息
//				tong.Put([]byte("lasthahsh"), blockHash)
//			}
//			//把创世区块存到桶中,并更新最新去看的信息
//			blockBytes := block.Serialize()
//			tong.Put(blockHash, blockBytes)
//			//更新表示最新区块的标志
//			tong.Put(block.Hash, blockBytes)
//			return nil
//		})
//		return
//	}
