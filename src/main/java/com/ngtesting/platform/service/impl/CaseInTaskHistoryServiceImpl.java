package com.ngtesting.platform.service.impl;

import com.ngtesting.platform.config.Constant;
import com.ngtesting.platform.dao.CaseInTaskHistoryDao;
import com.ngtesting.platform.dao.UserDao;
import com.ngtesting.platform.model.TstCaseInTask;
import com.ngtesting.platform.model.TstCaseInTaskHistory;
import com.ngtesting.platform.model.TstUser;
import com.ngtesting.platform.service.intf.CaseInTaskHistoryService;
import com.ngtesting.platform.utils.StringUtil;
import org.apache.commons.lang3.StringUtils;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

@Service
public class CaseInTaskHistoryServiceImpl extends BaseServiceImpl implements CaseInTaskHistoryService {
    @Autowired
    UserDao userDao;

    @Autowired
    CaseInTaskHistoryDao caseInTaskHistoryDao;

    @Override
    public void saveHistory(TstUser user, Constant.EntityAct act, Integer caseInTaskId, String field) {
	    String action = act.msg;

        String msg = "用户" + StringUtil.highlightDict(user.getNickname()) + action;
        if (StringUtils.isNotEmpty(field)) {
            msg += " " + field;
        } else {
//            msg += "信息";
        }
        TstCaseInTaskHistory his = new TstCaseInTaskHistory();
        his.setTitle(msg);
        his.setCaseInTaskId(caseInTaskId);
        caseInTaskHistoryDao.save(his);
    }

    @Override
    public void saveHistory(Integer caseId, Integer caseInTaskId, Constant.EntityAct act, TstUser user,
                            String status, String result) {
        String action = act.msg;

        String msg = "用户" + StringUtil.highlightDict(user.getNickname()) + action
                + "为\"" + Constant.ExeStatus.get(status) + "\"";
        if (!StringUtils.isEmpty(result)) {
            msg += ", 结果内容：" + result;
        }

        TstCaseInTaskHistory his = new TstCaseInTaskHistory();
        his.setTitle(msg);
        his.setCaseId(caseId);
        his.setCaseInTaskId(caseInTaskId);

        caseInTaskHistoryDao.save(his);
    }

}
