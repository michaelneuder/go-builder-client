package v2_test

import (
	"encoding/json"
	"fmt"
	"testing"

	v2 "github.com/attestantio/go-builder-client/api/v2"
	"github.com/stretchr/testify/require"
)

func TestSubmitBlockHeaderRequestSSZ(t *testing.T) {
	tests := []struct {
		name  string
		input []byte
		err   string
	}{
		{
			name:  "ssz header message",
			input: []byte(`{"message":{"slot":"1","parent_hash":"0xae694782c41219774f46891c6365243b97d63d66aeb7827023b8336161615652","block_hash":"0x6462e48cff39c6e4e02e5fe1aa97bf03b23a1aa588f07cfd6296d2b9bb909ce4","builder_pubkey":"0x8efc1675ffb449abc00a6ad8a2808cdf798d96fbb979cf00956012f3983577c9afe69495411a89385421f1cff47dfc98","proposer_pubkey":"0xb7da036d8aedf726e2b3439f95bdf0e68519bb55ab83d5d97a70a5b8510f612ad45a6ecc58b8b5b9b09c6b445491a02b","proposer_fee_recipient":"0x9427A30991170f917d7b83dEf6e44d26577871Ed","gas_limit":"30000000","gas_used":"7675443","value":"22135875749231725"},"execution_payload_header":{"parent_hash":"0x17f4eeae822cc81533016678413443b95e34517e67f12b4a3a92ff6b66f972ef","fee_recipient":"0x58E809C71e4885cB7B3f1D5c793AB04eD239d779","state_root":"0x3d6e230e6eceb8f3db582777b1500b8b31b9d268339e7b32bba8d6f1311b211d","receipts_root":"0xea760203509bdde017a506b12c825976d12b04db7bce9eca9e1ed007056a3f36","logs_bloom":"0x0c803a8d3c6642adee3185bd914c599317d96487831dabda82461f65700b2528781bdadf785664f9d8b11c4ee1139dfeb056125d2abd67e379cabc6d58f1c3ea304b97cf17fcd8a4c53f4dedeaa041acce062fc8fbc88ffc111577db4a936378749f2fd82b4bfcb880821dd5cbefee984bc1ad116096a64a44a2aac8a1791a7ad3a53d91c584ac69a8973daed6daee4432a198c9935fa0e5c2a4a6ca78b821a5b046e571a5c0961f469d40e429066755fec611afe25b560db07f989933556ce0cea4070ca47677b007b4b9857fc092625f82c84526737dc98e173e34fe6e4d0f1a400fd994298b7c2fa8187331c333c415f0499836ff0eed5c762bf570e67b44","prev_randao":"0x76ff751467270668df463600d26dba58297a986e649bac84ea856712d4779c00","block_number":"2983837628677007840","gas_limit":"6738255228996962210","gas_used":"5573520557770513197","timestamp":"1744720080366521389","extra_data":"0xc648a0a0a0a0a0a0a0a0a0a0a0a0a0a0c648a0a0a0a0a0a0a0a0a0a0a0a0a0a0","base_fee_per_gas":"88770397543877639215846057887940126737648744594802753726778414602657613619599","block_hash":"0x42c294e902bfc9884c1ce5fef156d4661bb8f0ff488bface37f18c3e7be64b0f","transactions_root":"0x8457d0eb7611a621e7a094059f087415ffcfc91714fc184a1f3c48db06b4d08b","withdrawals_root":"0x000102030405060708090a0b0c0d0e0f101112131415161718191a1b1c1d1e1f"},"signature":"0x010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101"}`),
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var res v2.SubmitBlockHeaderRequest
			err := json.Unmarshal(test.input, &res)
			require.NoError(t, err)

			out, err := res.MarshalSSZ()
			require.NoError(t, err)
			fmt.Printf("out bytes = %v\n", len(out))

			var new v2.SubmitBlockHeaderRequest
			err = new.UnmarshalSSZ(out)
			require.NoError(t, err)
		})
	}
}
