FROM ghcr.io/astral-sh/uv:bookworm-slim AS exporter

WORKDIR /app
COPY pyproject.toml uv.lock ./
RUN uv export --frozen --no-dev -o requirements.txt

FROM python:3.14-slim-bookworm

WORKDIR /app

RUN groupadd --gid 1000 appgroup && \
    useradd --uid 1000 --gid 1000 --shell /bin/bash appuser

COPY --from=exporter /app/requirements.txt ./
RUN pip install --no-cache-dir -r requirements.txt

COPY --chown=appuser:appgroup main.py ./

ENV HOST=0.0.0.0
ENV PORT=8080

EXPOSE 8080

USER appuser

CMD ["python", "main.py"]
