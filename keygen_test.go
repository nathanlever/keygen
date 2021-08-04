package keygen

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"
	"testing"
)

func TestWallets(t *testing.T) {
	var testWallets = []Wallet{
		GenerateWallet("~doznec-marbud", 1, "", 0, false),
		GenerateWallet("~marbud-tidsev-litsut-hidfep", 65012, "", 0, true),
		GenerateWallet("~wacfus-dabpex-danted-mosfep-pasrud-lavmer-nodtex-taslus-pactyp-milpub-pildeg-fornev-ralmed-dinfeb-fopbyr-sanbet-sovmyl-dozsut-mogsyx-mapwyc-sorrup-ricnec-marnys-lignex", 222, "froot loops", 6, false),
		GenerateWallet("~doznec-marbud", 0, "", 0, false),
		GenerateWallet("~doznec-marbud", 0x00ffffff, "", 0, false),
	}
	for i := 0; i < 5; i++ {
		assetWalletData, err := os.ReadFile("assets/wallet" + fmt.Sprint(i) + ".json")
		if err != nil {
			panic(err)
		}
		var assetWallet Wallet
		err = json.Unmarshal(assetWalletData, &assetWallet)
		if err != nil {
			panic(err)
		}

		assetWallet.Meta.Generator = Generator{}
		testWallets[i].Meta.Generator = Generator{}

		if !reflect.DeepEqual(assetWallet, testWallets[i]) {
			t.Error("error at wallet" + fmt.Sprint(i) + ".json")
			t.Logf("%#v\n", assetWallet)
			t.Logf("%#v\n", testWallets[i])
		}
	}
}
