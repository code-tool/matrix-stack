ARG SYNAPSE_PKG_VER="v1.142.1"
ARG PYTHON_PKG_VER="3.13"

# stage 1 : build pip
FROM matrixdotorg/synapse:${SYNAPSE_PKG_VER} AS pipwheels
RUN apt-get update \
  && apt-get install -y --no-install-recommends \
      git \
      build-essential \
      patch \
      libpq-dev\
  && rm -rf /var/cache/apt/*
WORKDIR /wheels
RUN pip install psycopg2
RUN pip wheel --no-cache-dir --wheel-dir /wheels \
      git+https://github.com/matrix-org/synapse-s3-storage-provider.git@2dbba1020fe6fb84d79f1a23c925e508d9ba15e1 \
      git+https://gitlab.com/zerodotfive/googlesamlhandler.git@cb239e5d98bbe00465945c49661a9ca9ca134328

# stage 2 : apply workers.patch
FROM pipwheels as workers_patch
ARG PYTHON_PKG_VER
COPY workers.patch /tmp/workers.patch
RUN patch $(ls /usr/local/lib/python$PYTHON_PKG_VER/site-packages/synapse/config/workers.py) \
    < /tmp/workers.patch \
    && rm /tmp/workers.patch

# stage 3 : apply e2e_room_keys.patch
FROM pipwheels as e2e_patch
ARG PYTHON_PKG_VER
COPY e2e_room_keys.patch /tmp/e2e_room_keys.patch
RUN patch $(ls /usr/local/lib/python${PYTHON_PKG_VER}/site-packages/synapse/storage/databases/main/e2e_room_keys.py) \
    < /tmp/e2e_room_keys.patch \
    && rm /tmp/e2e_room_keys.patch

# stage 4 : base version with workers names
FROM matrixdotorg/synapse:$SYNAPSE_PKG_VER as main
ARG PYTHON_PKG_VER
RUN --mount=type=bind,from=pipwheels,source=/wheels,target=/wheels \
    pip install --no-cache-dir /wheels/*
COPY --from=workers_patch \
  /usr/local/lib/python${PYTHON_PKG_VER}/site-packages/synapse/config/workers.py \
  /usr/local/lib/python${PYTHON_PKG_VER}/site-packages/synapse/config/workers.py

# stage 5 : base version with e2e_room_keys table optimization
FROM main as e2e
ARG PYTHON_PKG_VER
COPY --from=e2e_patch \
  /usr/local/lib/python${PYTHON_PKG_VER}/site-packages/synapse/storage/databases/main/e2e_room_keys.py \
  /usr/local/lib/python${PYTHON_PKG_VER}/site-packages/synapse/storage/databases/main/e2e_room_keys.py
