package com.ngtesting.platform.service.impl;

import com.ngtesting.platform.config.Constant;
import com.ngtesting.platform.dao.IssueHistoryDao;
import com.ngtesting.platform.dao.UserDao;
import com.ngtesting.platform.model.IsuAttachment;
import com.ngtesting.platform.model.IsuHistory;
import com.ngtesting.platform.model.IsuIssue;
import com.ngtesting.platform.model.TstUser;
import com.ngtesting.platform.service.intf.IssueHistoryService;
import com.ngtesting.platform.utils.StringUtil;
import org.apache.commons.lang3.StringUtils;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.util.List;

@Service
public class IssueHistoryServiceImpl extends BaseServiceImpl implements IssueHistoryService {

	@Autowired
	UserDao userDao;

	@Autowired
	IssueHistoryDao issueHistoryDao;

	@Override
	public List<IsuHistory> query(Integer issueId) {
		return issueHistoryDao.query(issueId);
	}

	@Override
	public void saveHistory(TstUser user, Constant.EntityAct act, Integer issueId, String field) {
		String action = act.msg;

		String msg = "用户" + StringUtil.highlightDict(user.getNickname()) + action;
		if (StringUtils.isNotEmpty(field)) {
			msg += " " + field;
		} else {
//            msg += "信息";
		}
		IsuHistory his = new IsuHistory();
		his.setTitle(msg);
		his.setIssueId(issueId);
		issueHistoryDao.save(his);
	}

}

