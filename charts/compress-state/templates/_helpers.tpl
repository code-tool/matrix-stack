{{- define "synapse-compress-state.fullname" -}}
{{ printf "%s-%s" .Release.Name "compress-state" }}
{{- end }}
