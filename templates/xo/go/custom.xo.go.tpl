{{/* custom template for Emoine_R */}}

{{/* all: use in typedef func */}}
{{ define "all" }}
{{- $t := .Data -}}
// {{ $t.GoName }}s retrieves all rows from '{{ schema $t.SQLName }}' as a [{{ $t.GoName }}].
func {{ $t.GoName }}s(ctx context.Context, db DB, limit, int, offset int) ([]{{ $t.GoName }}, error) {
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
