package com.ngtesting.platform.service.impl;

import com.ngtesting.platform.config.Constant;
import com.ngtesting.platform.dao.CaseInTaskHistoryDao;
import com.ngtesting.platform.dao.UserDao;
import com.ngtesting.platform.model.TstCaseInTaskHistory;
import com.ngtesting.platform.model.TstUser;
import com.ngtesting.platform.service.intf.CaseInTaskHistoryService;
import com.ngtesting.platform.utils.MsgUtil;
import org.apache.commons.lang3.StringUtils;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.text.MessageFormat;

@Service
public class CaseInTaskHistoryServiceImpl extends BaseServiceImpl implements CaseInTaskHistoryService {
    @Autowired
    UserDao userDao;

    @Autowired
    CaseInTaskHistoryDao caseInTaskHistoryDao;

    @Override
    public void saveHistory(TstUser user, MsgUtil.MsgAction act, Integer caseInTaskId, String field) {
	    String action = act.msg;

        String fieldMsg = StringUtils.isNotEmpty(field)? "字段 "  + field: "";
        String msg = MessageFormat.format(MsgUtil.HistoryMsgTemplate.opt_entity.msg, user.getNickname(), action, fieldMsg);

        TstCaseInTaskHistory his = new TstCaseInTaskHistory();
        his.setTitle(msg);
        his.setCaseInTaskId(caseInTaskId);
        caseInTaskHistoryDao.save(his);
    }

    @Override
    public void saveHistory(Integer caseId, Integer caseInTaskId, MsgUtil.MsgAction act, TstUser user,
                            String status, String result) {
        String action = act.msg;

        String resultMsg = StringUtils.isNotEmpty(result)? Constant.ExeStatus.get(status): "";
        String msg = MessageFormat.format(MsgUtil.HistoryMsgTemplate.exe_case.msg, user.getNickname(), action, resultMsg);

        TstCaseInTaskHistory his = new TstCaseInTaskHistory();
        his.setTitle(msg);
        his.setCaseId(caseId);
        his.setCaseInTaskId(caseInTaskId);

        caseInTaskHistoryDao.save(his);
    }

}
