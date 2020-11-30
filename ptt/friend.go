package ptt

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"math/rand"
	"os"

	"github.com/PichuChen/go-bbs/names"
	"github.com/PichuChen/go-bbs/path"
	"github.com/PichuChen/go-bbs/ptttype"
	"github.com/PichuChen/go-bbs/types"
	log "github.com/sirupsen/logrus"
)

func friendDeleteAll(userID *[ptttype.IDLEN + 1]byte, friendType int) error {
	filename, err := path.SetHomeFile(userID, ptttype.FriendFile[friendType])
	if err != nil { //unable to get the file. assuming not-exists
		return err
	}

	file, err := os.Open(filename)
	if err != nil { //unable to open the file. assuming not-exists
		return nil
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	friendID := &[ptttype.IDLEN + 1]byte{}
	for line, err := reader.ReadBytes('\n'); err == nil; line, err = reader.ReadBytes('\n') {
		copy(friendID[:], line)
		if !names.IsValidUserID(friendID) {
			continue
		}

		// XXX race-condition.
		deleteUserFriend(friendID, userID, friendType) // remove me from my friend.
	}
	return nil
}

func deleteUserFriend(userID *[ptttype.IDLEN + 1]byte, friendID *[ptttype.IDLEN + 1]byte, friendType int) {
	if bytes.Equal(userID[:], friendID[:]) {
		return
	}

	filename, err := path.SetHomeFile(userID, ptttype.FN_ALOHA)
	if err != nil {
		return
	}

	_ = deleteFriendFromFile(filename, friendID, false)
}

func deleteFriendFromFile(filename string, friend *[ptttype.IDLEN + 1]byte, isCaseSensitive bool) bool {
	// XXX race-condition
	randNum := rand.Intn(0xfff)
	randStr := fmt.Sprintf("%3.3X", randNum)
	new_filename := filename + "." + randStr

	file, err := os.OpenFile(filename, os.O_RDONLY, 0600)
	if err != nil { // no file
		return false
	}
	defer file.Close()

	new_file, err := os.OpenFile(new_filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil { // unable to create new file.
		return false
	}
	defer new_file.Close()

	userIDInFile := &[ptttype.IDLEN + 1]byte{}
	reader := bufio.NewReader(file)
	writer := bufio.NewWriter(new_file)

	var line []byte
	for {
		line, err = reader.ReadBytes('\n')
		if err != nil {
			if err == io.EOF {
				err = nil // make it clear about err here
				break
			}

			log.Errorf("friend.deleteFriendFromFile: unable to read file: filename: %v new_filename: %v e: %v", filename, new_filename, err)
			return false
		}
		copy(userIDInFile[:], line[:])
		if isCaseSensitive && !bytes.Equal(friend[:], userIDInFile[:]) ||
			!isCaseSensitive && !bytes.EqualFold(friend[:], userIDInFile[:]) {
			sanitizedUserIDInFile := types.CstrToBytes(userIDInFile[:])
			_, err = writer.Write(sanitizedUserIDInFile)
			if err != nil { // unable to write-bytes into tmp-file.
				log.Errorf("friend.deleteFriendFromFile: unable write to tmp-file (possible zombie-tmp-file): filename: %v new_filename: %v e: %v", filename, new_filename, err)
				return false
			}
			err = writer.WriteByte('\n')
			if err != nil { // unable to write new-line into tmp-file.
				log.Errorf("friend.deleteFriendFromFile: unable write to tmp-file (possible zombie-tmp-file): filename: %v new_filename: %v e: %v", filename, new_filename, err)
				return false
			}

		}
	}

	err = os.Rename(new_filename, filename)
	if err != nil {
		return false
	}

	return true
}
