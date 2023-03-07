package postgres

import (
	example "app-module/internal/app/example/domain"
	"app-module/pkg/postgres"
	"database/sql"
	"fmt"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	repo := New(nil, &postgres.Postgres{
		DB: db,
	})

	tests := []struct {
		input  string
		want   string
		err    error
		expect bool
	}{
		{
			input:  "tester",
			want:   "tester",
			err:    nil,
			expect: true,
		},
		{
			input:  "builder",
			want:   "tester",
			err:    fmt.Errorf("example unique"),
			expect: false,
		},
		{
			input:  "",
			want:   "tester",
			err:    fmt.Errorf("example constraint"),
			expect: false,
		},
	}

	for i := range tests {
		result := false
		willReturn := sqlmock.NewResult(1, 1)
		if tests[i].err != nil {
			willReturn = sqlmock.NewErrorResult(tests[i].err)
		}

		mock.ExpectExec("INSERT INTO example").WithArgs(tests[i].want).WillReturnResult(willReturn)

		if err := repo.Create(&example.Instance{
			Test: tests[i].input,
		}); err == nil {
			result = true
		}

		assert.Equal(t, result, tests[i].expect, "TestCase # %d", i+1)
	}
}

func TestGetOne(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	repo := New(nil, &postgres.Postgres{
		DB: db,
	})

	tests := []struct {
		input  int
		want   int
		result *sqlmock.Rows
		err    error
		expect bool
	}{
		{
			input:  1,
			want:   1,
			result: sqlmock.NewRows([]string{"name"}).AddRow("tester"),
			err:    nil,
			expect: true,
		},
		{
			input:  2,
			want:   1,
			result: new(sqlmock.Rows),
			err:    sql.ErrNoRows,
			expect: false,
		},
		{
			input:  3,
			want:   2,
			result: new(sqlmock.Rows),
			err:    fmt.Errorf("some error"),
			expect: false,
		},
	}

	for i := range tests {
		var (
			err    error
			result bool
		)

		mock.ExpectQuery("SELECT name FROM example").WithArgs(tests[i].want).WillReturnRows(tests[i].result).WillReturnError(tests[i].err)

		if _, err = repo.GetOne(tests[i].input); err == nil {
			result = true
		}

		assert.Equal(t, result, tests[i].expect, "TestCase # %d", i+1)
	}
}

func TestGetMany(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	repo := New(nil, &postgres.Postgres{
		DB: db,
	})

	tests := []struct {
		input  []int
		want   []int
		result *sqlmock.Rows
		err    error
		expect bool
	}{
		{
			input:  []int{1, 2, 3, 4, 5},
			want:   []int{1, 2, 3, 4, 5},
			result: sqlmock.NewRows([]string{"id", "name"}).AddRow(1, "tester"),
			err:    nil,
			expect: true,
		},
		{
			input:  nil,
			want:   []int{1, 2},
			result: new(sqlmock.Rows),
			err:    fmt.Errorf("some error"),
			expect: false,
		},
	}

	for i := range tests {
		var (
			err    error
			result bool
		)

		mock.ExpectQuery("SELECT id, name FROM example").WithArgs(pq.Array(tests[i].want)).WillReturnRows(tests[i].result).WillReturnError(tests[i].err)

		if _, err = repo.GetMany(tests[i].input); err == nil {
			result = true
		}

		assert.Equal(t, result, tests[i].expect, "TestCase # %d", i+1)
	}
}

func TestUpdate(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	repo := New(nil, &postgres.Postgres{
		DB: db,
	})

	tests := []struct {
		input  example.Instance
		want   example.Instance
		err    error
		expect bool
	}{
		{
			input:  example.Instance{ID: 1, Test: "tester"},
			want:   example.Instance{ID: 1, Test: "tester"},
			err:    nil,
			expect: true,
		},
		{
			input:  example.Instance{ID: 2, Test: "tester"},
			want:   example.Instance{ID: 3, Test: "tester"},
			err:    fmt.Errorf("example unique"),
			expect: false,
		},
		{
			input:  example.Instance{},
			want:   example.Instance{ID: 1},
			err:    fmt.Errorf("example constraint"),
			expect: false,
		},
	}

	for i := range tests {
		result := false
		willReturn := sqlmock.NewResult(1, 1)
		if tests[i].err != nil {
			willReturn = sqlmock.NewErrorResult(tests[i].err)
		}

		mock.ExpectExec("UPDATE example SET").WithArgs(tests[i].want.ID, tests[i].want.Test).WillReturnResult(willReturn)

		if err := repo.Update(&tests[i].input); err == nil {
			result = true
		}

		assert.Equal(t, result, tests[i].expect, "TestCase # %d", i+1)
	}
}

func TestDelete(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	repo := New(nil, &postgres.Postgres{
		DB: db,
	})

	tests := []struct {
		input  int
		want   int
		err    error
		expect bool
	}{
		{
			input:  1,
			want:   1,
			err:    nil,
			expect: true,
		},
		{
			input:  2,
			want:   3,
			err:    fmt.Errorf("example unique"),
			expect: false,
		},
		{
			input:  0,
			want:   1,
			err:    fmt.Errorf("example constraint"),
			expect: false,
		},
	}

	for i := range tests {
		result := false
		willReturn := sqlmock.NewResult(1, 1)
		if tests[i].err != nil {
			willReturn = sqlmock.NewErrorResult(tests[i].err)
		}

		mock.ExpectExec("DELETE FROM example").WithArgs(tests[i].want).WillReturnResult(willReturn)

		if err := repo.Delete(tests[i].input); err == nil {
			result = true
		}

		assert.Equal(t, result, tests[i].expect, "TestCase # %d", i+1)
	}
}
