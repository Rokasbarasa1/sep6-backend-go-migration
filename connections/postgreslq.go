package connections

import (
	"context"
	"fmt"
	"log"
	"reflect"
	"sync"

	"github.com/jackc/pgx/v4/pgxpool"
)

var PostgreSLQPool *pgxpool.Pool

func ConnectToPostgres() {
	var err error
	// databaseURL := os.Getenv("DATABASE_URL")                   // Get the database URL from environment variables
	databaseURL := "host=localhost port=4000 user=postgres password=postgres dbname=postgres sslmode=disable" // Get the database URL from environment variables

	if databaseURL == "" {
		log.Fatal("DATABASE_URL is not set.")
	}

	PostgreSLQPool, err = pgxpool.Connect(context.Background(), databaseURL)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}

	fmt.Println("Connected to PostgreSQL!")
}

var (
	fieldMapCache sync.Map // Safe for concurrent use
)

// getFieldMapping retrieves or builds the field mapping for a given type and columns.
func getFieldMapping(t reflect.Type, columns []string) ([]int, error) {
	if cached, ok := fieldMapCache.Load(t); ok {
		return cached.([]int), nil
	}

	mapping := make([]int, len(columns))
	for i, col := range columns {
		found := false
		for j := 0; j < t.NumField(); j++ {
			field := t.Field(j)
			if field.Tag.Get("db") == col {
				mapping[i] = j
				found = true
				break
			}
		}
		if !found {
			return nil, fmt.Errorf("no field corresponding to column %s", col)
		}
	}

	fieldMapCache.Store(t, mapping)
	return mapping, nil
}

// ExecuteQuery executes a SQL query and fills a slice of type T with the results.
func PostgreSQLQuery[T any](dest *[]T, query string, args []interface{}) error {
	// Make the query using the query string and the arguments
	rows, err := PostgreSLQPool.Query(context.Background(), query, args...)
	if err != nil {
		return err
	}
	defer rows.Close() // Basically "final" tag to perform code right before return

	columns := rows.FieldDescriptions() // Correct usage without error handling

	columnNames := make([]string, len(columns))
	for i, col := range columns {
		columnNames[i] = string(col.Name)
	}

	t := reflect.TypeOf(*new(T))
	mapping, err := getFieldMapping(t, columnNames)
	if err != nil {
		return err
	}

	for rows.Next() {
		elemPtr := reflect.New(t).Elem()
		scanArgs := make([]interface{}, len(columns))

		for i, idx := range mapping {
			scanArgs[i] = elemPtr.Field(idx).Addr().Interface()
		}

		if err := rows.Scan(scanArgs...); err != nil {
			return err
		}

		*dest = append(*dest, elemPtr.Interface().(T))
	}

	return rows.Err()
}

func init() {
	ConnectToPostgres()
}
