package main_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/Null-nil/cosmos-sdk/client/flags"
	svrcmd "github.com/Null-nil/cosmos-sdk/server/cmd"
	"github.com/Null-nil/cosmos-sdk/x/genutil/client/cli"

	"github.com/Null-nil/ethermint/app"
	ethermintd "github.com/Null-nil/ethermint/cmd/ethermintd"
)

func TestInitCmd(t *testing.T) {
	rootCmd, _ := ethermintd.NewRootCmd()
	rootCmd.SetArgs([]string{
		"init",          // Test the init cmd
		"etherminttest", // Moniker
		fmt.Sprintf("--%s=%s", cli.FlagOverwrite, "true"), // Overwrite genesis.json, in case it already exists
		fmt.Sprintf("--%s=%s", flags.FlagChainID, "ethermint_9000-1"),
	})

	err := svrcmd.Execute(rootCmd, "", app.DefaultNodeHome)
	require.NoError(t, err)
}
