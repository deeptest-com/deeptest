package com.ngtesting.platform.action;

import com.alibaba.fastjson.JSONObject;
import com.ngtesting.platform.config.Constant;
import com.ngtesting.platform.model.TstProject;
import com.ngtesting.platform.model.TstUser;
import com.ngtesting.platform.service.ProjectService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Controller;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.ResponseBody;

import javax.servlet.http.HttpServletRequest;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

@Controller
@RequestMapping(value = Constant.API_PATH_CLIENT + "/project")
public class ProjectAction {
    @Autowired
    private ProjectService projectService;

    @ResponseBody
    @PostMapping("/list")
    public Object list(HttpServletRequest request, @RequestBody JSONObject json) {
        Map<String, Object> ret = new HashMap<String, Object>();

        TstUser user = (TstUser) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_KEY);
        Integer orgId = user.getDefaultOrgId();

        String keywords = json.getString("keywords");
        Boolean disabled = json.getBoolean("disabled");

        List<TstProject> vos = projectService.list(orgId, user.getId(), keywords, disabled);

        ret.put("data", vos);
        ret.put("code", Constant.RespCode.SUCCESS.getCode());

        return ret;
    }

}
