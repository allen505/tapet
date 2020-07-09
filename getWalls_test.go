package main

import "testing"

func TestSum(t *testing.T){
	total := successSum(1,2)
	if total != 3 {
       t.Errorf("Sum was incorrect, got: %d, want: %d.", total, 3)
	}
}

func TestSumWillFail(t *testing.T){
	total := failSum(1,2)
	if total != 3 {
       t.Errorf("Sum was incorrect, got: %d, want: %d.", total, 3)
	}
}

func TestValidURL(t *testing.T){
	resp := validURL("http://i.imgur.com/Z6kdWmA.jpg")
	if resp == false {
		t.Errorf("Not a valid URL")
	}
}

func TestPrepareDirectory(t *testing.T){
	resp := prepareDirectory("/Pictures/Wallpapers/Reddit")
	if resp == false {
		t.Errorf("Funtion failed to create directory")
	}
}

func TestVerifySubreddit(t *testing.T){
	resp := verifySubreddit("wallpapers")
	if resp == false {
		t.Errorf("Subreddit invalid")
	}
}