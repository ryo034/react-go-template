package error

import "fmt"

type Conflicted struct {
	table string
	key   interface{}
}

func NewConflicted(table string, key interface{}) *Conflicted {
	return &Conflicted{table: table, key: key}
}

func (c *Conflicted) Table() string {
	return c.table
}

func (c *Conflicted) Key() interface{} {
	return c.key
}

func (c *Conflicted) Error() string {
	return fmt.Sprintf("insert for table:%s has Conflicted on key:%#v", c.Table(), c.Key())
}

func (c *Conflicted) MessageKey() MessageKey {
	return ConflictedMessageKey
}
