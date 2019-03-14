package com.ngtesting.platform.action.client;

import com.alibaba.fastjson.JSONObject;
import com.ngtesting.platform.action.BaseAction;
import com.ngtesting.platform.config.Constant;
import com.ngtesting.platform.dao.CaseInTaskHistoryDao;
import com.ngtesting.platform.dao.CaseInTaskIssueDao;
import com.ngtesting.platform.model.TstCaseInTaskIssue;
import com.ngtesting.platform.model.TstUser;
import com.ngtesting.platform.service.intf.CaseInTaskIssueService;
import com.ngtesting.platform.servlet.PrivPrj;
import org.apache.shiro.SecurityUtils;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.*;

import javax.servlet.http.HttpServletRequest;
import java.util.HashMap;
import java.util.List;
import java.util.Map;


@RestController
@RequestMapping(Constant.API_PATH_CLIENT + "case_in_task_issue/")
public class CaseInTaskIssueAction extends BaseAction {
	@Autowired
	CaseInTaskIssueService caseInTaskIssueService;
    @Autowired
	CaseInTaskHistoryDao caseInTaskHistoryDao;
	@Autowired
	CaseInTaskIssueDao caseInTaskIssueDao;

	@RequestMapping(value = "save", method = RequestMethod.POST)
	@PrivPrj(perms = {"test_task-exe"})
	public Map<String, Object> save(HttpServletRequest request, @RequestBody JSONObject json) {
		Map<String, Object> ret = new HashMap<String, Object>();

        TstUser user = (TstUser) SecurityUtils.getSubject().getPrincipal();

        Integer caseInTaskId = json.getInteger("caseInTaskId");
		Integer issueId = json.getInteger("issueId");

		Boolean result = caseInTaskIssueService.save(caseInTaskId, issueId, user);

        List<TstCaseInTaskIssue> issues = caseInTaskIssueDao.query(caseInTaskId);

        ret.put("issues", issues);
		ret.put("code", Constant.RespCode.SUCCESS.getCode());
		return ret;
	}


	@RequestMapping(value = "remove", method = RequestMethod.POST)
	@PrivPrj(perms = {"test_task-exe"})
	public Map<String, Object> remove(HttpServletRequest request, @RequestBody JSONObject json) {
		Map<String, Object> ret = new HashMap<String, Object>();

        TstUser user = (TstUser) SecurityUtils.getSubject().getPrincipal();

        Integer caseInTaskId = json.getInteger("caseInTaskId");
		Integer id = json.getInteger("id");

		Boolean result = caseInTaskIssueService.remove(caseInTaskId, id, user);

		List<TstCaseInTaskIssue> issues = caseInTaskIssueDao.query(caseInTaskId);

		ret.put("issues", issues);
		ret.put("code", Constant.RespCode.SUCCESS.getCode());
		return ret;
	}

}
