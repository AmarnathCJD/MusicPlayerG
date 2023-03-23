package main

import (
	"log"
	_ "main/server"
	"main/util"
	"net/http"
)

func main() {
	log.Fatal(http.ListenAndServe(util.GetSpecifiedPort(), nil))
}

// Copyright (c) 2023 by RoseLoverX. All rights reserved.
