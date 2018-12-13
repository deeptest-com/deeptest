package com.ngtesting.platform.service.intf;

import com.alibaba.fastjson.JSONObject;
import com.ngtesting.platform.model.TstCaseComments;
import com.ngtesting.platform.model.TstUser;
import org.springframework.transaction.annotation.Transactional;

public interface CaseCommentsService extends BaseService {

	TstCaseComments save(JSONObject vo, TstUser userVo);

    @Transactional
    TstCaseComments save(TstCaseComments vo, TstUser user);

    Boolean delete(Integer id, TstUser user);

    TstCaseComments genVo(TstCaseComments po);
}
