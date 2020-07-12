package chatwords

import (
	"reflect"
	"testing"
)

func tSplit(t *testing.T, text string, limit int, expArr []string, expText string) {
	actArr, actText := Split(text, limit)
	if !(reflect.DeepEqual(expArr, actArr) && (expText == actText)) {
		t.Errorf("Split(\"%s\", %d) = (%s, \"%s\"), want (%s, \"%s\")", text, limit, actArr, actText, expArr, expText)
	}
}

func TestSplit(t *testing.T) {
	tSplit(t, "", -1, []string{}, "")
	tSplit(t, "foo", -1, []string{"foo"}, "")
	tSplit(t, "foo ", -1, []string{"foo"}, "")
	tSplit(t, "foo bar", -1, []string{"foo", "bar"}, "")
	tSplit(t, "foo\tbar", -1, []string{"foo", "bar"}, "")
	tSplit(t, "foo", 0, []string{}, "foo")
	tSplit(t, "foo bar ", 1, []string{"foo"}, " bar ")
	tSplit(t, "foo 'hello world'", -1, []string{"foo", "hello world"}, "")
	tSplit(t, "foo \"hello world\"", -1, []string{"foo", "hello world"}, "")
	tSplit(t, "foo 'hello \\'world'", -1, []string{"foo", "hello 'world"}, "")
	tSplit(t, "foo \"hello \\\"world\"", -1, []string{"foo", "hello \"world"}, "")
	tSplit(t, "foo 'hello world'bar", -1, []string{"foo", "'hello", "world'bar"}, "")
	tSplit(t, "foo \"hello world\"bar", -1, []string{"foo", "\"hello", "world\"bar"}, "")
	tSplit(t, "<UXXXXX> echo 'hello world'", -1, []string{"<UXXXXX>", "echo", "hello world"}, "")
	tSplit(t, "<UXXXXX> echo 'hello world'", 2, []string{"<UXXXXX>", "echo"}, " 'hello world'")
	tSplit(t, "'foo'' bar'", -1, []string{"'foo''", "bar'"}, "")
	tSplit(t, "'foo'\" bar\"", -1, []string{"'foo'\"", "bar\""}, "")
	tSplit(t, "foo\\\\ bar", -1, []string{"foo\\", "bar"}, "")
}
