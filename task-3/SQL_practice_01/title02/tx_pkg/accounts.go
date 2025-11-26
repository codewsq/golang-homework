package tx_pkg

import (
	"fmt"
	mydb "github.com/codewsq/sql/title02/my_db"
)

/*
要求 ：
编写一个事务，实现从账户 A 向账户 B 转账 100 元的操作。在事务中，需要先检查账户 A 的余额是否足够，
如果足够则从账户 A 扣除 100 元，向账户 B 增加 100 元，并在 transactions 表中记录该笔转账信息。
如果余额不足，则回滚事务。
*/
type Accounts struct {
	ID      uint `gorm:"primary_key;auto_increment;column:id" json:"id"`
	Balance float64
	Name    string
}

func TransferAToB100() error {
	// 定义转账金额
	tranBalance := 100.00

	// 开始事务
	tx := mydb.OpenDB.Begin()
	a := Accounts{}
	b := Accounts{}
	// 查询A和B账户信息
	tx.Where(&Accounts{Name: "A"}).Order("id asc").First(&a)
	tx.Where(&Accounts{Name: "B"}).Order("id asc").First(&b)
	fmt.Println("交易前 - 账户A信息：", a)
	fmt.Println("交易前 - 账户B信息：", b)
	if a.Balance < tranBalance {
		fmt.Println("账户A余额不足，交易失败")
		return tx.Commit().Error
	}
	// A账户-100
	updateA := tx.Model(&a).Update("balance", a.Balance-tranBalance)
	if updateA.Error != nil {
		tx.Rollback()
		fmt.Println("A账户更新失败，事务回滚")
		return updateA.Error
	}
	// B账户+100
	updateB := tx.Model(&b).Update("balance", b.Balance+tranBalance)
	if updateB.Error != nil {
		tx.Rollback()
		fmt.Println("B账户更新失败，事务回滚")
		return updateB.Error
	}
	// 添加交易记录
	createT := TransactionInsert(a.ID, b.ID, tranBalance, mydb.OpenDB)
	if createT != nil {
		tx.Rollback()
		fmt.Println("日志记录失败，事务回滚")
		return createT
	}
	er := tx.Commit().Error
	if er != nil {
		fmt.Println("事务提交失败。。")
	} else {
		fmt.Println("事务提交成功,交易完成！")
	}
	// 提交事务
	return er
}

func init() {
	mydb.OpenDB.AutoMigrate(&Accounts{})
	mydb.OpenDB.AutoMigrate(&Transaction{})
	accounts := []Accounts{{Balance: 230, Name: "A"}, {Balance: 0, Name: "B"}}
	result := mydb.OpenDB.Create(&accounts)
	if result.Error != nil {
		panic(result.Error)
	}
	fmt.Println("原始数据插入成功！")

}

/*

// 开始事务
tx := db.Begin()

// 在事务中执行一些 db 操作（从这里开始，您应该使用 'tx' 而不是 'db'）
tx.Create(...)

// ...

// 遇到错误时回滚事务
tx.Rollback()

// 否则，提交事务
tx.Commit()
*/
