package com.ngtesting.platform.service.impl;

import com.ngtesting.platform.config.Constant;
import com.ngtesting.platform.dao.CaseHistoryDao;
import com.ngtesting.platform.dao.UserDao;
import com.ngtesting.platform.model.TstCase;
import com.ngtesting.platform.model.TstCaseHistory;
import com.ngtesting.platform.model.TstUser;
import com.ngtesting.platform.service.intf.CaseHistoryService;
import com.ngtesting.platform.utils.StringUtil;
import org.apache.commons.lang3.StringUtils;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

@Service
public class CaseHistoryServiceImpl extends BaseServiceImpl implements CaseHistoryService {
    @Autowired
    UserDao userDao;

    @Autowired
    CaseHistoryDao caseHistoryDao;

    @Override
    public void saveHistory(TstUser user, Constant.EntityAct act, Integer caseId, String field) {
	    String action = act.msg;

        String msg = "用户" + StringUtil.highlightDict(user.getNickname()) + action;
        if (StringUtils.isNotEmpty(field)) {
            msg += " " + field;
        } else {
//            msg += "信息";
        }
        TstCaseHistory his = new TstCaseHistory();
        his.setTitle(msg);
        his.setCaseId(caseId);
        caseHistoryDao.save(his);
    }

}
