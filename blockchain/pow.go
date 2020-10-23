package blockchain

import (
	"DataCertProject_Me/util"
	"bytes"
	"crypto/sha256"
	"fmt"
	"math/big"
)

const  DIFFFICULTY =19


//工作量证明结构体
type ProofOfWork struct {
   //目标值
	Target *big.Int
	//工作量证明算法对应的哪个区块
    Block   Block
}


//实例一个算法实例
func NewPow(block Block)  ProofOfWork{
	target:=big.NewInt(1)//初始值
	target.Lsh(target,255-DIFFFICULTY)
	pow:=ProofOfWork{
       Target:target,
       Block: block,
	}
	return pow
}


/*
*pow算法:寻找符合条件的nonce值*/
func (p ProofOfWork)Run() ([]byte,int64) {
	var nonce int64
	//var bigBlock *big.Int
	bigBlock:=new(big.Int)
	var block256hash []byte
	for  {

		block:=p.Block

		heightBytes,_:=util.IntToBytes(block.Height)
        timeBytes,_:=util.IntToBytes(block.TimeStamp)
        versionBytes:=util.StringToByte(block.Version)
        nonceBytes,_:=util.IntToBytes(nonce)
        blockBytes:=bytes.Join([][]byte{
        	heightBytes,
        	timeBytes,
        	block.Data,
        	block.PreHash,
        	versionBytes,
        	nonceBytes,
		},[]byte{})
        //fmt.Println("blockBytes:",blockBytes)


        //sha256 : sha256（A+nonce)
        sha256Hash := sha256.New()
        sha256Hash.Write(blockBytes)
		block256hash = sha256Hash.Sum(nil)
		//fmt.Println("block256Hash:",block256hash)


		//fmt.Println("挖矿中,当前尝试nonce值:\n",nonce)
        //sha256hash(区块+nonce值) 对应的大整数
        bigBlock=bigBlock.SetBytes(block256hash)
        fmt.Printf("目标值:%x\n",p.Target)
        fmt.Printf("Hash值:%x\n",bigBlock)
		if p.Target.Cmp(bigBlock) ==1 {//如果满足条件,退出循环
			break
		}
		nonce++//如果条件不满足,nonce值+1,继续下次循环

	}
	return block256hash,nonce
}