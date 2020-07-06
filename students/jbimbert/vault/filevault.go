package vault

import (
	"Gophercizes/secret/students/jbimbert/encrypt"
	"errors"
	"io/ioutil"
	"os"
	"strings"
)

const kvSep = "\t" // key/value separator in the encrypted file

type FileVault struct {
	Key  string // encoding key for the file
	File string // the encrypted file name
}

// open the encrypted file,
// read all,
// decrypt it,
// return the decrypted slice of key/value
func (fv *FileVault) readAll() ([]string, error) {
	if _, err := os.Stat(fv.File); os.IsNotExist(err) {
		return nil, err
	}
	bytes, err := ioutil.ReadFile(fv.File)
	if err != nil {
		return nil, err
	}
	if len(bytes) == 0 {
		return make([]string, 0), nil
	}
	decrypted, err := encrypt.Decrypt(fv.Key, string(bytes))
	if err != nil {
		return nil, err
	}
	kv := strings.Split(decrypted, "\n")
	return kv, nil
}

// join a key and a value with the const kv separator
func (fv *FileVault) joinkv(key, value string) string {
	return key + kvSep + value
}

// split a kv into a key and a value, using the const kv separator
func (fv *FileVault) splitkv(kv string) (key, value string) {
	a := strings.Split(kv, kvSep)
	return a[0], a[1]
}

// delete a key from the vault
// return value corresponding to the deleted key
func (fv *FileVault) Delete(key string) (string, error) {
	kvs, err := fv.readAll()
	if err != nil {
		return "", err
	}
	var res []string
	var value string
	for _, kv := range kvs { // kv = key kvSep value
		k, v := fv.splitkv(kv)
		if k != key {
			res = append(res, kv)
		} else {
			value = v
		}
	}
	if value == "" {
		return "", errors.New("Key not found. Nothing deleted.")
	}
	encrypted, err := encrypt.Encrypt(fv.Key, strings.Join(res, "\n"))
	err = ioutil.WriteFile(fv.File, []byte(encrypted), 0755)
	return value, err
}

// list all keys in the vault
func (fv *FileVault) List() ([]string, error) {
	kvs, err := fv.readAll()
	if err != nil {
		return nil, err
	}
	var res []string
	for _, kv := range kvs { // kv = key kvSep value
		k, _ := fv.splitkv(kv)
		res = append(res, k)
	}
	return res, nil
}

// open up the file provided,
// decrypt it,
// set a new secret key/value pair,
// then save the file again in an encrypted manner
func (fv *FileVault) Set(key, value string) error {
	var bytes []byte
	_, err := os.Stat(fv.File)
	if os.IsNotExist(err) {
		return err
	}
	if err == nil {
		bytes, err = ioutil.ReadFile(fv.File)
		if err != nil {
			return err
		}
	}
	var newstring string = fv.joinkv(key, value)
	if len(bytes) > 0 {
		decrypted, err := encrypt.Decrypt(fv.Key, string(bytes))
		if err != nil {
			return err
		}
		// check that the given key does not already exist
		kvs := strings.Split(decrypted, "\n")
		for _, kv := range kvs {
			k, _ := fv.splitkv(kv)
			if k == key {
				return errors.New("This key already exist : use \"update\" command to modify it.")
			}
		}
		newstring = decrypted + "\n" + newstring
	}
	encrypted, err := encrypt.Encrypt(fv.Key, newstring)
	err = ioutil.WriteFile(fv.File, []byte(encrypted), 0755)
	return err
}

// open up the file provided,
// decrypt it,
// update a secret key/value pair,
// then save the file again in an encrypted manner
func (fv *FileVault) Update(key, value string) error {
	var bytes []byte
	if _, err := os.Stat(fv.File); err == nil {
		bytes, err = ioutil.ReadFile(fv.File)
		if err != nil {
			return err
		}
	}
	if len(bytes) == 0 {
		return errors.New("Can not update because the vault is empty : use \"set\" instead.")
	}
	decrypted, err := encrypt.Decrypt(fv.Key, string(bytes))
	if err != nil {
		return err
	}
	// check that the given key exist and update it
	kvs := strings.Split(decrypted, "\n")
	var res []string
	found := false
	for _, kv := range kvs {
		k, _ := fv.splitkv(kv)
		if k == key {
			res = append(res, fv.joinkv(key, value))
			found = true
		} else {
			res = append(res, kv)
		}
	}
	if !found {
		return errors.New("Key not found")
	}
	encrypted, err := encrypt.Encrypt(fv.Key, strings.Join(res, "\n"))
	err = ioutil.WriteFile(fv.File, []byte(encrypted), 0755)
	return err
}

// open up the file provided,
// decrypt it,
// output the value for the provided key
func (fv *FileVault) Get(key string) (string, error) {
	kvs, err := fv.readAll()
	if err != nil {
		return "", err
	}
	for _, kv := range kvs { // kv = key kvSep value
		k, v := fv.splitkv(kv)
		if k == key {
			return v, nil
		}
	}
	return "", errors.New("Key not found")
}
