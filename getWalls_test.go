package main

import "testing"

func TestValidURL(t *testing.T) {
	resp := validURL("http://i.imgur.com/Z6kdWmA.jpg")
	if resp == false {
		t.Errorf("Not a valid URL")
	}
}

func TestPrepareDirectory(t *testing.T) {
	resp := prepareDirectory("/Pictures/Wallpapers/Reddit")
	if resp == false {
		t.Errorf("Funtion failed to create directory")
	}
}

func TestVerifySubreddit(t *testing.T) {
	resp := verifySubreddit("wallpapers")
	if resp == false {
		t.Errorf("Subreddit invalid")
	}
}
