package main

import (
  _ "embed"
)

//go:embed version.txt
var currentVersionStr string

//go:embed flaarum-logo.ico
var flaarumLogoBytes []byte
