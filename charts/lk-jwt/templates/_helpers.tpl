{{/*
Expand the name of the chart.
*/}}
{{- define "element-call.name" -}}
{{- default .Chart.Name .Values.nameOverride | trunc 63 | trimSuffix "-" }}
{{- end }}

{{/*
Create a default fully qualified app name.
We truncate at 63 chars because some Kubernetes name fields are limited to this (by the DNS naming spec).
If release name contains chart name it will be used as a full name.
*/}}
{{- define "element-call.fullname" -}}
{{- if .Values.fullnameOverride }}
{{- .Values.fullnameOverride | trunc 63 | trimSuffix "-" }}
{{- else }}
{{- $name := default .Chart.Name .Values.nameOverride }}
{{- if contains $name .Release.Name }}
{{- .Release.Name | trunc 63 | trimSuffix "-" }}
{{- else }}
{{- printf "%s-%s" .Release.Name $name | trunc 63 | trimSuffix "-" }}
{{- end }}
{{- end }}
{{- end }}

{{/*
Create chart name and version as used by the chart label.
*/}}
{{- define "element-call.chart" -}}
{{- printf "%s-%s" .Chart.Name .Chart.Version | replace "+" "_" | trunc 63 | trimSuffix "-" }}
{{- end }}

{{/*
Common labels
*/}}
{{- define "element-call.labels" -}}
helm.sh/chart: {{ include "element-call.chart" . }}
{{ include "element-call.selectorLabels" . }}
{{- if .Chart.AppVersion }}
app.kubernetes.io/version: {{ .Chart.AppVersion | quote }}
{{- end }}
app.kubernetes.io/managed-by: {{ .Release.Service }}
{{- end }}

{{/*
Selector labels
*/}}
{{- define "element-call.selectorLabels" -}}
app.kubernetes.io/name: {{ include "element-call.name" . }}
app.kubernetes.io/instance: {{ .Release.Name }}
{{- end }}

{{/*
Create the name of the service account to use
*/}}
{{- define "element-call.serviceAccountName" -}}
{{- $ := get . "root" }}
{{- $suffix := get . "suffix" }}
{{- with get . "ctx" }}
{{- if .serviceAccount.create }}
{{- if $suffix }}
{{- default (printf "%s-%s" (include "element-call.fullname" $) $suffix) .serviceAccount.name }}
{{- else }}
{{- default (include "element-call.fullname" $) .serviceAccount.name }}
{{- end }}
{{- else }}
{{- default "default" .serviceAccount.name }}
{{- end }}
{{- end }}
{{- end }}
