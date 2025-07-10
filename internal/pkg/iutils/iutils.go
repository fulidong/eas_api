package iutils

import "reflect"

func GetDistinctFields[T any, Y int | int64 | int32 | string](entities []T, getField func(T) Y) []Y {
	m := make(map[Y]struct{}, len(entities))
	for _, re := range entities {
		m[getField(re)] = struct{}{}
	}
	if len(m) == 0 {
		// 没有有效的 ID，直接返回空 map 即可
		return nil
	}
	// 转换为切片用于查询
	res := make([]Y, 0, len(m))
	for id := range m {
		res = append(res, id)
	}
	return res
}

// 计算两个集合中哪些是新增，修改和删除
func DiffEntities[T1 any, T2 any, Y int | int32 | int64 | string](
	dbList []T1,
	inputList []T2,
	getId1 func(T1) Y,
	getId2 func(T2) Y,
	instanceFunc func() T1,
) (toCreate, toUpdate []T1, toDelete []Y) {
	dbMap := make(map[Y]T1, len(dbList))
	inputMap := make(map[Y]T2, len(inputList))

	// 构建 map
	for _, dbEnt := range dbList {
		dbMap[getId1(dbEnt)] = dbEnt
	}
	for _, inputEnt := range inputList {
		id := getId2(inputEnt)
		var zero Y
		if id == zero {
			dbEnt := instanceFunc()
			toCreate = append(toCreate, MapTo[T2, T1](dbEnt, inputEnt))
		} else {
			inputMap[getId2(inputEnt)] = inputEnt
		}
	}
	if len(dbMap) > 0 {
		// 查找需要更新和删除的
		for id, dbEnt := range dbMap {
			if inputEnt, exists := inputMap[getId1(dbEnt)]; exists {
				toUpdate = append(toUpdate, MapTo[T2, T1](dbEnt, inputEnt))
				delete(inputMap, id) // 已处理，避免重复判断
			} else {
				// 数据库中有，输入中无 → 删除
				toDelete = append(toDelete, getId1(dbEnt))
			}
		}
	}
	return toCreate, toUpdate, toDelete
}

// MapTo 将 src 结构体字段映射到 dst 结构体字段中，仅映射同名同类型的字段
func MapTo[T1, T2 any](dst T2, src T1) T2 {

	srcVal := reflect.ValueOf(src).Elem()
	dstVal := reflect.ValueOf(dst).Elem()
	srcType := srcVal.Type()
	dstType := dstVal.Type()

	for i := 0; i < srcType.NumField(); i++ {
		srcField := srcType.Field(i)
		srcName := srcField.Name

		// 查找目标结构体是否有同名字段
		dstField, ok := dstType.FieldByName(srcName)
		if !ok || dstField.Type != srcField.Type {
			continue // 字段不存在或类型不一致，跳过
		}

		// 赋值
		dstVal.FieldByName(srcName).Set(srcVal.Field(i))
	}

	return dst
}
