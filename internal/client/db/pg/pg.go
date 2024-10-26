package pg

import (
	"context"
	"fmt"
	"log"
	"reflect"
	"strings"

	"github.com/Dnlbb/auth/internal/client/db"
	"github.com/Dnlbb/auth/internal/client/db/prettier"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

type key string

// TxKey ключ транзакции.
const (
	TxKey key = "tx"
)

type pg struct {
	dbc *pgxpool.Pool
}

// NewDB конструктор для базы
func NewDB(dbc *pgxpool.Pool) db.DB {
	return &pg{
		dbc: dbc,
	}
}

func (p *pg) ScanOneContext(ctx context.Context, dest interface{}, query db.Query, args ...interface{}) error {
	logQuery(ctx, query, args...)

	rows, err := p.QueryContext(ctx, query, args...)
	if err != nil {
		return err
	}
	defer rows.Close()

	// Проверяем, что в rows есть хотя бы одна строка
	if !rows.Next() {
		return fmt.Errorf("no rows in result set")
	}

	// Получаем описание полей (колонок) из rows
	columns := rows.FieldDescriptions()
	numColumns := len(columns)

	// Подготавливаем сканируемые значения для каждой колонки
	scanArgs := make([]interface{}, numColumns)
	v := reflect.ValueOf(dest).Elem() // Получаем значение, на которое указывает dest

	for i, column := range columns {
		// Находим поле структуры, соответствующее имени колонки
		field := v.FieldByNameFunc(func(name string) bool {
			return strings.EqualFold(name, string(column.Name))
		})
		if field.IsValid() && field.CanAddr() {
			scanArgs[i] = field.Addr().Interface()
		} else {
			// Используем временное значение, если поле не найдено
			var ignore interface{}
			scanArgs[i] = &ignore
		}
	}

	// Сканируем строку в подготовленные аргументы
	if err := rows.Scan(scanArgs...); err != nil {
		return err
	}

	// Проверяем наличие дополнительных строк — если они есть, возвращаем ошибку
	if rows.Next() {
		return fmt.Errorf("more than one row in result set")
	}

	return rows.Err()
}

func (p *pg) QueryContext(ctx context.Context, q db.Query, args ...interface{}) (pgx.Rows, error) {
	logQuery(ctx, q, args...)

	tx, err := ctx.Value(TxKey).(pgx.Tx)
	if err {
		return tx.Query(ctx, q.QueryRow, args...)
	}

	return p.dbc.Query(ctx, q.QueryRow, args...)
}

func (p *pg) ScanAllContext(ctx context.Context, dest interface{}, query db.Query, args ...interface{}) error {
	logQuery(ctx, query, args...)

	rows, err := p.QueryContext(ctx, query, args...)
	if err != nil {
		return err
	}
	defer rows.Close()

	// Получаем описание колонок
	columns := rows.FieldDescriptions()
	numColumns := len(columns)

	// Отражаем тип получаемой структуры
	v := reflect.ValueOf(dest)
	if v.Kind() != reflect.Ptr || v.Elem().Kind() != reflect.Slice {
		return fmt.Errorf("dest должен быть указателем на срез структур")
	}
	sliceValue := v.Elem()

	// Итерация по строкам
	for rows.Next() {
		// Создаем новый экземпляр элемента среза
		elem := reflect.New(sliceValue.Type().Elem()).Elem()
		scanArgs := make([]interface{}, numColumns)

		// Устанавливаем указатели на поля структуры
		for i, column := range columns {
			field := elem.FieldByNameFunc(func(name string) bool {
				return strings.EqualFold(name, string(column.Name))
			})
			if field.IsValid() && field.CanAddr() {
				scanArgs[i] = field.Addr().Interface()
			} else {
				// Если нет соответствующего поля, используем временную переменную
				var ignore interface{}
				scanArgs[i] = &ignore
			}
		}

		// Сканируем значения строки в подготовленные указатели
		if err := rows.Scan(scanArgs...); err != nil {
			return err
		}

		// Добавляем элемент в срез
		sliceValue.Set(reflect.Append(sliceValue, elem))
	}

	return rows.Err()
}

func (p *pg) ExecContext(ctx context.Context, q db.Query, args ...interface{}) (pgconn.CommandTag, error) {
	logQuery(ctx, q, args...)

	tx, err := ctx.Value(TxKey).(pgx.Tx)
	if err {
		return tx.Exec(ctx, q.QueryRow, args...)
	}

	return p.dbc.Exec(ctx, q.QueryRow, args...)
}

func (p *pg) QueryRowContext(ctx context.Context, q db.Query, args ...interface{}) pgx.Row {
	logQuery(ctx, q, args...)

	tx, err := ctx.Value(TxKey).(pgx.Tx)
	if err {
		return tx.QueryRow(ctx, q.QueryRow, args...)
	}

	return p.dbc.QueryRow(ctx, q.QueryRow, args...)
}

func logQuery(ctx context.Context, q db.Query, args ...interface{}) {
	prettyQuery := prettier.Pretty(q.QueryRow, prettier.PlaceholderDollar, args...)
	log.Println(
		ctx,
		fmt.Sprintf("sql: %s", q.Name),
		fmt.Sprintf("query: %s", prettyQuery),
	)
}

func (p *pg) Ping(ctx context.Context) error {
	return p.dbc.Ping(ctx)
}

func (p *pg) Close() {
	p.dbc.Close()
}

func (p *pg) BeginTx(ctx context.Context, txOptions pgx.TxOptions) (pgx.Tx, error) {
	return p.dbc.BeginTx(ctx, txOptions)
}

// MakeContextTx добавляем в контекст ключ транзакций
func MakeContextTx(ctx context.Context, tx pgx.Tx) context.Context {
	return context.WithValue(ctx, TxKey, tx)
}
