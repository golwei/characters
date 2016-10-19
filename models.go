package models


type 历史人物 struct{
name string
关系人物 []name
参与事件 []历史事件
}



type 历史事件 struct{
起始时间 time
起因 string
经过 string
结果 string
}
