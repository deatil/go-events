package events

import (
    "fmt"
    "reflect"
)

/**
 * Pool
 *
 * @create 2024-7-26
 * @author deatil
 */
type Pool struct {}

func NewPool() *Pool {
    return &Pool{}
}

// Call Func
func (this *Pool) CallFunc(fn any, args []any) any {
    val := reflect.ValueOf(fn)
    if val.Kind() != reflect.Func {
        panic("go-events: func type error")
    }

    return this.Call(val, args)
}

// listen struct
func (this *Pool) CallStructMethod(in any, method string, args []any) any {
    val := reflect.ValueOf(in)
    if val.Kind() != reflect.Pointer && val.Kind() != reflect.Struct {
        panic("go-events: struct type error")
    }

    newMethod := val.MethodByName(method)
    return this.Call(newMethod, args)
}

// Call Func
func (this *Pool) Call(fn reflect.Value, args []any) any {
    if fn.Kind() != reflect.Func {
        panic("go-events: call func type error")
    }

    if !fn.IsValid() {
        panic("go-events: call func valid error")
    }

    fnType := fn.Type()

    // 参数
    params := this.bindParams(fnType, args)

    res := fn.Call(params)
    if len(res) == 0 {
        return nil
    }

    return res[0].Interface()
}

// is Struct
func (this *Pool) IsStruct(in any) bool {
    val := reflect.ValueOf(in)
    if val.Kind() == reflect.Pointer || val.Kind() == reflect.Struct {
        return true
    }

    return false
}

// is Func
func (this *Pool) IsFunc(in any) bool {
    val := reflect.ValueOf(in)
    if val.Kind() == reflect.Func {
        return true
    }

    return false
}

// bind params
func (this *Pool) bindParams(fnType reflect.Type, args []any) []reflect.Value {
    numIn := fnType.NumIn()
    if len(args) != numIn {
        err := fmt.Sprintf("go-events: func params error (args %d, func args %d)", len(args), numIn)
        panic(err)
    }

    // 参数
    params := make([]reflect.Value, 0)
    for i := 0; i < numIn; i++ {
        dataValue := this.convertTo(fnType.In(i), args[i])
        params = append(params, dataValue)
    }

    return params
}

// src convert type to new typ
func (this *Pool) convertTo(typ reflect.Type, src any) reflect.Value {
    dataKey := getTypeKey(typ)

    fieldType := reflect.TypeOf(src)
    if !fieldType.ConvertibleTo(typ) {
        return reflect.New(typ).Elem()
    }

    fieldValue := reflect.ValueOf(src)

    if dataKey != getTypeKey(fieldType) {
        fieldValue = fieldValue.Convert(typ)
    }

    return fieldValue
}
