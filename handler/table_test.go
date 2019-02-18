package handler

import (
	"fmt"
	"reflect"
)

func ExampleTable_init() {
	t := Table{Name: "scores"}
	t.init(reflect.TypeOf(Score{}))
	fmt.Printf("%s\n%s\n%s\n%s\n", t.Name, t.Columns, t.LoadSql, t.CheckColumns)
	// Output:
	// scores
	// student_id,subject,score
	// select student_id,subject,score from scores
	// student_id,subject,score
}
