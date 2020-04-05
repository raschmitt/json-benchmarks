package json_benchmark

import (
	"encoding/json"
	"github.com/mailru/easyjson"
	"github.com/smartystreets/assertions"
	"github.com/smartystreets/assertions/should"
	"testing"
)

func Test_should_encode_and_decode_UnloadCheckResponse_with_easyjson(t *testing.T) {
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

	//ITEM CHECK
	itemCheckId := "F7FA6B06-299D-4540-BB05-ACCA4674ED0C"
	itemCode := "564161"
	itemBatch := "ASDAD87989"
	itemDescription := "GFA VIDRO 1L"
	itemUnitsOfMeasurement := []string {"PL", "UN", "CX"}
	itemUnitOfMeasurement := "UN"
	itemManufactureDate := "04/04/2020"

	//DOCUMENT
	documentId := orderId
	documentCode := orderCode

	//BLITZED ITEM
	blitzedItemCode := itemCode
	blitzedItemAmount := "2"
	blitzedItemUnitOfMeasurement := "PL"

	initialUnloadCheckResponse := UnloadCheckResponse{
		Id:                  unloadId,
		Orders:              []OrderCheckResponse{{
			Id:   orderId,
			Code: orderCode,
			Items: []ItemCheckResponse{{
				CheckId:            itemCheckId,
				Code:               itemCode,
				Description:        itemDescription,
				Batch:              itemBatch,
				UnitsOfMeasurement: itemUnitsOfMeasurement,
				UnitOfMeasurement:  itemUnitOfMeasurement,
				ManufactureDate:    itemManufactureDate,
				Document:           DocumentResponse{
					Code: documentCode,
					Id:   documentId,
				},
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
	assert.So(order.Type, should.Equal, orderType)

	item := order.Items[0]

	assert.So(item, should.HaveSameTypeAs, ItemCheckResponse{})

	assert.So(item.UnitsOfMeasurement, should.Resemble, itemUnitsOfMeasurement)
	assert.So(item.Code, should.Equal, itemCode)
	assert.So(item.ManufactureDate, should.Equal, itemManufactureDate)
	assert.So(item.UnitOfMeasurement, should.Equal, itemUnitOfMeasurement)
	assert.So(item.Batch, should.Equal, itemBatch)
	assert.So(item.Description, should.Equal, itemDescription)
	assert.So(item.CheckId, should.Equal, itemCheckId)

	document := item.Document

	assert.So(document, should.HaveSameTypeAs, DocumentResponse{})

	assert.So(document.Code, should.Equal, documentCode)
	assert.So(document.Id, should.Equal, documentId)

	blitzedItem := order.BlitzedItems[0]

	assert.So(blitzedItem, should.HaveSameTypeAs, BlitzedItemResponse{})

	assert.So(blitzedItem.ItemCode, should.Equal, blitzedItemCode)
	assert.So(blitzedItem.Amount, should.Equal, blitzedItemAmount)
	assert.So(blitzedItem.UnitOfMeasurement, should.Equal, blitzedItemUnitOfMeasurement)
}

func Test_should_encode_and_decode_UnloadCheckResponse_with_json(t *testing.T) {
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

	//ITEM CHECK
	itemCheckId := "F7FA6B06-299D-4540-BB05-ACCA4674ED0C"
	itemCode := "564161"
	itemBatch := "ASDAD87989"
	itemDescription := "GFA VIDRO 1L"
	itemUnitsOfMeasurement := []string {"PL", "UN", "CX"}
	itemUnitOfMeasurement := "UN"
	itemManufactureDate := "04/04/2020"

	//DOCUMENT
	documentId := orderId
	documentCode := orderCode

	//BLITZED ITEM
	blitzedItemCode := itemCode
	blitzedItemAmount := "2"
	blitzedItemUnitOfMeasurement := "PL"

	initialUnloadCheckResponse := UnloadCheckResponse{
		Id:                  unloadId,
		Orders:              []OrderCheckResponse{{
			Id:   orderId,
			Code: orderCode,
			Items: []ItemCheckResponse{{
				CheckId:            itemCheckId,
				Code:               itemCode,
				Description:        itemDescription,
				Batch:              itemBatch,
				UnitsOfMeasurement: itemUnitsOfMeasurement,
				UnitOfMeasurement:  itemUnitOfMeasurement,
				ManufactureDate:    itemManufactureDate,
				Document:           DocumentResponse{
					Code: documentCode,
					Id:   documentId,
				},
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
	assert.So(order.Type, should.Equal, orderType)

	item := order.Items[0]

	assert.So(item, should.HaveSameTypeAs, ItemCheckResponse{})

	assert.So(item.UnitsOfMeasurement, should.Resemble, itemUnitsOfMeasurement)
	assert.So(item.Code, should.Equal, itemCode)
	assert.So(item.ManufactureDate, should.Equal, itemManufactureDate)
	assert.So(item.UnitOfMeasurement, should.Equal, itemUnitOfMeasurement)
	assert.So(item.Batch, should.Equal, itemBatch)
	assert.So(item.Description, should.Equal, itemDescription)
	assert.So(item.CheckId, should.Equal, itemCheckId)

	document := item.Document

	assert.So(document, should.HaveSameTypeAs, DocumentResponse{})

	assert.So(document.Code, should.Equal, documentCode)
	assert.So(document.Id, should.Equal, documentId)

	blitzedItem := order.BlitzedItems[0]

	assert.So(blitzedItem, should.HaveSameTypeAs, BlitzedItemResponse{})

	assert.So(blitzedItem.ItemCode, should.Equal, blitzedItemCode)
	assert.So(blitzedItem.Amount, should.Equal, blitzedItemAmount)
	assert.So(blitzedItem.UnitOfMeasurement, should.Equal, blitzedItemUnitOfMeasurement)
}