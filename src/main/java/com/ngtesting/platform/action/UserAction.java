package com.ngtesting.platform.action;

import com.alibaba.fastjson.JSONObject;
import com.ngtesting.platform.config.Constant;
import com.ngtesting.platform.model.TstUser;
import com.ngtesting.platform.service.UserService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Controller;
import org.springframework.web.bind.annotation.*;

import javax.servlet.http.HttpServletRequest;
import java.util.HashMap;
import java.util.Map;

@Controller
@RequestMapping(value = Constant.API_PATH_CLIENT + "/user")
public class UserAction {
    @Autowired
    private UserService userService;

    @RequestMapping(value = "getProfile", method = RequestMethod.POST)
    @ResponseBody
    public Map<String, Object> getProfile(HttpServletRequest request, @RequestBody JSONObject json) {
        Map<String, Object> ret = new HashMap<String, Object>();

        TstUser TstUser = (TstUser) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_KEY);
//        Long orgId = TstUser.getDefaultOrgId();
//        Long prjId = TstUser.getDefaultPrjId();
//
//        Long orgIdNew = json.getLong("orgId");
//        Long prjIdNew = json.getLong("prjId");
//
//        if (orgIdNew != null && orgIdNew.longValue() != orgId.longValue()) { // org不能为空
//            orgService.setDefaultPers(orgId, TstUser);
//        }
//        if (prjIdNew != null && (prjId == null || prjIdNew.longValue() != prjId.longValue())) { // prj可能为空
//            projectService.viewPers(prjIdNew, TstUser);
//        }
//
//        Long userId = TstUser.getId();
//
//        Map<String, Boolean> sysPrivileges = sysPrivilegeService.listByUser(userId);
//        ret.put("sysPrivileges", sysPrivileges);
//
//        List<OrgVo> orgs = orgService.listVo(null, "false", userId);
//        ret.put("myOrgs", orgs);
//
//        Map<String, Boolean> orgPrivileges = orgRolePrivilegeService.listByUser(TstUser.getId(), orgId);
//        ret.put("orgPrivileges", orgPrivileges);
//
//        Map<String,Map<String,String>> casePropertyMap = casePropertyService.getMap(orgId);
//        ret.put("casePropertyMap", casePropertyMap);
//
//        List<TestProjectAccessHistoryVo> recentProjects = projectService.listRecentProjectVo(orgId, userId);
//        ret.put("recentProjects", recentProjects);
//        TstUser.setDefaultPrjId(recentProjects.size() > 0?recentProjects.get(0).getProjectId(): null);
//
//        Map<String, Boolean> prjPrivileges = projectPrivilegeService.listByUserPers(userId, prjId, orgId);
//        ret.put("prjPrivileges", prjPrivileges);

        ret.put("profile", TstUser);
        ret.put("code", Constant.RespCode.SUCCESS.getCode());

        return ret;
    }

    @ResponseBody
    @PostMapping("/query")
    public Object query(
            @RequestParam(name = "pageNum", required = false, defaultValue = "1")
                    int pageNum,
            @RequestParam(name = "pageSize", required = false, defaultValue = "10")
                    int pageSize){
        return userService.query(pageNum, pageSize);
    }

    @ResponseBody
    @RequestMapping("/get")
    public Object get(@RequestBody Integer id) {
        Map<String, Object> ret = new HashMap();
        TstUser po = userService.get(id);

        ret.put("code", Constant.RespCode.SUCCESS.getCode());
        ret.put("data", po);
        return ret;
    }

}
