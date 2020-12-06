#!/usr/bin/env bash

## Add folders want to gen mock here
mockery --case=underscore --recursive=true --inpackage --name=I.+ --dir=domain
mockery --case=underscore --recursive=true --inpackage --name=I.+ --dir=repository