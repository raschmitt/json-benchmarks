package json_benchmark

import (
	"encoding/json"
	"github.com/mailru/easyjson"
	"github.com/smartystreets/assertions"
	"github.com/smartystreets/assertions/should"
	"testing"
)

func Test_should_decode_UnloadCheckResponse_with_easyjson(t *testing.T) {
	//Arrange
	assert := assertions.New(t)

	//UNLOAD CHECK
	unloadId := "8F769CC4-6E51-4FD6-93F6-AAE1EF4856EB"
	unloadStatus := "Pending"
	transportIdentifier := "AVC7557"
	locationCode := "BR23"

	//ORDER CHECK
	orderId := "AA568145-EE86-4508-B34B-D901DAF5DADF"
	orderCode := "0004"
	orderType := "OD"

	//BLITZED ITEM
	blitzedItemCode := "564161"
	blitzedItemAmount := "2"
	blitzedItemUnitOfMeasurement := "PL"

	jsn := []byte(`{
					"id": "`+ unloadId +`", 
					"status": "`+ unloadStatus +`",
					"transportIdentifier": "`+ transportIdentifier +`",
					"locationCode": "`+ locationCode +`",
					"orders": [{
						"id": "`+ orderId +`", 
						"code": "`+ orderCode +`", 
						"type": "`+ orderType +`",
						"blitzedItems": [{
 							 "itemCode": "`+ blitzedItemCode +`",
 							 "amount": "`+ blitzedItemAmount +`",
							 "unitOfMeasurement": "`+ blitzedItemUnitOfMeasurement +`"
						}]
					}]
				  }`)

	unloadCheckResponse := UnloadCheckResponse{}

	//Act
	err := easyjson.Unmarshal(jsn, &unloadCheckResponse)

	//Assert
	assert.So(err, should.Resemble, error(nil))

	assert.So(unloadCheckResponse, should.HaveSameTypeAs, UnloadCheckResponse{})

	assert.So(unloadCheckResponse.Id, should.Equal, "bacon")
	assert.So(unloadCheckResponse.Status, should.Equal, unloadStatus)
	assert.So(unloadCheckResponse.TransportIdentifier, should.Equal, transportIdentifier)
	assert.So(unloadCheckResponse.LocationCode, should.Equal, locationCode)

	order := unloadCheckResponse.Orders[0]

	assert.So(order, should.HaveSameTypeAs, OrderCheckResponse{})

	assert.So(order.Id, should.Equal, orderId)
	assert.So(order.Code, should.Equal, orderCode)
	assert.So(order.Items, should.BeEmpty)
	assert.So(order.Type, should.Equal, orderType)

	blitzedItem := order.BlitzedItems[0]

	assert.So(blitzedItem, should.HaveSameTypeAs, BlitzedItemResponse{})

	assert.So(blitzedItem.ItemCode, should.Equal, blitzedItemCode)
	assert.So(blitzedItem.Amount, should.Equal, blitzedItemAmount)
	assert.So(blitzedItem.UnitOfMeasurement, should.Equal, blitzedItemUnitOfMeasurement)
}

func Test_should_decode_UnloadCheckResponse_with_json(t *testing.T) {
	//Arrange
	assert := assertions.New(t)

	//UNLOAD CHECK
	unloadId := "8F769CC4-6E51-4FD6-93F6-AAE1EF4856EB"
	unloadStatus := "Pending"
	transportIdentifier := "AVC7557"
	locationCode := "BR23"

	//ORDER CHECK
	orderId := "AA568145-EE86-4508-B34B-D901DAF5DADF"
	orderCode := "0004"
	orderType := "OD"

	//BLITZED ITEM
	blitzedItemCode := "564161"
	blitzedItemAmount := "2"
	blitzedItemUnitOfMeasurement := "PL"

	jsn := []byte(`{
					"id": "`+ unloadId +`", 
					"status": "`+ unloadStatus +`",
					"transportIdentifier": "`+ transportIdentifier +`",
					"locationCode": "`+ locationCode +`",
					"orders": [{
						"id": "`+ orderId +`", 
						"code": "`+ orderCode +`", 
						"type": "`+ orderType +`",
						"blitzedItems": [{
 							 "itemCode": "`+ blitzedItemCode +`",
 							 "amount": "`+ blitzedItemAmount +`",
							 "unitOfMeasurement": "`+ blitzedItemUnitOfMeasurement +`"
						}]
					}]
				  }`)

	unloadCheckResponse := UnloadCheckResponse{}

	//Act
	err := json.Unmarshal(jsn, &unloadCheckResponse)

	//Assert
	assert.So(err, should.Resemble, error(nil))

	assert.So(unloadCheckResponse, should.HaveSameTypeAs, UnloadCheckResponse{})

	assert.So(unloadCheckResponse.Id, should.Equal, unloadId)
	assert.So(unloadCheckResponse.Status, should.Equal, unloadStatus)
	assert.So(unloadCheckResponse.TransportIdentifier, should.Equal, transportIdentifier)
	assert.So(unloadCheckResponse.LocationCode, should.Equal, locationCode)

	order := unloadCheckResponse.Orders[0]

	assert.So(order, should.HaveSameTypeAs, OrderCheckResponse{})

	assert.So(order.Id, should.Equal, orderId)
	assert.So(order.Code, should.Equal, orderCode)
	assert.So(order.Items, should.BeEmpty)
	assert.So(order.Type, should.Equal, orderType)

	blitzedItem := order.BlitzedItems[0]

	assert.So(blitzedItem, should.HaveSameTypeAs, BlitzedItemResponse{})

	assert.So(blitzedItem.ItemCode, should.Equal, blitzedItemCode)
	assert.So(blitzedItem.Amount, should.Equal, blitzedItemAmount)
	assert.So(blitzedItem.UnitOfMeasurement, should.Equal, blitzedItemUnitOfMeasurement)
}

func Test_should_encode_UnloadCheckResponse_with_easyjson(t *testing.T) {
	//Arrange
	assert := assertions.New(t)

	//UNLOAD CHECK
	unloadId := "8F769CC4-6E51-4FD6-93F6-AAE1EF4856EB"
	unloadStatus := "Pending"
	transportIdentifier := "AVC7557"
	locationCode := "BR23"

	//ORDER CHECK
	orderId := "AA568145-EE86-4508-B34B-D901DAF5DADF"
	orderCode := "0004"
	orderType := "OD"

	//BLITZED ITEM
	blitzedItemCode := "564161"
	blitzedItemAmount := "2"
	blitzedItemUnitOfMeasurement := "PL"

	initialUnloadCheckResponse := UnloadCheckResponse{
		Id:                  unloadId,
		Orders:              []OrderCheckResponse{{
			Id:   orderId,
			Code: orderCode,
			Items: []ItemCheckResponse{{
				CheckId:            "",
				Code:               "",
				Description:        "",
				Batch:              "",
				UnitsOfMeasurement: nil,
				UnitOfMeasurement:  "",
				ManufactureDate:    "",
				Document:           DocumentResponse{},
			}},
			Type: orderType,
			BlitzedItems: []BlitzedItemResponse{{
				ItemCode:          blitzedItemCode,
				Amount:            blitzedItemAmount,
				UnitOfMeasurement: blitzedItemUnitOfMeasurement,
			}},
		}},
		Status:              unloadStatus,
		TransportIdentifier: transportIdentifier,
		LocationCode:        locationCode,
	}

	assertUnloadCheckResponse := UnloadCheckResponse{}

	//Act
	jsn, err := easyjson.Marshal(&initialUnloadCheckResponse)

	easyjson.Unmarshal(jsn, &assertUnloadCheckResponse)

	//Assert
	assert.So(err, should.Resemble, error(nil))

	assert.So(assertUnloadCheckResponse, should.HaveSameTypeAs, UnloadCheckResponse{})

	assert.So(assertUnloadCheckResponse.Id, should.Equal, unloadId)
	assert.So(assertUnloadCheckResponse.Status, should.Equal, unloadStatus)
	assert.So(assertUnloadCheckResponse.TransportIdentifier, should.Equal, transportIdentifier)
	assert.So(assertUnloadCheckResponse.LocationCode, should.Equal, locationCode)

	order := assertUnloadCheckResponse.Orders[0]

	assert.So(order, should.HaveSameTypeAs, OrderCheckResponse{})

	assert.So(order.Id, should.Equal, orderId)
	assert.So(order.Code, should.Equal, orderCode)
	assert.So(order.Items, should.BeEmpty)
	assert.So(order.Type, should.Equal, orderType)

	blitzedItem := order.BlitzedItems[0]

	assert.So(blitzedItem, should.HaveSameTypeAs, BlitzedItemResponse{})

	assert.So(blitzedItem.ItemCode, should.Equal, blitzedItemCode)
	assert.So(blitzedItem.Amount, should.Equal, blitzedItemAmount)
	assert.So(blitzedItem.UnitOfMeasurement, should.Equal, blitzedItemUnitOfMeasurement)
}

func Test_should_encode_UnloadCheckResponse_with_json(t *testing.T) {
	//Arrange
	assert := assertions.New(t)

	//UNLOAD CHECK
	unloadId := "8F769CC4-6E51-4FD6-93F6-AAE1EF4856EB"
	unloadStatus := "Pending"
	transportIdentifier := "AVC7557"
	locationCode := "BR23"

	//ORDER CHECK
	orderId := "AA568145-EE86-4508-B34B-D901DAF5DADF"
	orderCode := "0004"
	orderType := "OD"

	//BLITZED ITEM
	blitzedItemCode := "564161"
	blitzedItemAmount := "2"
	blitzedItemUnitOfMeasurement := "PL"

	initialUnloadCheckResponse := UnloadCheckResponse{
		Id:                  unloadId,
		Orders:              []OrderCheckResponse{{
			Id:   orderId,
			Code: orderCode,
			Items: []ItemCheckResponse{{
				CheckId:            "",
				Code:               "",
				Description:        "",
				Batch:              "",
				UnitsOfMeasurement: nil,
				UnitOfMeasurement:  "",
				ManufactureDate:    "",
				Document:           DocumentResponse{},
			}},
			Type: orderType,
			BlitzedItems: []BlitzedItemResponse{{
				ItemCode:          blitzedItemCode,
				Amount:            blitzedItemAmount,
				UnitOfMeasurement: blitzedItemUnitOfMeasurement,
			}},
		}},
		Status:              unloadStatus,
		TransportIdentifier: transportIdentifier,
		LocationCode:        locationCode,
	}

	assertUnloadCheckResponse := UnloadCheckResponse{}

	//Act
	jsn, err := json.Marshal(&initialUnloadCheckResponse)

	json.Unmarshal(jsn, &assertUnloadCheckResponse)

	//Assert
	assert.So(err, should.Resemble, error(nil))

	assert.So(assertUnloadCheckResponse, should.HaveSameTypeAs, UnloadCheckResponse{})

	assert.So(assertUnloadCheckResponse.Id, should.Equal, unloadId)
	assert.So(assertUnloadCheckResponse.Status, should.Equal, unloadStatus)
	assert.So(assertUnloadCheckResponse.TransportIdentifier, should.Equal, transportIdentifier)
	assert.So(assertUnloadCheckResponse.LocationCode, should.Equal, locationCode)

	order := assertUnloadCheckResponse.Orders[0]

	assert.So(order, should.HaveSameTypeAs, OrderCheckResponse{})

	assert.So(order.Id, should.Equal, orderId)
	assert.So(order.Code, should.Equal, orderCode)
	assert.So(order.Items, should.BeEmpty)
	assert.So(order.Type, should.Equal, orderType)

	blitzedItem := order.BlitzedItems[0]

	assert.So(blitzedItem, should.HaveSameTypeAs, BlitzedItemResponse{})

	assert.So(blitzedItem.ItemCode, should.Equal, blitzedItemCode)
	assert.So(blitzedItem.Amount, should.Equal, blitzedItemAmount)
	assert.So(blitzedItem.UnitOfMeasurement, should.Equal, blitzedItemUnitOfMeasurement)
}