package main

import (
	"os"
	"testing"
)

func TestMain(t *testing.T) {
	solFlag = str("SimpleMultiSig.sol")
	outFlag = str("multisig_abi.go.gen")
	pkgFlag = str("geth")
	tplgoFlag = str("mobile_abi_helper.tpl")
	_ = os.Remove(*outFlag)
	main()
}

func str(s string) *string { return &s }