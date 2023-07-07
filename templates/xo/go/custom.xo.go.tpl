{{/* custom template for Emoine_R */}}

{{/* all: get all rows */}}
{{/* use in typedef func */}}
{{ define "all" }}
{{- $t := .Data -}}
// {{ $t.GoName }}s retrieves all rows from '{{ schema $t.SQLName }}' as a [{{ $t.GoName }}].
func {{ $t.GoName }}s(ctx context.Context, db DB, limit int, offset int) ([]{{ $t.GoName }}, error) {
	// query
	const sqlstr = `SELECT ` +
		`{{ range $i, $f := $t.Fields }}{{ if $i }}, {{ end }}{{ $f.SQLName }}{{ end }} ` +
		`FROM {{ schema $t.SQLName }} ` +
		`LIMIT ? OFFSET ?`
	// run
	logf(sqlstr, limit, offset)

	rows, err := db.QueryContext(ctx, sqlstr, limit, offset)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()
	// process
	var res []{{ $t.GoName }}
	for rows.Next() {
		{{ short $t }} := {{ $t.GoName }}{
		{{- if $t.PrimaryKeys }}
			_exists: true,
		{{ end -}}
		}
		// scan
		if err := rows.Scan({{ names_ignore (print "&" (short $t) ".") $t }}); err != nil {
			return nil, logerror(err)
		}
		res = append(res, {{ short $t }})
	}
	if err := rows.Err(); err != nil {
		return nil, logerror(err)
	}
	return res, nil
}
{{ end }}

{{/* count: get count of rows */}}
{{/* use in typedef func */}}
{{ define "count" }}
{{- $t := .Data -}}
// {{ $t.GoName }}Count retrieves the number of rows in '{{ schema $t.SQLName }}'.
func {{ $t.GoName }}Count(ctx context.Context, db DB) (int, error) {
	// query
	const sqlstr = `SELECT COUNT(*) FROM {{ schema $t.SQLName }}`
	// run
	logf(sqlstr)

	var count int
	if err := db.QueryRowContext(ctx, sqlstr).Scan(&count); err != nil {
		return 0, logerror(err)
	}
	return count, nil
}
{{ end }}
