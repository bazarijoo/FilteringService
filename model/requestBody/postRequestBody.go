package requestBody

import "FilteringService/model"

// postRequest represents a data model of a request sent from client.
type PostRequest struct {
	Main  model.Rectangle
	Input []model.Rectangle
}
