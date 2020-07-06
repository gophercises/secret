package main

import (
	"Gophercizes/secret/students/jbimbert/encrypt"
	"Gophercizes/secret/students/jbimbert/vault"
	"os"
	"testing"
)

const (
	key      = "someKey"
	value    = "someValue"
	filename = "aFile"
)

var (
	keys   = []string{"key1", "key2", "key3", "key4", "key5"}
	values = []string{"value1", "value2", "value3", "value4", "value5"}
)

func createEmptyFile(t *testing.T) *os.File {
	f, err := os.Create(filename)
	if err != nil {
		t.Error(err)
	}
	return f
}

func createFile(t *testing.T) *os.File {
	f, err := os.Create(filename)
	if err != nil {
		t.Error(err)
	}
	fv := &vault.FileVault{Key: "encodingKey", File: filename}
	err = fv.Set(key, value)
	if err != nil {
		t.Error(err)
	}
	return f
}

func createFileMany(t *testing.T) *os.File {
	f, err := os.Create(filename)
	if err != nil {
		t.Error(err)
	}
	fv := &vault.FileVault{Key: "encodingKey", File: filename}
	for i, k := range keys {
		err := fv.Set(k, values[i])
		if err != nil {
			t.Error(err)
		}
	}
	return f
}

func Test_deleteExistingKey(t *testing.T) {
	createFileMany(t)
	defer os.Remove(filename)
	fv := &vault.FileVault{Key: "encodingKey", File: filename}
	_, err := fv.Delete(keys[0])
	if err != nil {
		t.Error(err)
	}
	for i, k := range keys[1:] {
		v, err := fv.Get(k)
		if err != nil {
			t.Error(err)
		}
		if v != values[i+1] {
			t.Errorf("Bad retrieved value.\nFound: %s\nWanted: %s", v, values[i])
		}
	}
}

func Test_deleteNotExistingKey(t *testing.T) {
	createFileMany(t)
	defer os.Remove(filename)
	fv := &vault.FileVault{Key: "encodingKey", File: filename}
	_, err := fv.Delete("badKey")
	if err.Error() != "Key not found. Nothing deleted." {
		t.Errorf("This test should raise an error : Key not found. Nothing deleted.")
	}
	for i, k := range keys {
		v, err := fv.Get(k)
		if err != nil {
			t.Error(err)
		}
		if v != values[i] {
			t.Errorf("Bad retrieved value.\nFound: %s\nWanted: %s", v, values[i])
		}
	}
}

func Test_updateNonExistingKey(t *testing.T) {
	createFile(t)
	defer os.Remove(filename)
	fv := &vault.FileVault{Key: "encodingKey", File: filename}
	err := fv.Update("badKey", "anything")
	if err.Error() != "Key not found" {
		t.Errorf("This test should raise an error : Key not found")
	}
}

func Test_updateExistingKey(t *testing.T) {
	createFile(t)
	defer os.Remove(filename)
	fv := &vault.FileVault{Key: "encodingKey", File: filename}
	newValue := "newValue"
	err := fv.Update(key, newValue)
	if err != nil {
		t.Error(err)
	}
	v, err := fv.Get(key)
	if err != nil {
		t.Error(err)
	}
	if v != newValue {
		t.Errorf("Bad retrieved value.\nFound: %s\nWanted: %s", v, newValue)
	}
}

func Test_setExistingKey(t *testing.T) {
	createFile(t)
	defer os.Remove(filename)
	fv := &vault.FileVault{Key: "encodingKey", File: filename}
	err := fv.Set(key, "someOtherValue")
	if err.Error() != "This key already exist : use \"update\" command to modify it." {
		t.Errorf("This test should raise an error : This key already exist : use \"update\" command to modify it.")
	}
}

func Test_getSeveralValues(t *testing.T) {
	createFileMany(t)
	defer os.Remove(filename)
	fv := &vault.FileVault{Key: "encodingKey", File: filename}
	for i, k := range keys {
		v, err := fv.Get(k)
		if err != nil {
			t.Error(err)
		}
		if v != values[i] {
			t.Errorf("Bad retrieved value.\nFound: %s\nWanted: %s", v, values[i])
		}
	}
}

func Test_listSeveralKeys(t *testing.T) {
	createFileMany(t)
	defer os.Remove(filename)
	fv := &vault.FileVault{Key: "encodingKey", File: filename}
	ks, err := fv.List()
	if err != nil {
		t.Error(err)
	}
	for i, k := range ks {
		if k != keys[i] {
			t.Errorf("Bad retrieved value.\nFound: %s\nWanted: %s", k, keys[i])
		}
	}
}

func Test_getExistingKey(t *testing.T) {
	createFile(t)
	defer os.Remove(filename)
	fv := &vault.FileVault{Key: "encodingKey", File: filename}
	v, err := fv.Get(key)
	if err != nil {
		t.Error(err)
	}
	if value != v {
		t.Errorf("Bad retrieved value.\nFound: %s\nWanted: %s", v, value)
	}
}

func Test_setEmptyFile(t *testing.T) {
	createEmptyFile(t)
	defer os.Remove(filename)
	fv := &vault.FileVault{Key: "encodingKey", File: filename}
	err := fv.Set("someKey", "someValue")
	if err != nil {
		t.Error(err)
	}
}

func Test_setBadFile(t *testing.T) {
	fv := &vault.FileVault{Key: "encodingKey", File: "badFile"}
	err := fv.Set("someKey", "someValue")
	if err.Error() != "stat badFile: no such file or directory" {
		t.Error("This test should raise an error : stat badFile: no such file or directory")
	}
}

func Test_getEmptyFile(t *testing.T) {
	createEmptyFile(t)
	defer os.Remove(filename)
	fv := &vault.FileVault{Key: "encodingKey", File: filename}
	_, err := fv.Get("someKey")
	if err.Error() != "Key not found" {
		t.Error("This test should raise an error : Key not found")
	}
}

func Test_getBadFile(t *testing.T) {
	fv := &vault.FileVault{Key: "encodingKey", File: "badFile"}
	_, err := fv.Get("someKey")
	if err.Error() != "stat badFile: no such file or directory" {
		t.Error("This test should raise an error : stat badFile: no such file or directory")
	}
}

func Test_encryptDecrypt(t *testing.T) {
	text1 := "un court text avec des espaces"
	key := "toto"
	encrypted, err := encrypt.Encrypt(key, text1)
	if err != nil {
		t.Error(err)
	}
	decrypted, err := encrypt.Decrypt(key, encrypted)
	if err != nil {
		t.Error(err)
	}
	if text1 != decrypted {
		t.Errorf("Bad encrypted <-> decrypted values.\nEncrypted : %s\nDecrypted : %s", text1, decrypted)
	}
}
