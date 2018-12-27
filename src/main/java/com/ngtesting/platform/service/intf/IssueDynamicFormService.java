package com.ngtesting.platform.service.intf;

import com.ngtesting.platform.model.IsuField;

import java.util.List;
import java.util.Map;

public interface IssueDynamicFormService extends BaseService {
  List<IsuField> listNotUsedField(Integer orgId, Integer projectId, Integer pageId);
  Map<String, List<Map>> genIssuePropMap(Integer orgId, Integer projectId);

  Map<String, Object> genIssueBuldInPropValMap(Integer orgId, Integer projectId);

  // TODO: cached
  List<Map> fetchOrgField(Integer orgId, Integer projectId, String sort);
}
