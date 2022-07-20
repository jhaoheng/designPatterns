package builder2

import (
	"reflect"
	"testing"
)

func TestDog(t *testing.T) {
	dog := NewDirector(&DogBuilder{})
	dog.Construct("Snoppy", 100, "WoofWoof")
	dogObj := dog.builder.getResult()
	checkData := map[string]interface{}{
		"age": 100, "name": "Snoppy", "sound": "WoofWoof",
	}

	eq := reflect.DeepEqual(dogObj, checkData)
	if !eq {
		t.Fatalf("Builder1 fail expect 123 acture %s", dogObj)
	}
}

func TestCat(t *testing.T) {
	cat := NewDirector(&CatBuilder{})
	cat.Construct("Kitty", 10110, "MeowMeow")
	catObj := cat.builder.getResult()
	checkData := map[string]interface{}{
		"age": 10110, "name": "Kitty", "sound": "MeowMeow",
	}

	eq := reflect.DeepEqual(catObj, checkData)
	if !eq {
		t.Fatalf("Builder1 fail expect 123 acture %s", catObj)
	}
}
