#!/bin/bash
cat a.json | gpg -e -r $USER > a.json.gpg
