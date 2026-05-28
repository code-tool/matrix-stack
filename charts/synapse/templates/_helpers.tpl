{{/*
Selector labels
*/}}
{{- define "synapse-client-reader.selectorLabels" -}}
app: synapse
component: synapse-client-reader
{{- end }}

{{/*
Selector labels
*/}}
{{- define "synapse-room.selectorLabels" -}}
app: synapse
component: synapse-room
{{- end }}

{{/*
Selector labels
*/}}
{{- define "synapse-sync.selectorLabels" -}}
app: synapse
component: synapse-sync
{{- end }}

{{/*
Selector labels
*/}}
{{- define "synapse-federation-reader.selectorLabels" -}}
app: synapse
component: synapse-federation-reader
{{- end }}

{{/*
Selector labels
*/}}
{{- define "synapse-client-reader-envoy.selectorLabels" -}}
app: synapse
component: synapse-client-reader-envoy
{{- end }}

{{/*
Selector labels
*/}}
{{- define "matrix-authentication.selectorLabels" -}}
app: synapse
component: matrix-authentication
{{- end }}

{{/*
Workers annotations
*/}}
{{- define "synapse-workers.annotations" -}}
prometheus.io/port: "9092"
prometheus.io/scrape: "true"
prometheus.io/path: "/_synapse/metrics"
checksum/secret: {{ include (print $.Template.BasePath "/synapse-secret.yaml") $ | sha256sum }}
{{- end }}

{{/*
Selector labels
*/}}
{{- define "synapse-workers.selectorLabels" -}}
app: synapse
component: synapse-{{ . }}
{{- end }}

{{/*
Base helper for Synapse cache size limits.
Falls back to .fallback when no limit is configured or suffix is unrecognised.
Supported Kubernetes suffixes: Gi, Mi, G, M.
*/}}
{{- define "synapse.workerCacheMemory" -}}
{{- $resources := .options.resources | default .defaults -}}
{{- $memLimit := "" -}}
{{- if and $resources $resources.limits $resources.limits.memory -}}
  {{- $memLimit = $resources.limits.memory | toString -}}
{{- end -}}
{{- $num := .num -}}
{{- $denom := .denom -}}
{{- if $memLimit -}}
  {{- if hasSuffix "Gi" $memLimit -}}
    {{- $val := trimSuffix "Gi" $memLimit | int64 -}}
    {{- printf "%dM" (div (mul (mul $val 1024) $num) $denom) -}}
  {{- else if hasSuffix "Mi" $memLimit -}}
    {{- $val := trimSuffix "Mi" $memLimit | int64 -}}
    {{- printf "%dM" (div (mul $val $num) $denom) -}}
  {{- else if hasSuffix "G" $memLimit -}}
    {{- $val := trimSuffix "G" $memLimit | int64 -}}
    {{- printf "%dM" (div (mul (mul $val 1024) $num) $denom) -}}
  {{- else if hasSuffix "M" $memLimit -}}
    {{- $val := trimSuffix "M" $memLimit | int64 -}}
    {{- printf "%dM" (div (mul $val $num) $denom) -}}
  {{- else -}}
    {{- .fallback -}}
  {{- end -}}
{{- else -}}
  {{- .fallback -}}
{{- end -}}
{{- end -}}

{{/* max_cache_memory_usage = 90% of memory limit */}}
{{- define "synapse.workerMaxCacheMemory" -}}
{{- include "synapse.workerCacheMemory" (merge (dict "num" 9 "denom" 10) .) -}}
{{- end -}}

{{/* target_cache_memory_usage = 80% of max = 72% of memory limit */}}
{{- define "synapse.workerTargetCacheMemory" -}}
{{- include "synapse.workerCacheMemory" (merge (dict "num" 72 "denom" 100) .) -}}
{{- end -}}

{{/*
Workers containers
*/}}
{{- define "synapse-workers.containers" -}}
containers:
- name: synapse
  image: {{ .image.repository }}:{{ .image.tag }}
  imagePullPolicy: {{ .image.pullPolicy }}
  resources: {{ .options.resources | default .resourcesDefaults | toYaml | nindent 4 }}
  {{- if ne .worker "master" }}
  env:
  - name: "SYNAPSE_WORKER"
    {{- if has .worker (list "media_repository" "media_repository_background_jobs") }}
    value: "synapse.app.media_repository"
    {{- else }}
    value: "synapse.app.generic_worker"
    {{- end }}
  {{- end }}
  ports:
    - containerPort: 8008
      name: http
      protocol: TCP
    - containerPort: 9092
      name: metrics
      protocol: TCP
  {{- if not (has .worker (list "background_worker" "event_persister" "pusher")) }}
  startupProbe:
    httpGet:
      path: /health
      port: http
    failureThreshold: 180
    periodSeconds: 15
  livenessProbe:
    httpGet:
      path: /health
      port: http
    failureThreshold: 5
    periodSeconds: 15
  readinessProbe:
    httpGet:
      path: /health
      port: http
    periodSeconds: 15
  {{- end }}
  volumeMounts:
  - name: synapse-{{ .name }}-secret
    mountPath: /data
lifecycle:
  preStop:
    exec:
      command: ["sleep", "15"]
terminationGracePeriodSeconds: 90
{{- if .nodeSelector }}
nodeSelector:
  {{ toYaml .nodeSelector | nindent 2 }}
{{- end }}
{{- if .tolerations }}
tolerations:
  {{ toYaml .tolerations | nindent 2 }}
{{- end }}
{{- if .affinity }}
affinity:
  {{ toYaml .affinity | nindent 2 }}
{{- end }}
volumes:
- name: synapse-{{ .name }}-secret
  secret:
    secretName: synapse-{{ .name }}-secret
{{- end }}
