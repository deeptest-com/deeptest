package com.ngtesting.platform.action.client;

import com.ngtesting.platform.action.BaseAction;
import com.ngtesting.platform.action.admin.CasePriorityAdmin;
import com.ngtesting.platform.config.Constant;
import org.apache.commons.logging.Log;
import org.apache.commons.logging.LogFactory;
import org.springframework.stereotype.Controller;
import org.springframework.web.bind.annotation.RequestMapping;


@Controller
@RequestMapping(Constant.API_PATH_CLIENT + "issue_query/")
public class IssueQueryAction extends BaseAction {
	private static final Log log = LogFactory.getLog(CasePriorityAdmin.class);


}
