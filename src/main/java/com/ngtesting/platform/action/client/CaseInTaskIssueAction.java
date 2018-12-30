package com.ngtesting.platform.action.client;

import com.alibaba.fastjson.JSONObject;
import com.ngtesting.platform.action.BaseAction;
import com.ngtesting.platform.config.Constant;
import com.ngtesting.platform.dao.CaseInTaskHistoryDao;
import com.ngtesting.platform.dao.CaseInTaskIssueDao;
import com.ngtesting.platform.model.*;
import com.ngtesting.platform.service.intf.CaseInTaskIssueService;
import com.ngtesting.platform.servlet.PrivPrj;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Controller;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestMethod;
import org.springframework.web.bind.annotation.ResponseBody;

import javax.servlet.http.HttpServletRequest;
import java.util.HashMap;
import java.util.List;
import java.util.Map;


@Controller
@RequestMapping(Constant.API_PATH_CLIENT + "case_in_task_issue/")
public class CaseInTaskIssueAction extends BaseAction {
	@Autowired
	CaseInTaskIssueService caseInTaskIssueService;
    @Autowired
	CaseInTaskHistoryDao caseInTaskHistoryDao;
	@Autowired
	CaseInTaskIssueDao caseInTaskIssueDao;

	@RequestMapping(value = "save", method = RequestMethod.POST)
	@ResponseBody
	@PrivPrj(perms = {"test_task-exe"})
	public Map<String, Object> save(HttpServletRequest request, @RequestBody JSONObject json) {
		Map<String, Object> ret = new HashMap<String, Object>();

        TstUser user = (TstUser) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_PROFILE);

        Integer caseInTaskId = json.getInteger("caseInTaskId");
		Integer issueId = json.getInteger("issueId");

		Boolean result = caseInTaskIssueService.save(caseInTaskId, issueId, user);

        List<TstCaseInTaskIssue> issues = caseInTaskIssueDao.query(caseInTaskId);

        ret.put("issues", issues);
		ret.put("code", Constant.RespCode.SUCCESS.getCode());
		return ret;
	}


	@RequestMapping(value = "remove", method = RequestMethod.POST)
	@ResponseBody
	@PrivPrj(perms = {"test_task-exe"})
	public Map<String, Object> remove(HttpServletRequest request, @RequestBody JSONObject json) {
		Map<String, Object> ret = new HashMap<String, Object>();

        TstUser user = (TstUser) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_PROFILE);

        Integer caseInTaskId = json.getInteger("caseInTaskId");
		Integer id = json.getInteger("id");

		Boolean result = caseInTaskIssueService.remove(caseInTaskId, id, user);

		List<TstCaseInTaskIssue> issues = caseInTaskIssueDao.query(caseInTaskId);

		ret.put("issues", issues);
		ret.put("code", Constant.RespCode.SUCCESS.getCode());
		return ret;
	}

}
