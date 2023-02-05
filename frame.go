package pandas

import (
	"reflect"
	"sync"
)

type Frame[T GenericType] interface {
	// Name 取得series名称
	Name() string
	// ReName renames the series.
	ReName(name string)
	// Type returns the type of data the series holds.
	// 返回series的数据类型
	Type() Type
	// Len 获得行数
	Len() int
	// Values 获得全部数据集
	Values() []T // 如果确定类型, 后面可能无法自动调整
}

type GenericFrame[T GenericType] struct {
	lock      sync.RWMutex    // 读写锁
	formatter StringFormatter // 字符串格式化工具
	name      string          // 帧名称
	type_     Type            // values元素类型
	nilCount  int             // nil和nan的元素有多少, 这种统计在bool和int64类型中不会大于0, 只对float64及string有效
	rows      int             // 行数
	values    []T             // 只能是一个一维slice, 在所有的运算中, values强制转换成float64切片
}

func NewFrame[T GenericType](name string, values ...any) Frame[T] {
	frame := GenericFrame[T]{
		formatter: DefaultFormatter,
		name:      name,
		type_:     SERIES_TYPE_INVAILD,
		nilCount:  0,
		rows:      0,
		values:    nil,
	}
	// 确定泛型的具体类型, 以便后面创建slice
	kind := checkoutRawType(&frame)
	if kind == SERIES_TYPE_INVAILD {
		return &frame
	}
	frame.type_ = kind
	if frame.type_ == SERIES_TYPE_BOOL {
		// bool
		frame.values = reflect.MakeSlice(typeBool, 0, 0).Interface().([]T)
	} else if frame.type_ == SERIES_TYPE_INT64 {
		// int64
		frame.values = reflect.MakeSlice(typeInt64, 0, 0).Interface().([]T)
	} else if frame.type_ == SERIES_TYPE_FLOAT32 {
		// float32
		frame.values = reflect.MakeSlice(typeFloat32, 0, 0).Interface().([]T)
	} else if frame.type_ == SERIES_TYPE_FLOAT64 {
		// float64
		frame.values = reflect.MakeSlice(typeFloat64, 0, 0).Interface().([]T)
	} else {
		// string, 字符串最后容错使用
		frame.values = reflect.MakeSlice(typeString, 0, 0).Interface().([]T)
	}
	size := 0
	for idx, v := range values {
		vv := reflect.ValueOf(v)
		vk := vv.Kind()
		switch vk {
		case reflect.Invalid: // {interface} nil
			frame.assign(idx, size, nil)
		case reflect.Slice, reflect.Array: // 切片或者数组
			for i := 0; i < vv.Len(); i++ {
				tv := vv.Index(i).Interface()
				frame.assign(idx, size, tv)
			}
		default:
			// 默认为基础数据类型
			tv := vv.Interface()
			frame.assign(idx, size, tv)
		}
	}
	return &frame
}

func (self *GenericFrame[T]) Name() string {
	//TODO implement me
	panic("implement me")
}

func (self *GenericFrame[T]) ReName(name string) {
	//TODO implement me
	panic("implement me")
}

func (self *GenericFrame[T]) Type() Type {
	//TODO implement me
	panic("implement me")
}

func (self *GenericFrame[T]) Len() int {
	//TODO implement me
	panic("implement me")
}

func (self *GenericFrame[T]) Values() []T {
	//TODO implement me
	panic("implement me")
}
