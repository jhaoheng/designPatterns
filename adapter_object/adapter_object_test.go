package adapter_object

import "testing"

func Test_adapter_object(t *testing.T) {
	client := NewService()
	client.ConnectDB()
	client.SetData()
	client.Adapter.ConvertToJson(client.MyData)
	client.Adapter.Adapee.DoSomething()
}
