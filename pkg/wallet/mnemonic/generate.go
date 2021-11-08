package mnemonic

import (
	"crypto/rand"
	"crypto/sha256"
	"math/big"
	"strings"
	wallet_cli "wallet-cli/pkg/wallet"
)

type MnemonicService struct {
}

func NewService() *MnemonicService {
	return  &MnemonicService{}
}

func (s *MnemonicService) Generate() (*wallet_cli.Mnemonic, error) {

	// ---------------------------------------
	// 1a. Generate Random Bytes (for Entropy)
	// ---------------------------------------

	// Create empty byte slice and fill with bytes from crypto/rand
	bytes := make([]byte, 16) // [0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0]
	rand.Read(bytes)

	// ----------------------------------------
	// 1b. Use Hexadecimal String (for Entropy)
	// ----------------------------------------

	// entropy := "7f7f7f7f7f7f7f7f7f7f7f7f7f7f7f7f"
	// bytes, _ := hex.DecodeString(entropy)
	// fmt.Println("Entropy:", entropy)
	// fmt.Println("Entropy:", bytes)
	// fmt.Printf("Entropy: %08b\n", bytes)
	// fmt.Println()

	// ------------------
	// 2. Create Checksum
	// ------------------
	// SHA256 the entropy
	h := sha256.New()                        // hash function
	h.Write(bytes)                           // write bytes to hash function
	checksum := h.Sum(nil)                   // get result as a byte slice

	// Get specific number of bits from the checksum
	size := len(bytes) * 8 / 32       // 1 bit for every 32 bits of entropy (1 byte = 8 bits)

	// Convert the byte slice to a big integer (so you can use arithmetic to add bits to the end)
	// Note: You can only add bytes (and not individual bits) to a byte slice, so you need to do bitwise arithmetic instead
	//
	//  entropy ([]byte)
	//     |
	//  762603471227441019646259032069709348664 (big.Int)
	//
	dataBigInt := new(big.Int).SetBytes(bytes)

	// Run through the number of bits you want from the checksum, manually adding each bit to the entropy (through arithmetic)
	for i := uint8(0); i < uint8(size); i++ {
		// Add a zero bit to the end for every bit of checksum we add
		//
		//          --->
		//          01001101
		// |entropy|0|
		//
		dataBigInt.Mul(dataBigInt, big.NewInt(2)) // multiplying an integer by two is like adding a 0 bit to the end

		// Use bitwise AND mask to check if each bit of the checksum is set
		//
		// checksum[0] = 01001101
		//           AND 10000000 = 0
		//           AND  1000000 = 1000000
		//           AND   100000 = 0
		//           AND    10000 = 0
		//
		mask := 1 << (7 - i)
		set := uint8(checksum[0]) & uint8(mask) // e.g. 100100100 AND 10000000 = 10000000

		if set > 0 {
			// If the bit is set, change the last zero bit to a 1 bit
			//          10001101
			// |entropy|1|
			//
			dataBigInt.Or(dataBigInt, big.NewInt(1)) // Use bitwise OR to toggle last bit (basically adds 1 to the integer)
		}
	}

	// TIP: Adding bits manually using bitwise arithemtic (can't use bitwise AND or SHIFT on an integer, but maths can do the same)
	// dataBigInt.Mul(dataBigInt, big.NewInt(2)) // add 0 bit to end         = 0  (multiplying by two is like adding a 0 bit)
	// dataBigInt.Or(dataBigInt, big.NewInt(1))  // change 0 bit on end to 1 = 1
	// dataBigInt.Mul(dataBigInt, big.NewInt(2)) // add 0 bit to end         = 10
	// dataBigInt.Mul(dataBigInt, big.NewInt(2)) // add 0 bit to end         = 100
	// dataBigInt.Mul(dataBigInt, big.NewInt(2)) // add 0 bit to end         = 1000

	// ---------------------------------------
	// 3. Convert Entropy+Checksum to Mnemonic
	// ---------------------------------------

	// Get array of words from wordlist file
	// fmt.Println(wordlist[2047])

	// Work out expected number of word in mnemonic based on how many 11-bit groups there are in entropy+checksum
	pieces := ((len(bytes) * 8) + size) / 11

	// Create an array of strings to hold words
	words := make([]string, pieces)

	// Loop through every 11 bits of entropy+checksum and convert to corresponding word from wordlist
	for i := pieces - 1; i >= 0; i-- {

		// Use bit mask (bitwise AND) to split integer in to 11-bit pieces
		//
		//            right to left          big.NewInt(2047) = bit mask
		//          <----------------          <--------->
		// 11111111111|11111111111|11111111111|11111111111
		//
		word := big.NewInt(0)                  // hold result of 11 bit mask
		word.And(dataBigInt, big.NewInt(2047)) // bit mask last 11 bits (2047 = 0b11111111111)

		// Add corresponding word to array
		//
		// 11100111000 = 1848 = train
		//
		words[i] = wordlist[word.Int64()] // insert word from wordlist in to array (need to convert big.Int to int64)

		// Remove those 11 bits from end of big integer by bit shifting
		//
		// 11111111111|11111111111|11111111111|11111111111
		//                                    /            - dividing is the same as bit shifting
		//                                    100000000000 = big.NewInt(2048)
		// 11111111111|11111111111|11111111111|
		//
		dataBigInt.Div(dataBigInt, big.NewInt(2048)) // dividing by 2048 is the same as bit shifting 11 bits

	}

	// Convert array of words in to mnemonic string
	mnemonic := strings.Join(words, " ")

	// ------------------------------------
	// 4. Convert Mnemonic Sentence to Seed
	// ------------------------------------
	//passphrase := "TREZOR"
	//seed := pbkdf2.Key([]byte(mnemonic), []byte("mnemonic"+passphrase), 2048, 64, sha512.New)
	//fmt.Println("Seed:", hex.EncodeToString(seed))

	return wallet_cli.NewMnemonic(mnemonic), nil
}
