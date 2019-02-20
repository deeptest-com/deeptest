package com.ngtesting.platform.service.impl;

import com.ngtesting.platform.dao.CaseHistoryDao;
import com.ngtesting.platform.dao.UserDao;
import com.ngtesting.platform.model.TstCaseHistory;
import com.ngtesting.platform.model.TstUser;
import com.ngtesting.platform.service.intf.CaseHistoryService;
import com.ngtesting.platform.utils.MsgUtil;
import org.apache.commons.lang3.StringUtils;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.text.MessageFormat;

@Service
public class CaseHistoryServiceImpl extends BaseServiceImpl implements CaseHistoryService {
    @Autowired
    UserDao userDao;

    @Autowired
    CaseHistoryDao caseHistoryDao;

    @Override
    public void saveHistory(TstUser user, MsgUtil.MsgAction act, Integer caseId, String field) {
	    String action = act.msg;

        String fieldMsg = StringUtils.isNotEmpty(field)? "字段 "  + field: "";
        String msg = MessageFormat.format(MsgUtil.HistoryMsgTemplate.opt_entity.msg, user.getNickname(), action, fieldMsg);

        TstCaseHistory his = new TstCaseHistory();
        his.setTitle(msg);
        his.setCaseId(caseId);
        caseHistoryDao.save(his);
    }

}
