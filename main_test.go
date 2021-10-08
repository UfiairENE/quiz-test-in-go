package main

import (
	"strings"
	"testing"
	"time"

	"gotest.tools/assert"
)

func testEachQuestion(t *testing.T) {
	timer := time.NewTimer(time.Duration(2) * time.Second).C
	done := make(chan string)
	var quest Question
	quest.question = "2+7"
	quest.answer = "9"
	var ans int
	var err error
	allDone := make(chan bool)
	go func() {
		ans, err = eachQuestion(quest.question, quest.answer, timer, done)
		allDone <- true
	}()
	done <- "9"

	<-allDone
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, ans, 1)
}

func testReadCSV(t *testing.T) {
	str := "2+7,9\n3+4,7\n5+2,7\n1+8,9\n7-1,6\n"
	quest, err := readCSV(strings.NewReader(str))
	if err != nil {
		t.Error(err)
	}
	var que [5]Question
	que[0].answer = "9"
	que[1].answer = "7"
	que[2].answer = "7"
	que[3].answer = "9"
	que[4].answer = "6"
	que[0].question = "2+7"
	que[1].question = "3+4"
	que[2].question = "5+2"
	que[3].question = "1+8"
	que[4].question = "7-1"

	assert.Equal(t, que[0], quest[0])
	assert.Equal(t, que[1], quest[1])
	assert.Equal(t, que[2], quest[2])
	assert.Equal(t, que[3], quest[3])
	assert.Equal(t, que[4], quest[4])

}

func TestEachQuestion(t *testing.T) {
	t.Run("test eachQuestion", testEachQuestion)
}

func TestReadCSV(t *testing.T) {
	t.Run("test ReadCSV", testReadCSV)
}
