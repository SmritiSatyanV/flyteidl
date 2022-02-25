# Flyte IDL

This is one of the core repositories of Flyte that contains specification of the Flyte Lanugage using protobuf messages, the Backend API specification in gRPC and Swagger REST. The repo contains generated clients and protocol message structures in multiple languages. In addition to the generated code, the repository also contains the Golang clients for Flyte's backend APIs (the services grouped under FlyteAdmin).

[![Slack](https://img.shields.io/badge/slack-join_chat-white.svg?logo=slack&style=social)](https://slack.flyte.org)

* [flyte.org](https://flyte.org)
* [Flyte Docs](http://docs.flyte.org)
* [FlyteIDL API reference documentation](https://docs.flyte.org/projects/flyteidl/en/stable/index.html)

## Contributing to FlyteIDL

## Tooling for FlyteIDL

1. Run ``make download_tooling`` to install generator dependencies

```bash
   make download_tooling
```

2. Make sure Docker is installed locally.
3. Once installed, run ``make generate`` to generate all the code and mock client for FlyteAdmin Service and the docs for it.

```bash
    make generate
```

4. To add new dependencies for documentation generation, modify ``doc-requirements.in`` and run

```bash
   make doc-requirements.txt
```

## Docs structure

The index.rst files for protos are kept in a parallel folder structure under the docs folder.
All the proto definitions are within protos/flyteidl and there corresponding docs are kept in protos/docs.

```
docs
├── admin
│   ├── admin.rst
│   └── index.rst
├── core
│   ├── core.rst
│   └── index.rst
├── datacatalog
│   ├── datacatalog.rst
│   └── index.rst
├── event
│   ├── event.rst
│   └── index.rst
├── plugins
│   ├── index.rst
│   └── plugins.rst
├── service
│   ├── index.rst
│   └── service.rst
```

Each module in protos has a module in docs with the same name.
E.g., protos/flyteidl/core has its corresponding docs in protos/docs/core.


## Docs Generation

* If a new module is added, follow the structure for core files in `generate_protos.sh` file which helps generate the core documentation from its proto files.
```
     core_proto_files=`ls protos/flyteidl/core/*.proto |xargs`
     # Remove any currently generated file
     ls -d protos/docs/core/* | grep -v index.rst | xargs rm
     protoc --doc_out=protos/docs/core --doc_opt=restructuredtext,core.rst -I=protos `echo $core_proto_files`
```

* ``make generate`` would have already generated the modified rst files.

* ``make html`` Generates the Sphinx documentation from the docs folder.

