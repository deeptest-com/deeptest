package com.ngtesting.platform.action.client;

import com.alibaba.fastjson.JSONObject;
import com.ngtesting.platform.action.BaseAction;
import com.ngtesting.platform.config.Constant;
import com.ngtesting.platform.service.*;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Controller;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.ResponseBody;

import javax.servlet.http.HttpServletRequest;
import java.util.HashMap;
import java.util.Map;

@Controller
@RequestMapping(value = Constant.API_PATH_CLIENT + "/project_config")
public class ProjectConfigAction extends BaseAction {
    @Autowired
    private ProjectService projectService;
    @Autowired
    private ProjectConfigService projectConfigService;

    @Autowired
    private ProjectRoleService projectRoleService;
    @Autowired
    private ProjectRoleEntityRelationService projectRoleEntityRelationService;
    @Autowired
    private HistoryService historyService;

    @Autowired
    private PushSettingsService pushSettingsService;

    @Autowired
    AuthService authService;

    @ResponseBody
    @PostMapping("/get")
    public Map<String, Object> get(HttpServletRequest request, @RequestBody JSONObject json) {
        Map<String, Object> ret = new HashMap<String, Object>();
//        TstUser user = (TstUser) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_PROFILE);
//        Integer orgId = user.getDefaultOrgId();
//
//        Integer projectId = json.getInteger("id");
//
//        if (projectId != null) {
//            TstProject project = projectService.get(projectId);
//            if (authService.noProjectAndProjectGroupPrivilege(user.getId(), project)) {
//                return authFail();
//            }
//
//            TstProject vo = projectService.genVo(project, null);
//
//            if (TstProject.ProjectType.group.equals(project.getType())) {
//                vo.setLastestProjectGroup(projectService.isLastestProjectGroup(orgId, projectId));
//            }
//
//            ret.put("data", vo);
//        }
//        List<TstProject> groups = projectService.listProjectGroups(orgId);
//        ret.put("groups", groups);
//
//        ret.put("code", Constant.RespCode.SUCCESS.getCode());
        return ret;
    }

}
