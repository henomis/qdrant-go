#!/bin/sh
docker run -p 6333:6333 --name qdrant --rm -v $(pwd)/config.yaml:/qdrant/config/production.yaml qdrant/qdrant
