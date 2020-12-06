package repo

type Technology struct {
	tableName struct{} `sql:"technologies"`

	Name    string `sql:"name"`
	Details string `sql:"details"`
}
