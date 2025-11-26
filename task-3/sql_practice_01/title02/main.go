package main

import (
	"fmt"
	_ "github.com/codewsq/sql/title02/my_db"
	pkg "github.com/codewsq/sql/title02/tx_pkg"
	"time"
)

/*
题目2：事务语句
假设有两个表： accounts 表（包含字段 id 主键， balance 账户余额）和 transactions 表（包含字段 id 主键， from_account_id 转出账户ID，
to_account_id 转入账户ID， amount 转账金额）。
要求 ：
编写一个事务，实现从账户 A 向账户 B 转账 100 元的操作。在事务中，需要先检查账户 A 的余额是否足够，
如果足够则从账户 A 扣除 100 元，向账户 B 增加 100 元，并在 transactions 表中记录该笔转账信息。
如果余额不足，则回滚事务。
*/

func main() {
	fmt.Println("=====================第1次交易===================")
	err1 := pkg.TransferAToB100()
	if err1 != nil {
		fmt.Println(err1)
	}
	time.Sleep(5 * time.Second)

	fmt.Println("=====================第2次交易===================")
	err2 := pkg.TransferAToB100()
	if err2 != nil {
		fmt.Println(err2)
	}
	time.Sleep(5 * time.Second)

	fmt.Println("=====================第3次交易===================")
	err3 := pkg.TransferAToB100()
	if err3 != nil {
		fmt.Println(err3)
	}
	time.Sleep(5 * time.Second)
}
