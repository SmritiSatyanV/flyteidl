DIR=`pwd`
rm -rf $DIR/gen
LYFT_IMAGE="lyft/protocgenerator:8d6ef5ddf4b858a90a53044a1b2dc7b034e83c88"

docker run -v $DIR:/defs $LYFT_IMAGE -i ./protos -d protos/flyteidl/service --with_gateway -l go --go_source_relative
docker run -v $DIR:/defs $LYFT_IMAGE -i ./protos -d protos/flyteidl/admin --with_gateway -l go --go_source_relative --validate_out
docker run -v $DIR:/defs $LYFT_IMAGE -i ./protos -d protos/flyteidl/core --with_gateway -l go --go_source_relative --validate_out
docker run -v $DIR:/defs $LYFT_IMAGE -i ./protos -d protos/flyteidl/event --with_gateway -l go --go_source_relative --validate_out
docker run -v $DIR:/defs $LYFT_IMAGE -i ./protos -d protos/flyteidl/plugins -l go --go_source_relative --validate_out
docker run -v $DIR:/defs $LYFT_IMAGE -i ./protos -d protos/flyteidl/legacydiscovery -l go --go_source_relative --validate_out

docker run -v $DIR:/defs $LYFT_IMAGE -i ./protos -d protos/flyteidl/service -l python
docker run -v $DIR:/defs $LYFT_IMAGE -i ./protos -d protos/flyteidl/admin -l python
docker run -v $DIR:/defs $LYFT_IMAGE -i ./protos -d protos/flyteidl/core -l python
docker run -v $DIR:/defs $LYFT_IMAGE -i ./protos -d protos/flyteidl/event -l python
docker run -v $DIR:/defs $LYFT_IMAGE -i ./protos -d protos/flyteidl/plugins -l python

# Docs generated
docker run -e REPO_BLOB_SHA=master -e PROJECT_ANNOTATION_PREFIX=flyte.interface -v $DIR:/defs $LYFT_IMAGE -i ./protos -d protos/flyteidl/service -l protodoc
docker run -e REPO_BLOB_SHA=master -e PROJECT_ANNOTATION_PREFIX=flyte.interface -v $DIR:/defs $LYFT_IMAGE -i ./protos -d protos/flyteidl/admin -l protodoc
docker run -e REPO_BLOB_SHA=master -e PROJECT_ANNOTATION_PREFIX=flyte.interface -v $DIR:/defs $LYFT_IMAGE -i ./protos -d protos/flyteidl/core -l protodoc
docker run -e REPO_BLOB_SHA=master -e PROJECT_ANNOTATION_PREFIX=flyte.interface -v $DIR:/defs $LYFT_IMAGE -i ./protos -d protos/flyteidl/event -l protodoc
docker run -e REPO_BLOB_SHA=master -e PROJECT_ANNOTATION_PREFIX=flyte.interface -v $DIR:/defs $LYFT_IMAGE -i ./protos -d protos/flyteidl/plugins -l protodoc

# Generate JS code
docker run -v $DIR:/defs schottra/docker-protobufjs:latest --module-name flyteidl -d protos/flyteidl/core  -d protos/flyteidl/event -d protos/flyteidl/admin -d protos/flyteidl/service  -- --root flyteidl -t static-module -w es6 --no-delimited --force-long --no-convert -p /defs/protos

# Unfortunately, the `--grpc-gateway-out` plugin doesn’t yet support the `source_relative` option. Until it does, we need to move the files from the autogenerated location to the source_relative location.
sudo cp -r gen/pb-go/github.com/lyft/flyteidl/gen/* gen/
sudo rm -rf gen/pb-go/github.com

# Copy the validate.py protos.
sudo mkdir -p gen/pb_python/validate
sudo cp -r validate/* gen/pb_python/validate/

if [ -n "$DELTA_CHECK" ]; then
  DIRTY=$(git status --porcelain)
  if [ -n "$DIRTY" ]; then
    echo "FAILED: Protos updated without commiting generated code."
    echo "Ensure make generate has run and all changes are committed."
    DIFF=$(git diff)
    echo "dif detected: $DIFF"
    exit 1
  else
    echo "SUCCESS: Generated code is up to date."
  fi
fi
