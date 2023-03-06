package postgres

import (
	example "app-module/internal/app/example/domain"
	"app-module/pkg/postgres"
	"database/sql"
	"fmt"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
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
		want   error
		expect bool
	}{
		{
			input:  "tester",
			want:   nil,
			expect: true,
		},
		{
			input:  "builder",
			want:   fmt.Errorf("example unique"),
			expect: false,
		},
		{
			input:  "",
			want:   fmt.Errorf("example constraint"),
			expect: false,
		},
	}

	for i := range tests {
		result := false
		willReturn := sqlmock.NewResult(1, 1)
		if tests[i].want != nil {
			willReturn = sqlmock.NewErrorResult(tests[i].want)
		}

		mock.ExpectExec("INSERT INTO example").WithArgs(tests[i].input).WillReturnResult(willReturn)

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
		want   error
		expect bool
	}{
		{
			input:  1,
			want:   nil,
			expect: true,
		},
		{
			input:  2,
			want:   sql.ErrNoRows,
			expect: false,
		},
		{
			input:  3,
			want:   fmt.Errorf("some error"),
			expect: false,
		},
	}

	for i := range tests {
		var (
			err    error
			result bool
			rows   = new(sqlmock.Rows)
		)

		if tests[i].want != nil {
			err = tests[i].want
		}

		mock.ExpectQuery("SELECT name FROM example").WithArgs(tests[i].input).WillReturnRows(rows).WillReturnError(err)

		if _, err = repo.GetOne(tests[i].input); err == nil {
			result = true
		}

		assert.Equal(t, result, tests[i].expect, "TestCase # %d", i+1)
	}
}
