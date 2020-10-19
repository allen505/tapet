package main

import "testing"

func TestValidURL(t *testing.T) {
	resp := validURL("http://i.imgur.com/Z6kdWmA.jpg")
	if resp == false {
		t.Errorf("Suppose to return True as a reply")
	}

	resp = validURL("invalidUrl.com")
	if resp == true {
		t.Errorf("Suppose to return a false value")
	}
}

func TestPrepareDirectory(t *testing.T) {
	resp := prepareDirectory("/Pictures/Wallpapers/Reddit")
	if resp == "FAIL" {
		t.Errorf("Funtion failed to create directory")
	}
}

func TestVerifySubreddit(t *testing.T) {
	resp := verifySubreddit("wallpapers")
	if resp == false {
		t.Errorf("Suppose to return True as a reply")
	}

	resp = verifySubreddit("wallpapaer_unknown")
	if resp == true {
		t.Errorf("Suppose to return a false value")
	}
}

func TestIsImg(t * testing.T){
	resp := isImg("http://i.imgur.com/5yeBVeM.jpg")
	if resp == false {
		t.Errorf("Suppose to return True as a reply")
	}

	resp = isImg("http://reddit.com")
	if resp == true {
		t.Errorf("Suppose to return a false value")
	}
}

func TestIsHD(t * testing.T){
	resp := isHD("http://i.imgur.com/5yeBVeM.jpg")
	if resp == false {
		t.Errorf("Suppose to return True as a reply")
	}

	resp = isImg("http://reddit.com")
	if resp == true {
		t.Errorf("Suppose to return a false value")
	}
}

func TestIsLandscape(t * testing.T){
	resp := isLandscape("http://i.imgur.com/5yeBVeM.jpg")
	if resp == false {
		t.Errorf("Suppose to return True as a reply")
	}

	resp = isImg("http://reddit.com")
	if resp == true {
		t.Errorf("Suppose to return a false value")
	}
}