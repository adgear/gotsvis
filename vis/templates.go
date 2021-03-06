package vis

//go:generate include_suffix -package=vis -folder=. -output=templates.go -suffix=tmpl

const(
buttonTMPL = `{{ define "hide show button" }}
<input type='button' id='hideshow' value='hide/show' onclick='hide_show("#{{ jq .Key}}")'>
{{ end }}

{{ define "hide show script" }}
<script>
    function hide_show(id) {
        console.log(this);
        console.log(id);

        $(id).toggle('show');
    }
</script>
{{ end }}
`
chartTMPL = `{{ define "chart" }}
{{ template "hide show button" . }} {{ TimeSeriesTagJS . }}
<div id="{{ TimeSeriesTagJS . }}"></div>
<script type="text/javascript">
    c3.generate({
        bindto: "#{{ jq .Key }}",
        data: {
            x: 'x',
            xFormat: '%%Y%%m%%d %%H:%%M:%%S',
            columns: [
                {{ Chart . }}
            ]
        },
        axis: {
            x: {
                type: 'timeseries',
                tick: {
                    format: '%%Y-%%m-%%d %%H:%%M:%%S',
                    count: 12
                }
            }
        }
    });
</script>
{{ end }}
`
linksTMPL = `
{{ define "c3_cdn" }}
<link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/c3/0.4.10/c3.min.css">
<script src="https://cdnjs.cloudflare.com/ajax/libs/d3/3.5.12/d3.min.js"></script>
<script src="https://cdnjs.cloudflare.com/ajax/libs/c3/0.4.11/c3.min.js"></script>
<script src="https://code.jquery.com/jquery-2.1.4.min.js"></script>
{{ end }}

{{ define "bootstrap_cdn" }}
<link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.6/css/bootstrap.min.css" integrity="sha384-1q8mTJOASx8j1Au+a5WDVnPi2lkFfwwEAa8hDDdjZlpLegxhjVME1fgjWPGmkzs7" crossorigin="anonymous">
<link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.6/css/bootstrap-theme.min.css" integrity="sha384-fLW2N01lMqjakBkx3l/M9EahuwpSfeNvV63J5ezn3uZzapT0u7EYsXMjQV+0En5r" crossorigin="anonymous">
<script src="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.6/js/bootstrap.min.js" integrity="sha384-0mSbJDEHialfmuBBQP6A4Qrprq5OVfW37PRR3j5ELqxss1yVqOtnepnHVP9aJ7xS" crossorigin="anonymous"></script>
{{ end }}
`
)

func GetTMPL() []string {
	return []string{ buttonTMPL, chartTMPL, linksTMPL }
}