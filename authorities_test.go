// Copyright 2016 Factom Foundation
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package factom_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"

	. "github.com/FactomProject/factom"

	"testing"
)

func TestUnmarshalAuthorities(t *testing.T) {
	js := []byte(`{"authorities":[{"chainid":"8888881541fc5bc1bcc0597e71eed5df7de8d47c8eb97f867d16ebb20781f38b","manageid":"0000000000000000000000000000000000000000000000000000000000000000","matroyshka":"1ce468172d6408643a8931838a935733f6fa97d02a8b44a741a1376da8829152","signingkey":"34ffc2a7f6e35e503fd2d4259113d4d9b131e8e56d63a1c277ab5064d58d9826","status":"audit","anchorkeys":[{"blockchain":"BTC","level":0,"keytype":0,"key":"010c53bd5e4a863cf8e7df48f567e3f2e492aba9"}]},{"chainid":"8888889585051d7117d217a55a366d56826eda35c951f02428b976524dbfc7f9","manageid":"0000000000000000000000000000000000000000000000000000000000000000","matroyshka":"914ab0fd1905f3ef19e54f94dd3caee1055793eb8cd5ce7f982cd15ea393bcd7","signingkey":"2001c69d076a5bf43335d41f49ad7626f1d79d8e1dfe9d9f9c8cc9a0d99efd5b","status":"audit","anchorkeys":[{"blockchain":"BTC","level":0,"keytype":0,"key":"c568a1206e29c7c8fed15aee12515833434b4eb4"}]},{"chainid":"888888a5ce32a3a257c1ff29033b6c01dd20239e7c68ebaf06d690e4ae2b7e83","manageid":"0000000000000000000000000000000000000000000000000000000000000000","matroyshka":"611fb3b711629ee6964f6e6d7a7a389ab275b4b14c8eafaaa72930f2b9c12303","signingkey":"13d42208f7a7699c7976dc19424872268e503779850fb72aecae4b5341dd40c7","status":"audit","anchorkeys":[{"blockchain":"BTC","level":0,"keytype":0,"key":"412945af7b4ec2ff17285b22631be19f3201d572"}]},{"chainid":"888888bf5e39211db27b2d2b1b57606b4d68cf57e908971949a233d8eb734156","manageid":"0000000000000000000000000000000000000000000000000000000000000000","matroyshka":"002762ccf5948b8e1c29a9c3df4748cf3efe6567eb3046e6361f353079e55344","signingkey":"646f6bf2eaa80a803f1ffd3286945c4d6ddfdf5974177a52141c6906153f5237","status":"federated","anchorkeys":[{"blockchain":"BTC","level":0,"keytype":0,"key":"7c6b5121835d148932c75ce773208ffc17a4144f"}]},{"chainid":"888888c1fd1cf7ca3e0c4e2e9a6462aa8ac4e537563ee54ff41eb6b617a1ec37","manageid":"0000000000000000000000000000000000000000000000000000000000000000","matroyshka":"96fa0827f28ced76f18e42b8ef836d96c5c5adde4b8c98a406ad006109985628","signingkey":"b9a4837383cf11d818f1c1931f5586f840967fe0931d9b733394f75bf39fcd17","status":"federated","anchorkeys":[{"blockchain":"BTC","level":0,"keytype":0,"key":"1f605e0d687dbb731e6961cdf8c30e24195889d0"}]},{"chainid":"8888886ff14cef50365b785eb3cefab5bc30175d022be06ed412391a82645376","manageid":"0000000000000000000000000000000000000000000000000000000000000000","matroyshka":"fe21f1320ff7eaaab9ceb9551833078ab79b5b0dfe86097a88ca26d74e48b354","signingkey":"0d6a22b9bf17851c830189fb324ba7d1ea8d6a15eea3adf671109825a1332147","status":"audit","anchorkeys":[{"blockchain":"BTC","level":0,"keytype":0,"key":"b4db03e03da3555f630aef3900897e67247c8477"}]},{"chainid":"888888a8da713519881065d90f73f498b36d956e3390c5a6c06747922395075f","manageid":"0000000000000000000000000000000000000000000000000000000000000000","matroyshka":"36108e2fd7ba67a25886c14408db1bc2a1d0098a23f2b64e4734ff80b772def0","signingkey":"ffb9efd4d490535e3b5041622354f5c440524b0d1976582e0c9ba6cb1649279b","status":"audit","anchorkeys":[{"blockchain":"BTC","level":0,"keytype":0,"key":"3d5ffebea388ce494cd7d24ff03165117561ef90"}]},{"chainid":"888888b4eecb6868615e1875120e855529b4e372e2887cdec7185b46abfcfb35","manageid":"0000000000000000000000000000000000000000000000000000000000000000","matroyshka":"86400145400bf22a717d1bd4fc7f15e5de2872d21e815bc0a4916c15de2e6eb7","signingkey":"c2bbab9d274415765eae5c3ee3b94ff3c38dd5c9b02c8f842e2770a6de0b5068","status":"federated","anchorkeys":[{"blockchain":"BTC","level":0,"keytype":0,"key":"e0e135c1ee0c2131b2dac5fcb353863ac21fff62"}]},{"chainid":"888888dda15d7ad44c3286d66cc4f82e6fc07ed88de4d13ac9a182199593cac1","manageid":"0000000000000000000000000000000000000000000000000000000000000000","matroyshka":"7c45e29fd0c7e09428e7ea60ed5042e8a0d6a091cc576e255eb10b7e899d3c03","signingkey":"07f339e556ee999cc7e33500753ea0933381b09f5c2bca26e224d716e61a8862","status":"audit","anchorkeys":[{"blockchain":"BTC","level":0,"keytype":0,"key":"6788c85b7963c8527900a2a2ad2c24d15f347d89"}]},{"chainid":"8888882fa588e8ad6e73555a9b9ff3d84b468601b81328ec09d91051369d7373","manageid":"0000000000000000000000000000000000000000000000000000000000000000","matroyshka":"a5f91355b6c8a1a9b38d378434886caea05cc73e544416ec4c9b7f219f23c497","signingkey":"296d08be4a741d6c328ab47d80a55590dceef6550066a0a76e4816a3f51eefee","status":"federated","anchorkeys":[{"blockchain":"BTC","level":0,"keytype":0,"key":"850fd39e1841b29c12f4ace379380a467489dba8"}]},{"chainid":"88888870cf06fb3ed94af3ba917dbe9fab73391915c57812bd6606e6ced76d53","manageid":"0000000000000000000000000000000000000000000000000000000000000000","matroyshka":"151253cf6f9ad8db3f1bd7116a6ec894851fff4268ad1c14fe3ce8f3933a9b08","signingkey":"5413e626ce80d90276b5b2388d13f4a4dce2faffce6bb76b9290fcd11dd700dc","status":"audit","anchorkeys":[{"blockchain":"BTC","level":0,"keytype":0,"key":"80b560002d85154fa1c255531c232f84b4293c86"}]},{"chainid":"888888b2ddad8c24fdf3033bdf6bd9c393ffe074b6f5d5741c84afea27c1656e","manageid":"0000000000000000000000000000000000000000000000000000000000000000","matroyshka":"74055ead8eb83d34515c66bb7824dfda3659e1193dd31f6f38eed6e2cdc4e592","signingkey":"b11d2c22e96af34946810c816ada60a7027ed3d7c98aac72283ed348fc58cf73","status":"federated","anchorkeys":[{"blockchain":"BTC","level":0,"keytype":0,"key":"72f4aa05adc0b5284602bd744858106c618b932e"}]},{"chainid":"888888f05308313f6e8f5619cacfb32e0dcba25b4741de9c0fc3b127e8ba2a6b","manageid":"0000000000000000000000000000000000000000000000000000000000000000","matroyshka":"8fcab189bbb2f97249d05b0b31adeaef23b7aaca326673e16fc901022f8285c8","signingkey":"6ceeb261cc19b14f6c89bb0bd937f195ffc9e6adaa5618e432752b01a00792c7","status":"federated","anchorkeys":[{"blockchain":"BTC","level":0,"keytype":0,"key":"57b3621913fd321c4c4f07cef3468bf04b0baf59"}]},{"chainid":"88888841ac82c501a300def3e95d724b4b5e31f729f3b6d9d9736dca0f0edc34","manageid":"0000000000000000000000000000000000000000000000000000000000000000","matroyshka":"52103541ebcd32f5a55dc3c5037fd6396bbe3d65d22f8c06026a9ad97440d8cd","signingkey":"667a53519cab0365d1a1ac625b6cd64d86695e8ae38d280ea6d3dbe8191acf34","status":"audit","anchorkeys":[{"blockchain":"BTC","level":0,"keytype":0,"key":"5ba2689c372fdf712e477a83059a5da313e07bf0"}]},{"chainid":"8888884a0acbf1a23e3291b99681b80a91ca51914d64e39de65645868e0b4714","manageid":"0000000000000000000000000000000000000000000000000000000000000000","matroyshka":"35b100ead1d81fe3a3e6b1a656c127b14a2ef9d520adec6ea0d7b9d1d5488268","signingkey":"93f6aca96b011fc31fd655fee9556b459509308eaaa63c02e9ebff8f384c72e0","status":"audit","anchorkeys":[{"blockchain":"BTC","level":0,"keytype":0,"key":"58e737d93cb52102d78ee7b918bd33a4412f901e"}]},{"chainid":"8888886043746fe47dcf55952b20d8aee6ae842024010fd3f42bc0076a502f42","manageid":"0000000000000000000000000000000000000000000000000000000000000000","matroyshka":"b566f30f2013dc3cf7960268da70efb76534ce710f270c1b3ae08781f9faae1b","signingkey":"847ef7a9d15df05940a97030a7b783fad54622bdb81f5698f948b94e127eb6e5","status":"federated","anchorkeys":[{"blockchain":"BTC","level":0,"keytype":0,"key":"e2a977f66a529d3746727f390c429298f6daef68"}]},{"chainid":"888888655866a003faabd999c7b0a7c908af17d63fd2ac2951dc99e1ad2a14f4","manageid":"0000000000000000000000000000000000000000000000000000000000000000","matroyshka":"86e2f9073dfafb461888955166c12c6b1d9aa98504af1cccb08f0ad53fbbb666","signingkey":"f8139f98fadc948b254d0dea29c55fab7fa14f1fd97ef78ef7bb99d2d82bd6f1","status":"federated","anchorkeys":[{"blockchain":"BTC","level":0,"keytype":0,"key":"5bf09c36ebb93643acf41e716261357583ee7281"}]},{"chainid":"888888b1255ea1cc0b3ab3b9425d3239643ae1f3c00ec634690eda784f05bda7","manageid":"0000000000000000000000000000000000000000000000000000000000000000","matroyshka":"1cbf54effa547cf89751e3a02d8980ea7e9325e591ff8f1d360bbe323da8fa5a","signingkey":"e3b88b704533612f69b5d6390737481694d7d8acb71e532cac3e8dd2d11ca691","status":"federated","anchorkeys":[{"blockchain":"BTC","level":0,"keytype":0,"key":"dcb4dcd7e5a518854eadd0ec48955101d9fbac35"}]}]}`)

	ret := new(struct {
		Authorities []Authority `json:"authorities"`
	})
	err := json.Unmarshal(js, ret)
	if err != nil {
		t.Error(err)
	}
	for _, a := range ret.Authorities {
		t.Log(a)
	}
}

func TestGetAuthorities(t *testing.T) {
	factomdResponse := `{
	    "jsonrpc": "2.0",
	    "id": 1,
	    "result": {
	        "Authorities": [{
	            "chainid": "8888883a40c004ba51834dd2599f271b30e3251180295f099886754d1b993667",
	            "manageid": "888888705202c41e75e22fc726933c2f6c74e43cc349e1c971e2672d0af74ecb",
	            "matroyshka": "c007809e5f0e1497dfc8f0d0a1e7bc130639e325e4832b5d2353638b413287bc",
	            "signingkey": "b6dae874df4eb179426afdc822a57190a9dadbfc84c3cd69cd5e7d05a2c3190e",
	            "status": "federated",
	            "anchorkeys": [{
	                "blockchain": "BTC",
	                "level": 0,
	                "keytype": 0,
	                "key": "3f55a831a82c408da54faed323af2859bb5d7b1f"
	            }]
	        }, {
	            "chainid": "8888887529d62b6d3d702bafb06f11ef825ec2fd54c978c1e1809a7eedba1514",
	            "manageid": "8888883c4446520dbaa80ff0f637e223a70a56c0dd91cb12d8bd35e0a9ce6659",
	            "matroyshka": "3f24854527710e9f3d47b92ac332982c2a34922f4409e13754b8779a07acd2e5",
	            "signingkey": "58d076db196351954d2bc717a8518bc5cd311e0496cb7ea9bec18ead8d553faf",
	            "status": "federated",
	            "anchorkeys": [{
	                "blockchain": "BTC",
	                "level": 0,
	                "keytype": 0,
	                "key": "f0c4c31c826c724405a0bcf1788327f86257751e"
	            }]
	        }, {
	            "chainid": "8888887f5125bfc597a05eca2db64298b88a9233dafdeb44bc0db7d55ee035aa",
	            "manageid": "88888873b24be7aa3cfc281b5d391e4619346699f308983533692c98d755c35f",
	            "matroyshka": "c19b16d9d9eb4b0feb9520e0bfa85b7fd8aa91b3bb9ee6f3a7a9e82201911329",
	            "signingkey": "8fbbfd86dddc384095e5792c011572511b338943b1d1d3b4697e8ead47afaa28",
	            "status": "federated",
	            "anchorkeys": [{
	                "blockchain": "BTC",
	                "level": 0,
	                "keytype": 0,
	                "key": "2bb2499aad2182db63c2a0a14e6f3b5310ef82b8"
	            }]
	        }, {
	            "chainid": "888888b4a7105ec297396592b8e9448b504a8fb41b82ee26e23068ff0e4549d0",
	            "manageid": "888888686e9eab2d3f919d641a17a0d7befb352a5ff0ab40058423c36de77a7c",
	            "matroyshka": "699c3a0e21b0b962ef2f7ff9a7669b548f8b260264fd3e19b3c384233e2a9143",
	            "signingkey": "93b409885d6d205e53d436fdd64c97bfe903167a2d733554d13073d5d57e5755",
	            "status": "federated",
	            "anchorkeys": [{
	                "blockchain": "BTC",
	                "level": 0,
	                "keytype": 0,
	                "key": "d09ff1c371b5ff6cc8cdffdf77217de7d8bfbdcd"
	            }]
	        }, {
	            "chainid": "888888f4d59308deaa587498e5e1c4e0228a190eba50c9ad23b604da1cbd8c77",
	            "manageid": "88888845e79eee6709318fdd47ed42f455e3438fb48e14c05df0736f34fbe3d1",
	            "matroyshka": "c9f7b68fe5e7867fed0545fce27779cd16e540d0af5ac046606f609a85808b1b",
	            "signingkey": "cb58d32e11c5dd07f37c3780307ed45db672afd72b98874f8ea6bb9bee36dd77",
	            "status": "federated",
	            "anchorkeys": [{
	                "blockchain": "BTC",
	                "level": 0,
	                "keytype": 0,
	                "key": "7f4e785f9f543f7f220eb27e61fa8a699a74509a"
	            }]
	        }, {
				"chainid": "8888889e4fbbcc0032e6a2ce517d39fc90cce1189a46d7cebfff4b8bc230744c",
	            "manageid": "88888876417375a4d4121a65fd20519e92b80f0f2b85d61c14b67b419214b690",
	            "matroyshka": "6b11882219b473b087721c758390af8ba24e8fbb2699461e8b9edbd178c97dff",
	            "signingkey": "333b82e2e71cdf4616dda9dfc64486018d3277f45f9a1010bc2e9d849f656bcf",
	            "status": "audit",
	            "anchorkeys": [{
	                "blockchain": "BTC",
	                "level": 0,
	                "keytype": 0,
	                "key": "59e285ec55f90783db1362494f166ffd60314d5e"
	            }]
	        }]
	    }
	}`
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, factomdResponse)
	}))
	defer ts.Close()

	SetFactomdServer(ts.URL[7:])

	as, err := GetAuthorities()
	if err != nil {
		t.Error(err)
	}
	t.Log(as)
}
