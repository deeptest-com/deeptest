package com.ngtesting.platform.action.client;

import com.alibaba.fastjson.JSONObject;
import com.ngtesting.platform.action.BaseAction;
import com.ngtesting.platform.config.Constant;
import com.ngtesting.platform.model.TstUser;
import com.ngtesting.platform.service.intf.IssueSearchService;
import com.ngtesting.platform.servlet.PrivPrj;
import org.apache.shiro.SecurityUtils;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestMethod;
import org.springframework.web.bind.annotation.RestController;

import javax.servlet.http.HttpServletRequest;
import java.util.HashMap;
import java.util.List;
import java.util.Map;


@RestController
@RequestMapping(Constant.API_PATH_CLIENT + "issue_search/")
public class IssueSearchAction extends BaseAction {
	@Autowired
    IssueSearchService issueSearchService;

    @RequestMapping(value = "idAndTitleSearch", method = RequestMethod.POST)
    @PrivPrj(perms = {"issue-view"})
    public Map<String, Object> idAndTitleSearch(HttpServletRequest request, @RequestBody JSONObject json) {
        Map<String, Object> ret = new HashMap<String, Object>();

        TstUser user = (TstUser) SecurityUtils.getSubject().getPrincipal();
        Integer prjId = user.getDefaultPrjId();

        String text = json.getString("text");
        List<Integer> exceptIds = json.getObject("exceptIds", List.class);

        List<Map> ls = issueSearchService.idAndTitleSearch(text, exceptIds, prjId);

        ret.put("data", ls);
        ret.put("code", Constant.RespCode.SUCCESS.getCode());
        return ret;
    }

}
