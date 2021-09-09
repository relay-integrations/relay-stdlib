#!/bin/bash
set -euo pipefail

MESSAGE=$(ni get -p '{.message}')
echo "${MESSAGE}"
