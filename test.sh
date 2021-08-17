#!/bin/bash
curl -v --form file'=@songs/test.mp3' --header "Content-Type: multipart/form-data" localhost:8080/api/file
