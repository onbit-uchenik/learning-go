/*
	Author - Anubhav Vats
	Description - Calculate SHA256 Digest of your message. This file written to learn closure in go.
	Last Modified - 21/04/21
*/
package main

import (
	"fmt"
	"crypto/sha256"
)

/*
	Have you ever used hash.update and hash.digest function of Nodejs implementation?
	https://nodejs.org/dist/latest-v14.x/docs/api/crypto.html#crypto_hash_digest_encoding
	https://nodejs.org/dist/latest-v14.x/docs/api/crypto.html#crypto_hash_update_data_inputencoding
	If yes and wants to use the same behaviour (the behaviour is same but performance is not gauranteed),
	but there is no such implementation in golang standard library

	Here is this polyfill can be used. This function returns two functions one to add more data to the
	existing input and other function to calculate sha256 digest of complete data till now.

*/
func hash (data string) (func (moreData string), func () string) {

	addMoreDataToHash := func (moreData string) {
		data = data + moreData
	}

	getDigest := func () string {
		digest := sha256.Sum256([]byte(data));

		return fmt.Sprintf("%x", digest)
	}

	return addMoreDataToHash, getDigest
}



func main() {
	
	add, getDigest := hash("anubhav")
	// adding more data to the existing input
	add("vats")
	// calculate the digest of input till now
	fmt.Println(getDigest())
	// add more data
	add("hello")
	// now again calculate the hash.
	fmt.Println(getDigest())
}