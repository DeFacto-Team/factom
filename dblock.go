package factom

import (
	"fmt"

	"github.com/FactomProject/factomd/wsapi"
)

/*
func GetDBlockHeight() (int, error) {
	resp, err := http.Get(
		fmt.Sprintf("http://%s/v1/directory-block-height/", server))
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}
	if resp.StatusCode != 200 {
		return 0, fmt.Errorf(string(body))
	}
	type dbh struct {
		Height int
	}
	d := new(dbh)
	if err := json.Unmarshal(body, d); err != nil {
		return 0, fmt.Errorf("%s: %s\n", err, body)
	}

	return d.Height, nil
}*/
/*
type DBlock struct {
	DBHash string
	Header struct {
		PrevBlockKeyMR string
		Timestamp      uint64
		SequenceNumber int
	}
	EntryBlockList []struct {
		ChainID string
		KeyMR   string
	}
}

*/

func GetDBlock(keymr string) (*wsapi.DirectoryBlockResponse, error) {
	resp, err := CallV2("directory-block-by-keymr", false, keymr)
	if err != nil {
		return nil, err
	}

	if resp.Error != nil {
		return nil, fmt.Errorf(resp.Error.Message)
	}

	return resp.Result.(*wsapi.DirectoryBlockResponse), nil
}

func GetDBlockHead() (string, error) {
	resp, err := CallV2("directory-block-head", false, nil)
	if err != nil {
		return "", err
	}

	if resp.Error != nil {
		return "", fmt.Errorf(resp.Error.Message)
	}

	return resp.Result.(*wsapi.DirectoryBlockHeadResponse).KeyMR, nil
}
