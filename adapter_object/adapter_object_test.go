package adapter_object

import "testing"

func Test_adapter_object(t *testing.T) {
	client := NewService()
	client.ConnectDB()
	client.SetXMLData()
	client.Adapter.ConvertToJson()
	client.Adapter.DoSomething()
}
