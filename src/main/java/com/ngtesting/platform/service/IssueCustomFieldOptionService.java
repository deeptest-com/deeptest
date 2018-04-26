package com.ngtesting.platform.service;

import com.ngtesting.platform.entity.TestCustomFieldOption;
import com.ngtesting.platform.vo.CustomFieldOptionVo;

import java.util.List;

public interface IssueCustomFieldOptionService extends BaseService {
    List<CustomFieldOptionVo> listVos(Long fieldId);
    TestCustomFieldOption save(CustomFieldOptionVo option);
    boolean delete(Long id);
    boolean changeOrderPers(Long id, String act, Long fieldId);

	List<CustomFieldOptionVo> genVos(List<TestCustomFieldOption> pos);
	CustomFieldOptionVo genVo(TestCustomFieldOption po);
}
