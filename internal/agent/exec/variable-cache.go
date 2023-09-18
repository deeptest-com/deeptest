package agentExec

//func GetCachedVariableMapInContext(processorId uint) (ret domain.VarKeyValuePair) {
//	ret = domain.VarKeyValuePair{}
//
//	variables := listCachedVariable(processorId)
//
//	for _, item := range variables {
//		valMap, isMap := item.Sample.(domain.VarKeyValuePair)
//
//		if isMap {
//			for propKey, v := range valMap {
//				ret[fmt.Sprintf("%s.%s", item.Name, propKey)] = v
//			}
//		} else {
//			ret[item.Name] = item.Sample
//		}
//	}
//
//	return
//}
//
//func listCachedVariable(processorId uint) (variables []domain.ExecVariable) {
//	effectiveScopeIds := ScopeHierarchy[processorId]
//
//	if effectiveScopeIds == nil {
//		effectiveScopeIds = &[]uint{uint(0)}
//	}
//
//	if effectiveScopeIds == nil {
//		return
//	}
//
//	for _, id := range *effectiveScopeIds {
//		for _, vari := range ScopedVariables[id] {
//			if vari.Scope == consts.Public || (vari.Scope == consts.Private && id == processorId) {
//				variables = append(variables, vari)
//			}
//		}
//	}
//
//	return
//}
