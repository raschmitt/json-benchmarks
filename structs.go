package json_benchmark

type UnloadCheckResponse struct {
	Id		 		  						 string			`json:"id"`
	Orders 	 	  			   []OrderCheckResponse 		`json:"orders"`
	Status 	   								 string 		`json:"status"`
	TransportIdentifier 					 string 		`json:"transportIdentifier"`
	LocationCode 	   						 string 		`json:"locationCode"`
}

type OrderCheckResponse struct {
	Id		 		  						 string			`json:"id"`
	Code 	   			    				 string 		`json:"code"`
	Items   	  			    []ItemCheckResponse 		`json:"items"`
	Type				 					 string 		`json:"type"`
	BlitzedItems 			  []BlitzedItemResponse 		`json:"blitzedItems"`
}

type ItemCheckResponse struct {
	CheckId           				  		 string     	`json:"id"`
	Code         					   		 string     	`json:"code"`
	Description        				  	 	 string     	`json:"description"`
	Batch         					 	 	 string     	`json:"batch,omitempty"`
	UnitsOfMeasurement  				   []string     	`json:"unitsOfMeasurement"`
	UnitOfMeasurement  				   		 string     	`json:"unitOfMeasurement"`
	ManufactureDate  				   	     string     	`json:"manufactureDate"`
	Document  				   	   DocumentResponse     	`json:"document"`
}

type DocumentResponse struct {
	Code		 		  				 	 string			`json:"code"`
	Id		 		  				 		 string			`json:"id"`
}

type BlitzedItemResponse struct {
	ItemCode		 		  				 string			`json:"itemCode"`
	Amount 	   			    				 string 		`json:"amount"`
	UnitOfMeasurement   	   				 string 		`json:"unitOfMeasurement"`
}
