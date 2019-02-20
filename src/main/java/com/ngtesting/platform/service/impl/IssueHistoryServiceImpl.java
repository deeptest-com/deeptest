package com.ngtesting.platform.service.impl;

import com.ngtesting.platform.dao.IssueHistoryDao;
import com.ngtesting.platform.dao.UserDao;
import com.ngtesting.platform.model.IsuHistory;
import com.ngtesting.platform.model.TstUser;
import com.ngtesting.platform.service.intf.IssueHistoryService;
import com.ngtesting.platform.utils.MsgUtil;
import org.apache.commons.lang3.StringUtils;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.text.MessageFormat;
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
	public void saveHistory(TstUser user, MsgUtil.MsgAction act, Integer issueId, String field) {
		String action = act.msg;

		String fieldMsg = StringUtils.isNotEmpty(field)? "字段 "  + field: "";
		String msg = MessageFormat.format(MsgUtil.HistoryMsgTemplate.opt_entity.msg, user.getNickname(), action, fieldMsg);

		IsuHistory his = new IsuHistory();
		his.setTitle(msg);
		his.setIssueId(issueId);
		issueHistoryDao.save(his);
	}

}

