package main
//
//import (
//	"fmt"
//	"github.com/boltdb/bolt"
//)
//
//const BUCKET_NAME = "class3"
//
////草稿,请别在意
//func main() {
//	db, err := bolt.Open("qiqicancanqiqi.db", 0600, nil)
//	if err != nil {
//     fmt.Println(err.Error())
//		return
//	}
//	defer db.Close()
//
//	//只读模式
//	db.View(func(tx *bolt.Tx) error {
//		bucket:=tx.Bucket([]byte(BUCKET_NAME))
//		if bucket!=nil {
//			bucket.Get([]byte("0"))
//		}
//		return nil
//	})
//
//	//写操作
//	db.Update(func(tx *bolt.Tx) error {
//		bucket:=tx.Bucket([]byte(BUCKET_NAME))
//		if bucket==nil {//名字class3的桶不在
//			bk,err:=tx.CreateBucket([]byte(BUCKET_NAME))
//			if err!=nil {
//				return err
//			}
//			bk.Put([]byte("0"),[]byte)
//		}
//		return nil
//	})
//	fmt.Println(db)
//}
//
//type Block struct {
//	Height int64
//	TimeStamp int64
//
//}