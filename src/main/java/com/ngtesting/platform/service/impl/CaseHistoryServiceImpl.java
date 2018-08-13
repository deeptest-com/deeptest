package com.ngtesting.platform.service.impl;

import com.ngtesting.platform.config.Constant;
import com.ngtesting.platform.dao.CaseHistoryDao;
import com.ngtesting.platform.dao.UserDao;
import com.ngtesting.platform.model.TstCase;
import com.ngtesting.platform.model.TstCaseHistory;
import com.ngtesting.platform.model.TstUser;
import com.ngtesting.platform.service.CaseHistoryService;
import com.ngtesting.platform.utils.StringUtil;
import org.apache.commons.lang3.StringUtils;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.util.List;

@Service
public class CaseHistoryServiceImpl extends BaseServiceImpl implements CaseHistoryService {
    @Autowired
    UserDao userDao;

    @Autowired
    CaseHistoryDao caseHistoryDao;

    @Override
    public List<TstCaseHistory> findHistories(Integer testCaseId) {
//        DetachedCriteria dc = DetachedCriteria.forClass(TstCaseHistory.class);
//        dc.add(Restrictions.eq("testCaseId", testCaseId));
//
//        dc.add(Restrictions.eq("deleted", Boolean.FALSE));
//        dc.add(Restrictions.eq("disabled", Boolean.FALSE));
//
//        dc.addOrder(Order.desc("createTime"));
//
//        List<TstCaseHistory> ls = findAllByCriteria(dc);
//        return ls;

        return null;
    }

    @Override
    public void saveHistory(TstUser user, Constant.CaseAct act, TstCase testCase, String field) {
	    String action = act.msg;

        String msg = "用户" + StringUtil.highlightDict(user.getNickname()) + action;
        if (StringUtils.isNotEmpty(field)) {
            msg += " " + field;
        } else {
//            msg += "信息";
        }
        TstCaseHistory his = new TstCaseHistory();
        his.setTitle(msg);
        his.setCaseId(testCase.getId());
        caseHistoryDao.save(his);
    }

}
