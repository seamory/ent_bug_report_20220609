package bug

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"strconv"
	"testing"

	"entgo.io/ent/dialect"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"

	"entgo.io/bug/ent"
	"entgo.io/bug/ent/enttest"
)

func TestBugSQLite(t *testing.T) {
	client := enttest.Open(t, dialect.SQLite, "file:ent?mode=memory&cache=shared&_fk=1")
	defer client.Close()
	test(t, client)
}

func TestBugMySQL(t *testing.T) {
	for version, port := range map[string]int{"56": 3306, "57": 3307, "8": 3308} {
		addr := net.JoinHostPort("localhost", strconv.Itoa(port))
		t.Run(version, func(t *testing.T) {
			client := enttest.Open(t, dialect.MySQL, fmt.Sprintf("root:pass@tcp(%s)/test?parseTime=True", addr))
			defer client.Close()
			test(t, client)
		})
	}
}

func TestBugPostgres(t *testing.T) {
	for version, port := range map[string]int{"10": 5430, "11": 5431, "12": 5432, "13": 5433, "14": 5434} {
		t.Run(version, func(t *testing.T) {
			client := enttest.Open(t, dialect.Postgres, fmt.Sprintf("host=localhost port=%d user=postgres dbname=test password=pass sslmode=disable", port))
			defer client.Close()
			test(t, client)
		})
	}
}

func TestBugMaria(t *testing.T) {
	for version, port := range map[string]int{"10.5": 4306, "10.2": 4307, "10.3": 4308} {
		t.Run(version, func(t *testing.T) {
			addr := net.JoinHostPort("localhost", strconv.Itoa(port))
			client := enttest.Open(t, dialect.MySQL, fmt.Sprintf("root:pass@tcp(%s)/test?parseTime=True", addr))
			defer client.Close()
			test(t, client)
		})
	}
}

func test(t *testing.T, client *ent.Client) {
	ctx := context.Background()
	// add basic records
	item1 := client.Hierarchy.Create().SetName("test1").SetDescribe("test1").SaveX(ctx)
	item2 := client.Hierarchy.Create().SetName("test2").SetDescribe("test2").SaveX(ctx)
	item3 := client.Hierarchy.Create().SetName("test3").SetDescribe("test3").SaveX(ctx)
	// switch to debug mode
	data := client.Debug().Hierarchy.GetX(ctx, item2.ID)
	// when do add child ids
	data.Update().AddChildIDs(item3.ID).Save(ctx)
	// when do add parent ids
	data.Update().AddParentIDs(item1.ID).Save(ctx)
	// print result
	// should be item1 is item2 parent, item2 is item3 parent. but not, item3 is also item2 parent.
	result := client.Hierarchy.Query().WithChild().WithParent().AllX(ctx)
	if blob, err := json.Marshal(result); err != nil {
		log.Fatalln(err)
	} else {
		fmt.Println(string(blob))
	}
}
