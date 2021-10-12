package test

import (
	"../interface/pkg/MYSQL"
	"testing"
)

func TestNames(t *testing.T)  {
	MYSQL.ExecSQLFile()
}