package mq

import (
	"encoding/json"
	"errors"
	"fmt"
	"gin_gomicro/app/user/service"
	"gin_gomicro/idl/pb"
	"github.com/streadway/amqp"
	"reflect"
)

type MessageConsumer struct {
	Method string
	Params interface{}
}

func ConsumeUser() error {
	mq := NewUserMq()
	if mq.conn == nil {
		fmt.Println("conn is nil")
		return errors.New("conn is nil")
	}
	if mq.ch == nil {
		fmt.Println("ch is nil")
		return errors.New("ch is nil")
	}
	msgs, err := mq.ch.Consume(
		"userChannel",
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	go func() {
		for {
			select {
			case msg := <-msgs:
				handleMsg(msg)
			}
		}
	}()
	fmt.Println("consume user success")
	return nil
}

func handleMsg(msg amqp.Delivery) {
	message, err := parseMessage(msg.Body)
	if err != nil {
		fmt.Println(err)
	}
	result, err := callMethod(message.Method, message.Params)
	if err != nil {
		fmt.Println(err)
	}
	if result.Code != 200 {
	}
}

// 解析方法名
func parseMessage(body []byte) (*MessageConsumer, error) {
	msg := &MessageConsumer{}
	err := json.Unmarshal(body, msg)
	if err != nil {
		return msg, err
	}
	return msg, nil
}

// 根据方法名调用方法
func callMethod(methodName string, args interface{}) (*pb.UserRes, error) {
	res := &pb.UserRes{}
	err := errors.New("")
	method := reflect.ValueOf(service.GetUserSrv()).MethodByName(methodName)
	if method.IsValid() {
		return res, errors.New("找不到方法" + methodName)
	}
	methodType := method.Type()
	numArgs := methodType.NumIn()
	if numArgs != 1 {
		return res, errors.New("参数不对")
	}
	arg := reflect.ValueOf(args)
	if arg.Type() != methodType.In(0) {
		return res, errors.New("参数不对")
	}

	//for i := 0; i < numArgs; i++ {
	//	paramName := methodType.In(i).Name()
	//
	//	//根据参数转换为反射值
	//	paramValue, ok := args[paramName]
	//	if !ok {
	//		return res, errors.New("缺少请求参数"+paramName)
	//	}
	//
	//	//如果参数是结构体
	//	if methodType.In(i).Kind() == reflect.Struct {
	//		structValue := reflect.New(methodType.In(i).Elem())
	//		for j := 0; j < methodType.NumIn(); j++ {
	//			field := structValue.Field(j)
	//			fieldName := methodType.In(j).Name()
	//
	//			mapValue, ok := paramValue.(map[string]interface{})
	//			if !ok {
	//				return res, errors.New("结构体参数不对"+fieldName)
	//			}
	//			field.Set(reflect.ValueOf(mapValue[fieldName]))
	//		}
	//		params = append(params, structValue)
	//	} else {
	//		param := reflect.ValueOf(paramValue)
	//		//检查参数类型是否匹配方法的参数类型
	//		if param.Type() != reflect.TypeOf(i) {
	//			return res, errors.New("参数类型有误"+paramName)
	//		}
	//		params = append(params, param)
	//	}
	//}

	result := method.Call([]reflect.Value{arg})
	if len(result) > 0 {
		res = result[0].Interface().(*pb.UserRes)
		err = result[1].Interface().(error)
	}
	return res, err
}
