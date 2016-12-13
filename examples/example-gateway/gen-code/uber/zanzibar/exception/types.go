// Code generated by thriftrw v1.1.0
// @generated

package exception

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"go.uber.org/thriftrw/wire"
	"math"
	"strconv"
	"strings"
)

type BadRequest struct {
	Code    BadRequestCode `json:"code"`
	Message string         `json:"message"`
}

func (v *BadRequest) ToWire() (wire.Value, error) {
	var (
		fields [2]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)
	w, err = v.Code.ToWire()
	if err != nil {
		return w, err
	}
	fields[i] = wire.Field{ID: 1, Value: w}
	i++
	w, err = wire.NewValueString(v.Message), error(nil)
	if err != nil {
		return w, err
	}
	fields[i] = wire.Field{ID: 2, Value: w}
	i++
	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

func _BadRequestCode_Read(w wire.Value) (BadRequestCode, error) {
	var v BadRequestCode
	err := v.FromWire(w)
	return v, err
}

func (v *BadRequest) FromWire(w wire.Value) error {
	var err error
	codeIsSet := false
	messageIsSet := false
	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 1:
			if field.Value.Type() == wire.TI32 {
				v.Code, err = _BadRequestCode_Read(field.Value)
				if err != nil {
					return err
				}
				codeIsSet = true
			}
		case 2:
			if field.Value.Type() == wire.TBinary {
				v.Message, err = field.Value.GetString(), error(nil)
				if err != nil {
					return err
				}
				messageIsSet = true
			}
		}
	}
	if !codeIsSet {
		return errors.New("field Code of BadRequest is required")
	}
	if !messageIsSet {
		return errors.New("field Message of BadRequest is required")
	}
	return nil
}

func (v *BadRequest) String() string {
	var fields [2]string
	i := 0
	fields[i] = fmt.Sprintf("Code: %v", v.Code)
	i++
	fields[i] = fmt.Sprintf("Message: %v", v.Message)
	i++
	return fmt.Sprintf("BadRequest{%v}", strings.Join(fields[:i], ", "))
}

func (v *BadRequest) Error() string {
	return v.String()
}

type BadRequestCode int32

const (
	BadRequestCodeZanzibarBadRequest BadRequestCode = 0
)

func (v BadRequestCode) ToWire() (wire.Value, error) {
	return wire.NewValueI32(int32(v)), nil
}

func (v *BadRequestCode) FromWire(w wire.Value) error {
	*v = (BadRequestCode)(w.GetI32())
	return nil
}

func (v BadRequestCode) String() string {
	w := int32(v)
	switch w {
	case 0:
		return "zanzibar__bad_request"
	}
	return fmt.Sprintf("BadRequestCode(%d)", w)
}

func (v BadRequestCode) MarshalJSON() ([]byte, error) {
	switch int32(v) {
	case 0:
		return ([]byte)("\"zanzibar__bad_request\""), nil
	}
	return ([]byte)(strconv.FormatInt(int64(v), 10)), nil
}

func (v *BadRequestCode) UnmarshalJSON(text []byte) error {
	d := json.NewDecoder(bytes.NewReader(text))
	d.UseNumber()
	t, err := d.Token()
	if err != nil {
		return err
	}
	switch w := t.(type) {
	case json.Number:
		x, err := w.Int64()
		if err != nil {
			return err
		}
		if x > math.MaxInt32 {
			return fmt.Errorf("enum overflow from JSON %q for %q", text, "BadRequestCode")
		}
		if x < math.MinInt32 {
			return fmt.Errorf("enum underflow from JSON %q for %q", text, "BadRequestCode")
		}
		*v = (BadRequestCode)(x)
		return nil
	case string:
		switch w {
		case "zanzibar__bad_request":
			*v = BadRequestCodeZanzibarBadRequest
			return nil
		default:
			return fmt.Errorf("unknown enum value %q for %q", w, "BadRequestCode")
		}
	default:
		return fmt.Errorf("invalid JSON value %q (%T) to unmarshal into %q", t, t, "BadRequestCode")
	}
}

type InternalServerErrorCode int32

const (
	InternalServerErrorCodeZanzibarInternalServerError InternalServerErrorCode = 0
)

func (v InternalServerErrorCode) ToWire() (wire.Value, error) {
	return wire.NewValueI32(int32(v)), nil
}

func (v *InternalServerErrorCode) FromWire(w wire.Value) error {
	*v = (InternalServerErrorCode)(w.GetI32())
	return nil
}

func (v InternalServerErrorCode) String() string {
	w := int32(v)
	switch w {
	case 0:
		return "zanzibar__internal_server_error"
	}
	return fmt.Sprintf("InternalServerErrorCode(%d)", w)
}

func (v InternalServerErrorCode) MarshalJSON() ([]byte, error) {
	switch int32(v) {
	case 0:
		return ([]byte)("\"zanzibar__internal_server_error\""), nil
	}
	return ([]byte)(strconv.FormatInt(int64(v), 10)), nil
}

func (v *InternalServerErrorCode) UnmarshalJSON(text []byte) error {
	d := json.NewDecoder(bytes.NewReader(text))
	d.UseNumber()
	t, err := d.Token()
	if err != nil {
		return err
	}
	switch w := t.(type) {
	case json.Number:
		x, err := w.Int64()
		if err != nil {
			return err
		}
		if x > math.MaxInt32 {
			return fmt.Errorf("enum overflow from JSON %q for %q", text, "InternalServerErrorCode")
		}
		if x < math.MinInt32 {
			return fmt.Errorf("enum underflow from JSON %q for %q", text, "InternalServerErrorCode")
		}
		*v = (InternalServerErrorCode)(x)
		return nil
	case string:
		switch w {
		case "zanzibar__internal_server_error":
			*v = InternalServerErrorCodeZanzibarInternalServerError
			return nil
		default:
			return fmt.Errorf("unknown enum value %q for %q", w, "InternalServerErrorCode")
		}
	default:
		return fmt.Errorf("invalid JSON value %q (%T) to unmarshal into %q", t, t, "InternalServerErrorCode")
	}
}

type NotAvailable struct {
	Message string           `json:"message"`
	Code    NotAvailableCode `json:"code"`
}

func (v *NotAvailable) ToWire() (wire.Value, error) {
	var (
		fields [2]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)
	w, err = wire.NewValueString(v.Message), error(nil)
	if err != nil {
		return w, err
	}
	fields[i] = wire.Field{ID: 1, Value: w}
	i++
	w, err = v.Code.ToWire()
	if err != nil {
		return w, err
	}
	fields[i] = wire.Field{ID: 2, Value: w}
	i++
	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

func _NotAvailableCode_Read(w wire.Value) (NotAvailableCode, error) {
	var v NotAvailableCode
	err := v.FromWire(w)
	return v, err
}

func (v *NotAvailable) FromWire(w wire.Value) error {
	var err error
	messageIsSet := false
	codeIsSet := false
	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 1:
			if field.Value.Type() == wire.TBinary {
				v.Message, err = field.Value.GetString(), error(nil)
				if err != nil {
					return err
				}
				messageIsSet = true
			}
		case 2:
			if field.Value.Type() == wire.TI32 {
				v.Code, err = _NotAvailableCode_Read(field.Value)
				if err != nil {
					return err
				}
				codeIsSet = true
			}
		}
	}
	if !messageIsSet {
		return errors.New("field Message of NotAvailable is required")
	}
	if !codeIsSet {
		return errors.New("field Code of NotAvailable is required")
	}
	return nil
}

func (v *NotAvailable) String() string {
	var fields [2]string
	i := 0
	fields[i] = fmt.Sprintf("Message: %v", v.Message)
	i++
	fields[i] = fmt.Sprintf("Code: %v", v.Code)
	i++
	return fmt.Sprintf("NotAvailable{%v}", strings.Join(fields[:i], ", "))
}

func (v *NotAvailable) Error() string {
	return v.String()
}

type NotAvailableCode int32

const (
	NotAvailableCodeZanzibarUsersNotAvailable NotAvailableCode = 0
)

func (v NotAvailableCode) ToWire() (wire.Value, error) {
	return wire.NewValueI32(int32(v)), nil
}

func (v *NotAvailableCode) FromWire(w wire.Value) error {
	*v = (NotAvailableCode)(w.GetI32())
	return nil
}

func (v NotAvailableCode) String() string {
	w := int32(v)
	switch w {
	case 0:
		return "zanzibar__users__not_available"
	}
	return fmt.Sprintf("NotAvailableCode(%d)", w)
}

func (v NotAvailableCode) MarshalJSON() ([]byte, error) {
	switch int32(v) {
	case 0:
		return ([]byte)("\"zanzibar__users__not_available\""), nil
	}
	return ([]byte)(strconv.FormatInt(int64(v), 10)), nil
}

func (v *NotAvailableCode) UnmarshalJSON(text []byte) error {
	d := json.NewDecoder(bytes.NewReader(text))
	d.UseNumber()
	t, err := d.Token()
	if err != nil {
		return err
	}
	switch w := t.(type) {
	case json.Number:
		x, err := w.Int64()
		if err != nil {
			return err
		}
		if x > math.MaxInt32 {
			return fmt.Errorf("enum overflow from JSON %q for %q", text, "NotAvailableCode")
		}
		if x < math.MinInt32 {
			return fmt.Errorf("enum underflow from JSON %q for %q", text, "NotAvailableCode")
		}
		*v = (NotAvailableCode)(x)
		return nil
	case string:
		switch w {
		case "zanzibar__users__not_available":
			*v = NotAvailableCodeZanzibarUsersNotAvailable
			return nil
		default:
			return fmt.Errorf("unknown enum value %q for %q", w, "NotAvailableCode")
		}
	default:
		return fmt.Errorf("invalid JSON value %q (%T) to unmarshal into %q", t, t, "NotAvailableCode")
	}
}

type NotFound struct {
	Code    NotFoundCode `json:"code"`
	Message string       `json:"message"`
}

func (v *NotFound) ToWire() (wire.Value, error) {
	var (
		fields [2]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)
	w, err = v.Code.ToWire()
	if err != nil {
		return w, err
	}
	fields[i] = wire.Field{ID: 1, Value: w}
	i++
	w, err = wire.NewValueString(v.Message), error(nil)
	if err != nil {
		return w, err
	}
	fields[i] = wire.Field{ID: 2, Value: w}
	i++
	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

func _NotFoundCode_Read(w wire.Value) (NotFoundCode, error) {
	var v NotFoundCode
	err := v.FromWire(w)
	return v, err
}

func (v *NotFound) FromWire(w wire.Value) error {
	var err error
	codeIsSet := false
	messageIsSet := false
	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 1:
			if field.Value.Type() == wire.TI32 {
				v.Code, err = _NotFoundCode_Read(field.Value)
				if err != nil {
					return err
				}
				codeIsSet = true
			}
		case 2:
			if field.Value.Type() == wire.TBinary {
				v.Message, err = field.Value.GetString(), error(nil)
				if err != nil {
					return err
				}
				messageIsSet = true
			}
		}
	}
	if !codeIsSet {
		return errors.New("field Code of NotFound is required")
	}
	if !messageIsSet {
		return errors.New("field Message of NotFound is required")
	}
	return nil
}

func (v *NotFound) String() string {
	var fields [2]string
	i := 0
	fields[i] = fmt.Sprintf("Code: %v", v.Code)
	i++
	fields[i] = fmt.Sprintf("Message: %v", v.Message)
	i++
	return fmt.Sprintf("NotFound{%v}", strings.Join(fields[:i], ", "))
}

func (v *NotFound) Error() string {
	return v.String()
}

type NotFoundCode int32

const (
	NotFoundCodeZanzibarNotFound NotFoundCode = 0
)

func (v NotFoundCode) ToWire() (wire.Value, error) {
	return wire.NewValueI32(int32(v)), nil
}

func (v *NotFoundCode) FromWire(w wire.Value) error {
	*v = (NotFoundCode)(w.GetI32())
	return nil
}

func (v NotFoundCode) String() string {
	w := int32(v)
	switch w {
	case 0:
		return "zanzibar__not_found"
	}
	return fmt.Sprintf("NotFoundCode(%d)", w)
}

func (v NotFoundCode) MarshalJSON() ([]byte, error) {
	switch int32(v) {
	case 0:
		return ([]byte)("\"zanzibar__not_found\""), nil
	}
	return ([]byte)(strconv.FormatInt(int64(v), 10)), nil
}

func (v *NotFoundCode) UnmarshalJSON(text []byte) error {
	d := json.NewDecoder(bytes.NewReader(text))
	d.UseNumber()
	t, err := d.Token()
	if err != nil {
		return err
	}
	switch w := t.(type) {
	case json.Number:
		x, err := w.Int64()
		if err != nil {
			return err
		}
		if x > math.MaxInt32 {
			return fmt.Errorf("enum overflow from JSON %q for %q", text, "NotFoundCode")
		}
		if x < math.MinInt32 {
			return fmt.Errorf("enum underflow from JSON %q for %q", text, "NotFoundCode")
		}
		*v = (NotFoundCode)(x)
		return nil
	case string:
		switch w {
		case "zanzibar__not_found":
			*v = NotFoundCodeZanzibarNotFound
			return nil
		default:
			return fmt.Errorf("unknown enum value %q for %q", w, "NotFoundCode")
		}
	default:
		return fmt.Errorf("invalid JSON value %q (%T) to unmarshal into %q", t, t, "NotFoundCode")
	}
}

type PayloadLimited struct {
	Code    PayloadLimitedCode `json:"code"`
	Message string             `json:"message"`
}

func (v *PayloadLimited) ToWire() (wire.Value, error) {
	var (
		fields [2]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)
	w, err = v.Code.ToWire()
	if err != nil {
		return w, err
	}
	fields[i] = wire.Field{ID: 1, Value: w}
	i++
	w, err = wire.NewValueString(v.Message), error(nil)
	if err != nil {
		return w, err
	}
	fields[i] = wire.Field{ID: 2, Value: w}
	i++
	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

func _PayloadLimitedCode_Read(w wire.Value) (PayloadLimitedCode, error) {
	var v PayloadLimitedCode
	err := v.FromWire(w)
	return v, err
}

func (v *PayloadLimited) FromWire(w wire.Value) error {
	var err error
	codeIsSet := false
	messageIsSet := false
	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 1:
			if field.Value.Type() == wire.TI32 {
				v.Code, err = _PayloadLimitedCode_Read(field.Value)
				if err != nil {
					return err
				}
				codeIsSet = true
			}
		case 2:
			if field.Value.Type() == wire.TBinary {
				v.Message, err = field.Value.GetString(), error(nil)
				if err != nil {
					return err
				}
				messageIsSet = true
			}
		}
	}
	if !codeIsSet {
		return errors.New("field Code of PayloadLimited is required")
	}
	if !messageIsSet {
		return errors.New("field Message of PayloadLimited is required")
	}
	return nil
}

func (v *PayloadLimited) String() string {
	var fields [2]string
	i := 0
	fields[i] = fmt.Sprintf("Code: %v", v.Code)
	i++
	fields[i] = fmt.Sprintf("Message: %v", v.Message)
	i++
	return fmt.Sprintf("PayloadLimited{%v}", strings.Join(fields[:i], ", "))
}

func (v *PayloadLimited) Error() string {
	return v.String()
}

type PayloadLimitedCode int32

const (
	PayloadLimitedCodeZanzibarPayloadTooLarge PayloadLimitedCode = 0
)

func (v PayloadLimitedCode) ToWire() (wire.Value, error) {
	return wire.NewValueI32(int32(v)), nil
}

func (v *PayloadLimitedCode) FromWire(w wire.Value) error {
	*v = (PayloadLimitedCode)(w.GetI32())
	return nil
}

func (v PayloadLimitedCode) String() string {
	w := int32(v)
	switch w {
	case 0:
		return "zanzibar__payload_too_large"
	}
	return fmt.Sprintf("PayloadLimitedCode(%d)", w)
}

func (v PayloadLimitedCode) MarshalJSON() ([]byte, error) {
	switch int32(v) {
	case 0:
		return ([]byte)("\"zanzibar__payload_too_large\""), nil
	}
	return ([]byte)(strconv.FormatInt(int64(v), 10)), nil
}

func (v *PayloadLimitedCode) UnmarshalJSON(text []byte) error {
	d := json.NewDecoder(bytes.NewReader(text))
	d.UseNumber()
	t, err := d.Token()
	if err != nil {
		return err
	}
	switch w := t.(type) {
	case json.Number:
		x, err := w.Int64()
		if err != nil {
			return err
		}
		if x > math.MaxInt32 {
			return fmt.Errorf("enum overflow from JSON %q for %q", text, "PayloadLimitedCode")
		}
		if x < math.MinInt32 {
			return fmt.Errorf("enum underflow from JSON %q for %q", text, "PayloadLimitedCode")
		}
		*v = (PayloadLimitedCode)(x)
		return nil
	case string:
		switch w {
		case "zanzibar__payload_too_large":
			*v = PayloadLimitedCodeZanzibarPayloadTooLarge
			return nil
		default:
			return fmt.Errorf("unknown enum value %q for %q", w, "PayloadLimitedCode")
		}
	default:
		return fmt.Errorf("invalid JSON value %q (%T) to unmarshal into %q", t, t, "PayloadLimitedCode")
	}
}

type PermissionDenied struct {
	Code    PermissionDeniedCode `json:"code"`
	Message string               `json:"message"`
}

func (v *PermissionDenied) ToWire() (wire.Value, error) {
	var (
		fields [2]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)
	w, err = v.Code.ToWire()
	if err != nil {
		return w, err
	}
	fields[i] = wire.Field{ID: 1, Value: w}
	i++
	w, err = wire.NewValueString(v.Message), error(nil)
	if err != nil {
		return w, err
	}
	fields[i] = wire.Field{ID: 2, Value: w}
	i++
	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

func _PermissionDeniedCode_Read(w wire.Value) (PermissionDeniedCode, error) {
	var v PermissionDeniedCode
	err := v.FromWire(w)
	return v, err
}

func (v *PermissionDenied) FromWire(w wire.Value) error {
	var err error
	codeIsSet := false
	messageIsSet := false
	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 1:
			if field.Value.Type() == wire.TI32 {
				v.Code, err = _PermissionDeniedCode_Read(field.Value)
				if err != nil {
					return err
				}
				codeIsSet = true
			}
		case 2:
			if field.Value.Type() == wire.TBinary {
				v.Message, err = field.Value.GetString(), error(nil)
				if err != nil {
					return err
				}
				messageIsSet = true
			}
		}
	}
	if !codeIsSet {
		return errors.New("field Code of PermissionDenied is required")
	}
	if !messageIsSet {
		return errors.New("field Message of PermissionDenied is required")
	}
	return nil
}

func (v *PermissionDenied) String() string {
	var fields [2]string
	i := 0
	fields[i] = fmt.Sprintf("Code: %v", v.Code)
	i++
	fields[i] = fmt.Sprintf("Message: %v", v.Message)
	i++
	return fmt.Sprintf("PermissionDenied{%v}", strings.Join(fields[:i], ", "))
}

func (v *PermissionDenied) Error() string {
	return v.String()
}

type PermissionDeniedCode int32

const (
	PermissionDeniedCodeZanzibarPermissionDenied PermissionDeniedCode = 0
)

func (v PermissionDeniedCode) ToWire() (wire.Value, error) {
	return wire.NewValueI32(int32(v)), nil
}

func (v *PermissionDeniedCode) FromWire(w wire.Value) error {
	*v = (PermissionDeniedCode)(w.GetI32())
	return nil
}

func (v PermissionDeniedCode) String() string {
	w := int32(v)
	switch w {
	case 0:
		return "zanzibar__permission_denied"
	}
	return fmt.Sprintf("PermissionDeniedCode(%d)", w)
}

func (v PermissionDeniedCode) MarshalJSON() ([]byte, error) {
	switch int32(v) {
	case 0:
		return ([]byte)("\"zanzibar__permission_denied\""), nil
	}
	return ([]byte)(strconv.FormatInt(int64(v), 10)), nil
}

func (v *PermissionDeniedCode) UnmarshalJSON(text []byte) error {
	d := json.NewDecoder(bytes.NewReader(text))
	d.UseNumber()
	t, err := d.Token()
	if err != nil {
		return err
	}
	switch w := t.(type) {
	case json.Number:
		x, err := w.Int64()
		if err != nil {
			return err
		}
		if x > math.MaxInt32 {
			return fmt.Errorf("enum overflow from JSON %q for %q", text, "PermissionDeniedCode")
		}
		if x < math.MinInt32 {
			return fmt.Errorf("enum underflow from JSON %q for %q", text, "PermissionDeniedCode")
		}
		*v = (PermissionDeniedCode)(x)
		return nil
	case string:
		switch w {
		case "zanzibar__permission_denied":
			*v = PermissionDeniedCodeZanzibarPermissionDenied
			return nil
		default:
			return fmt.Errorf("unknown enum value %q for %q", w, "PermissionDeniedCode")
		}
	default:
		return fmt.Errorf("invalid JSON value %q (%T) to unmarshal into %q", t, t, "PermissionDeniedCode")
	}
}

type RateLimited struct {
	Code    RateLimitedCode `json:"code"`
	Message string          `json:"message"`
}

func (v *RateLimited) ToWire() (wire.Value, error) {
	var (
		fields [2]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)
	w, err = v.Code.ToWire()
	if err != nil {
		return w, err
	}
	fields[i] = wire.Field{ID: 1, Value: w}
	i++
	w, err = wire.NewValueString(v.Message), error(nil)
	if err != nil {
		return w, err
	}
	fields[i] = wire.Field{ID: 2, Value: w}
	i++
	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

func _RateLimitedCode_Read(w wire.Value) (RateLimitedCode, error) {
	var v RateLimitedCode
	err := v.FromWire(w)
	return v, err
}

func (v *RateLimited) FromWire(w wire.Value) error {
	var err error
	codeIsSet := false
	messageIsSet := false
	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 1:
			if field.Value.Type() == wire.TI32 {
				v.Code, err = _RateLimitedCode_Read(field.Value)
				if err != nil {
					return err
				}
				codeIsSet = true
			}
		case 2:
			if field.Value.Type() == wire.TBinary {
				v.Message, err = field.Value.GetString(), error(nil)
				if err != nil {
					return err
				}
				messageIsSet = true
			}
		}
	}
	if !codeIsSet {
		return errors.New("field Code of RateLimited is required")
	}
	if !messageIsSet {
		return errors.New("field Message of RateLimited is required")
	}
	return nil
}

func (v *RateLimited) String() string {
	var fields [2]string
	i := 0
	fields[i] = fmt.Sprintf("Code: %v", v.Code)
	i++
	fields[i] = fmt.Sprintf("Message: %v", v.Message)
	i++
	return fmt.Sprintf("RateLimited{%v}", strings.Join(fields[:i], ", "))
}

func (v *RateLimited) Error() string {
	return v.String()
}

type RateLimitedCode int32

const (
	RateLimitedCodeZanzibarTooManyRequests RateLimitedCode = 0
)

func (v RateLimitedCode) ToWire() (wire.Value, error) {
	return wire.NewValueI32(int32(v)), nil
}

func (v *RateLimitedCode) FromWire(w wire.Value) error {
	*v = (RateLimitedCode)(w.GetI32())
	return nil
}

func (v RateLimitedCode) String() string {
	w := int32(v)
	switch w {
	case 0:
		return "zanzibar__too_many_requests"
	}
	return fmt.Sprintf("RateLimitedCode(%d)", w)
}

func (v RateLimitedCode) MarshalJSON() ([]byte, error) {
	switch int32(v) {
	case 0:
		return ([]byte)("\"zanzibar__too_many_requests\""), nil
	}
	return ([]byte)(strconv.FormatInt(int64(v), 10)), nil
}

func (v *RateLimitedCode) UnmarshalJSON(text []byte) error {
	d := json.NewDecoder(bytes.NewReader(text))
	d.UseNumber()
	t, err := d.Token()
	if err != nil {
		return err
	}
	switch w := t.(type) {
	case json.Number:
		x, err := w.Int64()
		if err != nil {
			return err
		}
		if x > math.MaxInt32 {
			return fmt.Errorf("enum overflow from JSON %q for %q", text, "RateLimitedCode")
		}
		if x < math.MinInt32 {
			return fmt.Errorf("enum underflow from JSON %q for %q", text, "RateLimitedCode")
		}
		*v = (RateLimitedCode)(x)
		return nil
	case string:
		switch w {
		case "zanzibar__too_many_requests":
			*v = RateLimitedCodeZanzibarTooManyRequests
			return nil
		default:
			return fmt.Errorf("unknown enum value %q for %q", w, "RateLimitedCode")
		}
	default:
		return fmt.Errorf("invalid JSON value %q (%T) to unmarshal into %q", t, t, "RateLimitedCode")
	}
}

type ServerError struct {
	Code    InternalServerErrorCode `json:"code"`
	Message string                  `json:"message"`
}

func (v *ServerError) ToWire() (wire.Value, error) {
	var (
		fields [2]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)
	w, err = v.Code.ToWire()
	if err != nil {
		return w, err
	}
	fields[i] = wire.Field{ID: 1, Value: w}
	i++
	w, err = wire.NewValueString(v.Message), error(nil)
	if err != nil {
		return w, err
	}
	fields[i] = wire.Field{ID: 2, Value: w}
	i++
	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

func _InternalServerErrorCode_Read(w wire.Value) (InternalServerErrorCode, error) {
	var v InternalServerErrorCode
	err := v.FromWire(w)
	return v, err
}

func (v *ServerError) FromWire(w wire.Value) error {
	var err error
	codeIsSet := false
	messageIsSet := false
	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 1:
			if field.Value.Type() == wire.TI32 {
				v.Code, err = _InternalServerErrorCode_Read(field.Value)
				if err != nil {
					return err
				}
				codeIsSet = true
			}
		case 2:
			if field.Value.Type() == wire.TBinary {
				v.Message, err = field.Value.GetString(), error(nil)
				if err != nil {
					return err
				}
				messageIsSet = true
			}
		}
	}
	if !codeIsSet {
		return errors.New("field Code of ServerError is required")
	}
	if !messageIsSet {
		return errors.New("field Message of ServerError is required")
	}
	return nil
}

func (v *ServerError) String() string {
	var fields [2]string
	i := 0
	fields[i] = fmt.Sprintf("Code: %v", v.Code)
	i++
	fields[i] = fmt.Sprintf("Message: %v", v.Message)
	i++
	return fmt.Sprintf("ServerError{%v}", strings.Join(fields[:i], ", "))
}

func (v *ServerError) Error() string {
	return v.String()
}

type TemporaryRedirect struct {
	Location    string                `json:"location"`
	Code        TemporaryRedirectCode `json:"code"`
	MessageType *string               `json:"messageType,omitempty"`
	URI         *string               `json:"uri,omitempty"`
}

func (v *TemporaryRedirect) ToWire() (wire.Value, error) {
	var (
		fields [4]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)
	w, err = wire.NewValueString(v.Location), error(nil)
	if err != nil {
		return w, err
	}
	fields[i] = wire.Field{ID: 1, Value: w}
	i++
	w, err = v.Code.ToWire()
	if err != nil {
		return w, err
	}
	fields[i] = wire.Field{ID: 2, Value: w}
	i++
	if v.MessageType != nil {
		w, err = wire.NewValueString(*(v.MessageType)), error(nil)
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 3, Value: w}
		i++
	}
	if v.URI != nil {
		w, err = wire.NewValueString(*(v.URI)), error(nil)
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 4, Value: w}
		i++
	}
	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

func _TemporaryRedirectCode_Read(w wire.Value) (TemporaryRedirectCode, error) {
	var v TemporaryRedirectCode
	err := v.FromWire(w)
	return v, err
}

func (v *TemporaryRedirect) FromWire(w wire.Value) error {
	var err error
	locationIsSet := false
	codeIsSet := false
	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 1:
			if field.Value.Type() == wire.TBinary {
				v.Location, err = field.Value.GetString(), error(nil)
				if err != nil {
					return err
				}
				locationIsSet = true
			}
		case 2:
			if field.Value.Type() == wire.TI32 {
				v.Code, err = _TemporaryRedirectCode_Read(field.Value)
				if err != nil {
					return err
				}
				codeIsSet = true
			}
		case 3:
			if field.Value.Type() == wire.TBinary {
				var x string
				x, err = field.Value.GetString(), error(nil)
				v.MessageType = &x
				if err != nil {
					return err
				}
			}
		case 4:
			if field.Value.Type() == wire.TBinary {
				var x string
				x, err = field.Value.GetString(), error(nil)
				v.URI = &x
				if err != nil {
					return err
				}
			}
		}
	}
	if !locationIsSet {
		return errors.New("field Location of TemporaryRedirect is required")
	}
	if !codeIsSet {
		return errors.New("field Code of TemporaryRedirect is required")
	}
	return nil
}

func (v *TemporaryRedirect) String() string {
	var fields [4]string
	i := 0
	fields[i] = fmt.Sprintf("Location: %v", v.Location)
	i++
	fields[i] = fmt.Sprintf("Code: %v", v.Code)
	i++
	if v.MessageType != nil {
		fields[i] = fmt.Sprintf("MessageType: %v", *(v.MessageType))
		i++
	}
	if v.URI != nil {
		fields[i] = fmt.Sprintf("URI: %v", *(v.URI))
		i++
	}
	return fmt.Sprintf("TemporaryRedirect{%v}", strings.Join(fields[:i], ", "))
}

func (v *TemporaryRedirect) Error() string {
	return v.String()
}

type TemporaryRedirectCode int32

const (
	TemporaryRedirectCodeZanzibarDatacenterRedirect TemporaryRedirectCode = 0
)

func (v TemporaryRedirectCode) ToWire() (wire.Value, error) {
	return wire.NewValueI32(int32(v)), nil
}

func (v *TemporaryRedirectCode) FromWire(w wire.Value) error {
	*v = (TemporaryRedirectCode)(w.GetI32())
	return nil
}

func (v TemporaryRedirectCode) String() string {
	w := int32(v)
	switch w {
	case 0:
		return "zanzibar__datacenter_redirect"
	}
	return fmt.Sprintf("TemporaryRedirectCode(%d)", w)
}

func (v TemporaryRedirectCode) MarshalJSON() ([]byte, error) {
	switch int32(v) {
	case 0:
		return ([]byte)("\"zanzibar__datacenter_redirect\""), nil
	}
	return ([]byte)(strconv.FormatInt(int64(v), 10)), nil
}

func (v *TemporaryRedirectCode) UnmarshalJSON(text []byte) error {
	d := json.NewDecoder(bytes.NewReader(text))
	d.UseNumber()
	t, err := d.Token()
	if err != nil {
		return err
	}
	switch w := t.(type) {
	case json.Number:
		x, err := w.Int64()
		if err != nil {
			return err
		}
		if x > math.MaxInt32 {
			return fmt.Errorf("enum overflow from JSON %q for %q", text, "TemporaryRedirectCode")
		}
		if x < math.MinInt32 {
			return fmt.Errorf("enum underflow from JSON %q for %q", text, "TemporaryRedirectCode")
		}
		*v = (TemporaryRedirectCode)(x)
		return nil
	case string:
		switch w {
		case "zanzibar__datacenter_redirect":
			*v = TemporaryRedirectCodeZanzibarDatacenterRedirect
			return nil
		default:
			return fmt.Errorf("unknown enum value %q for %q", w, "TemporaryRedirectCode")
		}
	default:
		return fmt.Errorf("invalid JSON value %q (%T) to unmarshal into %q", t, t, "TemporaryRedirectCode")
	}
}

type Unauthenticated struct {
	Code      UnauthenticatedCode `json:"code"`
	Message   string              `json:"message"`
	ErrorCode *string             `json:"errorCode,omitempty"`
}

func (v *Unauthenticated) ToWire() (wire.Value, error) {
	var (
		fields [3]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)
	w, err = v.Code.ToWire()
	if err != nil {
		return w, err
	}
	fields[i] = wire.Field{ID: 1, Value: w}
	i++
	w, err = wire.NewValueString(v.Message), error(nil)
	if err != nil {
		return w, err
	}
	fields[i] = wire.Field{ID: 2, Value: w}
	i++
	if v.ErrorCode != nil {
		w, err = wire.NewValueString(*(v.ErrorCode)), error(nil)
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 3, Value: w}
		i++
	}
	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

func _UnauthenticatedCode_Read(w wire.Value) (UnauthenticatedCode, error) {
	var v UnauthenticatedCode
	err := v.FromWire(w)
	return v, err
}

func (v *Unauthenticated) FromWire(w wire.Value) error {
	var err error
	codeIsSet := false
	messageIsSet := false
	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 1:
			if field.Value.Type() == wire.TI32 {
				v.Code, err = _UnauthenticatedCode_Read(field.Value)
				if err != nil {
					return err
				}
				codeIsSet = true
			}
		case 2:
			if field.Value.Type() == wire.TBinary {
				v.Message, err = field.Value.GetString(), error(nil)
				if err != nil {
					return err
				}
				messageIsSet = true
			}
		case 3:
			if field.Value.Type() == wire.TBinary {
				var x string
				x, err = field.Value.GetString(), error(nil)
				v.ErrorCode = &x
				if err != nil {
					return err
				}
			}
		}
	}
	if !codeIsSet {
		return errors.New("field Code of Unauthenticated is required")
	}
	if !messageIsSet {
		return errors.New("field Message of Unauthenticated is required")
	}
	return nil
}

func (v *Unauthenticated) String() string {
	var fields [3]string
	i := 0
	fields[i] = fmt.Sprintf("Code: %v", v.Code)
	i++
	fields[i] = fmt.Sprintf("Message: %v", v.Message)
	i++
	if v.ErrorCode != nil {
		fields[i] = fmt.Sprintf("ErrorCode: %v", *(v.ErrorCode))
		i++
	}
	return fmt.Sprintf("Unauthenticated{%v}", strings.Join(fields[:i], ", "))
}

func (v *Unauthenticated) Error() string {
	return v.String()
}

type UnauthenticatedCode int32

const (
	UnauthenticatedCodeZanzibarUnauthorized UnauthenticatedCode = 0
)

func (v UnauthenticatedCode) ToWire() (wire.Value, error) {
	return wire.NewValueI32(int32(v)), nil
}

func (v *UnauthenticatedCode) FromWire(w wire.Value) error {
	*v = (UnauthenticatedCode)(w.GetI32())
	return nil
}

func (v UnauthenticatedCode) String() string {
	w := int32(v)
	switch w {
	case 0:
		return "zanzibar__unauthorized"
	}
	return fmt.Sprintf("UnauthenticatedCode(%d)", w)
}

func (v UnauthenticatedCode) MarshalJSON() ([]byte, error) {
	switch int32(v) {
	case 0:
		return ([]byte)("\"zanzibar__unauthorized\""), nil
	}
	return ([]byte)(strconv.FormatInt(int64(v), 10)), nil
}

func (v *UnauthenticatedCode) UnmarshalJSON(text []byte) error {
	d := json.NewDecoder(bytes.NewReader(text))
	d.UseNumber()
	t, err := d.Token()
	if err != nil {
		return err
	}
	switch w := t.(type) {
	case json.Number:
		x, err := w.Int64()
		if err != nil {
			return err
		}
		if x > math.MaxInt32 {
			return fmt.Errorf("enum overflow from JSON %q for %q", text, "UnauthenticatedCode")
		}
		if x < math.MinInt32 {
			return fmt.Errorf("enum underflow from JSON %q for %q", text, "UnauthenticatedCode")
		}
		*v = (UnauthenticatedCode)(x)
		return nil
	case string:
		switch w {
		case "zanzibar__unauthorized":
			*v = UnauthenticatedCodeZanzibarUnauthorized
			return nil
		default:
			return fmt.Errorf("unknown enum value %q for %q", w, "UnauthenticatedCode")
		}
	default:
		return fmt.Errorf("invalid JSON value %q (%T) to unmarshal into %q", t, t, "UnauthenticatedCode")
	}
}

type Unauthorized struct {
	Code    UnauthorizedCode `json:"code"`
	Message string           `json:"message"`
}

func (v *Unauthorized) ToWire() (wire.Value, error) {
	var (
		fields [2]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)
	w, err = v.Code.ToWire()
	if err != nil {
		return w, err
	}
	fields[i] = wire.Field{ID: 1, Value: w}
	i++
	w, err = wire.NewValueString(v.Message), error(nil)
	if err != nil {
		return w, err
	}
	fields[i] = wire.Field{ID: 2, Value: w}
	i++
	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

func _UnauthorizedCode_Read(w wire.Value) (UnauthorizedCode, error) {
	var v UnauthorizedCode
	err := v.FromWire(w)
	return v, err
}

func (v *Unauthorized) FromWire(w wire.Value) error {
	var err error
	codeIsSet := false
	messageIsSet := false
	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 1:
			if field.Value.Type() == wire.TI32 {
				v.Code, err = _UnauthorizedCode_Read(field.Value)
				if err != nil {
					return err
				}
				codeIsSet = true
			}
		case 2:
			if field.Value.Type() == wire.TBinary {
				v.Message, err = field.Value.GetString(), error(nil)
				if err != nil {
					return err
				}
				messageIsSet = true
			}
		}
	}
	if !codeIsSet {
		return errors.New("field Code of Unauthorized is required")
	}
	if !messageIsSet {
		return errors.New("field Message of Unauthorized is required")
	}
	return nil
}

func (v *Unauthorized) String() string {
	var fields [2]string
	i := 0
	fields[i] = fmt.Sprintf("Code: %v", v.Code)
	i++
	fields[i] = fmt.Sprintf("Message: %v", v.Message)
	i++
	return fmt.Sprintf("Unauthorized{%v}", strings.Join(fields[:i], ", "))
}

func (v *Unauthorized) Error() string {
	return v.String()
}

type UnauthorizedCode int32

const (
	UnauthorizedCodeZanzibarForbidden UnauthorizedCode = 0
)

func (v UnauthorizedCode) ToWire() (wire.Value, error) {
	return wire.NewValueI32(int32(v)), nil
}

func (v *UnauthorizedCode) FromWire(w wire.Value) error {
	*v = (UnauthorizedCode)(w.GetI32())
	return nil
}

func (v UnauthorizedCode) String() string {
	w := int32(v)
	switch w {
	case 0:
		return "zanzibar__forbidden"
	}
	return fmt.Sprintf("UnauthorizedCode(%d)", w)
}

func (v UnauthorizedCode) MarshalJSON() ([]byte, error) {
	switch int32(v) {
	case 0:
		return ([]byte)("\"zanzibar__forbidden\""), nil
	}
	return ([]byte)(strconv.FormatInt(int64(v), 10)), nil
}

func (v *UnauthorizedCode) UnmarshalJSON(text []byte) error {
	d := json.NewDecoder(bytes.NewReader(text))
	d.UseNumber()
	t, err := d.Token()
	if err != nil {
		return err
	}
	switch w := t.(type) {
	case json.Number:
		x, err := w.Int64()
		if err != nil {
			return err
		}
		if x > math.MaxInt32 {
			return fmt.Errorf("enum overflow from JSON %q for %q", text, "UnauthorizedCode")
		}
		if x < math.MinInt32 {
			return fmt.Errorf("enum underflow from JSON %q for %q", text, "UnauthorizedCode")
		}
		*v = (UnauthorizedCode)(x)
		return nil
	case string:
		switch w {
		case "zanzibar__forbidden":
			*v = UnauthorizedCodeZanzibarForbidden
			return nil
		default:
			return fmt.Errorf("unknown enum value %q for %q", w, "UnauthorizedCode")
		}
	default:
		return fmt.Errorf("invalid JSON value %q (%T) to unmarshal into %q", t, t, "UnauthorizedCode")
	}
}

type UserBanned struct {
	Code    UserBannedCode `json:"code"`
	Message string         `json:"message"`
}

func (v *UserBanned) ToWire() (wire.Value, error) {
	var (
		fields [2]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)
	w, err = v.Code.ToWire()
	if err != nil {
		return w, err
	}
	fields[i] = wire.Field{ID: 1, Value: w}
	i++
	w, err = wire.NewValueString(v.Message), error(nil)
	if err != nil {
		return w, err
	}
	fields[i] = wire.Field{ID: 2, Value: w}
	i++
	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

func _UserBannedCode_Read(w wire.Value) (UserBannedCode, error) {
	var v UserBannedCode
	err := v.FromWire(w)
	return v, err
}

func (v *UserBanned) FromWire(w wire.Value) error {
	var err error
	codeIsSet := false
	messageIsSet := false
	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 1:
			if field.Value.Type() == wire.TI32 {
				v.Code, err = _UserBannedCode_Read(field.Value)
				if err != nil {
					return err
				}
				codeIsSet = true
			}
		case 2:
			if field.Value.Type() == wire.TBinary {
				v.Message, err = field.Value.GetString(), error(nil)
				if err != nil {
					return err
				}
				messageIsSet = true
			}
		}
	}
	if !codeIsSet {
		return errors.New("field Code of UserBanned is required")
	}
	if !messageIsSet {
		return errors.New("field Message of UserBanned is required")
	}
	return nil
}

func (v *UserBanned) String() string {
	var fields [2]string
	i := 0
	fields[i] = fmt.Sprintf("Code: %v", v.Code)
	i++
	fields[i] = fmt.Sprintf("Message: %v", v.Message)
	i++
	return fmt.Sprintf("UserBanned{%v}", strings.Join(fields[:i], ", "))
}

func (v *UserBanned) Error() string {
	return v.String()
}

type UserBannedCode int32

const (
	UserBannedCodeZanzibarUsersAccountBanned UserBannedCode = 0
)

func (v UserBannedCode) ToWire() (wire.Value, error) {
	return wire.NewValueI32(int32(v)), nil
}

func (v *UserBannedCode) FromWire(w wire.Value) error {
	*v = (UserBannedCode)(w.GetI32())
	return nil
}

func (v UserBannedCode) String() string {
	w := int32(v)
	switch w {
	case 0:
		return "zanzibar__users__account_banned"
	}
	return fmt.Sprintf("UserBannedCode(%d)", w)
}

func (v UserBannedCode) MarshalJSON() ([]byte, error) {
	switch int32(v) {
	case 0:
		return ([]byte)("\"zanzibar__users__account_banned\""), nil
	}
	return ([]byte)(strconv.FormatInt(int64(v), 10)), nil
}

func (v *UserBannedCode) UnmarshalJSON(text []byte) error {
	d := json.NewDecoder(bytes.NewReader(text))
	d.UseNumber()
	t, err := d.Token()
	if err != nil {
		return err
	}
	switch w := t.(type) {
	case json.Number:
		x, err := w.Int64()
		if err != nil {
			return err
		}
		if x > math.MaxInt32 {
			return fmt.Errorf("enum overflow from JSON %q for %q", text, "UserBannedCode")
		}
		if x < math.MinInt32 {
			return fmt.Errorf("enum underflow from JSON %q for %q", text, "UserBannedCode")
		}
		*v = (UserBannedCode)(x)
		return nil
	case string:
		switch w {
		case "zanzibar__users__account_banned":
			*v = UserBannedCodeZanzibarUsersAccountBanned
			return nil
		default:
			return fmt.Errorf("unknown enum value %q for %q", w, "UserBannedCode")
		}
	default:
		return fmt.Errorf("invalid JSON value %q (%T) to unmarshal into %q", t, t, "UserBannedCode")
	}
}
