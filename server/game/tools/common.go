package tools

import (
    "errors"
    "fmt"
    "reflect"
    "github.com/xiaonanln/goworld"
    "github.com/xiaonanln/goworld/engine/entity"
    "github.com/xiaonanln/goworld/engine/gwlog"
    "github.com/xiaonanln/goworld/newGame2/model"
    "github.com/json-iterator/go"
)

func GetStructStringField(input interface{}, key string) (value string, err error) {
    v, err := getStructField(input, key)
    if err != nil {
        return
    }
    value, ok := v.(string)
    if !ok {
        return value, errors.New("can't convert key'v to string")
    }
    return
}

func GetStructFloat64Field(input interface{}, key string) (value float64, err error) {
    v, err := getStructField(input, key)
    if err != nil {
        return
    }
    value, ok := v.(float64)
    if !ok {
        return value, errors.New("can't convert key'v to float64")
    }
    return
}

func GetStructInt64Field(input interface{}, key string) (value int64, err error) {
    v, err := getStructField(input, key)
    if err != nil {
        return
    }
    value, ok := v.(int64)
    if !ok {
        return value, errors.New("can't convert key'v to int64")
    }
    return
}

func GetStructArrayInt64Field(input interface{}, key string) (value []int64, err error) {
    v, err := getStructField(input, key)
    if err != nil {
        return
    }
    value, ok := v.([]int64)
    if !ok {
        return value, errors.New("can't convert key'v to []int64")
    }
    return
}

func getStructField(input interface{}, key string) (value interface{}, err error) {
    rv := reflect.ValueOf(input)
    rt := reflect.TypeOf(input)
    if rt.Kind() != reflect.Struct {
        return value, errors.New("input must be struct")
    }

    keyExist := false
    for i := 0; i < rt.NumField(); i++ {
        curField := rv.Field(i)
        if rt.Field(i).Name == key {
            switch curField.Kind() {
            case reflect.String, reflect.Int64, reflect.Int32, reflect.Int16, reflect.Int8, reflect.Int, reflect.Float64, reflect.Float32:
                keyExist = true
                value = curField.Interface()
            default:
                return value, errors.New("key must be int float or string")
            }
        }
    }
    if !keyExist {
        return value, errors.New(fmt.Sprintf("key %s not found in %s's field", key, rt))
    }
    return
}

func Contains(slice []string, s string) int {
	for index, value := range slice {
		if value == s {
			return index
		}
	}
	return -1
}

func ObjectJoinObjectByKey(obj *entity.MapAttr, obj2 interface{}, key string){
    switch v := obj2.(type) {
        case int64:
            obj.SetInt(key, obj.GetInt(key)+v)
        case float64:
            obj.SetFloat(key, obj.GetFloat(key)+v)
        case *entity.ListAttr:
            for _, v := range v.ToList(){
                switch vv := v.(type) {
                    case int64:
                        obj.GetListAttr(key).AppendInt(vv)
                    case float64:
                        obj.GetListAttr(key).AppendFloat(vv)
                    case string:
                        obj.GetListAttr(key).AppendStr(vv)
                    case *entity.ListAttr:
                        obj.GetListAttr(key).AppendListAttr(vv)
                    case *entity.MapAttr:
                        obj.GetListAttr(key).AppendMapAttr(vv)
                }
            }
        case *entity.MapAttr:
            obj.SetMapAttr(key, v)

    }
}

func Array(items... interface{}) *entity.ListAttr{
    a := goworld.ListAttr()
    for _,vv := range items{
        switch v := vv.(type) {
            case int:
                    a.AppendInt(int64(v))
     
            case string:
                    a.AppendStr(v)

            case int64:
                    a.AppendInt(v)

            case float64:
                    a.AppendFloat(v)

            case *entity.ListAttr:
                    a.AppendListAttr(v)

            case *entity.MapAttr:
                    a.AppendMapAttr(v)  
                    
            default:
                    gwlog.Debugf("Array %s create failed for tools/common/Array type:%s", items, reflect.TypeOf(vv))
        }
        
    }
    return a
}

func LaodConfigToRoleAttrModel(data string) []model.RoleAttrModel {
    var obj = []model.RoleAttrModel{}
    var jsonIterator = jsoniter.ConfigCompatibleWithStandardLibrary
    err := jsonIterator.Unmarshal( []byte(data), &obj)
    if err != nil {
        gwlog.Debugf("###############LaoConfigToRoleAttrModel error %s \n %s", err, data)
    }
    gwlog.Debugf("<<<<<<<<<<<<< LaoConfigToRoleAttrModel laod data done: %+v", obj)
    return obj
}

func LaodConfigToGoodsAttrModel(data string) []model.GoodsAttrModel {
    var obj = []model.GoodsAttrModel{}
    var jsonIterator = jsoniter.ConfigCompatibleWithStandardLibrary
    err := jsonIterator.Unmarshal( []byte(data), &obj)
    if err != nil {
        gwlog.Debugf("###############LaoConfigToGoodsAttrModel error %s \n %s", err, data)
    }
    gwlog.Debugf("<<<<<<<<<<<<< LaoConfigToGoodsAttrModel laod data done: %+v", obj)
    return obj
}

func LaodConfigToMissionModel(data string) []model.MissionModel {
    var obj = []model.MissionModel{}
    var jsonIterator = jsoniter.ConfigCompatibleWithStandardLibrary
    err := jsonIterator.Unmarshal( []byte(data), &obj)
    if err != nil {
        gwlog.Debugf("###############LaodConfigToMissionModel error %s \n %s", err, data)
    }
    gwlog.Debugf("<<<<<<<<<<<<< LaodConfigToMissionModel laod data done: %+v", obj)
    return obj
}

func LaodConfigToMonsterModel(data string) []model.MonsterModel {
    var obj = []model.MonsterModel{}
    var jsonIterator = jsoniter.ConfigCompatibleWithStandardLibrary
    err := jsonIterator.Unmarshal( []byte(data), &obj)
    if err != nil {
        gwlog.Debugf("###############LaodConfigToMonsterModel error %s \n %s", err, data)
    }
    gwlog.Debugf("<<<<<<<<<<<<< LaodConfigToMonsterModel laod data done: %+v", obj)
    return obj
}