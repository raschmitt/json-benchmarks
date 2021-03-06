package json_benchmark

import (
	"encoding/json"
	"github.com/mailru/easyjson"
	"testing"
)

func Benchmark_decode_UnloadCheckResponse_with_easyjson(b *testing.B) {
	for i := 0; i < b.N; i++ {
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
		itemUnitOfMeasurement := "UN"
		itemManufactureDate := "04/04/2020"

		//DOCUMENT
		documentId := orderId
		documentCode := orderCode

		//BLITZED ITEM
		blitzedItemCode := "564161"
		blitzedItemAmount := "2"
		blitzedItemUnitOfMeasurement := "PL"

		jsn := []byte(`{
					"id": "` + unloadId + `", 
					"status": "` + unloadStatus + `",
					"transportIdentifier": "` + transportIdentifier + `",
					"locationCode": "` + locationCode + `",
					"orders": [{
						"id": "` + orderId + `", 
						"code": "` + orderCode + `", 
						"type": "` + orderType + `",
						"items": [{
 							 "id": "` + itemCheckId + `",
 							 "code": "` + itemCode + `",
							 "description": "` + itemDescription + `",
							 "batch": "` + itemBatch + `",
							 "unitOfMeasurement": "` + itemUnitOfMeasurement + `",
							 "manufactureDate": "` + itemManufactureDate + `",
							 "document": {"code": "` + documentCode + `", "id": "` + documentId + `"},
						},
						"blitzedItems": [{
 							 "itemCode": "` + blitzedItemCode + `",
 							 "amount": "` + blitzedItemAmount + `",
							 "unitOfMeasurement": "` + blitzedItemUnitOfMeasurement + `"
						}]
					}]
				  }`)

		unloadCheckResponse := UnloadCheckResponse{}

		easyjson.Unmarshal(jsn, &unloadCheckResponse)
	}
}

func Benchmark_decode_UnloadCheckResponse_with_json(b *testing.B) {
	for i := 0; i < b.N; i++ {
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
		itemUnitOfMeasurement := "UN"
		itemManufactureDate := "04/04/2020"

		//DOCUMENT
		documentId := orderId
		documentCode := orderCode

		//BLITZED ITEM
		blitzedItemCode := "564161"
		blitzedItemAmount := "2"
		blitzedItemUnitOfMeasurement := "PL"

		jsn := []byte(`{
					"id": "` + unloadId + `", 
					"status": "` + unloadStatus + `",
					"transportIdentifier": "` + transportIdentifier + `",
					"locationCode": "` + locationCode + `",
					"orders": [{
						"id": "` + orderId + `", 
						"code": "` + orderCode + `", 
						"type": "` + orderType + `",
						"items": [{
 							 "id": "` + itemCheckId + `",
 							 "code": "` + itemCode + `",
							 "description": "` + itemDescription + `",
							 "batch": "` + itemBatch + `",
							 "unitOfMeasurement": "` + itemUnitOfMeasurement + `",
							 "manufactureDate": "` + itemManufactureDate + `",
							 "document": {"code": "` + documentCode + `", "id": "` + documentId + `"},
						},
						"blitzedItems": [{
 							 "itemCode": "` + blitzedItemCode + `",
 							 "amount": "` + blitzedItemAmount + `",
							 "unitOfMeasurement": "` + blitzedItemUnitOfMeasurement + `"
						}]
					}]
				  }`)

		unloadCheckResponse := UnloadCheckResponse{}

		json.Unmarshal(jsn, &unloadCheckResponse)
	}
}

func Benchmark_encode_UnloadCheckResponse_with_easyjson(b *testing.B) {
	for i := 0; i < b.N; i++ {
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

		easyjson.Marshal(&initialUnloadCheckResponse)
	}
}

func Benchmark_encode_UnloadCheckResponse_with_json(b *testing.B) {
	for i := 0; i < b.N; i++ {
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

		json.Marshal(&initialUnloadCheckResponse)
	}
}

