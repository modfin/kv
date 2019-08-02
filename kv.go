package kv

import (
	"encoding/json"
	"reflect"
)

//KV is a key/value entry and a struct which implements helper methods to help with retrial of data types from value.
type KV struct {
	key   string
	value interface{}
}

func New(key string, value interface{}) *KV {
	return &KV{
		key:   key,
		value: value,
	}
}

func (kv *KV) Bind(inf interface{}) error {
	data, err := json.Marshal(kv.Value)
	if err != nil {
		return err
	}

	return json.Unmarshal(data, inf)
}

var uintType = reflect.TypeOf(uint64(0))
var intType = reflect.TypeOf(int64(0))
var floatType = reflect.TypeOf(float64(0))
var stringType = reflect.TypeOf(string(""))
var boolType = reflect.TypeOf(false)

type converter func(in interface{}) (interface{}, bool)

func untypedFloat(in interface{}) (interface{}, bool) {
	return toFloat(in)
}
func untypedUint(in interface{}) (interface{}, bool) {
	return toUint(in)
}
func untypedInt(in interface{}) (interface{}, bool) {
	return toInt(in)
}
func untypedBool(in interface{}) (res interface{}, ok bool) {
	return toBool(in)
}
func untypedString(in interface{}) (res interface{}, ok bool) {
	return toString(in)
}

//NewParam creates a new key value param from input
func NewParam(key string, value interface{}) *KV {
	return &KV{key, value}
}

// Key returns the key of the key/value pair
func (kv *KV) Key() string {
	return kv.key
}

// Value returns the value of the key/value pair
func (kv *KV) Value() interface{} {
	return kv.value
}

// IsNil returns true if value is nil
func (kv *KV) IsNil() bool {
	return kv.value == nil
}

// IsSlice returns true if value is a array
func (kv *KV) IsSlice() bool {

	if kv.value == nil {
		return false
	}
	return reflect.TypeOf(kv.value).Kind() == reflect.Slice
}

// String returns value as a string, if possible
func (kv *KV) String() (string, bool) {
	return toString(kv.value)
}

// StringOr returns value as a string, otherwise the provided default
func (kv *KV) StringOr(defaultTo string) string {
	str, ok := kv.String()
	if ok {
		return str
	}
	return defaultTo
}

// StringSlice returns value as a []string, if possible
func (kv *KV) StringSlice() ([]string, bool) {
	if kv.value == nil {
		return nil, false
	}

	var res []string
	res, ok := kv.value.([]string)
	if ok {
		return res, true
	}

	r, ok := toSliceOf(kv.value, stringType, untypedString)
	if !ok {
		return nil, false
	}
	res, ok = r.([]string)
	if ok {
		return res, true
	}
	return nil, false
}

// StringSliceOr returns value as a []string, otherwise the provided default
func (kv *KV) StringSliceOr(defaultTo []string) []string {
	arr, ok := kv.StringSlice()

	if ok {
		return arr
	}
	return defaultTo
}

// Uint64 returns value as a uint64, if possible
func (kv *KV) Uint64() (uint64, bool) {
	return toUint(kv.value)
}

// Uint64Or returns value as a uint64, otherwise the provided default
func (kv *KV) Uint64Or(def uint64) uint64 {
	i, ok := kv.Uint64()
	if ok {
		return i
	}
	return def
}

// Uint32 returns value as a uint32, if possible
func (kv *KV) Uint32() (uint32, bool) {
	u, ok := toUint(kv.value)
	return uint32(u), ok
}

// Uint32Or returns value as a uint32, otherwise the provided default
func (kv *KV) Uint32Or(def uint32) uint32 {
	i, ok := kv.Uint32()
	if ok {
		return i
	}
	return def
}

// Uint16 returns value as a uint16, if possible
func (kv *KV) Uint16() (uint16, bool) {
	u, ok := toUint(kv.value)
	return uint16(u), ok
}

// Uint16Or returns value as a uint16, otherwise the provided default
func (kv *KV) Uint16Or(def uint16) uint16 {
	i, ok := kv.Uint16()
	if ok {
		return i
	}
	return def
}


// Uint8 returns value as a uint8, if possible
func (kv *KV) Uint8() (uint8, bool) {
	u, ok := toUint(kv.value)
	return uint8(u), ok
}


// Uint8Or returns value as a uint8, otherwise the provided default
func (kv *KV) Uint8Or(def uint8) uint8 {
	i, ok := kv.Uint8()
	if ok {
		return i
	}
	return def
}


// Uint returns value as a uint, if possible
func (kv *KV) Uint() (uint, bool) {
	u, ok := toUint(kv.value)
	return uint(u), ok
}


// UintOr returns value as a uint, otherwise the provided default
func (kv *KV) UintOr(def uint) uint {
	i, ok := kv.Uint()
	if ok {
		return i
	}
	return def
}


// UintSlice returns value as a []uint64, if possible
func (kv *KV) UintSlice() ([]uint64, bool) {
	if kv.value == nil {
		return nil, false
	}
	var res []uint64
	res, ok := kv.value.([]uint64)

	if ok {
		return res, true
	}

	r, ok := toSliceOf(kv.value, uintType, untypedUint)
	if !ok {
		return nil, false
	}
	res, ok = r.([]uint64)
	if ok {
		return res, true
	}
	return nil, false
}

// UintSliceOr returns value as a []uint64, otherwise the provided default
func (kv *KV) UintSliceOr(def []uint64) []uint64 {
	arr, ok := kv.UintSlice()
	if ok {
		return arr
	}
	return def
}

// Int64 returns value as a int64, if possible
func (kv *KV) Int64() (int64, bool) {
	return toInt(kv.value)
}

// Int64Or returns value as a int64, otherwise the provided default
func (kv *KV) Int64Or(def int64) int64 {
	i, ok := kv.Int64()
	if ok {
		return i
	}
	return def
}

// Int returns value as a int, if possible
func (kv *KV) Int() (int, bool) {
	i, ok := kv.Int64()
	return int(i), ok
}

// IntOr returns value as a int, otherwise the provided default
func (kv *KV) IntOr(def int) int {
	i := kv.Int64Or(int64(def))
	return int(i)
}

// Int32 returns value as a int, if possible
func (kv *KV) Int32() (int32, bool) {
	i, ok := kv.Int64()
	return int32(i), ok
}

// Int32Or returns value as a int, otherwise the provided default
func (kv *KV) Int32Or(def int32) int32 {
	i := kv.Int64Or(int64(def))
	return int32(i)
}

// Int16 returns value as a int16, if possible
func (kv *KV) Int16() (int16, bool) {
	i, ok := kv.Int64()
	return int16(i), ok
}

// Int32Or returns value as a int, otherwise the provided default
func (kv *KV) Int16Or(def int16) int16 {
	i := kv.Int64Or(int64(def))
	return int16(i)
}

// Int8 returns value as a int8, if possible
func (kv *KV) Int8() (int8, bool) {
	i, ok := kv.Int64()
	return int8(i), ok
}

// Int8Or returns value as a int, otherwise the provided default
func (kv *KV) Int8Or(def int8) int8 {
	i := kv.Int64Or(int64(def))
	return int8(i)
}

// IntSlice returns value as a []int64, if possible
func (kv *KV) IntSlice() ([]int64, bool) {
	if kv.value == nil {
		return nil, false
	}
	var res []int64
	res, ok := kv.value.([]int64)

	if ok {
		return res, ok
	}

	r, ok := toSliceOf(kv.value, intType, untypedInt)
	if !ok {
		return nil, false
	}
	res, ok = r.([]int64)
	if ok {
		return res, true
	}
	return nil, false
}

// IntSliceOr returns value as a []int64, otherwise the provided default
func (kv *KV) IntSliceOr(def []int64) []int64 {
	arr, ok := kv.IntSlice()

	if ok {
		return arr
	}
	return def
}

// Float returns value as a float64, if possible
func (kv *KV) Float() (float64, bool) {
	if kv.value == nil {
		return 0.0, false
	}

	return toFloat(kv.value)
}

// FloatOr returns value as a float64, otherwise the provided default
func (kv *KV) FloatOr(def float64) float64 {
	i, ok := kv.Float()
	if ok {
		return i
	}
	return def
}

// FloatSlice returns value as a []float64, if possible
func (kv *KV) FloatSlice() ([]float64, bool) {
	if kv.value == nil {
		return nil, false
	}

	var res []float64
	res, ok := kv.value.([]float64)

	if ok {
		return res, ok
	}

	r, ok := toSliceOf(kv.value, floatType, untypedFloat)
	if !ok {
		return nil, false
	}
	res, ok = r.([]float64)
	if ok {
		return res, true
	}
	return nil, false
}

// FloatSliceOr returns value as a []float64, otherwise the provided default
func (kv *KV) FloatSliceOr(def []float64) []float64 {
	arr, ok := kv.FloatSlice()

	if ok {
		return arr
	}
	return def
}

// Bool returns value as a bool, if possible
func (kv *KV) Bool() (bool, bool) {
	return toBool(kv.value)
}

// BoolOr returns value as a bool, otherwise the provided default
func (kv *KV) BoolOr(def bool) bool {
	i, ok := kv.Bool()
	if ok {
		return i
	}
	return def
}

// BoolSlice returns value as a []bool, if possible
func (kv *KV) BoolSlice() ([]bool, bool) {
	if kv.value == nil {
		return nil, false
	}

	var res []bool
	res, ok := kv.value.([]bool)

	if ok {
		return res, ok
	}

	r, ok := toSliceOf(kv.value, boolType, untypedBool)
	if !ok {
		return nil, false
	}
	res, ok = r.([]bool)
	if ok {
		return res, true
	}
	return nil, false
}

// BoolSliceOr returns value as a []bool, otherwise the provided default
func (kv *KV) BoolSliceOr(def []bool) []bool {
	arr, ok := kv.BoolSlice()

	if ok {
		return arr
	}
	return def
}

func toString(in interface{}) (res string, ok bool) {
	if in == nil {
		return "", false
	}

	switch in.(type) {
	case string:
		res, ok = in.(string)
	case []byte:
		var b []byte
		b, ok = in.([]byte)
		res = string(b)
	case []rune:
		var r []rune
		r, ok = in.([]rune)
		res = string(r)
	}
	return
}

func toBool(in interface{}) (res bool, ok bool) {
	if in == nil {
		return false, false
	}

	switch in.(type) {
	case bool:
		res, ok = in.(bool)
	}
	return
}

func toUint(num interface{}) (uint64, bool) {
	if num == nil {
		return 0, false
	}

	var i uint64
	var ok bool
	switch num.(type) {
	case int, int8, int16, int32, int64:
		a := reflect.ValueOf(num).Int() // a has type int64
		return uint64(a), true
	case uint, uint8, uint16, uint32, uint64:
		a := reflect.ValueOf(num).Uint() // a has type uint64
		return a, true
	case float64:
		f, ok := num.(float64)
		return uint64(f), ok
	case float32:
		f, ok := num.(float32)
		return uint64(f), ok
	}

	return i, ok

}

func toInt(num interface{}) (int64, bool) {
	if num == nil {
		return 0, false
	}

	var i int64
	var ok bool
	switch num.(type) {
	case int, int8, int16, int32, int64:
		a := reflect.ValueOf(num).Int() // a has type int64
		return a, true
	case uint, uint8, uint16, uint32, uint64:
		a := reflect.ValueOf(num).Uint() // a has type uint64
		return int64(a), true
	case float64, float32:
		a := reflect.ValueOf(num).Float()
		return int64(a), true
	}

	return i, ok
}

func toFloat(num interface{}) (float64, bool) {
	if num == nil {
		return 0, false
	}

	var i float64
	var ok bool

	switch num.(type) {
	case int, int8, int16, int32, int64:
		a := reflect.ValueOf(num).Int() // a has type int64
		return float64(a), true
	case uint, uint8, uint16, uint32, uint64:
		a := reflect.ValueOf(num).Uint() // a has type uint64
		return float64(a), true
	case float64:
		f, ok := num.(float64)
		return float64(f), ok
	case float32:
		f, ok := num.(float32)
		return float64(f), ok
	}

	return i, ok
}

func toSliceOf(value interface{}, typ reflect.Type, converter converter) (interface{}, bool) {
	if reflect.TypeOf(value).Kind() != reflect.Slice {
		return nil, false
	}

	slice := reflect.ValueOf(value)
	resSlice := reflect.MakeSlice(reflect.SliceOf(typ), slice.Len(), slice.Len())

	for i := 0; i < slice.Len(); i++ {
		val, ok := converter(slice.Index(i).Interface())
		if !ok {
			return nil, false
		}

		resSlice.Index(i).Set(reflect.ValueOf(val))
	}

	return resSlice.Interface(), true
}
