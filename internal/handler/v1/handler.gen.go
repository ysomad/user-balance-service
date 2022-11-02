// Package v1 provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.12.1 DO NOT EDIT.
package v1

import (
	"fmt"
	"net/http"
	"time"

	"github.com/deepmap/oapi-codegen/pkg/runtime"
	"github.com/go-chi/chi/v5"
	google_uuid "github.com/google/uuid"
)

// Defines values for SortOrder.
const (
	ASC  SortOrder = "ASC"
	DESC SortOrder = "DESC"
)

// Defines values for TransactionOperation.
const (
	DEPOSIT  TransactionOperation = "DEPOSIT"
	WITHDRAW TransactionOperation = "WITHDRAW"
)

// Account defines model for Account.
type Account struct {
	Balance string           `json:"balance"`
	ID      google_uuid.UUID `json:"id"`
	UserID  google_uuid.UUID `json:"user_id"`
}

// DeclareRevenueRequest defines model for DeclareRevenueRequest.
type DeclareRevenueRequest struct {
	Amount    string           `json:"amount"`
	OrderID   google_uuid.UUID `json:"order_id"`
	ServiceID google_uuid.UUID `json:"service_id"`
}

// DeclareRevenueResponse defines model for DeclareRevenueResponse.
type DeclareRevenueResponse struct {
	DeclaredAmount  string           `json:"declared_amount"`
	DeclaredAt      *time.Time       `json:"declared_at,omitempty"`
	Declared        bool             `json:"is_declared"`
	OrderID         google_uuid.UUID `json:"order_id"`
	RevenueReportID google_uuid.UUID `json:"revenue_report_id"`
	ServiceID       google_uuid.UUID `json:"service_id"`
}

// DepositFundsRequest defines model for DepositFundsRequest.
type DepositFundsRequest struct {
	Amount string `json:"amount"`
}

// Error defines model for Error.
type Error struct {
	Message string `json:"message"`
	Status  string `json:"status"`
}

// ReserveFundsRequest defines model for ReserveFundsRequest.
type ReserveFundsRequest struct {
	Amount    string           `json:"amount"`
	OrderID   google_uuid.UUID `json:"order_id"`
	ServiceID google_uuid.UUID `json:"service_id"`
}

// ReserveFundsResponse defines model for ReserveFundsResponse.
type ReserveFundsResponse struct {
	AccountBalance string           `json:"account_balance"`
	Declared       bool             `json:"is_declared"`
	OrderID        google_uuid.UUID `json:"order_id"`
	ReservedAmount string           `json:"reserved_amount"`
	ReservedAt     time.Time        `json:"reserved_at"`
	ServiceID      google_uuid.UUID `json:"service_id"`
}

// RevenueReportResponse defines model for RevenueReportResponse.
type RevenueReportResponse struct {
	ReportURL string `json:"report_url"`
}

// SortOrder defines model for SortOrder.
type SortOrder string

// Transaction defines model for Transaction.
type Transaction struct {
	AccountID  google_uuid.UUID     `json:"account_id"`
	Amount     string               `json:"amount"`
	Comment    string               `json:"comment"`
	CommitedAt time.Time            `json:"commited_at"`
	ID         uint32               `json:"id"`
	Operation  TransactionOperation `json:"operation"`
}

// TransactionList defines model for TransactionList.
type TransactionList struct {
	NextPageToken uint32        `json:"next_page_token"`
	Transactions  []Transaction `json:"transactions"`
}

// TransactionListRequest defines model for TransactionListRequest.
type TransactionListRequest struct {
	PageSize  uint64              `json:"page_size"`
	PageToken uint32              `json:"page_token"`
	Sort      TransactionListSort `json:"sort"`
}

// TransactionListSort defines model for TransactionListSort.
type TransactionListSort struct {
	Amount     SortOrder `json:"amount"`
	CommitTime SortOrder `json:"commit_time"`
}

// TransactionOperation defines model for TransactionOperation.
type TransactionOperation string

// TransferFundsRequest defines model for TransferFundsRequest.
type TransferFundsRequest struct {
	Amount   string           `json:"amount"`
	ToUserID google_uuid.UUID `json:"to_user_id"`
}

// DeclareRevenueRequestBody defines model for DeclareRevenueRequestBody.
type DeclareRevenueRequestBody = DeclareRevenueRequest

// DepositFundsRequestBody defines model for DepositFundsRequestBody.
type DepositFundsRequestBody = DepositFundsRequest

// ReserveFundsRequestBody defines model for ReserveFundsRequestBody.
type ReserveFundsRequestBody = ReserveFundsRequest

// TransactionListRequestBody defines model for TransactionListRequestBody.
type TransactionListRequestBody = TransactionListRequest

// TransferFundsRequestBody defines model for TransferFundsRequestBody.
type TransferFundsRequestBody = TransferFundsRequest

// DepositFundsJSONRequestBody defines body for DepositFunds for application/json ContentType.
type DepositFundsJSONRequestBody = DepositFundsRequest

// GetAccountTransactionsJSONRequestBody defines body for GetAccountTransactions for application/json ContentType.
type GetAccountTransactionsJSONRequestBody = TransactionListRequest

// TransferFundsJSONRequestBody defines body for TransferFunds for application/json ContentType.
type TransferFundsJSONRequestBody = TransferFundsRequest

// ReserveFundsJSONRequestBody defines body for ReserveFunds for application/json ContentType.
type ReserveFundsJSONRequestBody = ReserveFundsRequest

// DeclareRevenueJSONRequestBody defines body for DeclareRevenue for application/json ContentType.
type DeclareRevenueJSONRequestBody = DeclareRevenueRequest

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Get user account
	// (GET /accounts/{user_id})
	GetAccount(w http.ResponseWriter, r *http.Request, userId google_uuid.UUID)
	// Deposit funds to the user account
	// (PUT /accounts/{user_id})
	DepositFunds(w http.ResponseWriter, r *http.Request, userId google_uuid.UUID)
	// Get list of account balance transactions
	// (POST /accounts/{user_id}/transactions)
	GetAccountTransactions(w http.ResponseWriter, r *http.Request, userId google_uuid.UUID)
	// Transfer funds to user balance
	// (POST /accounts/{user_id}/transfer)
	TransferFunds(w http.ResponseWriter, r *http.Request, userId google_uuid.UUID)
	// Get revenue report
	// (GET /reports/{month})
	GetRevenueReport(w http.ResponseWriter, r *http.Request, month string)
	// Reserve funds from user account
	// (POST /reservations/{user_id})
	ReserveFunds(w http.ResponseWriter, r *http.Request, userId google_uuid.UUID)
	// Declare revenue
	// (POST /reservations/{user_id}/revenue)
	DeclareRevenue(w http.ResponseWriter, r *http.Request, userId google_uuid.UUID)
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler            ServerInterface
	HandlerMiddlewares []MiddlewareFunc
	ErrorHandlerFunc   func(w http.ResponseWriter, r *http.Request, err error)
}

type MiddlewareFunc func(http.Handler) http.Handler

// GetAccount operation middleware
func (siw *ServerInterfaceWrapper) GetAccount(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "user_id" -------------
	var userId google_uuid.UUID

	err = runtime.BindStyledParameterWithLocation("simple", false, "user_id", runtime.ParamLocationPath, chi.URLParam(r, "user_id"), &userId)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "user_id", Err: err})
		return
	}

	var handler http.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetAccount(w, r, userId)
	})

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// DepositFunds operation middleware
func (siw *ServerInterfaceWrapper) DepositFunds(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "user_id" -------------
	var userId google_uuid.UUID

	err = runtime.BindStyledParameterWithLocation("simple", false, "user_id", runtime.ParamLocationPath, chi.URLParam(r, "user_id"), &userId)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "user_id", Err: err})
		return
	}

	var handler http.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.DepositFunds(w, r, userId)
	})

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// GetAccountTransactions operation middleware
func (siw *ServerInterfaceWrapper) GetAccountTransactions(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "user_id" -------------
	var userId google_uuid.UUID

	err = runtime.BindStyledParameterWithLocation("simple", false, "user_id", runtime.ParamLocationPath, chi.URLParam(r, "user_id"), &userId)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "user_id", Err: err})
		return
	}

	var handler http.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetAccountTransactions(w, r, userId)
	})

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// TransferFunds operation middleware
func (siw *ServerInterfaceWrapper) TransferFunds(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "user_id" -------------
	var userId google_uuid.UUID

	err = runtime.BindStyledParameterWithLocation("simple", false, "user_id", runtime.ParamLocationPath, chi.URLParam(r, "user_id"), &userId)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "user_id", Err: err})
		return
	}

	var handler http.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.TransferFunds(w, r, userId)
	})

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// GetRevenueReport operation middleware
func (siw *ServerInterfaceWrapper) GetRevenueReport(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "month" -------------
	var month string

	err = runtime.BindStyledParameterWithLocation("simple", false, "month", runtime.ParamLocationPath, chi.URLParam(r, "month"), &month)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "month", Err: err})
		return
	}

	var handler http.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetRevenueReport(w, r, month)
	})

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// ReserveFunds operation middleware
func (siw *ServerInterfaceWrapper) ReserveFunds(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "user_id" -------------
	var userId google_uuid.UUID

	err = runtime.BindStyledParameterWithLocation("simple", false, "user_id", runtime.ParamLocationPath, chi.URLParam(r, "user_id"), &userId)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "user_id", Err: err})
		return
	}

	var handler http.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.ReserveFunds(w, r, userId)
	})

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// DeclareRevenue operation middleware
func (siw *ServerInterfaceWrapper) DeclareRevenue(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "user_id" -------------
	var userId google_uuid.UUID

	err = runtime.BindStyledParameterWithLocation("simple", false, "user_id", runtime.ParamLocationPath, chi.URLParam(r, "user_id"), &userId)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "user_id", Err: err})
		return
	}

	var handler http.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.DeclareRevenue(w, r, userId)
	})

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

type UnescapedCookieParamError struct {
	ParamName string
	Err       error
}

func (e *UnescapedCookieParamError) Error() string {
	return fmt.Sprintf("error unescaping cookie parameter '%s'", e.ParamName)
}

func (e *UnescapedCookieParamError) Unwrap() error {
	return e.Err
}

type UnmarshallingParamError struct {
	ParamName string
	Err       error
}

func (e *UnmarshallingParamError) Error() string {
	return fmt.Sprintf("Error unmarshalling parameter %s as JSON: %s", e.ParamName, e.Err.Error())
}

func (e *UnmarshallingParamError) Unwrap() error {
	return e.Err
}

type RequiredParamError struct {
	ParamName string
}

func (e *RequiredParamError) Error() string {
	return fmt.Sprintf("Query argument %s is required, but not found", e.ParamName)
}

type RequiredHeaderError struct {
	ParamName string
	Err       error
}

func (e *RequiredHeaderError) Error() string {
	return fmt.Sprintf("Header parameter %s is required, but not found", e.ParamName)
}

func (e *RequiredHeaderError) Unwrap() error {
	return e.Err
}

type InvalidParamFormatError struct {
	ParamName string
	Err       error
}

func (e *InvalidParamFormatError) Error() string {
	return fmt.Sprintf("Invalid format for parameter %s: %s", e.ParamName, e.Err.Error())
}

func (e *InvalidParamFormatError) Unwrap() error {
	return e.Err
}

type TooManyValuesForParamError struct {
	ParamName string
	Count     int
}

func (e *TooManyValuesForParamError) Error() string {
	return fmt.Sprintf("Expected one value for %s, got %d", e.ParamName, e.Count)
}

// Handler creates http.Handler with routing matching OpenAPI spec.
func Handler(si ServerInterface) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{})
}

type ChiServerOptions struct {
	BaseURL          string
	BaseRouter       chi.Router
	Middlewares      []MiddlewareFunc
	ErrorHandlerFunc func(w http.ResponseWriter, r *http.Request, err error)
}

// HandlerFromMux creates http.Handler with routing matching OpenAPI spec based on the provided mux.
func HandlerFromMux(si ServerInterface, r chi.Router) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{
		BaseRouter: r,
	})
}

func HandlerFromMuxWithBaseURL(si ServerInterface, r chi.Router, baseURL string) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{
		BaseURL:    baseURL,
		BaseRouter: r,
	})
}

// HandlerWithOptions creates http.Handler with additional options
func HandlerWithOptions(si ServerInterface, options ChiServerOptions) http.Handler {
	r := options.BaseRouter

	if r == nil {
		r = chi.NewRouter()
	}
	if options.ErrorHandlerFunc == nil {
		options.ErrorHandlerFunc = func(w http.ResponseWriter, r *http.Request, err error) {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	}
	wrapper := ServerInterfaceWrapper{
		Handler:            si,
		HandlerMiddlewares: options.Middlewares,
		ErrorHandlerFunc:   options.ErrorHandlerFunc,
	}

	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/accounts/{user_id}", wrapper.GetAccount)
	})
	r.Group(func(r chi.Router) {
		r.Put(options.BaseURL+"/accounts/{user_id}", wrapper.DepositFunds)
	})
	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/accounts/{user_id}/transactions", wrapper.GetAccountTransactions)
	})
	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/accounts/{user_id}/transfer", wrapper.TransferFunds)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/reports/{month}", wrapper.GetRevenueReport)
	})
	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/reservations/{user_id}", wrapper.ReserveFunds)
	})
	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/reservations/{user_id}/revenue", wrapper.DeclareRevenue)
	})

	return r
}
