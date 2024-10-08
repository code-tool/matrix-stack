{{/*
Selector labels
*/}}
{{- define "sliding-sync-proxy.selectorLabels" -}}
app: synapse
component: sliding-sync-proxy
{{- end }}

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
checksum/config: {{ include (print $.Template.BasePath "/synapse-configmap.yaml") $ | sha256sum }}
{{- end }}

{{/*
Selector labels
*/}}
{{- define "synapse-workers.selectorLabels" -}}
app: synapse
component: synapse-{{ . }}
{{- end }}

{{/*
Workers containers
*/}}
{{- define "synapse-workers.containers" -}}
containers:
- name: synapse
  image: {{ .image.repository }}:{{ .image.tag }}
  imagePullPolicy: {{ .image.pullPolicy }}
  resources: {{ .options.resources | default .resourcesDefaults | toYaml | nindent 4 }}
  {{- if has .worker (list "event_persister" "federation_sender" "client_reader" "event_creator" "account_data" "presence" "receipts" "keys" "typing" "background_worker" "pusher" "to_device") }}
  env:
  - name: "SYNAPSE_WORKER"
    value: "synapse.app.generic_worker"
  {{- end }}
  {{- if has .worker (list "media_repository" "media_repository_background_jobs") }}
  env:
  - name: "SYNAPSE_WORKER"
    value: "synapse.app.media_repository"
  {{- end }}
  ports:
    - containerPort: 8008
      name: http
      protocol: TCP
    - containerPort: 9092
      name: metrics
      protocol: TCP
  volumeMounts:
  - name: synapse-{{ .name }}-config
    mountPath: /data
terminationGracePeriodSeconds: 10
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
- name: synapse-{{ .name }}-config
  configMap:
    name: synapse-{{ .name }}-config
{{- end }}
