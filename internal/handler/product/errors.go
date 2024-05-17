package product

import (
	"net/http"

	"github.com/eulabs/back-end/api-product/internal/handler"
)

// Success Message Here

var SuccessHttpMsgToDeleteProduct handler.HttpMsg = handler.HttpMsg{
	Msg:  "Ok Product Deleted",
	Code: http.StatusOK,
}

// Erros Message Here

var ErroHttpMsgProductIdIsRequired handler.HttpMsg = handler.HttpMsg{
	Msg:  "Erro Product ID is required",
	Code: http.StatusBadRequest,
}

var ErroHttpMsgProductNameIsRequired handler.HttpMsg = handler.HttpMsg{
	Msg:  "Erro Product Name is required",
	Code: http.StatusBadRequest,
}

var ErroHttpMsgProductQuantityIsRequired handler.HttpMsg = handler.HttpMsg{
	Msg:  "Erro Product Quantity is required",
	Code: http.StatusBadRequest,
}

var ErroHttpMsgProductPriceIsRequired handler.HttpMsg = handler.HttpMsg{
	Msg:  "Erro Product Price is required",
	Code: http.StatusBadRequest,
}

var ErroHttpMsgProductNotFound handler.HttpMsg = handler.HttpMsg{
	Msg:  "Erro Product Not Found",
	Code: http.StatusNotFound,
}

var ErroHttpMsgToParseRequestProductToJson handler.HttpMsg = handler.HttpMsg{
	Msg:  "Erro to parse Request Product to JSON",
	Code: http.StatusInternalServerError,
}

var ErroHttpMsgToParseResponseProductToJson handler.HttpMsg = handler.HttpMsg{
	Msg:  "Erro to parse Response Product to JSON",
	Code: http.StatusInternalServerError,
}

var ErroHttpMsgToConvertingResponseProductListToJson handler.HttpMsg = handler.HttpMsg{
	Msg:  "Erro to converting Response Product List to JSON",
	Code: http.StatusInternalServerError,
}

var ErroHttpMsgToInsertProduct handler.HttpMsg = handler.HttpMsg{
	Msg:  "Erro to Insert the Product",
	Code: http.StatusInternalServerError,
}

var ErroHttpMsgToUpdateProduct handler.HttpMsg = handler.HttpMsg{
	Msg:  "Erro to Update the Product",
	Code: http.StatusInternalServerError,
}

var ErroHttpMsgToDeleteProduct handler.HttpMsg = handler.HttpMsg{
	Msg:  "Erro to Delete the Product",
	Code: http.StatusInternalServerError,
}
