package com.ngtesting.platform.service;

import com.ngtesting.platform.entity.TestCustomFieldOption;
import com.ngtesting.platform.vo.CustomFieldOptionVo;

import java.util.List;

public interface CustomFieldOptionService extends BaseService {

	List<CustomFieldOptionVo> genVos(List<TestCustomFieldOption> pos);
	CustomFieldOptionVo genVo(TestCustomFieldOption po);

}
