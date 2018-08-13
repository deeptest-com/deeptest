package com.ngtesting.platform.action;

import com.alibaba.fastjson.JSON;
import com.alibaba.fastjson.JSONArray;
import com.alibaba.fastjson.JSONObject;
import com.github.pagehelper.Page;
import com.github.pagehelper.PageHelper;
import com.ngtesting.platform.config.Constant;
import com.ngtesting.platform.model.TstOrg;
import com.ngtesting.platform.model.TstOrgGroupUserRelation;
import com.ngtesting.platform.model.TstProjectAccessHistory;
import com.ngtesting.platform.model.TstUser;
import com.ngtesting.platform.service.*;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Controller;
import org.springframework.web.bind.annotation.*;

import javax.servlet.http.HttpServletRequest;
import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

@Controller
@RequestMapping(value = Constant.API_PATH_CLIENT + "/user")
public class UserAction {
    @Autowired
    private UserService userService;

    @Autowired
    private OrgService orgService;

    @Autowired
    OrgGroupUserRelationService orgGroupUserRelationService;

    @Autowired
    private ProjectService projectService;

    @Autowired
    SysPrivilegeService sysPrivilegeService;
    @Autowired
    OrgPrivilegeService orgRolePrivilegeService;
    @Autowired
    CasePropertyService casePropertyService;
    @Autowired
    ProjectPrivilegeService projectPrivilegeService;

    @PostMapping(value = "list")
    @ResponseBody
    public Map<String, Object> list(HttpServletRequest request, @RequestBody JSONObject json) {
        Map<String, Object> ret = new HashMap<String, Object>();

        TstUser user = (TstUser) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_KEY);
        Integer orgId = user.getDefaultOrgId();

        String keywords = json.getString("keywords");
        Boolean disabled = json.getBoolean("disabled");
        Integer pageNum = json.getInteger("page");
        Integer pageSize = json.getInteger("pageSize");

        Page page = PageHelper.startPage(pageNum, pageSize);
        List<TstUser> users = userService.list(orgId, keywords, disabled, pageNum, pageSize);

        ret.put("total", page.getTotal());
        ret.put("data", users);
        ret.put("code", Constant.RespCode.SUCCESS.getCode());
        return ret;
    }

    @PostMapping(value = "getUsers")
    @ResponseBody
    public Map<String, Object> getUsers(HttpServletRequest request, @RequestBody JSONObject json) {
        Map<String, Object> ret = new HashMap<String, Object>();

        TstUser user = (TstUser) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_KEY);
        Integer orgId = user.getDefaultOrgId();

        Integer projectId = json.getInteger("projectId");

        List <TstUser> vos = userService.getProjectUsers(orgId, projectId);

        ret.put("data", vos);
        ret.put("code", Constant.RespCode.SUCCESS.getCode());
        return ret;
    }

    @RequestMapping(value = "get", method = RequestMethod.POST)
    @ResponseBody
    public Map<String, Object> get(HttpServletRequest request, @RequestBody JSONObject json) {
        Map<String, Object> ret = new HashMap<String, Object>();

        TstUser user = (TstUser) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_KEY);
        Integer orgId = user.getDefaultOrgId();
        Integer userId = json.getInteger("id");

        List<TstOrgGroupUserRelation> relations = orgGroupUserRelationService.listRelationsByUser(orgId, userId);

        if (userId == null) {
            ret.put("user", new TstUser());
            ret.put("relations", relations);
            ret.put("code", Constant.RespCode.SUCCESS.getCode());
            return ret;
        }

        TstUser po = userService.get(userId);

        ret.put("user", po);
        ret.put("relations", relations);
        ret.put("code", Constant.RespCode.SUCCESS.getCode());
        return ret;
    }

    @PostMapping(value = "getProfile")
    @ResponseBody
    public Map<String, Object> getProfile(HttpServletRequest request, @RequestBody JSONObject json) {
        Map<String, Object> ret = new HashMap<String, Object>();

        TstUser user = (TstUser) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_KEY);
        Integer orgId = user.getDefaultOrgId();
        Integer prjId = user.getDefaultPrjId();

        Integer orgIdNew = json.getInteger("orgId");
        Integer prjIdNew = json.getInteger("prjId");

        if (orgIdNew != null && orgIdNew.longValue() != orgId.longValue()) { // org不能为空
            userService.setDefaultOrg(user, orgId);
        }
        if (prjIdNew != null && (prjId == null || prjIdNew.longValue() != prjId.longValue())) { // prj可能为空
            projectService.viewPers(prjIdNew, user);
        }

        Integer userId = user.getId();

        Map<String, Boolean> sysPrivileges = sysPrivilegeService.listByUser(userId);
        ret.put("sysPrivileges", sysPrivileges);

        List<TstOrg> orgs = orgService.listByUser(userId);
        ret.put("myOrgs", orgs);

        Map<String, Boolean> orgPrivileges = orgRolePrivilegeService.listByUser(user.getId(), orgId);
        ret.put("orgPrivileges", orgPrivileges);

        Map<String,Map<String,String>> casePropertyMap = casePropertyService.getMap(orgId);
        ret.put("casePropertyMap", casePropertyMap);

        List<TstProjectAccessHistory> recentProjects = projectService.listRecentProject(orgId, userId);
        ret.put("recentProjects", recentProjects);

//        user.setDefaultPrjId(recentProjects.size() > 0?recentProjects.getDetail(0).getPrjId(): null);

        Map<String, Boolean> prjPrivileges = projectPrivilegeService.listByUser(userId, prjId, orgId);
        ret.put("prjPrivileges", prjPrivileges);

        ret.put("profile", user);
        ret.put("code", Constant.RespCode.SUCCESS.getCode());

        return ret;
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

    @RequestMapping(value = "invite")
    @ResponseBody
    public Map<String, Object> invite(HttpServletRequest request, @RequestBody JSONObject json) {
        Map<String, Object> ret = new HashMap<String, Object>();

        TstUser user = (TstUser) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_KEY);

        TstUser vo = JSON.parseObject(JSON.toJSONString(json.get("user")), TstUser.class);
        List<TstOrgGroupUserRelation> relations = (List<TstOrgGroupUserRelation>) json.get("relations");
        TstUser po = userService.invitePers(user, vo, relations);

        if (po == null) {
            ret.put("code", Constant.RespCode.BIZ_FAIL.getCode());
            ret.put("msg", "邮箱已加入当期组织");
            return ret;
        }

        ret.put("code", Constant.RespCode.SUCCESS.getCode());
        return ret;
    }

    @PostMapping(value = "update")
    @ResponseBody
    public Map<String, Object> update(HttpServletRequest request, @RequestBody JSONObject json) {
        Map<String, Object> ret = new HashMap<String, Object>();

        TstUser user = (TstUser) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_KEY);
        Integer orgId = user.getDefaultOrgId();

        TstUser vo = JSON.parseObject(JSON.toJSONString(json.get("user")), TstUser.class);

        TstUser existUser = userService.getByEmail(vo.getEmail());

        if (existUser != null && existUser.getId() != vo.getId()) {
            ret.put("code", Constant.RespCode.BIZ_FAIL.getCode());
            ret.put("msg", "邮箱已被占用");
            return ret;
        }

        userService.update(vo);

        List<TstOrgGroupUserRelation> relations = (List<TstOrgGroupUserRelation>) json.get("relations");
        orgGroupUserRelationService.saveRelationsForUser(orgId, vo.getId(), relations);
        ret.put("code", Constant.RespCode.SUCCESS.getCode());

        return ret;
    }

    @PostMapping(value = "search")
    @ResponseBody
    public Map<String, Object> search(HttpServletRequest request, @RequestBody JSONObject json) {
        Map<String, Object> ret = new HashMap<String, Object>();

        TstUser user = (TstUser) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_KEY);

        Integer orgId = json.getInteger("orgId");
        String keywords = json.getString("keywords");
        JSONArray exceptIds = json.getJSONArray("exceptIds");

        String ids = "";
        if (exceptIds != null && exceptIds.size() > 0) {
            int i = 0;
            for (Object item : exceptIds.toArray()) {
                if (i++ > 0) {
                    ids += ",";
                }
                ids += item.toString();
            }
        }

        List users = userService.search(orgId, keywords, ids);

        List<Object> vos = new ArrayList<>();
        vos.addAll(users);

        ret.put("data", vos);
        ret.put("code", Constant.RespCode.SUCCESS.getCode());
        return ret;
    }

    @PostMapping(value = "setLeftSize")
    @ResponseBody
    public Map<String, Object> setLeftSize(HttpServletRequest request, @RequestBody JSONObject json) {
        Map<String, Object> ret = new HashMap<String, Object>();

        TstUser user = (TstUser) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_KEY);

        Integer left = json.getInteger("left");
        String prop = json.getString("prop");

        user = userService.setLeftSizePers(user, left, prop);

        ret.put("data", user);
        ret.put("code", Constant.RespCode.SUCCESS.getCode());
        return ret;
    }

}
