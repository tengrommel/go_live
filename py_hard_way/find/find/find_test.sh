#!/usr/bin/env bash
./bin/find . ".*.py"
echo "--------"
./bin/find tests ".*.pyc"
echo "--------"
./bin/find . "tests/tests_*"
