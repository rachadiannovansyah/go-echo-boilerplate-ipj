package list

import (
	mysql "github.com/ShkrutDenis/go-migrations/builder"
	"github.com/jmoiron/sqlx"
)

type CreateCategoryTable struct{}

func (m *CreateCategoryTable) GetName() string {
	return "CreateCategoryTable"
}

func (m *CreateCategoryTable) Up(con *sqlx.DB) {
	table := mysql.NewTable("categories", con)
	table.Column("id").Type("int unsigned").Autoincrement()
	table.PrimaryKey("id")
	table.String("name", 250)
	table.String("description", 500).Nullable()
	table.Column("deleted_at").Type("datetime").Nullable()
	table.Column("user_id").Type("int unsigned")
	table.ForeignKey("user_id").
		Reference("users").
		On("id").
		OnDelete("cascade").
		OnUpdate("cascade")
	table.WithTimestamps()

	table.MustExec()
}

func (m *CreateCategoryTable) Down(con *sqlx.DB) {
	mysql.DropTable("categories", con).MustExec()
}
