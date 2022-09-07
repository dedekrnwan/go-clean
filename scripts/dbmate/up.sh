#!/bin/bash
DIR=$PWD/migrations
DATABASE_URL="postgres://postgres@127.0.0.1:5432/go-clean?sslmode=disable"
export DATABASE_URL
dbmate -d $DIR up