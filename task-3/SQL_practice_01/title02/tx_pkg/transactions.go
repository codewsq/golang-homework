package tx_pkg

import (
	"fmt"
	"gorm.io/gorm"
)

/*
记录表：transactions 表（包含字段 id 主键， from_account_id 转出账户ID，
to_account_id 转入账户ID， amount 转账金额）。
*/
type Transaction struct {
	ID            uint    `gorm:"primary_key;auto_increment;column:id" json:"id"`
	FromAccountId uint    `gorm:"not null" json:"fromAccountId"`
	ToAccountId   uint    `gorm:"not null" json:"toAccountId"`
	Amount        float64 `gorm:"not null" json:"amount"`
}

func TransactionInsert(fromAccountId uint, toAccountId uint, amount float64, db *gorm.DB) error {
	tran := Transaction{FromAccountId: fromAccountId, ToAccountId: toAccountId, Amount: amount}
	tx := db.Create(&tran)
	if tx.Error != nil {
		return tx.Error
	}

	//fmt.Println("转账记录保存成功，记录条数：", tx.RowsAffected, "条，记录信息：", tran)
	fmt.Printf("\n转账记录保存成功，记录条数：%d条，记录信息：转出账户ID：%d，转入账户ID: %d，转账金额：%d\n", tran.ID, tran.FromAccountId, tran.ToAccountId, int(tran.Amount))

	return nil
}
